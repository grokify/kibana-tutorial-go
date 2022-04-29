// Go code to create index mappings from the Kibana Getting Started Tutorial:
// https://www.elastic.co/guide/en/kibana/current/tutorial-load-dataset.html
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	elastirad "github.com/grokify/elastirad-go"
	"github.com/grokify/elastirad-go/models"
	v5 "github.com/grokify/elastirad-go/models/v5"
	"github.com/grokify/mogo/log/logutil"
)

const (
	shakespearePath     = "/shakespeare"
	shakespeareMappings = `{
 "mappings" : {
  "_default_" : {
   "properties" : {
    "speaker" : {"type": "keyword" },
    "play_name" : {"type": "keyword" },
    "line_id" : { "type" : "integer" },
    "speech_number" : { "type" : "integer" }
   }
  }
 }
}`
	logstashPath     = "/logstash-2015.05.18"
	logstashMappings = `{
  "mappings": {
    "log": {
      "properties": {
        "geo": {
          "properties": {
            "coordinates": {
              "type": "geo_point"
            }
          }
        }
      }
    }
  }
}`
)

func createMapping(esClient elastirad.Client, path string, mappingsBody string) error {
	esBody := v5.CreateIndexBody{}
	err := json.Unmarshal([]byte(mappingsBody), &esBody)
	if err != nil {
		return err
	}

	esReq := models.Request{
		Method: http.MethodPut,
		Path:   []interface{}{path},
		Body:   esBody}

	res, err := esClient.SendRequest(esReq)

	if err != nil {
		return fmt.Errorf("error creating mapping for path [%s] error [%s]", path, err)
	} else if res.StatusCode >= 300 {
		return fmt.Errorf("error creating mapp for path [%s] http status [%d]", path, res.StatusCode)
	}

	fmt.Printf("success creating [%s] mapping [%d]", path, res.StatusCode)
	return nil
}

func main() {
	client := elastirad.NewClient(url.URL{})

	err := createMapping(client, shakespearePath, shakespeareMappings)
	logutil.FatalErr(err)

	err = createMapping(client, logstashPath, logstashMappings)
	logutil.FatalErr(err)

	fmt.Println("DONE")
}
