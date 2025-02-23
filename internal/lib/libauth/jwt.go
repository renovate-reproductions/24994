package libauth

import (
	"context"
	"fmt"
	"time"

	"github.com/tuihub/librarian/internal/model"

	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	jwtv4 "github.com/golang-jwt/jwt/v4"
)

type Claims struct {
	InternalID       model.InternalID `json:"iid,string"`
	Type             ClaimsType       `json:"ct"`
	UserType         UserType         `json:"ut"`
	TransferMetadata any              `json:"tm,omitempty"`
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

func (a *Auth) KeyFunc(t ClaimsType) jwtv4.Keyfunc {
	return func(token *jwtv4.Token) (interface{}, error) {
		return a.generateSecret(t), nil
	}
}

func NewClaims() jwtv4.Claims {
	return new(Claims)
}

func FromContext(ctx context.Context) (*Claims, bool) {
	if token, ok := jwt.FromContext(ctx); ok {
		if claims, met := token.(*Claims); met {
			return claims, true
		}
	}
	return nil, false
}

func FromContextAssertUserType(ctx context.Context, userTypes ...UserType) bool {
	if c, ok := FromContext(ctx); ok {
		for _, ut := range userTypes {
			if c.UserType == ut {
				return true
			}
		}
	}
	return false
}

func (a *Auth) generateSecret(t ClaimsType) interface{} {
	return []byte(fmt.Sprintf("%s%d", a.config.JwtSecret, t))
}

func (a *Auth) GenerateToken(id model.InternalID, claimsType ClaimsType, userType UserType,
	transferMetadata any, expire time.Duration) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(expire)

	claims := Claims{
		InternalID:       id,
		Type:             claimsType,
		UserType:         userType,
		TransferMetadata: transferMetadata,
		RegisteredClaims: jwtv4.RegisteredClaims{
			Issuer:    a.config.Issuer,
			Subject:   "",
			Audience:  nil,
			ExpiresAt: jwtv4.NewNumericDate(expireTime),
			NotBefore: nil,
			IssuedAt:  nil,
			ID:        "",
		},
	}

	tokenClaims := jwtv4.NewWithClaims(jwtv4.SigningMethodHS256, claims)

	token, err := tokenClaims.SignedString(a.generateSecret(claimsType))
	return token, err
}
