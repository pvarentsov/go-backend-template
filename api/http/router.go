package http

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"

	"go-backend-template/internal/auth"
	"go-backend-template/internal/base/contexts"
	"go-backend-template/internal/base/errors"
	"go-backend-template/internal/user"
)

type router struct {
	*Server
}

func newRouter(server *Server) *router {
	return &router{
		Server: server,
	}
}

func (r *router) init() {
	r.engine.Use(r.trace())
	r.engine.Use(r.recover())
	r.engine.Use(r.logger())

	r.engine.POST("/login", r.login)

	r.engine.POST("/users", r.addUser)
	r.engine.GET("/users/me", r.authenticate, r.getMe)
	r.engine.PUT("/users/me", r.authenticate, r.updateMe)
	r.engine.PATCH("/users/me/password", r.authenticate, r.changeMyPassword)

	r.engine.NoRoute(r.methodNotFound)
}

// Auth methods

func (r *router) login(c *gin.Context) {
	var loginUserDto auth.LoginUserDto

	if err := bindBody(&loginUserDto, c); err != nil {
		errorResponse(err, nil).reply(c)
		return
	}

	user, err := r.authService.Login(c, loginUserDto)
	if err != nil {
		errorResponse(err, nil).reply(c)
		return
	}

	okResponse(user).reply(c)
}

func (r *router) authenticate(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")

	userId, err := r.authService.VerifyAccessToken(token)
	if err != nil {
		response := errorResponse(err, nil)
		c.AbortWithStatusJSON(response.Status, response)
	}

	setUserId(c, userId)
}

// User methods

func (r *router) addUser(c *gin.Context) {
	var addUserDto user.AddUserDto

	if err := bindBody(&addUserDto, c); err != nil {
		errorResponse(err, nil).reply(c)
		return
	}

	user, err := r.userUsecases.Add(contextWithReqInfo(c), addUserDto)
	if err != nil {
		errorResponse(err, nil).reply(c)
		return
	}

	okResponse(user).reply(c)
}

func (r *router) updateMe(c *gin.Context) {
	var updateUserDto user.UpdateUserDto

	reqInfo := getReqInfo(c)
	updateUserDto.Id = reqInfo.UserId

	if err := bindBody(&updateUserDto, c); err != nil {
		errorResponse(err, nil).reply(c)
		return
	}

	err := r.userUsecases.Update(contextWithReqInfo(c), updateUserDto)
	if err != nil {
		errorResponse(err, nil).reply(c)
		return
	}

	okResponse(nil).reply(c)
}

func (r *router) changeMyPassword(c *gin.Context) {
	var changeUserPasswordDto user.ChangeUserPasswordDto

	reqInfo := getReqInfo(c)
	changeUserPasswordDto.Id = reqInfo.UserId

	if err := bindBody(&changeUserPasswordDto, c); err != nil {
		errorResponse(err, nil).reply(c)
		return
	}

	err := r.userUsecases.ChangePassword(contextWithReqInfo(c), changeUserPasswordDto)
	if err != nil {
		errorResponse(err, nil).reply(c)
		return
	}

	okResponse(nil).reply(c)
}

func (r *router) getMe(c *gin.Context) {
	reqInfo := getReqInfo(c)

	user, err := r.userUsecases.GetById(contextWithReqInfo(c), reqInfo.UserId)
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
			traceId, _ = r.crypto.GenerateUUID()
		}

		setTraceId(c, traceId)
	}
}

func (r *router) logger() gin.HandlerFunc {
	return gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		var parsedReqInfo contexts.ReqInfo

		reqInfo, exists := param.Keys[reqInfoKey]
		if exists {
			parsedReqInfo = reqInfo.(contexts.ReqInfo)
		}

		return fmt.Sprintf("%s - [HTTP] TraceId: %s; UserId: %d; Method: %s; Path: %s; Status: %d, Latency: %s;\n\n",
			param.TimeStamp.Format(time.RFC1123),
			parsedReqInfo.TraceId,
			parsedReqInfo.UserId,
			param.Method,
			param.Path,
			param.StatusCode,
			param.Latency,
		)
	})
}
