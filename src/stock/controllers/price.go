package controllers

import (
	"bytes"
	"io/ioutil"
	"net/http"

	"github.com/antchfx/htmlquery"
	"github.com/labstack/echo"
)

type company struct {
	ticker string `json:"ticker" form:"ticker" query:"ticker`
}

func GrabPrice(c echo.Context) (err error) {

	var bodyBytes []byte
	if c.Request().Body != nil {
		bodyBytes, _ = ioutil.ReadAll(c.Request().Body)
	}
	c.Request().Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

	u := new(company)
	er := c.Bind(u)

	if er != nil {
		panic(er)
	}

	ticker := u.ticker
	baseURL := "https:///finance.yahoo.com/quote/"

	pricePath := "//*[@id=\"quote-header-info\"]"

	doc, err := htmlquery.LoadURL(baseURL + ticker)

	if err != nil {
		panic(err)
	}
	context := htmlquery.FindOne(doc, pricePath)

	price := string(htmlquery.InnerText(context))
	// fmt.Println(price)
	// c.Response().Header().Add("Content-Type", "application/json")
	return c.JSON(http.StatusOK, price)
}
