package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"ambridge-backend/database"
	"ambridge-backend/models"
)

// CrewRequest represents the request body for crew operations
type CrewRequest struct {
	Username string `json:"username" binding:"required"`
	Role     string `json:"role" binding:"required"`
	About    string `json:"about"`
	URLPhoto string `json:"urlphoto"`
}

// IsAdmin checks if the user is an admin
func isAdmin(c *gin.Context) bool {
	// Get the user ID from the context (set by the auth middleware)
	userID, exists := c.Get("user_id")
	if !exists {
		return false
	}

	// Get the user from the database
	var user models.User
	if result := database.DB.First(&user, userID); result.Error != nil {
		return false
	}

	// Check if the user is an admin
	return user.Role == "admin"
}

// CreateCrew handles the creation of a new crew member
// Only admins can create crew members
func CreateCrew(c *gin.Context) {
	// Check if the user is an admin
	if !isAdmin(c) {
		c.JSON(http.StatusForbidden, gin.H{"error": "Only admins can create crew members"})
		return
	}

	var request CrewRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	crew := models.Crew{
		Username: request.Username,
		Role:     request.Role,
		About:    request.About,
		URLPhoto: request.URLPhoto,
	}

	if result := database.DB.Create(&crew); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create crew member"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  "success",
		"message": "Crew member created successfully",
		"crew":    crew,
	})
}

// GetAllCrewMembers returns all crew members
func GetAllCrewMembers(c *gin.Context) {
	var crews []models.Crew
	if result := database.DB.Find(&crews); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve crew members"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"crews":  crews,
	})
}

// GetCrewMember returns a specific crew member by ID
func GetCrewMember(c *gin.Context) {
	id := c.Param("id")
	var crew models.Crew

	if result := database.DB.First(&crew, id); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Crew member not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"crew":   crew,
	})
}

// UpdateCrewMember updates a specific crew member
// Only admins can update crew members
func UpdateCrewMember(c *gin.Context) {
	// Check if the user is an admin
	if !isAdmin(c) {
		c.JSON(http.StatusForbidden, gin.H{"error": "Only admins can update crew members"})
		return
	}

	id := c.Param("id")
	var crew models.Crew

	// Check if crew member exists
	if result := database.DB.First(&crew, id); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Crew member not found"})
		return
	}

	// Parse request
	var request CrewRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update crew member fields
	crew.Username = request.Username
	crew.Role = request.Role
	crew.About = request.About
	crew.URLPhoto = request.URLPhoto

	// Save changes
	if result := database.DB.Save(&crew); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update crew member"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Crew member updated successfully",
		"crew":    crew,
	})
}

// DeleteCrewMember removes a crew member from the database
// Only admins can delete crew members
func DeleteCrewMember(c *gin.Context) {
	// Check if the user is an admin
	if !isAdmin(c) {
		c.JSON(http.StatusForbidden, gin.H{"error": "Only admins can delete crew members"})
		return
	}

	id := c.Param("id")

	// Check if the ID is a valid number
	crewID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid crew member ID"})
		return
	}

	// Delete the crew member
	result := database.DB.Delete(&models.Crew{}, crewID)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete crew member"})
		return
	}

	// Check if anything was deleted
	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Crew member not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Crew member deleted successfully",
	})
}
