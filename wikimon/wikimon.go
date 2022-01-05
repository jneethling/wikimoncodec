package wikimon

import (
	"encoding/json"

	"github.com/linkedin/goavro"
	"go.uber.org/zap"
)

// Wikimon defines wikimon structure
type Wikimon struct {
	Action      string   `json:"action,omitempty"`
	Change      int64    `json:"change_size,omitempty"`
	Geo         GeoIP    `json:"geo_ip,omitempty"`
	Hastags     []string `json:"hastags,omitempty"`
	Anon        bool     `json:"is_anon,omitempty"`
	Bot         bool     `json:"is_bot,omitempty"`
	Minor       bool     `json:"is_minor,omitempty"`
	New         bool     `json:"is_new,omitempty"`
	Unpatrolled bool     `json:"is_unpatrolled,omitempty"`
	Mentions    []string `json:"mentions,omitempty"`
	Ns          string   `json:"ns,omitempty"`
	PageTitle   string   `json:"page_title,omitempty"`
	ParentRevID string   `json:"parent_rev_id,omitempty"`
	RevID       string   `json:"rev_id,omitempty"`
	Summary     string   `json:"summary,omitempty"`
	URL         string   `json:"url,omitempty"`
	User        string   `json:"user,omitempty"`
}

// GeoIP defines geo_ip structure
type GeoIP struct {
	City      string  `json:"city,omitempty"`
	Country   string  `json:"country_name,omitempty"`
	Latitude  float64 `json:"latitude,omitempty"`
	Longitude float64 `json:"longitude,omitempty"`
	Region    string  `json:"region_name,omitempty"`
}

// AvroCodec defines avro codec structure
type AvroCodec struct {
	*goavro.Codec
}

func WikiAvroCodec() (*AvroCodec, error) {
	codec, err := goavro.NewCodec(AvroSchema)
	return &AvroCodec{Codec: codec}, err
}

func (c *AvroCodec) Unmarshal(msg []byte, logger *zap.Logger) (*Wikimon, error) {
	var wikimon *Wikimon
	native, _, err := c.NativeFromBinary(msg)
	if err != nil {
		logger.Error("Error converting from binary to native: " + err.Error())
		return nil, err
	}

	data, err := json.Marshal(native)
	if err != nil {
		logger.Error("Error marshalling native to JSON")
		return nil, err
	}

	err = json.Unmarshal(data, &wikimon)
	return wikimon, err
}
