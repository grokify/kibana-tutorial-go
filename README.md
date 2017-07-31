# Kibana Tutorial

This is a simple repo to get started using the Kibana Getting Started Tutorial.

1. Read the instructions here:
  1. https://www.elastic.co/guide/en/kibana/current/tutorial-load-dataset.html
  1. https://github.com/elastic/kibana/blob/master/docs/getting-started/tutorial-load-dataset.asciidoc
1. Run `step-1_create-indices.go`, e.g. `$ go run step-1_create-indices.go`
1. Run `step-2_bulk-load.sh`, e.g. `$ sh step-2_bulk-load.sh`
1. In Kibana, navigate to *Management > Index Patterns*. Then create the following index patterns:
  1. `shakes*`
  1. `ba*`
  1. `logstash-*` - Choose `@timestamp` as the Time Filter field name
1. In Kibana, navigate to *Discover* and you should see search results for each index filter.