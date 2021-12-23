package wikimon

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestGeoIp(t *testing.T) {
	logger, _ := zap.NewProduction()

	var wikimon = []byte(`
	{
		"action": "edit",
		"change_size": -12,
		"flags": null,
		"geo_ip": {
		  "city": "Salisbury",
		  "country_name": "United States",
		  "latitude": 38.3761,
		  "longitude": -75.6086,
		  "region_name": "Maryland"
		},
		"hashtags": [],
		"is_anon": true,
		"is_bot": false,
		"is_minor": false,
		"is_new": false,
		"is_unpatrolled": false,
		"mentions": [],
		"ns": "Main",
		"page_title": "Evanescence (Evanescence album)",
		"parent_rev_id": "775894800",
		"rev_id": "774995266",
		"summary": "/* Credits and personnel */ \"Personnel\" is sufficient",
		"url": "https://en.wikipedia.org/w/index.php?diff=775894800&oldid=774995266",
		"user": "71.200.123.192"
	  }
	`)

	var m = make(map[string]interface{})
	json.Unmarshal(wikimon, &m)

	var wikiExpected *Wikimon
	json.Unmarshal(wikimon, &wikiExpected)

	codec, err := WikiAvroCodec()
	assert.Nil(t, err)

	binary, err := codec.BinaryFromNative(nil, m)
	if err != nil {
		logger.Fatal("Failed to convert Go map to Avro binary: " + err.Error())
	}

	wikiObserved, err := codec.Unmarshal(binary, logger)
	if err != nil {
		logger.Fatal("Failed to convert Avro binary to Golang type")
	}

	assert.EqualValues(t, wikiExpected, wikiObserved)

}

func TestNoGeoIp(t *testing.T) {
	logger, _ := zap.NewProduction()

	var wikimon = []byte(`
	{
		"action": "edit",
		"change_size": 19,
		"flags": "M",
		"hashtags": [],
		"is_anon": false,
		"is_bot": false,
		"is_minor": true,
		"is_new": false,
		"is_unpatrolled": false,
		"mentions": [],
		"ns": "User talk",
		"page_title": "User talk:Manxruler",
		"parent_rev_id": "775894803",
		"rev_id": "775894650",
		"summary": "/* The battle of Kristiansand (1940) */",
		"url": "https://en.wikipedia.org/w/index.php?diff=775894803&oldid=775894650",
		"user": "Carsten R D"
	  }
	`)

	var m = make(map[string]interface{})
	json.Unmarshal(wikimon, &m)

	var wikiExpected *Wikimon
	json.Unmarshal(wikimon, &wikiExpected)

	codec, err := WikiAvroCodec()
	assert.Nil(t, err)

	binary, err := codec.BinaryFromNative(nil, m)
	if err != nil {
		logger.Fatal("Failed to convert Go map to Avro binary: " + err.Error())
	}

	wikiObserved, err := codec.Unmarshal(binary, logger)
	if err != nil {
		logger.Fatal("Failed to convert Avro binary to Golang type")
	}

	assert.EqualValues(t, wikiExpected, wikiObserved)

}
