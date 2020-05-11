package ESClientService

import (
	"context"
	"fmt"
	"log"

	elastic "gopkg.in/olivere/elastic.v7"
)

const indexString = `
{
	"settings":{
		"number_of_shards": 1,
		"number_of_replicas": 0
	},
	"mappings":{
		"_doc": {
		}
	}
}`

// type Doc struct {
// }

// // default _doc
// type Mapping struct {
// 	Doc *Doc `json:"_doc"`
// }

// type Setting struct {
// 	Number_of_shards   int `json:"number_of_shards"`
// 	Number_of_replicas int `json:"number_of_replicas"`
// }

// type Index struct {
// 	Mappings *Mapping `json:"mappings"`
// 	Settings *Setting `json:"settings"`
// }

// func makeIndexString() *Index {
// 	index := &Index{
// 		Settings: &Setting{
// 			Number_of_shards:   1,
// 			Number_of_replicas: 0,
// 		},
// 		Mappings: &Mapping{
// 			Doc: &Doc{},
// 		},
// 	}

// 	return index
// }

type ESClient struct {
	url       string
	indexName string
	typeName  string
	client    *elastic.Client
}

func NewESClient(url, indexName, typeName string) ESClientServiceIf {
	client, err := elastic.NewClient(elastic.SetURL(url),
		elastic.SetSniff(false),
		elastic.SetHealthcheck(false))
	if err != nil {
		log.Println("ESClient err", err)
	}
	es := &ESClient{
		url:       url,
		indexName: indexName,
		typeName:  typeName,
		client:    client,
	}

	// make index if not existed
	// indexByte, _ := json.Marshal(makeIndexString())
	// indexString := string(indexByte)
	// indexString = strings.ReplaceAll(indexString, "default1", typeName)
	es.checkExistedIndex(indexString)

	return es
}

// kiểm tra xem có index hay chưa (index chính là database , type ứng với 1 table, doc ứng với 1 item)
func (es *ESClient) checkExistedIndex(indexString string) {
	ctx := context.Background()
	esclient, _ := es.getESClient()
	if esclient == nil {
		return
	}
	exists, err := esclient.IndexExists(es.indexName).Do(ctx)
	if err != nil {
		// Handle error

		log.Printf("[checkExistedIndex] err = %v \n", err)
		return
	}
	if !exists {
		// Create a new index.
		createIndex, err := esclient.CreateIndex(es.indexName).BodyString(indexString).Do(ctx)
		if err != nil {
			log.Printf("[checkExistedIndex] err = %v \n", err)
			return
		}

		log.Printf("[checkExistedIndex] createIndex = %v, %v \n", createIndex, err)
	}
}

func (es *ESClient) getESClient() (*elastic.Client, error) {

	if es.client == nil {
		client, err := elastic.NewClient(elastic.SetURL(es.url),
			elastic.SetSniff(false),
			elastic.SetHealthcheck(false))
		fmt.Printf("[getESClient] ES initialized... err = %v \n", err)
		es.client = client
	}

	return es.client, nil
}

func (es *ESClient) PutDataToES(id string, dataJson string) (err error) {
	ctx := context.Background()
	esclient, err := es.getESClient()

	if err != nil || esclient == nil {
		fmt.Printf("[PutDataToES] Error initializing : %v", err)
		return err
	}

	ind, err := esclient.Index().
		Index(es.indexName).
		Type(es.typeName).
		Id(id).
		BodyJson(dataJson).
		Do(ctx)

	if err != nil {
		fmt.Printf("[PutDataToES] err = %v \n", err)
		return err
	}
	fmt.Printf("[PutDataToES] ind=%v, err=%v \n", ind, err)
	return nil
}

func (es *ESClient) DeleteIndexES() {
	ctx := context.Background()
	esclient, _ := es.getESClient()
	if esclient == nil {
		return
	}
	// Delete an index.
	deleteIndex, err := esclient.DeleteIndex(es.indexName).Do(ctx)
	if err != nil {
		// Handle error
		fmt.Printf("[deleteIndexES] err = %v \n", err)
		return
	}
	esclient.Search()
	fmt.Println("[deleteIndexES] = ", deleteIndex)
	return
}

func (es *ESClient) DeleteDataES(id string) {
	ctx := context.Background()
	esclient, _ := es.getESClient()
	if esclient == nil {
		return
	}
	// Delete an index.
	deleteIndex, err := esclient.Delete().Index(es.indexName).Type(es.typeName).Id(id).Do(ctx)
	if err != nil {
		// Handle error
		fmt.Printf("[deleteDataES] err = %v \n", err)
		return
	}

	fmt.Println("[deleteDataES] = ", deleteIndex)
	return
}

func (es *ESClient) UpdateDataES(id string, mapUpdate map[string]interface{}) {
	ctx := context.Background()
	esclient, _ := es.getESClient()
	if esclient == nil {
		return
	}
	update, err := esclient.Update().Index(es.indexName).Type(es.typeName).Id(id).
		Doc(mapUpdate).
		Do(ctx)
	if err != nil {
		fmt.Printf("[updateDataES] err = %v \n", err)
		return
	}

	fmt.Println("[updateDataES] = ", update)
}

func (es *ESClient) SearchESByQuery(mapSearch map[string]interface{}, sort map[string]bool) ([]*elastic.SearchHit, error) {
	fmt.Printf("[SearchESByQuery] mapSearch = %v, sort = %v \n", mapSearch, sort)
	ctx := context.Background()
	esclient, _ := es.getESClient()
	searchSource := elastic.NewSearchSource()
	for k, v := range mapSearch {
		searchSource.Query(elastic.NewMatchQuery(k, v))
	}

	for k, v := range sort {
		searchSource.Sort(k, v)
	}

	searchService := esclient.Search().Index(es.indexName).SearchSource(searchSource)

	searchResult, err := searchService.Do(ctx)
	if err != nil || searchResult == nil || searchResult.Hits == nil {
		fmt.Println("[SearchESByQuery] Error = ", err)
		return []*elastic.SearchHit{}, err
	}

	// for _, v := range searchResult.Hits.Hits {
	// 	fmt.Printf("%s \n", string(v.Source))
	// }

	return searchResult.Hits.Hits, nil
}
