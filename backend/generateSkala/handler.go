package generateSkala

import "github.com/gin-gonic/gin"

type Handler struct {
	Service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service}
}

func (h *Handler) GenerateSkalaRentalTab(c *gin.Context) {
	skala, status, err := h.Service.GenerateSkalaRentalTab()
	if err != nil {
		c.JSON(status, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(status, gin.H{
		"message": "success",
		"data":    skala,
	})
}
