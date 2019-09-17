package utils

/**
 * \* @author: WangChen
 * \* Date: 19-9-17
 * \* Time: 下午4:30
 */
import (
	"github.com/dgrijalva/jwt-go"
	"time"
)


type Claims struct {
	Identify string `json:"identify"`
	jwt.StandardClaims
}


/*  jwttest  就是一个服务器生成的字符串 能够标明是服务器签发Token */
var jwtSecret = "2337736075"

//为用户签发token
func GenerateUserToken(identify string) (string, error){
	return GenerateToken(identify, "ROLE_USER")
}

//为管理员签发token
func GenerateADMINToken(identify string) (string, error){
	return GenerateToken(identify, "ROLE_ADMIN")
}


func GenerateToken(identify string,sub string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(1 * time.Hour) //中间件有效期三个小时

	claims := Claims{
		identify,//用户为id 管理员为用户名
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Subject: sub,
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return nil, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}