package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/umerm-work/xm-exercise/interface/gateway/database"
	"github.com/umerm-work/xm-exercise/interface/presenter"
	inputdata "github.com/umerm-work/xm-exercise/usercase/input/data"
	inputport "github.com/umerm-work/xm-exercise/usercase/input/port"
	"github.com/umerm-work/xm-exercise/usercase/interactor"
	"net/http"
	"strconv"
)

const (
	ID      = "id"
	Name    = "name"
	Country = "country"
	Phone   = "phone"
	Website = "website"
)

// CompanyController ...
type CompanyController struct {
	inputport inputport.CompanyInputPort
}

// NewCompanyController passes database to usecase
func NewCompanyController() *CompanyController {
	return &CompanyController{
		inputport: interactor.NewCompanyInteractor(
			presenter.NewCompanyPresenter(),
			database.NewCompanyDatabase(),
		),
	}
}

func (r *CompanyController) Add(c *gin.Context) {
	var in inputdata.Company
	if err := c.BindJSON(&in); err != nil {
		return
	}
	res, err := r.inputport.Create(in)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusCreated, res)
}

func (r *CompanyController) List(c *gin.Context) {
	filter, value := "", ""
	for key, values := range c.Request.URL.Query() {
		filter = key
		value = values[0]
	}
	list, err := r.inputport.List(filter, value)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, list)
}

func (r *CompanyController) Remove(c *gin.Context) {
	companyId := c.Param("companyId")
	cid, err := strconv.Atoi(companyId)
	if err != nil {
		cid = 0
	}
	err = r.inputport.Remove(cid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, "")
}

func (r *CompanyController) Update(c *gin.Context) {
	var in inputdata.Company
	if err := c.BindJSON(&in); err != nil {
		return
	}
	err := r.inputport.Update(in)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, "")
}
