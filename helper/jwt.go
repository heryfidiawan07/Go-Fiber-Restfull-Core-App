package helper

import (
	"api-fiber-gorm/config"
	"api-fiber-gorm/model"
	"fmt"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

func CreateJwtToken(user *model.User) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["identity"] = user.Id
	claims["exp"] = time.Now().Add(time.Minute * 5).Unix()

	tokenString, err := token.SignedString([]byte(config.Config("SECRET")))
	if err != nil {
		return tokenString, err
	}

	return tokenString, nil
}

func CreateRefreshJwtToken(refreshToken *model.RefreshToken) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["identity"] = refreshToken.Id
	claims["exp"] = time.Now().Add(time.Minute * 120).Unix()

	tokenString, err := token.SignedString([]byte(config.Config("SECRET")))
	if err != nil {
		return tokenString, err
	}

	return tokenString, nil
}

func JwtParse(c *fiber.Ctx) (interface{}, string) {
	type Header struct {
		Authorization string
	}

	headers := new(Header)

	if err := c.ReqHeaderParser(headers); err != nil {
		// fmt.Println("ReqHeaderParser", err)
		return "", "Error ReqHeaderParser !"
	}

	authorization := strings.Split(headers.Authorization, " ")
	accessToken := authorization[1]

	// fmt.Println("authorization", authorization)
	// fmt.Println("accessToken", accessToken)

	token, _ := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			// return false, "Error JWT Parse"
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(config.Config("SECRET")), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// fmt.Println("USER ID ****************", claims["identity"])
		// fmt.Println("EXP ****************", claims["exp"])

		return claims["identity"], ""
	} else {
		return "", "Invalid token !"
	}
}

func RefreshTokenParse(c *fiber.Ctx, tokenString string) (interface{}, string) {
	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			// return false, "Error JWT Parse"
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(config.Config("SECRET")), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// fmt.Println("USER ID ****************", claims["identity"])
		// fmt.Println("EXP ****************", claims["exp"])

		return claims["identity"], ""
	} else {
		return "", "Invalid token !"
	}
}
