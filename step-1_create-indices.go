// Go code to create index mappings from the Kibana Getting Started Tutorial:
// https://www.elastic.co/guide/en/kibana/current/tutorial-load-dataset.html
package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"

	"github.com/grokify/elastirad-go"
	"github.com/grokify/elastirad-go/models"
	"github.com/grokify/elastirad-go/models/v5"
	"github.com/valyala/fasthttp"
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
		Method: "PUT",
		Path:   []interface{}{path},
		Body:   esBody}

	res, req, err := esClient.SendFastRequest(esReq)

	if err != nil {
		fmt.Sprintf("ERROR Creating %v Mapping: %v\n", path, err)
	} else if res.StatusCode() >= 400 {
		err = errors.New(fmt.Sprintf("ERROR Creating %v Mapping: %v\n", path, res.StatusCode()))
		fmt.Println(err)
	} else {
		fmt.Printf("SUCCESS Creating %v Mapping: %v\n", path, res.StatusCode())
	}

	fasthttp.ReleaseRequest(req)
	fasthttp.ReleaseResponse(res)
	return err
}

func main() {
	client := elastirad.NewClient(url.URL{})

	createMapping(client, shakespearePath, shakespeareMappings)
	createMapping(client, logstashPath, logstashMappings)

	fmt.Println("DONE")
}
