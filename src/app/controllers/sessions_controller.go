package controllers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"url_manager/app/repositories"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func NewSession(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{"title": "login"})
}

type LoginForm struct {
	LoginId  string `form:"login_id"`
	Password string `form:"password"`
}

// type SessionURI struct {
// 	Id int `uri:"id"`
// }

func CreateSession(c *gin.Context) {
	var form LoginForm
	c.ShouldBind(&form)

	// var uri SessionURI
	// c.ShouldBind(uri)

	log.Println("Success form binding.")
	log.Println(form.LoginId)
	log.Println(form.Password)
	log.Println("End")

	session := sessions.Default(c)

	if Authenticate(form.LoginId, form.Password) != nil {
		log.Println("Fail to authenticate")
		c.Redirect(http.StatusFound, "/users/new")
		return
	}

	log.Println("Success to authenticate")

	repo := repositories.UserRepository{}
	u, err := repo.FindByLoginId(form.LoginId)
	if err != nil {
		log.Println("Fail to findByLoginId")
		c.Redirect(http.StatusFound, "/users/new")
		return
	}

	log.Println("Success to findByLoginId")
	Login(session, u.LoginId)

	c.Redirect(302, "/users/"+strconv.Itoa(int(u.ID)))
}

func Authenticate(loginId string, password string) error {

	fmt.Println(loginId)
	repo := repositories.UserRepository{}
	u, err := repo.FindByLoginId(loginId)

	log.Println(u, err)

	if err != nil {
		return err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)); err != nil {
		return err
	}

	return nil
}

func Login(s sessions.Session, loginId string) {
	s.Set("login_id", loginId)
	s.Save()
}

func DestroySession(c *gin.Context) {
	session := sessions.Default(c)
	Logout(session)
	c.Redirect(302, "/about")
}

func Logout(s sessions.Session) {
	s.Clear()
	s.Save()
}
