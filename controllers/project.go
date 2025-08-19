package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"ambridge-backend/database"
	"ambridge-backend/models"
)

// ProjectRequest represents the request body for project operations
type ProjectRequest struct {
	ProjLink     string `json:"projlink"`
	Title        string `json:"title" binding:"required"`
	Type         string `json:"type"`
	Cover        string `json:"cover"`
	Logo         string `json:"logo"`
	ProfileName  string `json:"profilename"`
	ProfilePic   string `json:"profilepic"`
	AboutProject string `json:"aboutproject"`
	Technologies string `json:"technologies"`
	LinkedinLink string `json:"linkedin_link"`
	TelegramLink string `json:"telegram_link"`
	XLink        string `json:"x_link"`
	YoutubeLink  string `json:"youtube_link"`
	GithubLink   string `json:"github_link"`
	InstaLink    string `json:"insta_link"`
}

// CreateProject handles the creation of a new project
func CreateProject(c *gin.Context) {
	var request ProjectRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	project := models.Project{
		ProjLink:     request.ProjLink,
		Title:        request.Title,
		Type:         request.Type,
		Cover:        request.Cover,
		Logo:         request.Logo,
		ProfileName:  request.ProfileName,
		ProfilePic:   request.ProfilePic,
		AboutProject: request.AboutProject,
		Technologies: request.Technologies,
		LinkedinLink: request.LinkedinLink,
		TelegramLink: request.TelegramLink,
		XLink:        request.XLink,
		YoutubeLink:  request.YoutubeLink,
		GithubLink:   request.GithubLink,
		InstaLink:    request.InstaLink,
	}

	if result := database.DB.Create(&project); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create project"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  "success",
		"message": "Project created successfully",
		"project": project,
	})
}

// GetAllProjects returns all projects
func GetAllProjects(c *gin.Context) {
	var projects []models.Project
	if result := database.DB.Find(&projects); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve projects"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":   "success",
		"projects": projects,
	})
}

// GetProject returns a specific project by ID
func GetProject(c *gin.Context) {
	id := c.Param("id")
	var project models.Project

	if result := database.DB.First(&project, id); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Project not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"project": project,
	})
}

// UpdateProject updates a specific project
func UpdateProject(c *gin.Context) {
	id := c.Param("id")
	var project models.Project

	// Check if project exists
	if result := database.DB.First(&project, id); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Project not found"})
		return
	}

	// Parse request
	var request ProjectRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update project fields
	project.ProjLink = request.ProjLink
	project.Title = request.Title
	project.Type = request.Type
	project.Cover = request.Cover
	project.Logo = request.Logo
	project.ProfileName = request.ProfileName
	project.ProfilePic = request.ProfilePic
	project.AboutProject = request.AboutProject
	project.Technologies = request.Technologies
	project.LinkedinLink = request.LinkedinLink
	project.TelegramLink = request.TelegramLink
	project.XLink = request.XLink
	project.YoutubeLink = request.YoutubeLink
	project.GithubLink = request.GithubLink
	project.InstaLink = request.InstaLink

	// Save changes
	if result := database.DB.Save(&project); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update project"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Project updated successfully",
		"project": project,
	})
}

// DeleteProject removes a project from the database
func DeleteProject(c *gin.Context) {
	id := c.Param("id")

	// Check if the ID is a valid number
	projectID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid project ID"})
		return
	}

	// Delete the project
	result := database.DB.Delete(&models.Project{}, projectID)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete project"})
		return
	}

	// Check if anything was deleted
	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Project not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Project deleted successfully",
	})
}
