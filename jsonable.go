package bebe

import "encoding/json"

type JSON []byte

type JSONEncodable interface {
	ToJSON() (JSON, error)
}

type JSONDecodable interface {
	FromJSON(JSON) error
}

type JSONDecoderDecodable interface {
	FromJSONDecoder(*json.Decoder) error
}
