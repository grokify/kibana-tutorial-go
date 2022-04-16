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
		Method: http.MethodPut,
		Path:   []interface{}{path},
		Body:   esBody}

	res, req, err := esClient.SendFastRequest(esReq)

	if err != nil {
		return fmt.Errorf("error creating [%v] Mapping [%v]", path, err)
	} else if res.StatusCode() >= 400 {
		return fmt.Errorf("error creating [%v] mapping: [%v]", path, res.StatusCode())
	} else {
		fmt.Printf("success creating [%v] mapping [%v]", path, res.StatusCode())
	}

	fasthttp.ReleaseRequest(req)
	fasthttp.ReleaseResponse(res)
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
