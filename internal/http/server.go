package api

import (
	"net/http"
	"simpleOpenapi/internal/database/repository"
	"simpleOpenapi/internal/http/gen"
	mm "simpleOpenapi/pkg/middleware"

	om "github.com/deepmap/oapi-codegen/pkg/middleware"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

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
	// validatorの設定
	e.Validator = mm.NewValidator()

	// validator
	spec, err := gen.GetSwagger()
	if err != nil {
		panic(err)
	}
	e.Use(om.OapiRequestValidator(spec))

	//mysql connection
	//TODO 設定ファイルの利用と、database共通処理を作る
	dsn := "docker:docker@tcp(127.0.0.1:3306)/simple_openapi?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	//TODO 外からauto-migration対象を指定できる仕組みを作る
	if err := db.AutoMigrate(&repository.AmazonItems{}); err != nil {
		panic(err.Error())
	}

	gen.RegisterHandlers(e, NewApi(db))
	e.GET("/health", func(context echo.Context) error {
		return context.String(http.StatusOK, "ok")
	})
	e.Logger.Fatal(e.Start(":1232"))
}
