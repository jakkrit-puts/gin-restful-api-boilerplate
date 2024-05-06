package userctrl

import (
	"net/http"
	"os"
	"time"

	"example.com/jakkrit/ginbackendapi/configs"
	"example.com/jakkrit/ginbackendapi/models"
	"example.com/jakkrit/ginbackendapi/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/matthewhartstonge/argon2"
)

func Login(c *gin.Context) {

	var login LoginReq

	if err := c.ShouldBindJSON(&login); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User{
		Email:    login.Email,
		Password: login.Password,
	}

	// check has email
	userEmail := configs.DB.Where("email = ?", login.Email).First(&user)
	if userEmail.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "email not found."})
		return
	}

	// compare password
	ok, _ := argon2.VerifyEncoded([]byte(login.Password), []byte(user.Password))

	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "password invalid."})
		return
	}

	// create token
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 24 * 1).Unix(),
	})

	jwtSecretKey := os.Getenv("JWT_SECRET")
	token, _ := claims.SignedString([]byte(jwtSecretKey))

	c.JSON(http.StatusOK, gin.H{
		"message":      "login success.",
		"access_token": token,
	})
}

func Register(c *gin.Context) {

	var register RegisterReq

	if err := c.ShouldBindJSON(&register); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User{
		Fullname: register.Fullname,
		Email:    register.Email,
		Password: register.Password,
		ImageName: utils.UploadImage(register.ImageName),
	}

	// check email duplicate
	userExist := configs.DB.Where("email = ?", register.Email).First(&user)
	if userExist.RowsAffected == 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "email already exist."})
		return
	}

	result := configs.DB.Create(&user)

	// case error
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "register success.",
		"data":    &user,
	})
}

func GetAll(c *gin.Context) {
	var users []models.User
	// configs.DB.Order("id desc").Find(&users)

	configs.DB.Preload("Blogs").Find(&users)

	c.JSON(http.StatusOK, gin.H{
		"data": users,
	})
}

func GetById(c *gin.Context) {
	id := c.Param("id")

	var user models.User
	result := configs.DB.First(&user, id)

	if result.RowsAffected < 1 {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found."})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": user,
	})
}

func SearchByName(c *gin.Context) {
	fullname := c.Query("fullname")

	var users []models.User
	result := configs.DB.Where("fullname LIKE ?", "%"+fullname+"%").Scopes(utils.Paginate(c)).Find(&users)

	if result.RowsAffected < 1 {
		c.JSON(http.StatusNotFound, gin.H{"error": "data not found."})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": users,
	})
}

func GetProfile(c *gin.Context) {

	user := c.MustGet("user")

	c.JSON(http.StatusOK, gin.H{
		"data": user,
	})
}
