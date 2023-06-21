package controllers

import (
	"log"
	"net/http"
	"udemy/Golong/section18/app/models"
)

func signUp(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		// sessionがあるかどうかを確認
		_, err := session(w, r)
		if err != nil {
			// ログインしていない場合は、signupページを表示
			generateHTML(w, nil, "layout", "public_navbar", "signup")
		} else {
			// ログインしている場合は、todoページにリダイレクト
			http.Redirect(w, r, "/todos", 302)
		}
	} else if r.Method == "POST" {
		// データをFormから取得できるようにする
		err := r.ParseForm()
		if err != nil {
			log.Println(err)
		}
		user := models.User{
			Name:     r.PostFormValue("name"),
			Email:    r.PostFormValue("email"),
			PassWord: r.PostFormValue("password"),
		}
		if err := user.CreateUser(); err != nil {
			log.Println(err)
		}

		//　ユーザー登録後成功後、ログイン画面にリダイレクト
		http.Redirect(w, r, "/", 302)
	}
}

func login(w http.ResponseWriter, r *http.Request) {
	_, err := session(w, r)
	if err != nil {
		generateHTML(w, nil, "layout", "public_navbar", "login")
	} else {
		http.Redirect(w, r, "/todos", 302)
	}
}

func authenticate(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
	}
	user, err := models.GetUserByEmail(r.PostFormValue("email"))
	if err != nil {
		log.Println(err)
		http.Redirect(w, r, "/login", 302)
	}
	if user.PassWord == models.Encrypt(r.PostFormValue("password")) {
		session, err := user.CreateSession()
		if err != nil {
			log.Println(err)
		}

		// ! Cookieをセット
		cookie := http.Cookie{
			Name:     "_cookie",
			Value:    session.UUID,
			HttpOnly: true,
		}

		http.SetCookie(w, &cookie)
		http.Redirect(w, r, "/", 302)
	} else {
		http.Redirect(w, r, "/login", 302)
	}
}

func logout(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("_cookie")
	if err != nil {
		log.Println(err)
	}
	if err != http.ErrNoCookie {
		session := models.Session{UUID: cookie.Value}
		session.DeleteSessionByUUID()
	}
	http.Redirect(w, r, "/login", 302)
}
