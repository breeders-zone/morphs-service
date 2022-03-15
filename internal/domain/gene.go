package domain

import (
	"fmt"
	"strconv"
	"time"

	"github.com/breeders-zone/morphs-service/utils"
)

type Gene struct {
	Id           int       `json:"-" selector:"id"`
	Title        string    `json:"title,omitempty" selector:"title"`
	Type         string    `json:"type,omitempty" selector:"type"`
	ProducedName string    `json:"producedName,omitempty" selector:"producedName" db:"produced_name"`
	ProducedDate *time.Time `json:"producedDate,omitempty" selector:"producedDate" db:"produced_date"`
	Availability string    `json:"availability,omitempty" selector:"availability"`
	Description  string    `json:"description,omitempty" selector:"description"`
	History      string    `json:"history,omitempty" selector:"history"`
	Links        Links     `json:"links,omitempty" selector:"links"`

	CreatedAt *time.Time `json:"createdAt,omitempty" db:"created_at"`
}

func (g Gene) GetID() string {
	return fmt.Sprint(g.Id)
}

func (g *Gene) SetID(sId string) error {
	id, err := strconv.Atoi(sId)
	if err != nil {
		return err
	}

	g.Id = id
	return nil
}

type Links []string

func (r *Links) Scan(src interface{}) error {
	return utils.ParseJSONToModel(src, r)
}
