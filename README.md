[![Go](https://github.com/jneethling/wikimoncodec/actions/workflows/go.yml/badge.svg)](https://github.com/jneethling/wikimoncodec/actions/workflows/go.yml)
# wikimoncodec
A Golang Apache Avro Codec for the Hatnote Wikipedia monitor (https://github.com/hatnote/wikimon).

This module can be imported into other Golang projects in order to simplify marshalling/unmarshalling data consumed from the wikimon websocket service for use in Apache Kafka, for example. This is useful if you plan to have multiple seperate producers & consumers of serialised data, using a shared codec. A possible use case for this type of scenario might be to profile wikipedia editors and/or pages pro-actively for some or other machine learning application.

For a simple Kafka producer example see https://github.com/jneethling/wikimoningest.