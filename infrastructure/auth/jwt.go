package auth

import (
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

const identityKey = "identity"

type JWTMiddleware struct {
	*jwt.GinJWTMiddleware
}

type JWTMiddlewareOption func(*jwt.GinJWTMiddleware)

func NewJWTMiddleware(opts ...JWTMiddlewareOption) (*JWTMiddleware, error) {
	cfg := &jwt.GinJWTMiddleware{
		Realm:         "sonic odyssey",
		Timeout:       time.Hour,
		MaxRefresh:    time.Hour,
		IdentityKey:   identityKey,
		TokenLookup:   "header: Authorization, query: token, cookie: jwt",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
	}

	for _, opt := range opts {
		opt(cfg)
	}

	md, err := jwt.New(cfg)
	return &JWTMiddleware{md}, err
}

// Required option functions
func WithKey(key []byte) JWTMiddlewareOption {
	return func(cfg *jwt.GinJWTMiddleware) {
		cfg.Key = key
	}
}

func WithPayloadFunc(fn func(data interface{}) jwt.MapClaims) JWTMiddlewareOption {
	return func(cfg *jwt.GinJWTMiddleware) {
		cfg.PayloadFunc = fn
	}
}

func WithIdentityHandler(fn func(c *gin.Context) interface{}) JWTMiddlewareOption {
	return func(cfg *jwt.GinJWTMiddleware) {
		cfg.IdentityHandler = fn
	}
}

func WithAuthenticator(fn func(c *gin.Context) (interface{}, error)) JWTMiddlewareOption {
	return func(cfg *jwt.GinJWTMiddleware) {
		cfg.Authenticator = fn
	}
}

func WithAuthorizator(fn func(data interface{}, c *gin.Context) bool) JWTMiddlewareOption {
	return func(cfg *jwt.GinJWTMiddleware) {
		cfg.Authorizator = fn
	}
}

func WithUnauthorized(fn func(c *gin.Context, code int, message string)) JWTMiddlewareOption {
	return func(cfg *jwt.GinJWTMiddleware) {
		cfg.Unauthorized = fn
	}
}

// Optional option functions
func WithRealm(realm string) JWTMiddlewareOption {
	return func(cfg *jwt.GinJWTMiddleware) {
		cfg.Realm = realm
	}
}

func WithTimeout(d time.Duration) JWTMiddlewareOption {
	return func(cfg *jwt.GinJWTMiddleware) {
		cfg.Timeout = d
	}
}

func WithMaxRefresh(d time.Duration) JWTMiddlewareOption {
	return func(cfg *jwt.GinJWTMiddleware) {
		cfg.MaxRefresh = d
	}
}

func WithLoginResponse(fn func(c *gin.Context, code int, token string, expire time.Time)) JWTMiddlewareOption {
	return func(cfg *jwt.GinJWTMiddleware) {
		cfg.LoginResponse = fn
	}
}
