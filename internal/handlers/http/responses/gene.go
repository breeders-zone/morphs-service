package responses

import (
	"encoding/json"

	"github.com/breeders-zone/morphs-service/internal/domain"
	"github.com/manyminds/api2go/jsonapi"
)

type GeneResponse struct {
	Links    jsonapi.Links          `json:"links,omitempty" extensions:"x-omitempty"`
	Included []jsonapi.Data         `json:"included,omitempty" extensions:"x-omitempty"`
	Meta     map[string]interface{} `json:"meta,omitempty" extensions:"x-omitempty"`
	Data     struct {
		Type       string          `json:"type"`
		ID         string          `json:"id"`
		Attributes domain.Gene     `json:"attributes"`
		Links      jsonapi.Links   `json:"links,omitempty" extensions:"x-omitempty"`
		Meta       json.RawMessage `json:"meta,omitempty" extensions:"x-omitempty"`
	} `json:"data"`
}
