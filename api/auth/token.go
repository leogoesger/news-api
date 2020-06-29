package auth

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// CreateToken create the token
func CreateToken(userID uint32) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = userID
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix() //Token expires after 1 hour
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("API_SECRET")))

}

// TokenValid create the token
func TokenValid(r *http.Request) (*http.Request, error) {
	tokenString := ExtractToken(r)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("API_SECRET")), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		tokenObj, err := Pretty(claims)
		if err != nil {
			return nil, err
		}
		r2 := new(http.Request)
		*r2 = *r
		r2.Header.Set("user-id", strconv.Itoa(tokenObj.UserID))
		return r2, nil
	}
	return nil, nil
}

// ExtractToken get token
func ExtractToken(r *http.Request) string {
	keys := r.URL.Query()
	token := keys.Get("token")
	if token != "" {
		return token
	}
	for _, cookie := range r.Cookies() {
		if cookie.Name == "news-token"{
			return cookie.Value
		}
		break
	}
	return ""
}

// ExtractTokenID get token id
func ExtractTokenID(r *http.Request) (uint32, error) {

	tokenString := ExtractToken(r)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("API_SECRET")), nil
	})
	if err != nil {
		return 0, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		uid, err := strconv.ParseUint(fmt.Sprintf("%.0f", claims["user_id"]), 10, 32)
		if err != nil {
			return 0, err
		}
		return uint32(uid), nil
	}
	return 0, nil
}

// Token struct
type Token struct {
	authorized bool
	exp int
	UserID int `json:"user_id"`
}
//Pretty display the claims licely in the terminal
func Pretty(data interface{}) (Token, error) {
	var obj Token
	b, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		log.Println(err)
		return obj, err
	}

	if err := json.Unmarshal(b, &obj); err != nil {
		panic(err)
	}

	return obj, nil
}