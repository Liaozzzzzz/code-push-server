package service

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/liaozzzzzz/code-push-server/internal/config"
	"github.com/liaozzzzzz/code-push-server/internal/dao"
	"github.com/liaozzzzzz/code-push-server/internal/dto"
	"github.com/liaozzzzzz/code-push-server/internal/types"
	"github.com/liaozzzzzz/code-push-server/internal/utils/crypto"
	utilsErrors "github.com/liaozzzzzz/code-push-server/internal/utils/errors"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type LoginService struct {
	userDAO *dao.UserDAO
}

func NewLoginService() *LoginService {
	return &LoginService{
		userDAO: dao.NewUserDAO(),
	}
}

// Login 用户登录
func (s *LoginService) Login(req *dto.LoginForm) (*dto.LoginResult, error) {
	// 根据用户名查询用户
	user, err := s.userDAO.GetByUsername(req.Username)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, utilsErrors.NewBusinessError(utilsErrors.CodeResourceNotFound, "用户名或密码错误")
		}
		return nil, err
	}

	decryptedPassword, err := crypto.Decrypt(req.Password)
	if err != nil {
		return nil, err
	}

	// 验证密码
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(decryptedPassword)); err != nil {
		return nil, utilsErrors.NewBusinessError(utilsErrors.CodeInvalidParams, "密码错误")
	}

	// 检查用户状态
	if user.UserStatus != types.UserEnabled {
		return nil, utilsErrors.NewBusinessError(utilsErrors.CodePermissionDenied, "用户已停用")
	}

	// 生成token
	claims := jwt.MapClaims{
		"username": user.Username,
		"userId":   user.UserID,
		"ack":      user.AckCode,
		"exp":      time.Now().Add(time.Duration(config.C.Security.JWTExpiration) * time.Second).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(config.C.Security.JWTSecret))
	if err != nil {
		return nil, err
	}

	return &dto.LoginResult{
		Token: tokenString,
		User:  user,
	}, nil
}
