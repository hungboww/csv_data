package controllers

import (
	"ads/database"
	"ads/helpers"
	"ads/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"os"
	"time"
)

func RegisterUser(c *gin.Context) {
	var body struct {
		Name      string
		Email     string
		Password  string
		Reconfirm string
		Role      string
	}
	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": helpers.CheckValue})
		return
	}
	if body.Password != body.Reconfirm {
		c.JSON(http.StatusBadRequest, gin.H{"message": helpers.PasswordConfirmErr})
		return
	}
	// hash the password
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": helpers.PasswordErr})
		return
	}
	user := models.User{
		Email:    body.Email,
		Password: string(hash),
		UserName: body.Name,
		RoleId:   1,
		ISActive: true,
		IsStaff:  true,
	}
	checkEmail := database.DB.First(&user, "email = ?", body.Email)
	if checkEmail.Error == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": helpers.EmailErr})
		panic("failed to register user")
	}
	//SELECT * FROM users WHERE email=?"
	//checkEmail := database.DB.Where()
	result := database.DB.Create(&user)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": helpers.AccountErrRegister})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": helpers.Success})
}

func LoginAccount(c *gin.Context) {
	var body struct {
		Email    string
		Password string
	}
	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": helpers.CheckValue})
		return
	}
	var user models.User
	database.DB.First(&user, "Email=?", body.Email)
	if user.Id == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": helpers.AccountErr})
		return
	}
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": helpers.PasswordErr})
		return
	}
	// Sign and get the complete encoded token as a string using the secret
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.Id,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": helpers.TokenErr})
	}
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)
	c.Header("Bearer", tokenString)
	c.JSON(http.StatusOK, gin.H{
		"token":   tokenString,
		"message": helpers.LoginSuccess,
	})
}
