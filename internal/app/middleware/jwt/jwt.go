package jwt

import (
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// JWTConfig JWT配置
type JWTConfig struct {
	SecretKey     string
	TokenExpireIn time.Duration
}

// 默认配置
var defaultConfig = JWTConfig{
	SecretKey:     "tier_up_secret_key",
	TokenExpireIn: time.Hour * 24, // 1天
}

// CustomClaims 自定义JWT声明
type CustomClaims struct {
	UserID   int    `json:"user_id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// JWTService JWT服务
type JWTService struct {
	Config JWTConfig
}

// NewJWTService 创建JWT服务
func NewJWTService() *JWTService {
	return &JWTService{Config: defaultConfig}
}

// GenerateToken 生成JWT令牌
func (s *JWTService) GenerateToken(userID int, username string) (string, error) {
	// 创建自定义声明
	claims := &CustomClaims{
		UserID:   userID,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(s.Config.TokenExpireIn)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "tier_up",
			Subject:   username,
		},
	}

	// 创建JWT令牌
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 签名令牌
	tokenString, err := token.SignedString([]byte(s.Config.SecretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// ParseToken 解析JWT令牌
func (s *JWTService) ParseToken(tokenString string) (*CustomClaims, error) {
	// 解析令牌
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.Config.SecretKey), nil
	})

	if err != nil {
		return nil, err
	}

	// 验证令牌
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("无效的令牌")
}

// JWTAuthMiddleware JWT认证中间件
func (s *JWTService) JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取Authorization头
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "message": "未提供授权令牌"})
			c.Abort()
			return
		}

		// 检查Bearer前缀
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "message": "令牌格式错误"})
			c.Abort()
			return
		}

		// 解析令牌
		claims, err := s.ParseToken(parts[1])
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "message": "无效的令牌: " + err.Error()})
			c.Abort()
			return
		}

		// 将用户信息存储到上下文中
		c.Set("userID", claims.UserID)
		c.Set("username", claims.Username)

		c.Next()
	}
}
