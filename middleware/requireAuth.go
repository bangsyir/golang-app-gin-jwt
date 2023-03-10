package middleware

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/bangsyir/go-jwt/initializers"
	"github.com/bangsyir/go-jwt/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func RequireAuth(c *gin.Context) {
	// get the cookie from the request
	tokenString, err := c.Cookie("Authorization")
	if len(tokenString) == 0 {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	// decode/validate it
	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(os.Getenv("SECRET")), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// check exp
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		// find the user with token
		var user models.User
		initializers.DB.First(&user, claims["sub"])
		if user.ID == 0 {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		// attach to req
		c.Set("user", user)
		// continue
		c.Next()
		fmt.Println(claims["foo"], claims["nbf"])
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

}
