package ElasticSearchService

type Client interface {
	Index(indexName string, documentJson string) error

	Search(indexName string, query map[string]interface{}) (rawResult []byte, err error)

	Delete(indexName string, docID string) error

	Get(indexName string, id string) (rawResult []byte, err error)

	Update(indexName string, id string, documentJson string) error

	DeteleIndex(indexName string) error
}
