package http

import (
	"go-backend-template/internal/dto"
	"go-backend-template/internal/util/crypto"
	"go-backend-template/internal/util/errors"

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
	r.server.engine.PUT("/users/me", r.authenticate, r.updateMyInfo)
	r.server.engine.PATCH("/users/me/password", r.authenticate, r.changeMyPassword)

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

	setUserId(c, userId)
}

// User methods

func (r *router) addUser(c *gin.Context) {
	var addUserDTO dto.AddUser

	if err := bindBody(&addUserDTO, c); err != nil {
		errorResponse(err, nil).reply(c)
		return
	}

	user, err := r.server.usecases.User.Add(withReqInfo(c), addUserDTO)
	if err != nil {
		errorResponse(err, nil).reply(c)
		return
	}

	okResponse(user).reply(c)
}

func (r *router) updateMyInfo(c *gin.Context) {
	var updateUserDTO dto.UpdateUserInfo

	reqInfo := getReqInfo(c)
	updateUserDTO.Id = reqInfo.UserId

	if err := bindBody(&updateUserDTO, c); err != nil {
		errorResponse(err, nil).reply(c)
		return
	}

	err := r.server.usecases.User.UpdateInfo(withReqInfo(c), updateUserDTO)
	if err != nil {
		errorResponse(err, nil).reply(c)
		return
	}

	okResponse(nil).reply(c)
}

func (r *router) changeMyPassword(c *gin.Context) {
	var changeUserPasswordDTO dto.ChangeUserPassword

	reqInfo := getReqInfo(c)
	changeUserPasswordDTO.Id = reqInfo.UserId

	if err := bindBody(&changeUserPasswordDTO, c); err != nil {
		errorResponse(err, nil).reply(c)
		return
	}

	err := r.server.usecases.User.ChangePassword(withReqInfo(c), changeUserPasswordDTO)
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

	user, err := r.server.usecases.User.GetById(withReqInfo(c), userId)
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

		setTraceId(c, traceId)
	}
}
