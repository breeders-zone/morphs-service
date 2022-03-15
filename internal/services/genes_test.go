package services_test

import (
	"errors"
	"testing"
	"time"

	"github.com/breeders-zone/morphs-service/internal/domain"
	"github.com/breeders-zone/morphs-service/internal/services"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	mockRepos "github.com/breeders-zone/morphs-service/internal/repositories/mocks"
)

var errInternalServErr = errors.New("test: internal server error")

var gene = &domain.Gene{
	Id:           1,
	Title:        "test",
	Type:         "type",
	ProducedName: "name",
	ProducedDate: &time.Time{},
	Availability: "high",
	Description:  "desc",
	History:      "hist",
	Links:        []string{"1", "2"},
	CreatedAt:    &time.Time{},
}

func mockGenesService(t *testing.T) (*services.GenesService, *mockRepos.MockGenes) {
	t.Helper()

	mockCtl := gomock.NewController(t)
	defer mockCtl.Finish()

	mockGenesRepo := mockRepos.NewMockGenes(mockCtl)

	geneService := services.NewGenesService(
		mockGenesRepo,
	)

	return geneService, mockGenesRepo
}

func Test_GenesService_GetAll(t *testing.T) {
	genesService, mockGenesRepo := mockGenesService(t)

	mockGenesRepo.EXPECT().GetAll(gomock.Any()).Return([]domain.Gene{{Id: 5}, {Id: 6}}, nil)

	res, err := genesService.GetAll([]string{})

	require.NoError(t, err)
	require.Equal(t, []domain.Gene{{Id: 5}, {Id: 6}}, res)
}

func Test_GenesService_GetAll_Err(t *testing.T) {
	genesService, mockGenesRepo := mockGenesService(t)

	mockGenesRepo.EXPECT().GetAll(gomock.Any()).Return([]domain.Gene{}, errInternalServErr)

	res, err := genesService.GetAll([]string{})

	require.True(t, errors.Is(err, errInternalServErr))
	require.Equal(t, []domain.Gene{}, res)
}


func Test_GenesService_GetById(t *testing.T) {
	genesService, mockGenesRepo := mockGenesService(t)

	mockGenesRepo.EXPECT().GetById(gomock.Any()).Return(&domain.Gene{Id: 1}, nil)

	res, err := genesService.GetById(1)

	require.NoError(t, err)
	require.Equal(t, &domain.Gene{Id: 1}, res)
}


func Test_GenesService_GetById_Err(t *testing.T) {
	genesService, mockGenesRepo := mockGenesService(t)

	mockGenesRepo.EXPECT().GetById(gomock.Any()).Return(&domain.Gene{}, errInternalServErr)

	res, err := genesService.GetById(1)

	require.True(t, errors.Is(err, errInternalServErr))
	require.Equal(t, &domain.Gene{}, res)
}

func Test_GenesService_Create(t *testing.T) {
	genesService, mockGenesRepo := mockGenesService(t)
	
	mockGenesRepo.EXPECT().Create(gomock.Any()).Return(gene, nil)

	res, err := genesService.Create(gene)

	require.NoError(t, err)
	require.Equal(t, gene, res)
}


func Test_GenesService_Create_Err(t *testing.T) {
	genesService, mockGenesRepo := mockGenesService(t)

	mockGenesRepo.EXPECT().Create(gomock.Any()).Return(&domain.Gene{}, errInternalServErr)

	res, err := genesService.Create(gene)

	require.True(t, errors.Is(err, errInternalServErr))
	require.Equal(t, &domain.Gene{}, res)
}

func Test_GenesService_Update(t *testing.T) {
	genesService, mockGenesRepo := mockGenesService(t)

	mockGenesRepo.EXPECT().Update(gomock.Any()).Return(gene, nil)

	res, err := genesService.Update(gene)

	require.NoError(t, err)
	require.Equal(t, gene, res)
}


func Test_GenesService_Update_Err(t *testing.T) {
	genesService, mockGenesRepo := mockGenesService(t)

	mockGenesRepo.EXPECT().Update(gomock.Any()).Return(&domain.Gene{}, errInternalServErr)

	res, err := genesService.Update(gene)

	require.True(t, errors.Is(err, errInternalServErr))
	require.Equal(t, &domain.Gene{}, res)
}

func Test_GenesService_Delete(t *testing.T) {
	genesService, mockGenesRepo := mockGenesService(t)

	mockGenesRepo.EXPECT().Delete(gomock.Any())

	err := genesService.Delete(1)

	require.NoError(t, err)
}


func Test_GenesService_Delete_Err(t *testing.T) {
	genesService, mockGenesRepo := mockGenesService(t)

	mockGenesRepo.EXPECT().Delete(gomock.Any()).Return(errInternalServErr)

	err := genesService.Delete(1)

	require.True(t, errors.Is(err, errInternalServErr))
}