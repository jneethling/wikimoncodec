# wikimoncodec
A Golang Apache Avro Codec for the Hatnote Wikipedia monitor (https://github.com/hatnote/wikimon).

This module can be imported into other Golang projects in order to simplify marshalling/unmarshalling data originally consumed from the wikimon websocket service for use in Apache Kafka, for example. This is useful if you plan to have multiple seperate producers & consumers of binary data, using a shared codec. A possible use case for this type of scenario might be to profile wikipedia editors and/or edit patterns pro-actively for some or other machine learning application.