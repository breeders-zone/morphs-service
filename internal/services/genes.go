package services

import (
	"github.com/breeders-zone/morphs-service/internal/domain"
	"github.com/breeders-zone/morphs-service/internal/repositories"
)

type GenesService struct {
	repo repositories.Genes	
}

func NewGenesService(repo repositories.Genes) *GenesService {
	return &GenesService{
		repo,
	}
}

func (s *GenesService) GetById(id int) (*domain.Gene, error) {
	return s.repo.GetById(id)
}


func (s *GenesService) Create(g *domain.Gene) (*domain.Gene, error) {
	return s.repo.Create(g)
}


func (s *GenesService) Update(g *domain.Gene) (*domain.Gene, error) {
	return s.repo.Update(g)
}


func (s *GenesService) Delete(id int) error {
	return s.repo.Delete(id)
}
