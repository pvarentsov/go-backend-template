package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func okResponse(data interface{}) *response {
	return &response{
		Status:  http.StatusOK,
		Message: "ok",
		Data:    data,
	}
}

func errorResponse(err error, data interface{}) *response {
	status, message := parseError(err)

	return &response{
		Status:  status,
		Message: message,
		Data:    data,
	}
}

func (r *response) reply(c *gin.Context) {
	c.JSON(r.Status, r)
}
