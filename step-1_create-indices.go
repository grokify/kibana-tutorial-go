// Go code to create index mappings from the Kibana Getting Started Tutorial:
// https://www.elastic.co/guide/en/kibana/current/tutorial-load-dataset.html
package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/grokify/mogo/log/logutil"
	"github.com/grokify/mogo/net/httputilmore"
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

func createMapping(path string, mappingsBody io.Reader) error {
	req, err := http.NewRequest(http.MethodPut, path, mappingsBody)
	if err != nil {
		return err
	}
	req.Header.Add(httputilmore.HeaderContentType, httputilmore.ContentTypeAppJSONUtf8)

	client := &http.Client{}
	res, err := client.Do(req)

	if err != nil {
		return fmt.Errorf("error creating mapping for path [%s] error [%s]", path, err)
	} else if res.StatusCode >= 300 {
		return fmt.Errorf("error creating mapp for path [%s] http status [%d]", path, res.StatusCode)
	}

	fmt.Printf("success creating [%s] mapping [%d]", path, res.StatusCode)
	return nil
}

func main() {
	err := createMapping(shakespearePath, strings.NewReader(shakespeareMappings))
	logutil.FatalErr(err)

	err = createMapping(logstashPath, strings.NewReader(logstashMappings))
	logutil.FatalErr(err)

	fmt.Println("DONE")
}
