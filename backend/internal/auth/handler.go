package auth

import (
	"github.com/gin-gonic/gin"
)

func NewHandler(s Service) *Handler {
	return &Handler{
		Service: s,
	}
}

func (h *Handler) SignInHandler(c *gin.Context) {

	var req Request

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	token, err := h.Service.SignIn(c.Request.Context(), req.Email, req.Password)
	if err != nil {
		c.JSON(401, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"message": "Login successful",
		"token":   token,
	})

}

func (h *Handler) SignupHandler(c *gin.Context) {
	var req SignupRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	token, err := h.Service.Signup(
		c.Request.Context(),
		req.Email,
		req.Password,
		req.FirstName,
		req.LastName,
	)
	if err != nil {
		c.JSON(401, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"message": "Registration successful",
		"token":   token,
	})
}
