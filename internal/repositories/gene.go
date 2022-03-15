package repositories

import (
	"encoding/json"
	"fmt"

	"github.com/breeders-zone/morphs-service/internal/domain"
	"github.com/breeders-zone/morphs-service/utils"
	"github.com/jmoiron/sqlx"
)

type GenesRepo struct {
	db *sqlx.DB
}

func NewGenesRepo(db *sqlx.DB) *GenesRepo {
	return &GenesRepo{
		db,
	}
}

func (r *GenesRepo) GetAll(order []string) ([]domain.Gene, error) {

	orderStr := utils.GenerateOrderString(order)

	g := []domain.Gene{}
	err := r.db.Select(&g, "SELECT * FROM genes " + orderStr)

	if err != nil {
		return nil, err
	}

	return g, nil
}

func (r *GenesRepo) GetById(id int) (*domain.Gene, error) {
	g := domain.Gene{}
	err := r.db.Get(&g, "SELECT * FROM genes WHERE id=$1", id)

	if err != nil {
		return nil, err
	}

	return &g, nil
}

func (r *GenesRepo) Create(g *domain.Gene) (*domain.Gene, error) {
	links, err := json.Marshal(g.Links)
	if err != nil {
		return nil, err
	}

	var id int

	rows, err := r.db.NamedQuery(`
		INSERT INTO genes (
			title, 
			type, 
			produced_name, 
			produced_date,
			availability, 
			description, 
			history, 
			links
		) VALUES (
			:title,
			:type,
			:produced_name,
			:produced_date,
			:availability,
			:description,
			:history,
			:links
		) RETURNING id`,
		map[string]interface{}{
			"title":         g.Title,
			"type":          g.Type,
			"produced_name": g.ProducedName,
			"produced_date": g.ProducedDate,
			"availability":  g.Availability,
			"description":   g.Description,
			"history":       g.History,
			"links":         string(links),
		})

	if err != nil {
		return nil, err
	}

	if rows.Next() {
		rows.Scan(&id)
	}

	g.Id = id

	return g, nil
}

func (r *GenesRepo) Update(g *domain.Gene) (*domain.Gene, error) {
	links, err := json.Marshal(g.Links)
	if err != nil {
		return nil, err
	}

	rows, err := r.db.NamedQuery(`
		UPDATE genes SET
			title=:title, 
			type=:type, 
			produced_name=:produced_name, 
			produced_date=:produced_date,
			availability=:availability, 
			description=:description, 
			history=:history, 
			links=:links
		WHERE id = :id RETURNING *`,
		map[string]interface{}{
			"id":            g.Id,
			"title":         g.Title,
			"type":          g.Type,
			"produced_name": g.ProducedName,
			"produced_date": g.ProducedDate,
			"availability":  g.Availability,
			"description":   g.Description,
			"history":       g.History,
			"links":         string(links),
		})

	if err != nil {
		return nil, err
	}

	updatedG := new(domain.Gene)
	rows.Next()
	err = rows.Scan(&updatedG)
	if err != nil {
		return nil, err
	}

	return updatedG, nil
}

func (r *GenesRepo) Delete(id int) error {
	res, err := r.db.NamedExec(
		`DELETE FROM genes WHERE id = :id RETURNING *`,
		map[string]interface{}{
			"id": id,
		})

	if err != nil {
		return err
	}

	rowsCount, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsCount == 0 {
		return fmt.Errorf("node does not exist")
	}

	return nil
}
