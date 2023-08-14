package authcontrollers

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/hardzal/go-example/go-books/pkg/config"
	"github.com/hardzal/go-example/go-books/pkg/models"
	"github.com/hardzal/go-example/go-books/pkg/utils"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"encoding/json"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	// mengambil inputan json
	var userInput models.User
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&userInput); err != nil {
		response := map[string]string{"message": err.Error()}
		utils.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}
	defer r.Body.Close() // menutup proses json decoder

	// check user
	var user models.User
	// error handling
	if err := models.Db.Where("username = ?", userInput.Username).First(&user).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			response := map[string]string{"message": "Username atau password salah!"}
			utils.ResponseJSON(w, http.StatusUnauthorized, response)
			return
		default:
			response := map[string]string{"message": err.Error()}
			utils.ResponseJSON(w, http.StatusInternalServerError, response)
			return
		}
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userInput.Password)); err != nil {
		response := map[string]string{"message": "Username atau password salah"}
		utils.ResponseJSON(w, http.StatusUnauthorized, response)
		return
	}

	// proses pembuatan token jwt
	expTime := time.Now().Add(time.Minute * 1)
	claims := &config.JWTClaim{
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "go-books",
			ExpiresAt: jwt.NewNumericDate(expTime),
		},
	}

	// proses signing
	tokenAlgo := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// signed token
	token, err := tokenAlgo.SignedString(config.JWT_KEY)
	if err != nil {
		response := map[string]string{"message": err.Error()}
		utils.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}

	// menyimpan token ke cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Path:     "/",
		Value:    token,
		HttpOnly: true,
	})

	response := map[string]string{"message": "Login berhasil"}
	utils.ResponseJSON(w, http.StatusOK, response)
}

func Register(w http.ResponseWriter, r *http.Request) {
	// mengambil inputan json
	var userInput models.User
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&userInput); err != nil {
		response := map[string]string{"message": err.Error()}
		utils.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}
	defer r.Body.Close() // menutup proses json decoder

	// hash password
	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(userInput.Password), bcrypt.DefaultCost)
	userInput.Password = string(hashPassword)
	// currentTime := time.Now()
	// userInput.CreatedAt = currentTime
	// userInput.UpdatedAt = currentTime

	if err := models.Db.Create(&userInput).Error; err != nil {
		response := map[string]string{"message": err.Error()}
		utils.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}

	response, _ := json.Marshal(map[string]string{"message": "success"})
	w.Header().Add("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func Logout(w http.ResponseWriter, r *http.Request) {

}
