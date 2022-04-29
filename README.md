# Kibana Tutorial

This is a simple repo to get started using the Kibana Getting Started Tutorial.

1. Read the instructions [here](https://www.elastic.co/guide/en/kibana/current/tutorial-load-dataset.html) ([archive](https://web.archive.org/web/20190710200800/https://www.elastic.co/guide/en/kibana/current/tutorial-load-dataset.html)):
    1. [tutorial-load-dataset.asciidoc](tutorial-load-dataset.asciidoc) [`482cb4f`](https://github.com/elastic/kibana/blob/482cb4f603d56b06e3405aaebee95b571b2480fb/docs/getting-started/tutorial-load-dataset.asciidoc)
1. Run `step-1_create-indices.go`, e.g. `$ go run step-1_create-indices.go`
1. Run `step-2_bulk-load.sh`, e.g. `$ sh step-2_bulk-load.sh`
1. In Kibana, navigate to *Management > Index Patterns*. Then create the following index patterns:
    1. `shakes*`
    1. `ba*`
    1. `logstash-*` - Choose `@timestamp` as the Time Filter field name
1. In Kibana, navigate to *Discover* and you should see search results for each index filter.
