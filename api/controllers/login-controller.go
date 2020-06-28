package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/leogoesger/news-api/api/auth"
	"github.com/leogoesger/news-api/api/responses"
	"github.com/leogoesger/news-api/api/utils/formaterror"
	"github.com/leogoesger/news-api/db/models"
	"golang.org/x/crypto/bcrypt"
)

// Login logic
func (server *Server) Login(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	user := models.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	user.Prepare()
	err = user.Validate("login")
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	token, err := server.SignIn(user.Email, user.Password)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusUnprocessableEntity, formattedError)
		return
	}
	cookie := http.Cookie{Name: "news-token", Value: token, Path: "/", MaxAge: 0}
	http.SetCookie(w, &cookie)
	responses.JSON(w, http.StatusOK, token)
}

// SignIn sign in user
func (server *Server) SignIn(email, password string) (string, error) {

	var err error

	user := models.User{}

	err = server.DB.Debug().Model(models.User{}).Where("email = ?", email).Take(&user).Error
	if err != nil {
		return "", err
	}
	err = models.VerifyPassword(user.Password, password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}
	return auth.CreateToken(user.ID)
}