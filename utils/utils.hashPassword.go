package utils

import (
    "golang.org/x/crypto/bcrypt"
    "github.com/dgrijalva/jwt-go"
	"crud-app/config"
    "time"
    "fmt"
)

func HashPassword(password string) (string, error) {
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        return "", err
    }
    return string(hashedPassword), nil
}

func VerifyPassword(hashedPassword, password string) error {
    return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func GenerateJwtToken(userId string , email string) (string, error) {
    token := jwt.New(jwt.SigningMethodHS256)
    claims := token.Claims.(jwt.MapClaims)
    claims["id"] = userId
    claims["email"] = email
    claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
    tokenString, err := token.SignedString([]byte(config.Config().JWT_SECRET))
    return tokenString, err
}

func VerifyToken(tokenString string) (jwt.MapClaims, error) {
    secretKey := []byte(config.Config().JWT_SECRET)
    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, fmt.Errorf("invalid signing method")
        }
        return secretKey, nil
    })

    if err != nil {
        return nil, fmt.Errorf("error parsing token: %v", err)
    }

    if token.Valid {
        if claims, ok := token.Claims.(jwt.MapClaims); ok {
            return claims, nil
        }
    }

    return nil, fmt.Errorf("invalid token")
}
