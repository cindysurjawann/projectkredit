package drawdownReport

import (
	"kredit/backend/generateCustomer"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service}
}

func (h *Handler) GetDrawdownReport(c *gin.Context) {
	cdt, status, err := h.Service.GetDrawdownReport()
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

func (h *Handler) GetDrawdownReportByFilter(c *gin.Context) {
	branch := c.Query("branch")
	company := c.Query("company")
	var startDate, endDate time.Time
	var err error

	if startDate, err = generateCustomer.ConvertStringtoDate(c.Query("start_date")); err != nil {
		messageErr := []string{"gagal convert start_date to datetime"}
		c.JSON(http.StatusBadRequest, gin.H{"error": messageErr})
		return
	}

	if endDate, err = generateCustomer.ConvertStringtoDate(c.Query("end_date")); err != nil {
		messageErr := []string{"gagal convert end_date to datetime"}
		c.JSON(http.StatusBadRequest, gin.H{"error": messageErr})
		return
	}

	cdt, status, err := h.Service.GetDrawdownReportByFilter(branch, company, startDate, endDate)
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
