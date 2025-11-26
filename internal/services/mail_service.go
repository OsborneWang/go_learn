package services

import (
	"fmt"

	"gopkg.in/gomail.v2"

	"goServer/internal/config"
)

// MailService 负责统一发送邮件，可扩展模板等能力。
type MailService struct {
	dialer *gomail.Dialer
	from   string
}

// NewMailService 基于配置创建一个 gomail Dialer。
func NewMailService(cfg *config.Config) *MailService {
	dialer := gomail.NewDialer(cfg.MailHost, cfg.MailPort, cfg.MailUsername, cfg.MailPassword)
	return &MailService{
		dialer: dialer,
		from:   cfg.MailFrom,
	}
}

// SendPlainText 发送简单文本邮件，常用于注册确认或重置密码。
func (m *MailService) SendPlainText(to, subject, body string) error {
	msg := gomail.NewMessage()
	msg.SetHeader("From", m.from)
	msg.SetHeader("To", to)
	msg.SetHeader("Subject", subject)
	msg.SetBody("text/plain", body)

	if err := m.dialer.DialAndSend(msg); err != nil {
		return fmt.Errorf("发送邮件失败: %w", err)
	}
	return nil
}
