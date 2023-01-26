package mid

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/ezzycreative1/svc-blog-profile/config"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

// func ApiKeyMiddleware(apiKey string) middleware {
// 	return middleware.KeyAuthWithConfig(middleware.KeyAuthConfig{
// 		Skipper: func(fiber.Ctx) bool {
// 			return apiKey == ""
// 		},
// 		KeyLookup:  "header:X-API-Key",
// 		AuthScheme: "",
// 		Validator: func(key string, c *fiber.Ctx) (bool, error) {
// 			return key == apiKey, nil
// 		},
// 		ErrorHandler: func(err error, c *fiber.Ctx) error {
// 			response := map[string]interface{}{
// 				"message": "unauthorized",
// 				"data":    nil,
// 				"error":   err.Error(),
// 			}
// 			return c.JSON(response)
// 		},
// 	})
// }

type UserAuth struct {
	Id int64
}

// GenerateToken ..
// func GenerateToken(email string, userid int64) (string, error) {
// 	claims := jwt.MapClaims{}
// 	claims["authorized"] = true
// 	claims["user_id"] = userid
// 	claims["exp"] = time.Now().Add(72 * time.Hour).Unix()
// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
// 	return token.SignedString([]byte(os.Getenv("API_SECRET")))
// }

func GenerateToken(userID int64, tokenType string) (string, error) {
	var config config.Jwt

	claims := jwt.MapClaims{
		"user_id": userID,
		"type":    tokenType,
		"admin":   false,
		"iss":     "blog-api",
		"iat":     time.Now().Unix(),
	}
	if tokenType == "access" {
		claims["exp"] = time.Now().Add(time.Minute * time.Duration(config.AccessExpireMin)).Unix()
	} else if tokenType == "refresh" {
		claims["exp"] = time.Now().Add(time.Minute * time.Duration(config.RefreshExpireMin)).Unix()
	} else {
		return "Please pass access or refresh in tokenType", nil
	}
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := t.SignedString([]byte(config.Secret))
	if err != nil {
		return "", err
	}
	return token, nil
}

// TokenValid ..
func TokenValid(c *fiber.Ctx) error {
	tokenString := ExtractToken(c)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("API_SECRET")), nil
	})
	if err != nil {
		return err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		Pretty(claims)
	}
	return nil
}

// ExtractToken ..
func ExtractToken(c *fiber.Ctx) string {
	token := c.Get("token")
	if token != "" {
		return token
	}
	bearerToken := c.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}
	return ""
}

// ExtractTokenID ..
func ExtractTokenID(c *fiber.Ctx) (int64, error) {

	tokenString := ExtractToken(c)
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
		uid, err := strconv.ParseUint(fmt.Sprintf("%.0f", claims["user_id"]), 10, 64)
		if err != nil {
			return 0, err
		}
		return int64(uid), nil
	}
	return 0, nil
}

//Pretty display the claims
func Pretty(data interface{}) {
	b, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(string(b))
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// GenerateRandString ..
func GenerateRandString(lg int) string {

	var letterRunes = []rune("123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	b := make([]rune, lg)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)

}
