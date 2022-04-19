package user

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"
	"twof/blog-api/internal/constant/model"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

func (s *service) Registration(ctx context.Context, user *model.User) error {

	if err := s.ValidateUser(user); err != nil {
		return err
	}

	s.Hash(&user.Password)
	err := s.userPersist.CreateUser(user)
	if err != nil {
		return fmt.Errorf("failed to save new user %s", err.Error())
	}

	df, err := s.enforcer.GetEnforcer().AddRoleForUser(fmt.Sprint(user.ID), user.Role)

	fmt.Println(df)
	fmt.Println(err)

	return nil
}

func (s *service) Hash(pass *string) error {
	err := s.userRepo.Hash(pass)
	if err != nil {
		return fmt.Errorf("failed to Hash Password %s", err.Error())
	}

	return nil
}

func (s *service) ComparePassword(dbPass, pass string) bool {
	return bcrypt.CompareHashAndPassword([]byte(dbPass), []byte(pass)) == nil
}

func (s *service) GenerateToken(user model.User) (string, error) {
	claims := jwt.MapClaims{
		"exp":    time.Now().Add(time.Hour * 3).Unix(),
		"iat":    time.Now().Unix(),
		"userID": user.ID,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte("secret"))

	if err != nil {
		return t, err
	}

	return t, nil
}

func (s *service) ValidateToken(token string) (*jwt.Token, error) {

	return jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte("secret"), nil

	})
}

func (s *service) ValidateRefreshToken(refreshToken string) (uint, error) {
	var userId uint
	token, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("SECRET_KEY")), nil
	})

	if err != nil {
		return userId, errors.New("Refresh token expired")
	}

	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return userId, errors.New("Unautorized")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !(ok && token.Valid) {
		return userId, errors.New("refresh expired")
	}

	userId = claims["userID"].(uint)

	return userId, nil

}

func (s *service) CreateRefreshToken(user model.User) (string, error) {

	rtClaims := jwt.MapClaims{
		"exp":    time.Now().Add(time.Hour * 24).Unix(),
		"iat":    time.Now().Unix(),
		"userID": user.ID,
	}
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)

	rt, err := refreshToken.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}

	return rt, nil
}

func (s *service) GenerateTokenPair(user model.User) (model.Token, error) {
	var err error
	jwt := model.Token{}
	jwt.RefreshToken, err = s.GenerateToken(user)

	if err != nil {
		return jwt, err
	}
	jwt.AccessToken, err = s.CreateRefreshToken(user)
	if err != nil {
		return jwt, err
	}

	return jwt, err

}

func (s *service) GenerateResetToken(email string) (string, error) {

	rstClaims := jwt.MapClaims{
		"exp":       time.Now().Add(time.Minute * 140).Unix(),
		"iat":       time.Now().Unix(),
		"userEmail": email,
	}
	resetToken := jwt.NewWithClaims(jwt.SigningMethodHS256, rstClaims)

	rst, err := resetToken.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}

	return rst, nil
}

func (s *service) DecodeResetToken(resetToken string) (string, error) {

	var userEmail string
	token, err := jwt.Parse(resetToken, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte("secret"), nil
	})

	if err != nil {
		return userEmail, errors.New("Reset token expired")
	}

	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return userEmail, errors.New("In valid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !(ok && token.Valid) {
		return userEmail, errors.New("Reset token  expired")
	}

	userEmail = claims["userEmail"].(string)

	return userEmail, nil

}

func (s *service) ResetPassword(resetPassword model.PasswordReset, resetToken string) error {

	// Decode the token
	userEmail, _ := s.DecodeResetToken(resetToken)

	// Fetch the user
	user, err := s.userPersist.GetUserByEmail(userEmail)

	if err != nil {
		return err
	}
	if user.Email == "" {
		return fmt.Errorf("user do not exist")
	}

	s.Hash(&resetPassword.Password)

	err = s.userPersist.UpdateUserPass(userEmail, resetPassword.Password)

	if err != nil {
		return err
	}

	return nil
}
