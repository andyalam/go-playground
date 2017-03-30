package main

import (
	"fmt"
	"html/template"
	"net/http"

	uuid "github.com/satori/go.uuid"
)

type user struct {
	Id       string
	First    string
	Last     string
	Email    string
	Password string
}

var tpl *template.Template

var mUser = map[string]user{}      // key is userid, value is user
var mSession = map[string]string{} // key is sessionid, value is userid

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	fmt.Println(http.ListenAndServe(":4800", nil))
}

func index(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("mySession")
	if err != nil {
		u := uuid.NewV4()
		c = &http.Cookie{
			Name:     "mySession",
			Value:    u.String(),
			Path:     "/",
			HttpOnly: true,
		}
	}
	uid := mSession[c.Value]
	usr := mUser[uid]

	http.SetCookie(w, c)
	tpl.ExecuteTemplate(w, "index.gohtml", usr)
}

func login(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("mySession")
	if err != nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	if r.Method == "POST" {
		email := r.FormValue("email")
		password := r.FormValue("password")
		m := map[string]user{}
		for _, v := range mUser {
			m[v.Email] = v
		}
		usr, ok := m[email]
		if !ok || usr.Password != password {
			http.Redirect(w, r, "login", http.StatusForbidden)
			return
		}
		mSession[c.Value] = usr.Id
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	tpl.ExecuteTemplate(w, "login.gohtml", nil)
}

func signup(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("mySession")
	if err != nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		fmt.Println("redirected - no cookie")
		return
	}

	uid := mSession[c.Value]
	usr := mUser[uid]
	if usr.Email != "" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		fmt.Println("redirected - already signed up")
		return
	}

	if r.Method == "POST" {
		fn := r.FormValue("first")
		ln := r.FormValue("last")
		em := r.FormValue("email")
		pwd := r.FormValue("password")
		u := uuid.NewV4()
		usr = user{
			Id:       u.String(),
			First:    fn,
			Last:     ln,
			Email:    em,
			Password: pwd,
		}
		mUser[usr.Id] = usr
		mSession[c.Value] = usr.Id
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	tpl.ExecuteTemplate(w, "signup.gohtml", nil)
}

func logout(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("mySession")
	if err != nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	delete(mSession, c.Value)
	c.MaxAge = -1
	http.SetCookie(w, c)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
