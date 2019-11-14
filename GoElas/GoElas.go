package main

import (
	"log"
	"time"
)
main

import (
"encoding/json"
"fmt"
"log"
"net/http"
"strconv"
"time"

"github.com/gin-gonic/gin"
"github.com/olivere/elastic"
"github.com/teris-io/shortid"
)



var (
	elasticClient *elastic.Client
)

func main() {
	var err error
	for {
		elasticClient, err = elastic.NewClient(
			elastic.SetURL("http://elasticsearch:9200"),
			elastic.SetSniff(false),
		)
		if err != nil {
			log.Println(err)
			time.Sleep(3 * time.Second)
		} else {
			break
		}
	}
	// ...
}