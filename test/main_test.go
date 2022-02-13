package test

import (
	unitTest "github.com/Valiben/gin_unit_test"
	"github.com/gin-gonic/gin"
	"github.com/umerm-work/xm-exercise/interface/controller"
	"log"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	// initialize the router
	router := gin.Default()
	// Handlers for testing
	companyController := controller.NewCompanyController()
	router.GET("/company", companyController.List)
	router.POST("/company", companyController.Add)
	router.PUT("/company", companyController.Update)
	router.DELETE("/company/:companyId", companyController.Remove)

	// Setup the router
	unitTest.SetRouter(router)
	newLog := log.New(os.Stdout, "", log.Llongfile|log.Ldate|log.Ltime)
	unitTest.SetLog(newLog)
}
