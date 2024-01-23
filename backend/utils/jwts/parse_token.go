package jwts

import (
	"backend-go/global"
	"github.com/golang-jwt/jwt/v5"
)

// ParseToken token解析
func ParseToken(tokenstring string) (*CustomClaims, error) {
	t, err := jwt.ParseWithClaims(tokenstring, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(global.Config.Jwt.Secret), nil
	})

	if claims, ok := t.Claims.(*CustomClaims); ok && t.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}
