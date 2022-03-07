package repositories

import (
	"github.com/breeders-zone/morphs-service/internal/domain"
	"github.com/jmoiron/sqlx"
)

type Repositories struct {
	Genes Genes
}

type Genes interface {
	Create(g *domain.Gene) (*domain.Gene, error)
	Update(g *domain.Gene) (*domain.Gene, error)
	Delete(id int) (error)
	GetById(id int) (*domain.Gene, error)
}

func NewRepositories(db *sqlx.DB) *Repositories {
	return &Repositories{
		Genes: NewGenesRepo(db),
	}
}