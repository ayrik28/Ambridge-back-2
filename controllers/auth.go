package controllers

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"ambridge-backend/database"
	"ambridge-backend/models"
	"ambridge-backend/utils"
)

// RegisterRequest represents the request body for user registration
type RegisterRequest struct {
	Name            string `json:"name" binding:"required"`
	Surname         string `json:"surname" binding:"required"`
	Email           string `json:"email" binding:"required,email"`
	Password        string `json:"password" binding:"required,min=6"`
	ProfileImage    string `json:"profileImage"`
	Referral        string `json:"referral"`
	Company         string `json:"company"`
	CurrentPosition string `json:"currentPosition"`
}

// LoginRequest represents the request body for user login
type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

// RefreshTokenRequest represents the request body for token refresh
type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

// UpdateProfileRequest represents the request body for updating user profile
type UpdateProfileRequest struct {
	Name            *string `json:"name"`
	Surname         *string `json:"surname"`
	ProfileImage    *string `json:"profileImage"`
	Referral        *string `json:"referral"`
	Company         *string `json:"company"`
	CompanyEmail    *string `json:"companyEmail"`
	CompanyAddress  *string `json:"companyAddress"`
	CompanyPhone    *string `json:"companyPhone"`
	CurrentPosition *string `json:"currentPosition"`
	ResumeFile      *string `json:"resumeFile"`
}

// Register handles user registration
func Register(c *gin.Context) {
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if user already exists
	var existingUser models.User
	result := database.DB.Where("email = ?", strings.ToLower(req.Email)).First(&existingUser)
	if result.Error == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Email already registered"})
		return
	} else if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}

	// Hash password
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	// Create user
	user := models.User{
		Name:           req.Name,
		Surname:        req.Surname,
		Email:          strings.ToLower(req.Email),
		Password:       hashedPassword,
		Role:           "user",
		ProfileImage:   req.ProfileImage,
		ReferralSource: req.Referral,
		CompanyName:    req.Company,
		Position:       req.CurrentPosition,
	}

	// Set default profile image if not provided
	if user.ProfileImage == "" {
		user.ProfileImage = "/default-avatar.png"
	}

	if err := database.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User registered successfully",
		"user_id": user.ID,
	})
}

// Login handles user login
func Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Find user by email
	var user models.User
	result := database.DB.Where("email = ?", strings.ToLower(req.Email)).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		}
		return
	}

	// Check password
	if !utils.CheckPassword(user.Password, req.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	// Generate JWT token
	token, err := utils.GenerateJWT(user.ID, user.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	// Generate refresh token
	refreshToken, err := utils.GenerateRefreshToken()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate refresh token"})
		return
	}

	// Store refresh token in database
	user.RefreshToken = refreshToken

	if err := database.DB.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to store refresh token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token":         token,
		"refresh_token": refreshToken,
		"user": gin.H{
			"id":              user.ID,
			"name":            user.Name,
			"surname":         user.Surname,
			"email":           user.Email,
			"role":            user.Role,
			"profileImage":    user.ProfileImage,
			"referral":        user.ReferralSource,
			"company":         user.CompanyName,
			"currentPosition": user.Position,
		},
	})
}

// Logout handles user logout
func Logout(c *gin.Context) {
	// Get user ID from JWT token
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// Clear refresh token in database
	if err := database.DB.Model(&models.User{}).Where("id = ?", userID).Update("refresh_token", "").Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to logout"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Logged out successfully"})
}

// RefreshToken handles token refresh
func RefreshToken(c *gin.Context) {
	var req RefreshTokenRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Find user by refresh token
	var user models.User
	result := database.DB.Where("refresh_token = ?", req.RefreshToken).First(&user)
	if result.Error != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid refresh token"})
		return
	}

	// Generate new JWT token
	token, err := utils.GenerateJWT(user.ID, user.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	// Generate new refresh token
	refreshToken, err := utils.GenerateRefreshToken()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate refresh token"})
		return
	}

	// Store new refresh token in database
	user.RefreshToken = refreshToken

	if err := database.DB.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to store refresh token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token":         token,
		"refresh_token": refreshToken,
	})
}

// GetProfile retrieves the user's profile information
func GetProfile(c *gin.Context) {
	// Get user ID from JWT token (set by AuthMiddleware)
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// Find user by ID
	var user models.User
	result := database.DB.First(&user, userID)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		}
		return
	}

	// Return user profile data (excluding sensitive fields)
	c.JSON(http.StatusOK, gin.H{
		"user": gin.H{
			"name":            user.Name,
			"surname":         user.Surname,
			"email":           user.Email,
			"role":            user.Role,
			"profileImage":    user.ProfileImage,
			"referral":        user.ReferralSource,
			"company":         user.CompanyName,
			"companyEmail":    user.CompanyEmail,
			"companyAddress":  user.CompanyAddress,
			"companyPhone":    user.CompanyPhone,
			"currentPosition": user.Position,
			"resumeFile":      user.ResumeFile,
		},
	})
}

// AdminCheckRequest represents the request body for admin check
type AdminCheckRequest struct {
	Username string `json:"username" binding:"required"`
}

// IsAdmin checks if the current user has admin role
func IsAdmin(c *gin.Context) {
	// Get user ID from JWT token (set by AuthMiddleware)
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// Parse request body to get username
	var req AdminCheckRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Find token user by ID (the user making the request)
	var tokenUser models.User
	result := database.DB.First(&tokenUser, userID)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		}
		return
	}

	// Find user by username (the user to check)
	var targetUser models.User
	result = database.DB.Where("email = ?", strings.ToLower(req.Username)).First(&targetUser)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Username not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		}
		return
	}

	// Check if target user has admin role
	isAdmin := targetUser.Role == "admin"

	// Return result
	c.JSON(http.StatusOK, gin.H{
		"isAdmin": isAdmin,
		"user": gin.H{
			"email": targetUser.Email,
			"role":  targetUser.Role,
		},
	})
}

// UpdateProfile updates the user's profile information
func UpdateProfile(c *gin.Context) {
	// Get user ID from JWT token (set by AuthMiddleware)
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// Parse request body
	var req UpdateProfileRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Find user by ID
	var user models.User
	result := database.DB.First(&user, userID)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		}
		return
	}

	// Update user fields if provided in the request
	if req.Name != nil {
		user.Name = *req.Name
	}
	if req.Surname != nil {
		user.Surname = *req.Surname
	}
	if req.ProfileImage != nil {
		user.ProfileImage = *req.ProfileImage
	}
	if req.Referral != nil {
		user.ReferralSource = *req.Referral
	}
	if req.Company != nil {
		user.CompanyName = *req.Company
	}
	if req.CompanyEmail != nil {
		user.CompanyEmail = *req.CompanyEmail
	}
	if req.CompanyAddress != nil {
		user.CompanyAddress = *req.CompanyAddress
	}
	if req.CompanyPhone != nil {
		user.CompanyPhone = *req.CompanyPhone
	}
	if req.CurrentPosition != nil {
		user.Position = *req.CurrentPosition
	}
	if req.ResumeFile != nil {
		user.ResumeFile = *req.ResumeFile
	}

	// Save updated user to database
	if err := database.DB.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update profile"})
		return
	}

	// Return updated user profile
	c.JSON(http.StatusOK, gin.H{
		"message": "Profile updated successfully",
		"user": gin.H{
			"name":            user.Name,
			"surname":         user.Surname,
			"email":           user.Email,
			"role":            user.Role,
			"profileImage":    user.ProfileImage,
			"referral":        user.ReferralSource,
			"company":         user.CompanyName,
			"companyEmail":    user.CompanyEmail,
			"companyAddress":  user.CompanyAddress,
			"companyPhone":    user.CompanyPhone,
			"currentPosition": user.Position,
			"resumeFile":      user.ResumeFile,
		},
	})
}
