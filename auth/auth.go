package auth

import (
	"errors"
	"go_project_structure/config"
	"go_project_structure/database"
	"go_project_structure/utils"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)


type AuthRepository interface {
	GenerateToken(username string) (string,error)
	EncodingPassword(string) (string,error)
	CompareHashAndPassword(hash, password string) (bool)
	LoginHandler(username string,password string) (UserData,error)
	ValidateToken(token string) (error)
}

type authRepository struct {}

func NewAuthRepository() AuthRepository {
	return &authRepository{}
}

func (r *authRepository) GenerateToken(username string) (string,error){
	claim := &jwt.MapClaims{
		"Username" : string(username),
		"ExpiresAt": jwt.NewNumericDate(time.Now().Add(1 * time.Minute)),
		"IssuedAt": jwt.NewNumericDate(time.Now()),
		"NotBefore": jwt.NewNumericDate(time.Now().Add(1 * time.Minute)),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim);
	secret := config.GetSecret();
	ss,err := token.SignedString([]byte(secret));
	if err != nil {
		return "",err
	};
	return ss,nil
}

func (r *authRepository) ValidateToken(token string) (error){
	parsedToken, err := jwt.Parse(token, 
		func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("error validate token")
		}
		return []byte(config.GetSecret()), nil
	})
	if err != nil {
		return err
	}
	if !parsedToken.Valid || parsedToken.Claims.(jwt.MapClaims)["ExpiresAt"].(float64) < float64(time.Now().Unix()) {
		return errors.New("Unauthorized")
	}
	return nil
}

func (r *authRepository) EncodingPassword(password string) (string,error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost);
	return string(bytes),err	
}

func (r *authRepository) CompareHashAndPassword(hash, password string) (bool){
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (r *authRepository) LoginHandler(username string,	password string) (UserData,error) {
	var user User
	var userData UserData

	query := `SELECT * FROM users WHERE "Username" = ?  AND "IsActive" = true LIMIT 1`
	if err := database.DB.Raw(query, username).Scan(&user).Error; err != nil {
		return userData, utils.ErrInvalidCredentials
	}

	if !r.CompareHashAndPassword(user.Password,password){
		return userData, utils.ErrInvalidCredentials
	}

	token, err := r.GenerateToken(user.Username)
	if err != nil {
		return userData, utils.ErrTokenGeneration
	}

	userData.Username = user.Username
	userData.AccessToken = token
	return userData,nil
}