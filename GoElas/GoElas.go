package main


import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/teris-io/shortid"
	"github.com/olivere/elastic/v7"
	"log"
	"reflect"
	"time"
)


const (
	elasticIndexName = "documents"
)
var (
	elasticClient *elastic.Client
)
type DocumentRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type Document struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"created_at"`
	Content   string    `json:"content"`
}
func main() {
	var err error
	ctx := context.Background()
	elasticClient, err = elastic.NewClient(
		elastic.SetURL("http://127.0.0.1:9200"),
		elastic.SetSniff(false),
	)
	if err != nil {
		panic(err)
	}
	//_, err = elasticClient.CreateIndex(elasticIndexName).Do(ctx)
	if err != nil {
		fmt.Println(err)
	}
	_,err = elasticClient.DeleteIndex(elasticIndexName).Do(ctx)
	if err != nil {
		fmt.Println(err)
	}
	indexParams := `{
		"settings":{
			"number_of_shards":1,
			"number_of_replicas":0


		},
		"mappings":{
				"properties": {
					"content": {
"type":"text",
 "fielddata": true
					},
                    "title": {
"type":"text",
 "fielddata": true
					}
				}
			}
		}
	}`

	// Create an index
	 _,err = elasticClient.CreateIndex(elasticIndexName).BodyString(indexParams).Do(ctx)
	fmt.Println(indexParams)
	if err != nil {
		// Handle error
		// Get *elastic.Error which contains additional information
		e, ok := err.(*elastic.Error)
		if !ok {
		}
		log.Printf("Elastic failed with status %d and error %s.", e.Status, e.Details)
		panic(err)
	}
	if err != nil {
		fmt.Println(err)
	}
	var docs []DocumentRequest
	var doc1 DocumentRequest
	doc1.Content = "Content Bonjour"
	doc1.Title = "Title Bonjour"
	docs = append(docs,doc1)
fmt.Println(docs)

	for _, d := range docs {
		doc := Document{
			ID:        shortid.MustGenerate(),
			Title:     d.Title,
			CreatedAt: time.Now().UTC(),
			Content:   d.Content,
		}
		docJ,_ := json.Marshal(doc)
		var put2 *elastic.IndexResponse
		fmt.Println(string(docJ))
		put2, err = elasticClient.Index().
			Index(elasticIndexName ).
			Id(doc.ID).
			BodyString(string(docJ)).
		    Refresh("true").
			Do(ctx)
		if err != nil {
			panic(err)
		}
		fmt.Println("rajout : ")
		fmt.Println(put2)


	}
	//termQuery := elastic.NewTermQuery("title", "Bonjour")
	termQuery := elastic.NewMatchQuery("content","1 Rue Des fontaines 91150")
	var result *elastic.SearchResult
	result, err = elasticClient.Search().
		Index("adresses").            // search in index "documents"
		Query(termQuery).           // specify the query
		Sort("content", true). // sort by "user" field, ascending
		From(0).Size(10).           // take documents 0-9
		Pretty(true).               // pretty print request and response JSON
		Do(ctx)

	if err != nil {
		// Handle error
		// Get *elastic.Error which contains additional information
		e, ok := err.(*elastic.Error)
		if !ok {
		}
		log.Printf("Elastic failed with status %d and error %s.", e.Status, e.Details)
		panic(err)
	}
	fmt.Println("result")
	fmt.Println(result)
	fmt.Printf("Found a total of %d docs\n", result.TotalHits())
	var ttyp DocumentRequest
	for _, item := range result.Each(reflect.TypeOf(ttyp)) {
		if t, ok := item.(DocumentRequest); ok {
			fmt.Printf("Doc %s: %s\n", t.Title, t.Content)
		}
	}
}