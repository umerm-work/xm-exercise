package router

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/umerm-work/xm-exercise/interface/controller"
	"io/ioutil"
	"log"
	"net/http"
)

// Setup returns the app router.
func Setup() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	companyController := controller.NewCompanyController()
	r.GET("/company", companyController.List)
	r.Use(Middleware).POST("/company", companyController.Add)
	r.PUT("/company", companyController.Update)
	r.Use(Middleware).DELETE("/company/:companyId", companyController.Remove)

	return r
}

type IP struct {
	Ip                 string  `json:"ip"`
	Version            string  `json:"version"`
	City               string  `json:"city"`
	Region             string  `json:"region"`
	RegionCode         string  `json:"region_code"`
	CountryCode        string  `json:"country_code"`
	CountryCodeIso3    string  `json:"country_code_iso3"`
	CountryName        string  `json:"country_name"`
	CountryCapital     string  `json:"country_capital"`
	CountryTld         string  `json:"country_tld"`
	ContinentCode      string  `json:"continent_code"`
	InEu               bool    `json:"in_eu"`
	Postal             string  `json:"postal"`
	Latitude           float64 `json:"latitude"`
	Longitude          float64 `json:"longitude"`
	Timezone           string  `json:"timezone"`
	UtcOffset          string  `json:"utc_offset"`
	CountryCallingCode string  `json:"country_calling_code"`
	Currency           string  `json:"currency"`
	CurrencyName       string  `json:"currency_name"`
	Languages          string  `json:"languages"`
	CountryArea        float64 `json:"country_area"`
	CountryPopulation  int     `json:"country_population"`
	Asn                string  `json:"asn"`
	Org                string  `json:"org"`
	Hostname           string  `json:"hostname"`
}

func Middleware(c *gin.Context) {
	// panic with a string -- the custom middleware could save this to a database or report it to the user
	fmt.Println(c.ClientIP())
	//ip := "102.38.233.0"
	ip := c.ClientIP()
	ipapiClient := http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("https://ipapi.co/%s/json/", ip), nil)
	req.Header.Set("User-Agent", "ipapi.co/#go-v1.5")
	resp, err := ipapiClient.Do(req)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		c.Status(http.StatusInternalServerError)
		c.Abort()
	}
	var ipInfor IP
	err = json.Unmarshal(body, &ipInfor)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		c.Abort()
		return
	}
	//fmt.Println("country", ipInfor.CountryCode)
	if ipInfor.CountryCode != "CY" {
		c.Status(http.StatusUnauthorized)
		c.Abort()
	}
}
