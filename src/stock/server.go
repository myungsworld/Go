package main

import (
	"net/http"
	"stock/controllers"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()

	//CORS 웹서버에 도메인 간 액세스 제어를 제공해 안전한 도메인 간 데이터 전송을 가능하게 함
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello")
	})

	e.POST("/price", controllers.GrabPrice)

	e.Logger.Fatal(e.Start(":8000"))
}
