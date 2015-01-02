package routes

import (
	"fmt"
	"net/http"

	"github.com/gavruk/go-blog-example/session"

	"github.com/martini-contrib/render"
)

func GetLoginHandler(rnd render.Render) {
	rnd.HTML(200, "login", nil)
}

func PostLoginHandler(rnd render.Render, r *http.Request, s *session.Session) {
	username := r.FormValue("username")
	password := r.FormValue("password")

	fmt.Println(username)
	fmt.Println(password)

	s.Username = username
	s.IsAuthorized = true

	rnd.Redirect("/")
}

func LogoutHandler(rnd render.Render, r *http.Request, s *session.Session) {
	s.Username = ""
	s.IsAuthorized = false

	rnd.Redirect("/")
}
