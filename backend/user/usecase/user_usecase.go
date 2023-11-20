package usecase

import (
	"context"
	"log"
	"time"

	"github.com/PUArallelepiped/PUN-street-Universal-Access/domain"
	"github.com/golang-jwt/jwt"
	"github.com/spf13/viper"
)

type userUsecase struct {
	userRepo domain.UserRepo
}

func NewUserUsecase(userRepo domain.UserRepo) domain.UserUsecase {
	return &userUsecase{
		userRepo: userRepo,
	}
}

type Claims struct {
	Email     string `json:"email"`
	Authority int    `json:"authority"`
	jwt.StandardClaims
}

func init() {
	viper.SetConfigFile("../.env")
	viper.SetConfigType("dotenv")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}
}

func CreateToken(email string, authority int) (string, error) {
	expiresAt := time.Now().Add(24 * time.Hour).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, Claims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiresAt,
		},
		Email:     email,
		Authority: authority,
	})
	jwtSecret := []byte(viper.GetString("JWT_SECRET"))
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (uu *userUsecase) Login(ctx context.Context, email string, password string) (string, error) {
	authority, err := uu.userRepo.Login(ctx, email, password)
	if err != nil {
		return "", err
	}
	token, err := CreateToken(email, authority)
	if err != nil {
		return "", err
	}

	return token, nil
}
