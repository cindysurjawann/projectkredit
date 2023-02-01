package api

import (
	"kredit/backend/generateCustomer"
	"kredit/backend/generateSkala"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/robfig/cron/v3"
	"gorm.io/gorm"
)

type server struct {
	Router *gin.Engine
	DB     *gorm.DB
}

func MakeServer(db *gorm.DB) *server {
	s := &server{
		Router: gin.Default(),
		DB:     db,
	}

	c := cron.New()
	c.AddFunc("@every 30m", func() {
		generateCustomer := generateCustomer.NewRepository(s.DB)
		generateCustomer.ValidateStagingCustomer()
	})

	c.AddFunc("@every 15m", func() {
		generateSkala := generateSkala.NewRepository(s.DB)
		generateSkala.GenerateSkalaRentalTab()
	})
	c.Start()

	return s
}

func (s *server) RunServer() {
	s.SetupRouter()
	port := os.Getenv("PORT")
	if err := s.Router.Run(":" + port); err != nil {
		panic(err)
	}
}
