package main

import (
	"log"
	"net/http"
	"strings"

	"github.com/codegangsta/negroni"
	sessions "github.com/goincremental/negroni-sessions"
	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/gomniauth"
	"github.com/stretchr/gomniauth/providers/google"
	"github.com/stretchr/objx"
)

const (
	nextPageKey     = "nextPage"
	authSecurityKey = "authKey"
)

func init() {
	gomniauth.SetSecurityKey(authSecurityKey)
	gomniauth.WithProviders(
		google.New("636296155193-a9abes4mc1p81752l116qkr9do6oev3f.apps.googleusercontent.com", "EVvuy0Agv4jWflml0pvC6-vI", "http://127.0.0.1:3000/auth/callback/google"),
	)
}

func loginHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	action := ps.ByName("action")
	provider := ps.ByName("provider")

	s := sessions.GetSession(r)

	switch action {
	case "login":
		// gomniauth.Provider의 login 페이지로 이동
		p, err := gomniauth.Provider(provider)
		if err != nil {
			log.Fatalln(err)
		}
		loginUrl, err := p.GetBeginAuthURL(nil, nil)
		if err != nil {
			log.Fatalln(err)
		}
		http.Redirect(w, r, loginUrl, http.StatusFound)
	case "callback":
		// gomniauth 콜백 처리
		p, err := gomniauth.Provider(provider)
		if err != nil {
			log.Fatalln(err)
		}
		creds, err := p.CompleteAuth(objx.MustFromURLQuery(r.URL.RawQuery))
		if err != nil {
			log.Fatalln(err)
		}

		// 콜백 결과로부터 사용자 정보 확인
		user, err := p.GetUser(creds)
		if err != nil {
			log.Fatalln(err)
		}

		if err != nil {
			log.Fatalln(err)
		}

		u := &User{
			Uid:       user.Data().Get("id").MustStr(),
			Name:      user.Name(),
			Email:     user.Email(),
			AvatarUrl: user.AvatarURL(),
		}

		SetCurrentUser(r, u)
		http.Redirect(w, r, s.Get(nextPageKey).(string), http.StatusFound)
	default:
		http.Error(w, "Auth action '"+action+"' is not supported", http.StatusNotFound)
	}

}

func LoginRequired(ignore ...string) negroni.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		for _, s := range ignore {
			if strings.HasPrefix(r.URL.Path, s) {
				next(w, r)
				return // ignore할 url이면 다음거 실행
			}
		}

		u := GetCurrentUser(r)

		if u != nil && u.Valid() {
			SetCurrentUser(r, u)
			next(w, r)
			return
		}

		// 유저가 비 유효시 nil
		SetCurrentUser(r, nil)

		// 로그인 후 이동할 url을 세션에 저장(r)
		sessions.GetSession(r).Set(nextPageKey, r.URL.RequestURI())

		// 로그인 페이지로 리다이렉트
		http.Redirect(w, r, "/login", http.StatusFound)
	}

}
