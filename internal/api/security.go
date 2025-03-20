package api

import (
	"app/internal/dto"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

// handleLogin GoDoc
//
//	@Summary	Login
//	@Schemes
//	@Description	Authorization with help username and password
//	@Param			request	body	dto.UserLogin	true	"Model"
//	@Tags			Security
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	dto.UserResponseDto
//	@Router			/api/v1/login [POST]
func (s *Server) handleLogin(c *gin.Context) {
	var login dto.UserLogin

	if err := c.ShouldBindJSON(&login); err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}

	user, err := s.securityUseCase.LoginUser(login)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}

	session := sessions.Default(c)
	session.Set(UserKey, user.Username)
	err = session.Save()

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": "cannot save cookies",
		})
		return
	}

	c.JSON(http.StatusOK, user)

}

// handleLogout GoDoc
//
//	@Summary	Logout
//	@Schemes
//	@Description	Logout from account
//	@Tags			Security
//	@Accept			json
//	@Produce		json
//	@Success		200
//	@Router			/api/v1/logout [POST]
func (s *Server) handleLogout(c *gin.Context) {
	user := c.MustGet("user")

	if user == nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "unauthorized",
		})
		return
	}

	session := sessions.Default(c)
	session.Delete(UserKey)
	err := session.Save()

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": "cannot save cookies",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}
