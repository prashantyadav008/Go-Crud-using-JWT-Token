package handler

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"go-crud-msql/internal/models"
	"go-crud-msql/internal/service"
	"go-crud-msql/internal/utils"
)

type UserHandler struct {
	userService service.UserService
	db          *gorm.DB
}

func NewUserHandler(userService service.UserService, db *gorm.DB) *UserHandler {
	return &UserHandler{
		userService: userService,
		db:          db,
	}
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var user models.User

	err := c.ShouldBind(&user)

	body, _ := io.ReadAll(c.Request.Body)
	log.Println("Request Body:", string(body))

	// Bind the JSON request to the User struct
	if err != nil {
		log.Println("user -->>>", err)

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	// Handle image upload
	file, err := c.FormFile("image")
	if err == nil {
		timestamp := time.Now().Unix()
		originalName := file.Filename
		ext := filepath.Ext(originalName)             // eg: ".jpg"
		base := strings.TrimSuffix(originalName, ext) // eg: "my photo"
		base = strings.ReplaceAll(base, " ", "-")     // "my-photo"

		newFileName := fmt.Sprintf("%d-%s%s", timestamp, base, ext) // eg: "1715700000-my-photo.jpg"
		savePath := filepath.Join("internal/uploads/images", newFileName)

		if err := c.SaveUploadedFile(file, savePath); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"errMessage": err.Error(), "error": "Failed to save image"})
			return
		}

		user.ImageName = newFileName
	}

	// Hash the password before saving
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"errMessage": err.Error(), "error": "Failed to hash password"})
		return
	}
	user.Password = hashedPassword

	if err := h.userService.CreateUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"errMessage": err.Error(), "error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User Created successfully"})
}

func (h *UserHandler) GetAllUsers(c *gin.Context) {
	users, err := h.userService.FindAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"errMessage": err.Error(), "error": "Failed to fetch users"})
		return
	}

	c.JSON(http.StatusOK, users)
}

func (h *UserHandler) GetUserById(c *gin.Context) {
	idParams := c.Param("id")

	id, err := strconv.Atoi(idParams)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errMessage": err.Error(), "error": "Invalid user ID"})
		return
	}

	user, err := h.userService.FindById(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"errMessage": err.Error(), "error": "Failed to fetch user"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *UserHandler) UpdateUser(c *gin.Context) {
	idParams := c.Param("id")
	id, err := strconv.Atoi(idParams)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errMessage": err.Error(), "error": "Invalid user ID"})
		return
	}

	var user models.User

	// Bind the JSON request to the User struct
	if err := c.ShouldBind(&user); err != nil {
		log.Println("user -->>>", err)

		c.JSON(http.StatusBadRequest, gin.H{"errMessage": err.Error(), "error": "Invalid input"})

		return
	}

	// Hash the password before saving
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"errMessage": err.Error(), "error": "Failed to hash password"})
		return
	}
	user.Password = hashedPassword

	user.Id = uint(id)
	if err := h.userService.UpdateUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"errMessage": err.Error(), "error": "Failed to update user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}

func (h *UserHandler) DeleteUser(c *gin.Context) {
	idParams := c.Param("id")
	id, err := strconv.Atoi(idParams)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errMessage": err.Error(), "error": "Invalid user ID"})
		return
	}

	if err := h.userService.DeleteUser(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"errMessage": err.Error(), "error": "Failed to delete user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})

}

func (h *UserHandler) LoginUser(c *gin.Context) {
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errMessage": err.Error(), "error": "Invalid input"})
		return
	}

	var user models.User
	log.Println("input -->>>", input)
	log.Println("input.Email -->>>", input.Email)
	log.Println("input.Password -->>>", input.Password)

	if err := h.db.Where("email = ?", input.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	if !utils.CheckPasswordHash(input.Password, user.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	user.Password = ""
	token, err := utils.GenerateJWT(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Token generation failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
		"user":  user,
	})

}

func (h *UserHandler) LogoutUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Logout successful (remove token on client side)"})

}

func (h *UserHandler) GetUserProfile(c *gin.Context) {
	userData, exists := c.Get("userData")
	log.Println("userData Details -->>>", userData)

	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"errMessage": "Unauthorized", "error": "User not authenticated"})
		return
	}

	// remove password from userData
	delete(userData.(map[string]interface{}), "password")

	c.JSON(http.StatusOK, userData)

}
