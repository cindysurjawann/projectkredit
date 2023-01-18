package api

import (
	"kredit/backend/checklistPencairan"
	"kredit/backend/generateCustomer"
	"kredit/backend/generateSkala"
	"kredit/backend/login"

	"github.com/gin-contrib/cors"
)

func (s *server) SetupRouter() {
	s.Router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"POST", "GET", "DELETE", "PUT"},
		AllowHeaders: []string{"*"},
	}))
	loginRepo := login.NewRepository(s.DB)
	loginService := login.NewService(loginRepo)
	loginHandler := login.NewHandler(loginService)
	s.Router.GET("/findUser", loginHandler.FindUser)
	s.Router.POST("/register", loginHandler.Register)
	s.Router.POST("/login", loginHandler.Login)

	checklistPencairanRepo := checklistPencairan.NewRepository(s.DB)
	checklistPencairanService := checklistPencairan.NewService(checklistPencairanRepo)
	checklistPencairanHandler := checklistPencairan.NewHandler(checklistPencairanService)
	s.Router.GET("/getChecklistPengajuan", checklistPencairanHandler.FindPengajuanByApprovalStatus)

	generateCustomerRepo := generateCustomer.NewRepository(s.DB)
	generateCustomerService := generateCustomer.NewService(generateCustomerRepo)
	generateCustomerHandler := generateCustomer.NewHandler(generateCustomerService)
	s.Router.GET("/validateStagingCustomer", generateCustomerHandler.ValidateStagingCustomer)

	generateSkalaRepo := generateSkala.NewRepository(s.DB)
	generateSkalaService := generateSkala.NewService(generateSkalaRepo)
	generateSkalaHandler := generateSkala.NewHandler(generateSkalaService)
	s.Router.GET("/generateSkalaRentalTab", generateSkalaHandler.GenerateSkalaRentalTab)
}
