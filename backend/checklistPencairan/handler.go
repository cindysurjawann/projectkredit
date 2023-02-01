package checklistPencairan

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

func (h *Handler) FindPengajuanByFilter(c *gin.Context) {
	approval_status := c.Query("approval_status")
	branch := c.Query("branch")
	company := c.Query("company")
	var startDate, endDate time.Time
	var err error

	if approval_status == "" {
		messageErr := []string{"approval_status harus diisi"}
		c.JSON(http.StatusBadRequest, gin.H{"error": messageErr})
		return
	}

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

	cdt, status, err := h.Service.FindPengajuanByFilter(approval_status, branch, company, startDate, endDate)
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

func (h *Handler) GetBranchList(c *gin.Context) {
	bt, status, err := h.Service.GetBranchList()
	if err != nil {
		c.JSON(status, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(status, gin.H{
		"message":    "success",
		"branch_tab": bt,
	})
}

func (h *Handler) GetCompanyList(c *gin.Context) {
	mct, status, err := h.Service.GetCompanyList()
	if err != nil {
		c.JSON(status, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(status, gin.H{
		"message":         "success",
		"mst_company_tab": mct,
	})
}

func (h *Handler) UpdateApprovalStatus(c *gin.Context) {
	var Input UpdateApprovalStatusRequest

	if err := c.ShouldBindJSON(&Input); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	cdt, status, err := h.Service.UpdateApprovalStatus(Input.CustomerDataTab, Input.ApprovalStatus)
	if err != nil {
		c.JSON(status, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(status, gin.H{
		"message":         "success",
		"mst_company_tab": cdt,
	})
}
