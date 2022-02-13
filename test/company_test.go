package test

import (
	unitTest "github.com/Valiben/gin_unit_test"
	"github.com/Valiben/gin_unit_test/utils"
	inputdata "github.com/umerm-work/xm-exercise/usercase/input/data"
	ouputdata "github.com/umerm-work/xm-exercise/usercase/output/data"
	"testing"
)

func Test_CreateCompany(t *testing.T) {
	var output ouputdata.Company
	var input = inputdata.Company{
		Name:    "abc",
		Code:    "abc",
		Country: "abc",
		Website: "abc",
		Phone:   "abc",
	}

	err := unitTest.TestHandlerUnMarshalResp(utils.POST, "/company", "json", input, &output)
	if err != nil {
		t.Errorf("TestPostMappingClientNotFound: %v\n", err)
		return
	}

	checkResponseValues(t, input.Name, output.Name)
	checkResponseValues(t, input.Code, output.Code)
	checkResponseValues(t, input.Website, output.Website)
	checkResponseValues(t, input.Country, output.Country)
	checkResponseValues(t, input.Phone, output.Phone)

}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}
func checkResponseValues(t *testing.T, expected, actual string) {
	if expected != actual {
		t.Errorf("Expected response code %s. Got %s\n", expected, actual)
	}
}

func Test_UpdateCompany(t *testing.T) {
	var output ouputdata.Company
	var input = inputdata.Company{
		Name:    "abc",
		Code:    "abc",
		Country: "abc",
		Website: "abc",
		Phone:   "abc",
	}

	err := unitTest.TestHandlerUnMarshalResp(utils.PUT, "/company", "json", input, &output)
	if err != nil {
		t.Errorf("TestPostMappingClientNotFound: %v\n", err)
		return
	}

	checkResponseValues(t, input.Name, output.Name)
	checkResponseValues(t, input.Code, output.Code)
	checkResponseValues(t, input.Website, output.Website)
	checkResponseValues(t, input.Country, output.Country)
	checkResponseValues(t, input.Phone, output.Phone)

}
