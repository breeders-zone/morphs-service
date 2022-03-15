package selector_test

import (
	"encoding/json"
	"testing"

	"github.com/breeders-zone/morphs-service/pkg/selector"
	"github.com/stretchr/testify/assert"
)

type A struct {
	ID int `json:"id,omitempty" selector:"id"`
	Title string `json:"title,omitempty" selector:"title"`
	Type  string `json:"type,omitempty" selector:"type"`
}

func Test_SelectorTest(t *testing.T) {
	g := &A{ID: 2, Title: "123", Type: "321"}

	res := selector.SelectFields(g, "id", "title").(A)

	gj, _ := json.Marshal(res)

	assert.Equalf(t, string(gj), `{"id":2,"title":"123"}`, "Test selector")
}