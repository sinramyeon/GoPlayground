package main

import (
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/goincremental/negroni-sessions"
	"github.com/goincremental/negroni-sessions/cookiestore"
	"github.com/julienschmidt/httprouter"
	"github.com/unrolled/render"
)

var renderer *render.Render

const (
	sessionKey    = "session"
	sessionSecret = "secret"
)

func init() {
	// 렌더러 생성
	renderer = render.New()
}

func main() {
	// 라우터 생성
	router := httprouter.New()

	// 핸들러 정의
	router.GET("/", func(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
		// 렌더러를 사용하여 템플릿 렌더링
		renderer.HTML(w, http.StatusOK, "index", map[string]string{"title": "Simple Chat!"})
	})

	router.GET("/login", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

		renderer.HTML(w, http.StatusOK, "login", nil)

	})

	router.GET("/logout", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		// 세션에서 사용자 정보 제거 후 로그인 페이지로 이동

		sessions.GetSession(r).Delete(currentUserKey)
		http.Redirect(w, r, "/login", http.StatusFound)
	})

	// negroni 미들웨어 생성
	n := negroni.Classic()
	store := cookiestore.New([]byte(sessionSecret))
	n.Use(sessions.Sessions(sessionKey, store))

	// negroni에 router를 핸들러로 등록
	n.UseHandler(router)

	// 웹 서버 실행
	n.Run(":3000")
}
