package handler

import (
	"CompanyAPI/ent"
	"CompanyAPI/ent/company"
	"CompanyAPI/request"
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type companyHandler struct {
	db *ent.Client
}

func NewCompanyHandler(db *ent.Client) *companyHandler {
	return &companyHandler{db}
}

func (compHand *companyHandler) Create(c echo.Context) error {
	comp := new(request.CompanyRequest)
	if err := c.Bind(comp); err != nil {
		return err
	}
	if err := c.Validate(comp); err != nil {
		return err
	}
	ctx := context.Background()
	newComp := compHand.db.Company.Create().SetName(comp.Name).SetAddress(comp.Address).SetService(comp.Service).SaveX(ctx)
	return c.JSON(http.StatusOK, newComp)
}

func (compHand *companyHandler) Read(c echo.Context) error {
	ctx := context.Background()
	comps := compHand.db.Company.Query().AllX(ctx)
	return c.JSON(http.StatusOK, comps)
}

func (compHand *companyHandler) ReadByID(c echo.Context) error {
	ctx := context.Background()
	id, _ := strconv.Atoi(c.Param("id"))
	comp := compHand.db.Company.Query().Where(company.ID(id)).AllX(ctx)

	return c.JSON(http.StatusOK, comp)
}

func (compHand *companyHandler) Update(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	req := new(request.CompanyRequest)
	ctx := context.Background()
	c.Bind(req)

	name := req.Name
	if name == "" {
		name = compHand.db.Company.Query().Where(company.ID(id)).Select(company.FieldName).StringX(ctx)
	}

	address := req.Address
	if address == "" {
		address = compHand.db.Company.Query().Where(company.ID(id)).Select(company.FieldAddress).StringX(ctx)
	}

	service := req.Service
	if service == "" {
		service = compHand.db.Company.Query().Where(company.ID(id)).Select(company.FieldService).StringX(ctx)
	}

	comp := compHand.db.Company.UpdateOneID(id).SetAddress(address).SetName(name).SetService(service).SaveX(ctx)

	return c.JSON(http.StatusOK, comp)
}

func (compHand *companyHandler) Delete(c echo.Context) error {
	ctx := context.Background()
	id, _ := strconv.Atoi(c.Param("id"))

	comp := compHand.db.Company.Query().Where(company.ID(id)).AllX(ctx)
	del := compHand.db.Company.DeleteOneID(id).Exec(ctx)

	if del != nil {
		log.Fatal(del)
	}

	return c.JSON(http.StatusOK, comp)
}
