package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"jonopens/sitemapper/internal/repositories"
)

// UserHandler handles user-related HTTP requests
type UserHandler struct {
	db repositories.Database
}

// NewUserHandler creates a new user handler
func NewUserHandler(db repositories.Database) *UserHandler {
	return &UserHandler{db: db}
}

// Get retrieves a specific user
// GET /api/v1/users/:id
func (h *UserHandler) Get(c *gin.Context) {
	id := c.Param("id")
	user, err := h.db.Users().GetByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "User not found",
		})
		return
	}
	
	c.JSON(http.StatusOK, user)
}

// Create creates a new user
// POST /api/v1/users
func (h *UserHandler) Create(c *gin.Context) {
	// TODO: Implement user creation logic
	// 1. Parse request body
	// 2. Validate input
	// 3. Create user in database
	// 4. Return response
	c.JSON(http.StatusCreated, gin.H{
		"message": "User created successfully",
	})
}

