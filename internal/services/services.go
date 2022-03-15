package services

import (
	"github.com/breeders-zone/morphs-service/internal/domain"
	"github.com/breeders-zone/morphs-service/internal/repositories"
)

type Services struct {
	Genes Genes
}

type Genes interface {
	GetAll(order []string) ([]domain.Gene, error)
	GetById(id int) (*domain.Gene, error)
	Create(g *domain.Gene) (*domain.Gene, error)
	Update(g *domain.Gene) (*domain.Gene, error)
	Delete(id int) error
}


func NewServices(repos *repositories.Repositories) *Services {
	return &Services{
		Genes: NewGenesService(repos.Genes),
	}
}