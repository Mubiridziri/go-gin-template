package api

import (
	"app/internal/dto"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// handleProfile GoDoc
//
//	@Summary	Profile
//	@Schemes
//	@Description	You can check auth or get profile data
//	@Tags			Security
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	dto.UserResponseDto
//	@Router			/api/v1/login [GET]
func (s *Server) handleProfile(c *gin.Context) {
	user := c.MustGet("user")

	if user == nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "unauthorized",
		})
		return
	}

	c.JSON(http.StatusOK, user)
}

// handleListUser GoDoc
//
//	@Summary	Get Users List
//	@Schemes
//	@Description	List of users
//	@Param			page	query	int	false	"Page"
//	@Param			limit	query	int	false	"Limit of page"
//	@Tags			Users
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	[]dto.PaginatedUsersList
//	@Router			/api/v1/users [GET]
func (s *Server) handleListUser(c *gin.Context) {
	query, err := s.getListQuery(c)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "bad query params: " + err.Error(),
		})
		return
	}

	rows, err := s.getUsersListUseCase.ListUsers(query.Page, query.Limit)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": "Server error: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, rows)

}

// handleCreateUser GoDoc
//
//	@Summary	Create User
//	@Schemes
//	@Description	Creating user
//	@Param			user	body	dto.SaveUserDto	true	"User"
//	@Tags			Users
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	dto.UserResponseDto
//	@Router			/api/v1/users [POST]
func (s *Server) handleCreateUser(c *gin.Context) {
	var createUser dto.SaveUserDto

	if err := c.ShouldBindJSON(&createUser); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "Bind error: " + err.Error(),
		})
		return
	}

	user, err := s.createUserUseCase.CreateUser(createUser)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": "Server error: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, user)

}

// handleUpdateUser GoDoc
//
//	@Summary	Update User
//	@Schemes
//	@Description	Updating user
//	@Param			user	body	dto.SaveUserDto	true	"User"
//	@Tags			Users
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	dto.UserResponseDto
//	@Router			/api/v1/users [PUT]
func (s *Server) handleUpdateUser(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "Invalid ID param in path",
		})
		return
	}

	var createUser dto.SaveUserDto

	if err := c.ShouldBindJSON(&createUser); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "Bind error: " + err.Error(),
		})
		return
	}

	user, err := s.updateUserUseCase.UpdateUser(id, createUser)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": "Server error: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, user)
}

// handleDetailUser GoDoc
//
//	@Summary	Detail User
//	@Schemes
//	@Description	Get user info by user id
//	@Param			id	path	int	true	"User ID"
//	@Tags			Users
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	dto.UserResponseDto
//	@Router			/api/v1/users/{id} [GET]
func (s *Server) handleDetailUser(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "Invalid ID param in path",
		})
		return
	}

	user, err := s.getUserUseCase.GetUserById(id)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error": "Not found",
		})
		return
	}

	c.JSON(http.StatusOK, user)
}

// handleDeleteUser GoDoc
//
//	@Summary	Delete User
//	@Schemes
//	@Description	Deleting user
//	@Param			id	path	int	true	"User ID"
//	@Tags			Users
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	dto.UserResponseDto
//	@Router			/api/v1/users/{id} [DELETE]
func (s *Server) handleDeleteUser(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "Invalid ID param in path",
		})
		return
	}

	user, err := s.removeUserUseCase.RemoveUser(id)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error": "Not found",
		})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (s *Server) AddUserRoutes(g *gin.RouterGroup) {
	grp := g.Group("/users")
	grp.GET("", s.handleListUser)
	grp.POST("", s.handleCreateUser)
	grp.PUT("/:id", s.handleUpdateUser)
	grp.GET("/:id", s.handleDetailUser)
	grp.DELETE("/:id", s.handleDeleteUser)
}
