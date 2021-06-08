package api

import (
	"net/http"
	"simpleOpenapi/internal/http/gen"

	"github.com/labstack/echo/v4/middleware"

	"github.com/labstack/echo/v4"
)

func Run() {
	e := echo.New()
	// リクエストIDの設定
	e.Use(middleware.RequestID())
	// loggerの設定
	e.Use(middleware.Logger())
	// recoverの設定
	e.Use(middleware.Recover())

	//TODO what's?  "message": "no matching operation was found"
	//// validator
	//spec, err := gen.GetSwagger()
	//if err != nil {
	//	panic(err)
	//}
	//e.Use(middleware2.OapiRequestValidator(spec))

	gen.RegisterHandlers(e, NewApi())
	e.GET("/health", func(context echo.Context) error {
		return context.String(http.StatusOK, "ok")
	})
	e.Logger.Fatal(e.Start(":1232"))
}
