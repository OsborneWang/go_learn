package services

import (
	"errors"
	"fmt"
	"strings"

	"gorm.io/gorm"

	"goServer/internal/models"
)

// AuthService 处理注册与登录逻辑。
type AuthService struct {
	db          *gorm.DB
	mailService *MailService
}

func NewAuthService(db *gorm.DB, mailService *MailService) *AuthService {
	return &AuthService{db: db, mailService: mailService}
}

// Register 创建新用户，并发送欢迎邮件。
func (s *AuthService) Register(email, password, name string) (*models.User, error) {
	email = strings.TrimSpace(email)
	if email == "" || password == "" {
		return nil, errors.New("邮箱和密码不能为空")
	}

	user := &models.User{
		Email: email,
		Name:  name,
	}

	if err := user.SetPassword(password); err != nil {
		return nil, fmt.Errorf("设置密码失败: %w", err)
	}

	if err := s.db.Create(user).Error; err != nil {
		return nil, fmt.Errorf("创建用户失败: %w", err)
	}

	go func() {
		if s.mailService != nil {
			if err := s.mailService.SendPlainText(
				user.Email,
				"欢迎注册",
				fmt.Sprintf("您好 %s，欢迎加入我们的服务！", user.Name),
			); err != nil {
				// 仅日志提示，不阻断主流程
				fmt.Printf("发送欢迎邮件失败: %v\n", err)
			}
		}
	}()

	return user, nil
}

// Login 校验账号密码，返回用户实体。
func (s *AuthService) Login(email, password string) (*models.User, error) {
	var user models.User
	if err := s.db.Where("email = ?", strings.ToLower(strings.TrimSpace(email))).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("邮箱或密码错误")
		}
		return nil, fmt.Errorf("查询用户失败: %w", err)
	}

	if !user.CheckPassword(password) {
		return nil, errors.New("邮箱或密码错误")
	}

	return &user, nil
}
