package domain

import (
	"fmt"
	"strconv"
	"time"

	"github.com/breeders-zone/morphs-service/utils"
)

type Gene struct {
	Id           int       `json:"-"`
	Title        string    `json:"title"`
	Type         string    `json:"type"`
	ProducedName string    `json:"producedName" db:"produced_name"`
	ProducedDate time.Time `json:"producedDate" db:"produced_date"`
	Availability string    `json:"availability"`
	Description  string    `json:"description"`
	History      string    `json:"history"`
	Links        Links     `json:"links"`

	CreatedAt time.Time `json:"createdAt" db:"created_at"`
}

func (g *Gene) GetID() string {
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
