package checklistPencairan

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service}
}

func (h *Handler) FindPengajuanByApprovalStatus(c *gin.Context) {
	approval_status := c.Query("approval_status")
	if approval_status == "" {
		messageErr := []string{"approval_status harus diisi"}
		c.JSON(http.StatusBadRequest, gin.H{"error": messageErr})
		return
	}
	cdt, status, err := h.Service.FindPengajuanByApprovalStatus(approval_status)
	if err != nil {
		c.JSON(status, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(status, gin.H{
		"message":           "success",
		"customer_data_tab": cdt,
	})
}
