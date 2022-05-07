package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/lzszq/AssessmentRestfulApi/api/auth"
	"github.com/lzszq/AssessmentRestfulApi/api/models"
	"github.com/lzszq/AssessmentRestfulApi/api/responses"
	"github.com/lzszq/AssessmentRestfulApi/api/utils/formaterror"
	"golang.org/x/crypto/bcrypt"
)

type Result struct {
	Token string `json:"token"`
}

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
	result := Result{
		Token: token,
	}
	responses.JSON(w, http.StatusOK, result)
}

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
