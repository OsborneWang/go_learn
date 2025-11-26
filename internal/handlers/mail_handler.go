package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"goServer/internal/services"
)

// MailHandler 用于调试邮件发送能力。
type MailHandler struct {
	mailService *services.MailService
}

func NewMailHandler(mailService *services.MailService) *MailHandler {
	return &MailHandler{mailService: mailService}
}

type sendTestMailRequest struct {
	To      string `json:"to" binding:"required,email"`
	Subject string `json:"subject" binding:"required"`
	Message string `json:"message" binding:"required"`
}

// SendTest 接收目标邮箱并发送测试邮件。
func (h *MailHandler) SendTest(c *gin.Context) {
	var req sendTestMailRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.mailService.SendPlainText(req.To, req.Subject, req.Message); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "测试邮件已发送"})
}
