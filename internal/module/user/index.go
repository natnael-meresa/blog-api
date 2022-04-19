package user

import (
	"fmt"
	"twof/blog-api/internal/constant/model"
)

func (s *service) GetUserByQuery(user *model.User) (err error) {
	err = s.userPersist.GetUser(user)
	if err != nil {
		return fmt.Errorf("failed to save new user %s", err.Error())
	}

	return nil
}

func (s *service) GetAllUser(user *[]model.User) (err error) {
	err = s.userPersist.GetAllUsers(user)
	if err != nil {
		return fmt.Errorf("failed to Get user %s", err.Error())
	}

	return nil

}

func (s *service) GetUserById(userId uint, user *model.User) (err error) {
	err = s.userPersist.GetUserById(userId, user)

	if err != nil {
		return fmt.Errorf("failed to Get user %s", err.Error())
	}

	return nil

}

func (s *service) SendResetLink(Email string) error {
	user, err := s.userPersist.GetUserByEmail(Email)
	if err != nil {
		return fmt.Errorf("failed to save new user %s", err.Error())
	}

	if user.Email == "" || err != nil {
		return fmt.Errorf("user not found")
	}

	resetToken, _ := s.GenerateResetToken(user.Email)

	link := "http://localhost:5000/api/v1/password-reset?reset_token=" + resetToken

	body := "this is your reset  <a href='" + link + "'>Link</a>"
	html := "<strong>" + body + "</storng>"

	email := s.SendMail("Reset Password", body, user.Email, html, user.Name)

	if email != true {
		return err
	}

	return nil

}
