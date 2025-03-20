package api

import (
	"app/docs"
	"app/internal/config"
	"app/internal/dto"
	"app/internal/service/users"
	"net/http"
	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

const UserKey = "AUTH"

type ServerParams struct {
	Config              *config.Config
	SecurityUseCase     users.SecurityUseCase
	GetUserUseCase      users.GetUserUseCase
	GetUsersListUseCase users.GetUsersListUseCase
	CreateUserUseCase   users.CreateUserUseCase
	UpdateUserUseCase   users.UpdateUserUseCase
	RemoveUserUseCase   users.RemoveUserUseCase
}

type Server struct {
	Config              *config.Config
	Router              *gin.Engine
	securityUseCase     users.SecurityUseCase
	getUserUseCase      users.GetUserUseCase
	getUsersListUseCase users.GetUsersListUseCase
	createUserUseCase   users.CreateUserUseCase
	updateUserUseCase   users.UpdateUserUseCase
	removeUserUseCase   users.RemoveUserUseCase
}

func New(params ServerParams) *Server {
	s := Server{
		Config:              params.Config,
		Router:              gin.New(),
		securityUseCase:     params.SecurityUseCase,
		getUserUseCase:      params.GetUserUseCase,
		getUsersListUseCase: params.GetUsersListUseCase,
		createUserUseCase:   params.CreateUserUseCase,
		updateUserUseCase:   params.UpdateUserUseCase,
		removeUserUseCase:   params.RemoveUserUseCase,
	}

	s.registerRoutes(s.Config.AppSecret)
	return &s
}

func (s *Server) registerRoutes(appSecret string) {
	//Middleware
	s.Router.Use(gin.Logger())
	s.Router.Use(gin.Recovery())
	s.Router.Use(sessions.Sessions(UserKey, cookie.NewStore([]byte(appSecret))))

	//Swagger
	configureSwagger(s.Router)

	// K8s probe
	s.Router.GET("/healthz", func(c *gin.Context) { c.Status(http.StatusOK) })

	//API /api/v1
	mainGroup := s.Router.Group("/api/v1")
	mainGroup.POST("/login", s.handleLogin)

	mainGroup.Use(AuthRequired(s.getUserUseCase))
	{
		s.AddUserRoutes(mainGroup)

		mainGroup.GET("/login", s.handleProfile)
		mainGroup.GET("/logout", s.handleLogout)
	}
}

func configureSwagger(r *gin.Engine) {
	//Swagger
	docs.SwaggerInfo.Title = "app"
	docs.SwaggerInfo.Description = "AI Seller"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Schemes = []string{"http"}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

func AuthRequired(useCase users.GetUserUseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		userKey := session.Get(UserKey)

		if userKey == nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "unauthorized",
			})
			return
		}

		authUser, err := useCase.GetUserByUsername(userKey.(string))

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "unauthorized",
			})
			return
		}

		c.Set("user", authUser)
		c.Next()
	}

}

func (s *Server) getListQuery(c *gin.Context) (dto.ListQuery, error) {
	pageQuery := c.Query("page")
	limitQuery := c.Query("limit")
	simplifyQuery := c.Query("simplify")

	if simplifyQuery == "" {
		simplifyQuery = "false"
	}

	page, err := strconv.Atoi(pageQuery)
	if err != nil || page == 0 {
		page = 1
	}

	limit, err := strconv.Atoi(limitQuery)
	if err != nil || limit == 0 {
		limit = 10
	}

	simplify, err := strconv.ParseBool(simplifyQuery)

	if err != nil {
		return dto.ListQuery{}, err
	}

	query := dto.ListQuery{
		Page:     page,
		Limit:    limit,
		Simplify: simplify,
	}
	return query, nil
}
