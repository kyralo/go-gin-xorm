package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"net/http"
	"time"
)

type JWT struct {
	SigningKey []byte
}

var (
	TokenExpired     = errors.New("Token is expired")
	TokenNotValidYet = errors.New("Token not active yet")
	TokenMalformed   = errors.New("That's not even a token")
	TokenInvalid     = errors.New("Couldn't handle this token:")
	SignKey          = "2337736075"
)
// 自定义结构体参数
type CustomClaims struct {
	Id      string `form:"id"`
	UserName  string `form:"username"`
	jwt.StandardClaims
}

// JWT验证
func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := c.Cookie("claims")
		if err != nil {
			c.JSONP(
				http.StatusUnauthorized,gin.H{
					"message" : "401 unauthorized",
				},
			)
			return
		}
		j := NewJWT()
		// token := t.(string)
		// claims, err := j.ParseToken(token)
		claims, err := j.ParseToken(token)
		if err != nil {
			if err == TokenExpired {
				if token, err = j.RefreshToken(token); err == nil {
					c.JSON(http.StatusOK, gin.H{"error": 0, "message": "refresh token", "token": token})
					return
				}
			}

			c.JSONP(
				http.StatusForbidden,gin.H{
					"message" : "403 forbidden",
				},
			)
			return
		}
		c.Set("claims", claims)
		c.Next()
		return
	}
}

func NewJWT() *JWT {
	return &JWT{
		[]byte(GetSignKey()),
	}
}
func GetSignKey() string {
	return SignKey
}

func SetSignKey(key string) string {
	SignKey = key
	return SignKey
}

// parse
func (j *JWT) ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, TokenInvalid
}
func (j *JWT) RefreshToken(tokenString string) (string, error) {
	jwt.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		return "", err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		jwt.TimeFunc = time.Now
		claims.StandardClaims.ExpiresAt = time.Now().Add(1 * time.Hour).Unix()
		return j.CreateToken(*claims)
	}
	return "", TokenInvalid
}

// create
func (j *JWT) CreateToken(claims CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}
