package repository

import (
	"fmt"
	"time"
	"twof/blog-api/internal/constant/model"

	"github.com/golang-jwt/jwt"

	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

var validate *validator.Validate

type UserRepository interface {
	Hash(*string) error
	ValidateUser(*model.User) error
	ComparePassword(dbPass, pass string) bool
	GenerateToken(uint) string
	ValidateToken(string) (*jwt.Token, error)
}

type userRepository struct {
	secretKey string
}

func UserInit(secretKey string) UserRepository {
	return &userRepository{
		secretKey,
	}
}

func (u *userRepository) Hash(pass *string) (err error) {
	if *pass != "" {

		bytePass := []byte(*pass)

		hPass, _ := bcrypt.GenerateFromPassword(bytePass, bcrypt.DefaultCost)
		*pass = string(hPass)

	}

	return
}

func (u *userRepository) ComparePassword(dbPass, pass string) bool {
	return bcrypt.CompareHashAndPassword([]byte(dbPass), []byte(pass)) == nil
}

func (u *userRepository) GenerateToken(userid uint) string {
	claims := jwt.MapClaims{
		"exp":    time.Now().Add(time.Hour * 3).Unix(),
		"iat":    time.Now().Unix(),
		"userID": userid,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, _ := token.SignedString([]byte("secret"))

	return t
}

func (u *userRepository) ValidateToken(token string) (*jwt.Token, error) {

	return jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte("secret"), nil

	})
}

func (u *userRepository) ValidateUser(user *model.User) (err error) {
	validate = validator.New()

	err = validate.Struct(user)

	if err != nil {
		fmt.Println(err)
		var Msg string
		for _, err := range err.(validator.ValidationErrors) {
			Msg += err.Field() + " is " + err.Tag() + "\n"
		}
		return fmt.Errorf(Msg)
	}

	return nil
}
