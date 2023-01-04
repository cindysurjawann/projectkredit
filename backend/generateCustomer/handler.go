package generateCustomer

import "github.com/gin-gonic/gin"

type Handler struct {
	Service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service}
}

func (h *Handler) GetStagingCustomer(c *gin.Context) {
	customer, status, err := h.Service.GetStagingCustomer()
	if err != nil {
		c.JSON(status, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(status, gin.H{
		"message": "success",
		"data":    customer,
	})
}
