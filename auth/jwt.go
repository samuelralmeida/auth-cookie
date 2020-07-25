package auth

import (
	"Project/auth-cookie/config"
	"Project/auth-cookie/model"
	"errors"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// Claims store data to identify user, don't store password
type Claims struct {
	UserID uint   `json:"userid"`
	Name   string `json:name`
	jwt.StandardClaims
}

const CookieName string = "auth-cookie"

func CreateCookie(user *model.User) (*http.Cookie, error) {
	expirationTime := getExpireTime()

	claims := &Claims{
		UserID: user.ID,
		Name:   user.Name,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: expirationTime.Unix(),
		},
	}

	tokenString, err := createToken(claims)
	if err != nil {
		return nil, err
	}

	cookie := &http.Cookie{}
	SetBasicCookie(cookie)
	cookie.Value = tokenString
	cookie.Expires = expirationTime

	return cookie, nil
}

func UpdateCookie(cookie *http.Cookie, claims *Claims) error {
	err := claims.Valid()
	if err != nil {
		return err
	}

	expirationTime := getExpireTime()
	claims.ExpiresAt = expirationTime.Unix()

	// Update expiration time token
	tokenString, err := createToken(claims)
	if err != nil {
		return err
	}

	// Update cookie
	SetBasicCookie(cookie)
	cookie.Value = tokenString
	cookie.Expires = expirationTime

	return nil
}

func ValidCookie(r *http.Request) (*http.Cookie, *Claims, error) {
	cookie, err := r.Cookie(CookieName)
	if err == http.ErrNoCookie {
		return nil, nil, err
	}

	claims := &Claims{}

	token, err := jwt.ParseWithClaims(cookie.Value, claims, callbackKey)
	if err != nil {
		return nil, nil, err
	}

	if !token.Valid {
		return nil, nil, errors.New("Invalid cookie token")
	}

	return cookie, claims, nil
}

func SetBasicCookie(cookie *http.Cookie) {
	cookie.Name = CookieName
	cookie.HttpOnly = true
	cookie.Path = "/"
}

func getExpireTime() time.Time {
	return time.Now().Add(3 * time.Minute)
}

func createToken(claims *Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(config.JWTKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func callbackKey(token *jwt.Token) (interface{}, error) {
	return config.JWTKey, nil
}
