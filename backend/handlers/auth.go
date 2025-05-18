package handlers

import (
	"daily-worker-roster-management-system/handlers/response"
	"daily-worker-roster-management-system/usecase"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"strings"
	"time"
)

type AuthHandler struct {
	userUsecase usecase.UserUsecase // Assuming you have a user usecase for validating users
	jwtSecret   []byte
}

func NewAuthHandler(userUsecase usecase.UserUsecase, secret []byte) *AuthHandler {
	return &AuthHandler{userUsecase: userUsecase, jwtSecret: secret}
}

type loginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req loginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "invalid request")
		return
	}

	// Validate user credentials (you implement this)
	roleId, err := h.userUsecase.ValidateCredentials(req.Username, req.Password)
	if err != nil {
		response.Error(c, http.StatusUnauthorized, "invalid credentials")
		return
	}

	// Create JWT token
	claims := jwt.MapClaims{
		"role_id": roleId,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(h.jwtSecret)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "could not generate token")
		return
	}

	response.JSON(c, http.StatusOK, map[string]interface{}{"access_token": tokenString, "role_id": roleId})
}

func (h *AuthHandler) VerifyToken(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")

	if authHeader == "" {
		response.Error(c, http.StatusUnauthorized, "Authorization header is missing")
		return
	}

	authToken := strings.Split(authHeader, " ")
	if len(authToken) != 2 || authToken[0] != "Bearer" {
		response.Error(c, http.StatusUnauthorized, "Invalid token format")
		return
	}

	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(authToken[1], &claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return h.jwtSecret, nil
	})
	if err != nil || !token.Valid {
		response.Error(c, http.StatusUnauthorized, "Invalid or expired token")
		return
	}
	roleId, ok := claims["role_id"]
	if !ok {
		response.Error(c, http.StatusUnauthorized, "Invalid token")
		return
	}
	response.JSON(c, http.StatusOK, map[string]interface{}{"role_id": roleId})
	return
}
