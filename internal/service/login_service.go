package service

import (
	"errors"
	"fmt"

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
		fmt.Println(err)
		return nil, utilsErrors.NewBusinessError(utilsErrors.CodeInvalidParams, "密码错误")
	}

	// 检查用户状态
	fmt.Println(user.UserStatus)
	if user.UserStatus != types.UserEnabled {
		return nil, utilsErrors.NewBusinessError(utilsErrors.CodePermissionDenied, "用户已被禁用")
	}

	// 生成token（这里简化处理，实际应该使用JWT等）
	token := "mock_token_" + user.Username

	return &dto.LoginResult{
		Token: token,
		User:  *user,
	}, nil
}
