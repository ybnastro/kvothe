package interfaces

type IElasticsearch interface {
	IndexDocument(indexName string, documentID int64, document string) error
	Get(indexName string, documentID int64) (string, error)
}
