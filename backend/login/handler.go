package login

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Handler struct {
	Service Service
}

const (
	message = "Input data not suitable"
)

func NewHandler(service Service) *Handler {
	return &Handler{service}
}

func (h *Handler) FindUser(c *gin.Context) {
	user_id := c.Query("userId")
	if user_id == "" {
		messageErr := []string{"userId harus diisi"}
		c.JSON(http.StatusBadRequest, gin.H{"error": messageErr})
		return
	}
	user, status, err := h.Service.FindUser(user_id)
	if err != nil {
		c.JSON(status, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(status, gin.H{
		"message": "success",
		"user_id": user.UserId,
		"name":    user.Name,
		"email":   user.Email,
	})
}

func (h *Handler) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		messageErr := ParseError(err)
		if messageErr == nil {
			messageErr = []string{message}
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": messageErr})
		return
	}
	res, status, err := h.Service.Login(req)
	if err != nil {
		c.JSON(status, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(status, gin.H{
		"data": map[string]any{
			"name":    res.Name,
			"user_id": res.UserId,
		},
	})
}

func (h *Handler) Register(c *gin.Context) {
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		messageErr := ParseError(err)
		if messageErr == nil {
			messageErr = []string{message}
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": messageErr})
		return
	}
	_, status, err := h.Service.Register(req)
	if err != nil {
		c.JSON(status, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(status, gin.H{
		"message": "success",
	})
}

func ParseError(err error) []string {
	if validationErrs, ok := err.(validator.ValidationErrors); ok {
		fmt.Println("err valid", validationErrs)
		errorMessages := make([]string, len(validationErrs))
		for i, e := range validationErrs {
			errorMessages[i] = fmt.Sprintf("The field %s is %s", e.Field(), e.Tag())
		}
		return errorMessages
	} else if marshallingErr, ok := err.(*json.UnmarshalTypeError); ok {
		return []string{fmt.Sprintf("The field %s must be a %s", marshallingErr.Field, marshallingErr.Type.String())}
	}
	return nil
}
