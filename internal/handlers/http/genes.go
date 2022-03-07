package http

import (
	"fmt"
	"strconv"
	"time"

	"github.com/breeders-zone/morphs-service/internal/domain"
	"github.com/breeders-zone/morphs-service/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/manyminds/api2go"
	"github.com/manyminds/api2go/jsonapi"
)

type CreateGeneRequest struct {
	Title        string    `json:"title" validate:"required"`
	Type         string    `json:"type" validate:"required"`
	ProducedName string    `json:"producedName" validate:"required"`
	ProducedDate time.Time `json:"producedDate" validate:"required"`
	Availability string    `json:"availability" validate:"required"`
	Description  string    `json:"description" validate:"required"`
	History      string    `json:"history" validate:"required"`
	Links        []string  `json:"links" validate:"required"`
}

func (r *CreateGeneRequest) GetName() string {
	return "genes"
}

func (r *CreateGeneRequest) SetID(sId string) error {
	return nil
}

type UpdateGeneRequest struct {
	Id           int       `json:"-"`
	Title        string    `json:"title" validate:"required"`
	Type         string    `json:"type" validate:"required"`
	ProducedName string    `json:"producedName" validate:"required"`
	ProducedDate time.Time `json:"producedDate" validate:"required"`
	Availability string    `json:"availability" validate:"required"`
	Description  string    `json:"description" validate:"required"`
	History      string    `json:"history" validate:"required"`
	Links        []string  `json:"links" validate:"required"`
}

func (r *UpdateGeneRequest) GetName() string {
	return "genes"
}

func (r *UpdateGeneRequest) SetID(sId string) error {
	id, err := strconv.Atoi(sId)
	if err != nil {
		return err
	}

	r.Id = id

	return nil
}

// GetGene
// @Summary      Get gene by ID
// @Description  Get gene by ID
// @Tags         genes
// @Accept       json-api
// @Produce      json-api
// @Param        id   path      int  true  "Gene ID"
// @Success      200  {object} responses.GeneResponse
// @Failure		 404,500 {object} api2go.Error
// @Router       /genes/{id} [get]
func (h *Handler) GetGene(c *fiber.Ctx) error {
	c.Set(fiber.HeaderContentType, "application/vnd.api+json")

	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return err
	}

	g, err := h.services.Genes.GetById(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(&api2go.Error{
			Status: fmt.Sprint(fiber.StatusNotFound),
			Title:  "Not found",
		})
	}

	j, err := jsonapi.Marshal(g)
	if err != nil {
		return err
	}

	return c.Send(j)
}

// CreateGene
// @Summary      Create Gene
// @Description  Create Gene
// @Tags         genes
// @Accept       json-api
// @Produce      json-api
// @Param        input body responses.GeneResponse true "Gene request"
// @Success      200  {object} responses.GeneResponse
// @Failure		 422 {object} []api2go.Error
// @Failure		 500 {object} api2go.Error
// @Router       /genes [post]
func (h *Handler) CreateGene(c *fiber.Ctx) error {
	input := new(CreateGeneRequest)

	if err := jsonapi.Unmarshal(c.Body(), input); err != nil {
		return err
	}

	if valErrors := utils.Validate(input); valErrors != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(api2go.HTTPError{Errors: valErrors})
	}

	g, err := h.services.Genes.Create(&domain.Gene{
		Title:        input.Title,
		Type:         input.Type,
		ProducedName: input.ProducedName,
		ProducedDate: input.ProducedDate,
		Availability: input.Availability,
		Description:  input.Description,
		History:      input.History,
		Links:        input.Links,
	})

	if err != nil {
		return err
	}

	j, err := jsonapi.Marshal(g)
	if err != nil {
		return err
	}

	c.Set(fiber.HeaderContentType, "application/vnd.api+json")
	return c.Send(j)
}

// UpdateGene
// @Summary      Update Gene
// @Description  Update Gene
// @Tags         genes
// @Accept       json-api
// @Produce      json-api
// @Param        input body responses.GeneResponse true "Gene request"
// @Param        id   path      int  true  "Gene ID"
// @Success      200  {object} responses.GeneResponse
// @Failure		 422 {object} []api2go.Error
// @Failure		 404,500 {object} api2go.Error
// @Router       /genes/{id} [put]
func (h *Handler) UpdateGene(c *fiber.Ctx) error {
	c.Set(fiber.HeaderContentType, "application/vnd.api+json")

	input := new(UpdateGeneRequest)

	if err := jsonapi.Unmarshal(c.Body(), input); err != nil {
		return err
	}

	if valErrors := utils.Validate(input); valErrors != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(api2go.HTTPError{Errors: valErrors})
	}

	g, err := h.services.Genes.Update(&domain.Gene{
		Id:           input.Id,
		Title:        input.Title,
		Type:         input.Type,
		ProducedName: input.ProducedName,
		ProducedDate: input.ProducedDate,
		Availability: input.Availability,
		Description:  input.Description,
		History:      input.History,
		Links:        input.Links,
	})
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(&api2go.Error{
			Status: fmt.Sprint(fiber.StatusNotFound),
			Title:  "Not found",
		})
	}

	j, err := jsonapi.Marshal(g)
	if err != nil {
		return err
	}

	return c.Send(j)
}

// DeleteGene
// @Summary      Delete Gene
// @Description  Delete Gene
// @Tags         genes
// @Accept       json-api
// @Produce      json-api
// @Param        id   path      int  true  "Gene ID"
// @Success      204
// @Failure		 404,500 {object} api2go.Error
// @Router       /genes/{id} [delete]
func (h *Handler) DeleteGene(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return err
	}

	err = h.services.Genes.Delete(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(&api2go.Error{
			Status: fmt.Sprint(fiber.StatusNotFound),
			Title:  "Not found",
		})
	}

	return c.Status(fiber.StatusNoContent).SendString("")
}
