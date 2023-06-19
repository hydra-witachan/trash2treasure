package main

import (
	"fmt"
	"go-boilerplate/internal/controllers"
	"go-boilerplate/internal/middlewares"
	"go-boilerplate/internal/repositories"
	"go-boilerplate/internal/services"
	"go-boilerplate/pkg/databases"
	"log"
	"net/http"
	"os"

	"github.com/amacneil/dbmate/v2/pkg/dbmate"
	"github.com/goava/di"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

func ProvideDIContainer() (container *di.Container, err error) {
	err = godotenv.Load()
	if err != nil {
		return
	}

	container, err = di.New(
		di.Provide(databases.NewMariaDB),
		di.Provide(databases.NewMigration),
		di.Provide(databases.NewFirebaseBucket),
		di.Provide(echo.New),

		// Include controllers, services, and repositories.
		repositories.Module,
		services.Module,
		controllers.Module,

		// Register routes
		di.Invoke(SetupRoutes),
	)
	return
}

func main() {
	di.SetTracer(&di.StdTracer{})

	container, err := ProvideDIContainer()
	if err != nil {
		log.Fatal(err)
	}

	// Force DB to load and test the connection.
	var gorm *gorm.DB
	if err := container.Resolve(&gorm); err != nil {
		return
	}

	container.Invoke(func(db *dbmate.DB, e *echo.Echo) {
		err := db.Migrate()
		if err != nil {
			log.Fatal(err)
		}

		// Enable CORS for all routes
		e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins: []string{"http://localhost:4200"},
			AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete},
			AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization}, // Include echo.HeaderAuthorization in the allowed headers
		}))

		e.HTTPErrorHandler = middlewares.ErrorHandler()
		e.Start(fmt.Sprintf(":%s", os.Getenv("PORT")))
	})
}
