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

type ESClient struct {
	url       string
	indexName string
	typeName  string
}

func NewESClient(url, indexName, typeName string) ESClientServiceIf {
	es := &ESClient{
		url:       url,
		indexName: indexName,
		typeName:  typeName,
	}
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

	client, err := elastic.NewClient(elastic.SetURL(es.url),
		elastic.SetSniff(false),
		elastic.SetHealthcheck(false))

	fmt.Printf("[getESClient] ES initialized... err = %v \n", err)

	return client, err
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
