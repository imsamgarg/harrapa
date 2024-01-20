package utils

import (
	crypto "crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"harrapa/internal/database"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

var jwtSecret = os.Getenv("JWT_SECRET")

type JwtPayload struct {
	UserId   string
	UserRole string
	// ApiKey   string
	jwt.StandardClaims
}

func GenerateOTP(length int) string {

	otp := ""
	for i := 0; i < length; i++ {
		otp += fmt.Sprintf("%d", rand.Intn(10))
	}

	return otp
}

func GenerateJWT(id string, userRole database.UserRole) (string, error) {
	claims := JwtPayload{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 1).Unix(), // Token expires in 1 hour
		},
		// ApiKey:   apiKey,
		UserId:   id,
		UserRole: string(userRole),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	secretKey := []byte(jwtSecret)
	signedToken, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func ParseJwt(tokenString string) (*JwtPayload, error) {
	secretKey := []byte(jwtSecret)
	token, err := jwt.ParseWithClaims(tokenString, &JwtPayload{}, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*JwtPayload); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token")
}

func ExtractJwtFromHeader(header http.Header) (*JwtPayload, error) {
	tokenSplit := strings.Split(header.Get("authorization"), " ")

	if len(tokenSplit) != 2 {
		return nil, errors.New("invalid token")
	}

	return ParseJwt(tokenSplit[1])
}

func GenerateRandomString(length int) (string, error) {

	bytes := make([]byte, length)

	_, err := crypto.Read(bytes)

	return base64.StdEncoding.EncodeToString(bytes), err
}

func GenerateHashedPassword(val string) (string, error) {
	//TODO(sam): handle max pass length
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(val), bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}

	return string(hashedPass), nil

}

func CompareHashAndPassword(hashedPass string, pass string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPass), []byte(pass))

	if err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			return false, nil
		}

		return false, err
	}

	return true, nil

}
