package libauth

import (
	"context"
	"fmt"
	"time"

	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	jwtv4 "github.com/golang-jwt/jwt/v4"
)

type Claims struct {
	ID               int64             `json:"id"`
	Type             ClaimsType        `json:"ct"`
	UserType         UserType          `json:"ut"`
	TransferMetadata *TransferMetadata `json:"tm,omitempty"`
	jwtv4.RegisteredClaims
}

type ClaimsType int

const (
	ClaimsTypeUnspecified ClaimsType = iota
	ClaimsTypeAccessToken
	ClaimsTypeRefreshToken
	ClaimsTypeUploadToken
	ClaimsTypeDownloadToken
)

type UserType int

const (
	UserTypeUnspecified UserType = iota
	UserTypeAdmin
	UserTypeNormal
	UserTypeSentinel
)

type TransferMetadata struct {
	Name      string `json:"n"`
	Size      int64  `json:"s"`
	ChunkSize int64  `json:"cs"`
	CallBack  int    `json:"cb"`
}

func KeyFunc(key string, ty ClaimsType) jwtv4.Keyfunc {
	return func(token *jwtv4.Token) (interface{}, error) {
		return fmt.Sprintf("%s%d", key, ty), nil
	}
}

func NewClaims() jwtv4.Claims {
	return Claims{}
}

func FromContext(ctx context.Context) (*Claims, bool) {
	if token, ok := jwt.FromContext(ctx); ok {
		if claims, met := token.(Claims); met {
			return &claims, true
		}
	}
	return nil, false
}

func (a *Auth) GenerateToken(id int64, claimsType ClaimsType, userType UserType,
	transferMetadata *TransferMetadata, expire time.Duration) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(expire)

	claims := Claims{
		ID:               id,
		Type:             claimsType,
		UserType:         userType,
		TransferMetadata: transferMetadata,
		RegisteredClaims: jwtv4.RegisteredClaims{
			ExpiresAt: jwtv4.NewNumericDate(expireTime),
			Issuer:    a.config.Issuer,
		},
	}

	tokenClaims := jwtv4.NewWithClaims(jwtv4.SigningMethodHS256, claims)

	token, err := tokenClaims.SignedString([]byte(a.config.JwtSecret))
	return token, err
}
