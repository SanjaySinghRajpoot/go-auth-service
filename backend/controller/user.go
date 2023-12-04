package controller

import (
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/SanjaySinghRajpoot/backend/config"
	"github.com/SanjaySinghRajpoot/backend/models"
	"github.com/SanjaySinghRajpoot/backend/utils"
	tokens "github.com/SanjaySinghRajpoot/backend/utils/token"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// use godot package to load/read the .env file and
// return the value of the key
func goDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

// Get all Users from the Database
func GetUsers(ctx *gin.Context) {
	users := []models.User{}
	config.DB.Find(&users)
	ctx.JSON(200, &users)
}

// Create New user Admin Only Access
func PostUsers(ctx *gin.Context) {

	var postUser models.PostUser
	ctx.BindJSON(&postUser)

	var adminUser models.User // get the admin User profile
	config.DB.Where("email = ?", postUser.AdminEmail).First(&adminUser)

	// Password Encryption
	password := postUser.Password
	hash, err := utils.HashPassword(password)

	if err != nil {
		ctx.JSON(500, err)
	}

	if adminUser.Admin {
		user := models.User{Name: postUser.Name, Email: postUser.Email, Password: hash, Admin: false}
		config.DB.Create(&user)
		ctx.JSON(200, &user)
	} else {
		ctx.JSON(401, "You are not authorized to perform this action")
	}

}

// Delete a User admin only access
func DeleteUser(ctx *gin.Context) {
	var userId models.UserId // payload
	ctx.BindJSON(&userId)

	var user models.User // get the admin User profile
	config.DB.Where("email = ?", userId.AdminEmail).First(&user)

	var deleteUser models.User // get the User to be deleted
	config.DB.Where("id = ?", userId.Id).First(&deleteUser)

	if user.Admin { // Check if the user had admin previlages
		config.DB.Where("id = ?", userId.Id).Delete(&deleteUser)
		ctx.JSON(200, "User deleted succesfully")
	} else {
		ctx.JSON(401, "You are not authorized to perform this action")
	}
}

func LoginUser(ctx *gin.Context) {
	var user models.LoginUser // payload
	ctx.BindJSON(&user)

	var get_user models.User // get the admin User profile
	config.DB.Where("email = ?", user.Email).First(&get_user)

	err := utils.VerifyPassword(get_user.Password, user.Password)

	if err != nil {
		ctx.JSON(401, "Invalid Password")
	}

	userPayload := models.UserJWT{ID: strconv.Itoa(get_user.Id), Name: get_user.Email}

	accessToken, err := tokens.CreateAccessToken(userPayload, os.Getenv("ACCESS_TOKEN_SECRET"), 1)

	if err != nil {
		ctx.JSON(500, "Unable to generate Cookies")
	}

	expirationTime := time.Now().Add(60 * time.Minute)
	http.SetCookie(ctx.Writer, &http.Cookie{
		Name:    "token",
		Value:   accessToken,
		Expires: expirationTime,
	})

	ctx.JSON(200, accessToken)
}

func LogoutUser(ctx *gin.Context) {

	expirationTime := time.Now().Add(0 * time.Minute)

	// Clear cookies to logout existing user
	http.SetCookie(ctx.Writer, &http.Cookie{
		Name:    "token",
		Value:   "",
		Expires: expirationTime,
	})

	ctx.JSON(200, "logout Succesfull")
}

func RefreshUser(ctx *gin.Context) {

	var user models.UserJWT // payload
	ctx.BindJSON(&user)

	userPayload := models.UserJWT{ID: user.ID, Name: user.Name}

	refreshToken, err := tokens.CreateRefreshToken(userPayload, os.Getenv("REFRESH_TOKEN_SECRET"), 24)

	if err != nil {
		ctx.JSON(500, err)
	}

	expirationTime := time.Now().Add(24 * time.Minute)

	http.SetCookie(ctx.Writer, &http.Cookie{
		Name:    "token",
		Value:   refreshToken,
		Expires: expirationTime,
	})

	ctx.JSON(200, refreshToken)
}
