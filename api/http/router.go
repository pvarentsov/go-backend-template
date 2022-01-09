package http

import (
	"go-backend-template/internal/dto"
	"go-backend-template/internal/errors"
	"go-backend-template/internal/util/crypto"

	"github.com/gin-gonic/gin"
)

type router struct {
	server *Server
}

func newRouter(server *Server) *router {
	return &router{
		server: server,
	}
}

func (r *router) init() {
	r.server.engine.Use(r.trace())
	r.server.engine.Use(r.recover())

	r.server.engine.POST("/login", r.login)

	r.server.engine.POST("/users", r.addUser)
	r.server.engine.GET("/users/me", r.authenticate, r.getMe)
	r.server.engine.PUT("/users/me", r.authenticate, r.updateMe)
	r.server.engine.PATCH("/users/me/password", r.authenticate, r.updateMe)

	r.server.engine.NoRoute(r.methodNotFound)
}

// Auth methods

func (r *router) login(c *gin.Context) {
	var loginUserDTO dto.LoginUser

	if err := bindBody(&loginUserDTO, c); err != nil {
		errorResponse(err, nil).reply(c)
		return
	}

	user, err := r.server.usecases.Auth.Login(c, loginUserDTO)
	if err != nil {
		errorResponse(err, nil).reply(c)
		return
	}

	okResponse(user).reply(c)
}

func (r *router) authenticate(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")

	userId, err := r.server.usecases.Auth.VerifyAccessToken(token)
	if err != nil {
		response := errorResponse(err, nil)
		c.AbortWithStatusJSON(response.Status, response)
	}

	SetUserId(c, userId)
}

// User methods

func (r *router) addUser(c *gin.Context) {
	var addUserDTO dto.AddUser

	if err := bindBody(&addUserDTO, c); err != nil {
		errorResponse(err, nil).reply(c)
		return
	}

	user, err := r.server.usecases.User.Add(WithReqInfo(c), addUserDTO)
	if err != nil {
		errorResponse(err, nil).reply(c)
		return
	}

	okResponse(user).reply(c)
}

func (r *router) updateMe(c *gin.Context) {
	var updateUserDTO dto.UpdateUserInfo

	contextUserId, exists := c.Get("userId")
	if exists {
		updateUserDTO.Id = contextUserId.(int64)
	}
	if err := bindBody(&updateUserDTO, c); err != nil {
		errorResponse(err, nil).reply(c)
		return
	}

	err := r.server.usecases.User.UpdateInfo(WithReqInfo(c), updateUserDTO)
	if err != nil {
		errorResponse(err, nil).reply(c)
		return
	}

	okResponse(nil).reply(c)
}

func (r *router) getMe(c *gin.Context) {
	var userId int64

	contextUserId, exists := c.Get("userId")
	if exists {
		userId = contextUserId.(int64)
	}

	user, err := r.server.usecases.User.GetById(WithReqInfo(c), userId)
	if err != nil {
		errorResponse(err, nil).reply(c)
		return
	}

	okResponse(user).reply(c)
}

// System

func (r *router) methodNotFound(c *gin.Context) {
	err := errors.New(errors.NotFoundError, "method not found")
	errorResponse(err, nil).reply(c)
}

func (r *router) recover() gin.HandlerFunc {
	return gin.CustomRecovery(func(c *gin.Context, recovered interface{}) {
		response := errorResponse(nil, nil)
		c.AbortWithStatusJSON(response.Status, response)
	})
}

func (r *router) trace() gin.HandlerFunc {
	return func(c *gin.Context) {
		traceId := c.Request.Header.Get("Trace-Id")
		if traceId == "" {
			traceId, _ = crypto.GenerateUUID()
		}

		SetTraceId(c, traceId)
	}
}
