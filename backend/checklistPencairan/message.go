package checklistPencairan

import "kredit/backend/model"

type UpdateApprovalStatusRequest struct {
	CustomerDataTab []model.CustomerDataTab `json:"customer_data_tab" binding:"required"`
	ApprovalStatus  string                  `json:"approval_status" binding:"required"`
}
