package routes

import (
    "net/http"
    "strings"
    "time"

    "github.com/dgrijalva/jwt-go"
    "github.com/gin-gonic/gin"
)

var jwtSecret = []byte("seu-segredo-jwt") // Alterar para um segredo seguro

// GenerateJWT gera um novo token JWT com um tempo de expiração.
func GenerateJWT(username string) (string, error) {
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "username": username,
        "exp":      time.Now().Add(time.Hour * 72).Unix(), // Token expira em 72 horas
    })
    return token.SignedString(jwtSecret)
}

// AuthMiddleware é um middleware que valida o token JWT.
func AuthMiddleware() gin.HandlerFunc {
    return func(ctx *gin.Context) {
        authHeader := ctx.GetHeader("Authorization")
        if authHeader == "" {
            ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Token necessário"})
            return
        }

        // Extrai o token do cabeçalho "Authorization: Bearer <token>"
        tokenString := strings.TrimPrefix(authHeader, "Bearer ")

        // Valida o token
        token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
            if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
                return nil, gin.Error{Err: jwt.ErrSignatureInvalid}
            }
            return jwtSecret, nil
        })

        if err != nil || !token.Valid {
            ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Token inválido"})
            return
        }

        // Pega as reivindicações (claims) e armazena no contexto
        if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
            ctx.Set("username", claims["username"])
        } else {
            ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Token inválido"})
            return
        }

        ctx.Next()
    }
}
