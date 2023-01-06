package api

import (
	"kredit/backend/generateCustomer"
	"kredit/backend/generateSkala"

	"github.com/gin-contrib/cors"
)

func (s *server) SetupRouter() {
	s.Router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"POST", "GET", "DELETE", "PUT"},
		AllowHeaders: []string{"*"},
	}))

	generateCustomerRepo := generateCustomer.NewRepository(s.DB)
	generateCustomerService := generateCustomer.NewService(generateCustomerRepo)
	generateCustomerHandler := generateCustomer.NewHandler(generateCustomerService)
	s.Router.GET("/validateStagingCustomer", generateCustomerHandler.ValidateStagingCustomer)

	generateSkalaRepo := generateSkala.NewRepository(s.DB)
	generateSkalaService := generateSkala.NewService(generateSkalaRepo)
	generateSkalaHandler := generateSkala.NewHandler(generateSkalaService)
	s.Router.GET("/generateSkalaRentalTab", generateSkalaHandler.GenerateSkalaRentalTab)
}
