package service

import (
	"Pc_Build_Service/Internal/genericresponse"
	"Pc_Build_Service/Internal/userinterface/service/model"
	"context"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"os"
	"time"
)

func (s *impl) UserLogin(ctx context.Context, req *model.UserLoginRequest) (string, error) {
	// Get the user data from the database
	resp, err := s.mysqlStore.GetUser(ctx, req.Email)
	if err != nil {
		return "", &genericresponse.GenericResponse{
			StatusCode: http.StatusUnauthorized,
			Message:    "Invalid Email or Password",
		}
	}

	// Compare hashed password from the database with the plain text password from the request
	err = bcrypt.CompareHashAndPassword([]byte(resp.Password), []byte(req.Password))
	if err != nil {
		return "", &genericresponse.GenericResponse{
			StatusCode: http.StatusUnauthorized,
			Message:    "Invalid Email or Password",
		}
	}

	// Define JWT claims
	claims := jwt.MapClaims{}
	claims["email"] = resp.Email
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // Token expires in 24 hours

	// Create token
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign token with secret key
	secretKey := getSecretKey()
	if secretKey == "" {
		return "", &genericresponse.GenericResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    "Error generating token",
		}
	}

	tokenString, err := jwtToken.SignedString([]byte(secretKey))
	if err != nil {
		return "", fmt.Errorf("error generating token: %v", err)
	}

	// Return the generated token and no error
	return tokenString, nil
}
func getSecretKey() string {
	secretKey := os.Getenv("JWT_SECRET_KEY")
	if secretKey == "" {
		// Handle missing secret key
		fmt.Println("JWT_SECRET_KEY environment variable is not set")
		return ""
	}
	return secretKey
}
