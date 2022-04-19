package user

import (
	"context"
	"twof/blog-api/internal/constant/model"
	"twof/blog-api/internal/repository"
	"twof/blog-api/internal/storage/persistence"

	"twof/blog-api/internal/glue/enforcer"

	"github.com/golang-jwt/jwt"
)

type Usecase interface {
	Registration(context.Context, *model.User) error
	GetUserByQuery(*model.User) (err error)
	Hash(*string) error
	ValidateUser(*model.User) error
	GenerateToken(model.User) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
	ComparePassword(dbPass, pass string) bool
	GenerateTokenPair(model.User) (model.Token, error)
	CreateRefreshToken(user model.User) (string, error)
	ValidateRefreshToken(string) (uint, error)
	GetUserById(uint, *model.User) (err error)
	SendResetLink(Email string) error
	ResetPassword(resetPassword model.PasswordReset, resetToken string) error
}

type service struct {
	enforcer    enforcer.CasbinMiddleware
	userRepo    repository.UserRepository
	userPersist persistence.UserPersistence
}

func Initialize(
	enforcer enforcer.CasbinMiddleware,
	userRepo repository.UserRepository,
	userPersist persistence.UserPersistence,
) Usecase {
	return &service{
		enforcer,
		userRepo,
		userPersist,
	}
}
