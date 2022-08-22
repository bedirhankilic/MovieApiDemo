package businesslayer

import (
	"time"

	"github.com/bedirhankilic/movieapicase/constants"
	dataaccesslayer "github.com/bedirhankilic/movieapicase/data-access-layer"
	"github.com/bedirhankilic/movieapicase/models"
	"github.com/dgrijalva/jwt-go"
	"gorm.io/gorm"
)

type IAuthService interface {
	Login(username string, password string) (models.User, constants.ErrorCode)
	GenerateJWT(email string, username string) (string, constants.ErrorCode)
	ValidateToken(tokenValid string) constants.ErrorCode
}

var jwtKey = []byte("supersecretkey")

type AuthService struct {
	Db      *gorm.DB
	DbError error
}

func NewAuthService() *AuthService {
	db, err := dataaccesslayer.DbContext()
	return &AuthService{
		Db:      db,
		DbError: err,
	}
}

func (service AuthService) Login(username string, password string) (models.User, constants.ErrorCode) {
	if service.DbError != nil {
		return models.User{}, constants.DB_ERROR
	}

	if username == "" || password == "" {
		return models.User{}, constants.WRONG_REQ_FORMAT
	}
	var user models.User
	result := service.Db.Where("username = ? and password = ?", username, password).First(&user)
	if result.RowsAffected == 0 {
		return models.User{}, constants.NOT_FOUND
	}
	return user, constants.OK
}

func (service AuthService) GenerateJWT(email string, username string) (string, constants.ErrorCode) {
	expirationTime := time.Now().Add(1 * time.Hour)
	claims := &models.JWTClaim{
		Email:    email,
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", constants.JWT_NOT_GENERATED
	}
	return tokenString, constants.OK
}

func (service AuthService) ValidateToken(tokenValid string) constants.ErrorCode {
	token, err := jwt.ParseWithClaims(
		tokenValid,
		&models.JWTClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtKey), nil
		},
	)
	if err != nil {
		return constants.JWT_TOKEN_ERROR
	}
	claims, ok := token.Claims.(*models.JWTClaim)
	if !ok {
		return constants.JWT_TOKEN_ERROR
	}
	if claims.ExpiresAt < time.Now().Local().Unix() {
		return constants.JWT_TOKEN_EXPERIED
	}
	return constants.OK
}
