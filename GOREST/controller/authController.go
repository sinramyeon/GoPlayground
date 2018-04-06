package controller

import (
	"errors"
	c "hero0926-api-test/config"
	"hero0926-api-test/util"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/sessions"
	"github.com/vmihailenco/msgpack"
)

// User ...
type User struct {
	Name     string
	Password string
}

const HOST = "localhost:8090" // localhost.
const EXPIRE_TIME = 86400     // 1 day
const SecretKey = "hellofresh"

var token string

// store ...
// Secret key cookie store
var store = sessions.NewCookieStore([]byte(c.ENV.Secret.Key))

var cookies []*http.Cookie

// IsAuthenticated ...
// check the request is by authentificated user
func IsAuthenticated(f func(w http.ResponseWriter, req *http.Request)) func(w http.ResponseWriter, req *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {

		err := LoginHandler(w, req)

		if err != nil {
			c.Errorf(`[ERROR] authController.IsAuthenticated err `, err)

			util.WriteJSON(w, http.StatusUnauthorized, nil)
			return
		}

		f(w, req)
	}
}

// LoginHandler ...
// Check User Auth and Returns Error
func LoginHandler(w http.ResponseWriter, r *http.Request) error {

	w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)

	// 1. Setup Default User and Create Redis User Info
	u := User{
		Name:     "hellofresh",
		Password: "hellofresh",
	}

	// 2. Get Session and Parsing Form
	session, _ := store.Get(r, "sesseion.id")
	err := r.ParseForm()

	if err != nil {
		return err
	}

	// 3. Check User Auth Compare with Redis
	username, password, authOK := r.BasicAuth()

	if authOK == false {
		c.Errorf("[ERROR] authController.LoginHandler authOK == false")
		return errors.New("[ERROR] authController.LoginHandler authOK == false")
	}

	if password == u.Password && username == u.Name {

		CreateUserInfo(&User{
			Name:     username,
			Password: password,
		})

		session.Save(r, w)
		token = MakeJWTToken(&u)

	} else {
		Destroy(&u)
		DeleteCookie(&u)
		c.Errorf(`[ERROR] authController.LoginHandler password and username is wrong`)
		return errors.New(`[ERROR] authController.LoginHandler password and username is wrong`)
	}

	// 4. Check JWT Token
	_, err = jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})

	return err

}

// makeSessionName ...
// make Redis Session DB name
func makeSessionName(user *User) string {

	username := user.Name
	uuid := util.CreateUUID()
	return username + "-" + uuid
}

// CreateUserInfo ...
// create valid user info in redis for 1 day.
func CreateUserInfo(user *User) {

	b, _ := msgpack.Marshal(user)
	sessionName := makeSessionName(user)
	SetCookie(user)

	c.RedisSession.Set(sessionName, []byte(b))

	// 1day :  86400sec
	c.RedisSession.Expire(sessionName, EXPIRE_TIME)
}

// SetCookie ...
func SetCookie(user *User) *http.Cookie {
	jar, _ := cookiejar.New(nil)

	cookie := &http.Cookie{
		Name:    user.Name,
		Value:   user.Password,
		Path:    "/",
		Expires: time.Now().AddDate(0, 0, 1),
		MaxAge:  EXPIRE_TIME,
	}

	cookies = append(cookies, cookie)

	u, _ := url.Parse(HOST)
	jar.SetCookies(u, cookies)

	return cookie
}

// Destroy ...
func Destroy(user *User) {

	// 1. get rid of global redis user session
	c.RedisSession.Delete(user.Name)
	// 2. delete cookie
	DeleteCookie(user)
}

// DeleteCookie ...
func DeleteCookie(user *User) {
	jar, _ := cookiejar.New(nil)
	cookie := &http.Cookie{
		Name:    user.Name,
		Value:   user.Password,
		Path:    "/",
		Expires: time.Unix(0, 0),
		MaxAge:  -1,
	}
	cookies = append(cookies, cookie)

	u, _ := url.Parse(HOST)
	jar.SetCookies(u, cookies)
}

// MakeJWTToken ...
// Make own JWT token for authorise
func MakeJWTToken(user *User) string {

	// Create JWT token
	token := jwt.New(jwt.GetSigningMethod("HS256"))
	token.Claims = jwt.MapClaims{ // can't use Claims index https://github.com/dgrijalva/jwt-go/issues/143
		"name": user.Name,
		"exp":  time.Now().Add(time.Minute * 5).Unix(), // expires in 5 min
	}

	tokenstring, err := token.SignedString([]byte(SecretKey))
	if err != nil {
		c.Errorf(`[ERROR] authController.MakeJWTToken : tokenstring, err := token.SignedString([]byte("foobar"))`)
	}

	return tokenstring

}
