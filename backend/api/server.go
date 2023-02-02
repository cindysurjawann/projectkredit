package api

import (
	"fmt"
	"kredit/backend/generateCustomer"
	"kredit/backend/generateSkala"
	"os"
	"time"

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
	//30 min
	c.AddFunc("@every 30m", func() {
		generateCustomer := generateCustomer.NewRepository(s.DB)
		generateCustomer.ValidateStagingCustomer()
		t := time.Now()
		fmt.Println("jalanin fungsi validasi staging_customer pada ", t)
	})

	//15min
	c.AddFunc("@every 15m", func() {
		generateSkala := generateSkala.NewRepository(s.DB)
		generateSkala.GenerateSkalaRentalTab()
		t := time.Now()
		fmt.Println("jalanin fungsi generate skala_rental_tab pada ", t)
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
