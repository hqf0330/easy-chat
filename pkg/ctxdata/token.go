package ctxdata

import "github.com/golang-jwt/jwt/v4"

const Identify = "binghu"

func GetJwtToken(secretKey, uid string, iat, seconds int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims[Identify] = uid
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}
