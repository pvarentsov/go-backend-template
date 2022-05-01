package http

import (
	"github.com/gin-gonic/gin"

	"go-backend-template/internal/base/errors"
)

func bindBody(payload interface{}, c *gin.Context) error {
	err := c.BindJSON(payload)

	if err != nil {
		return errors.New(errors.BadRequestError, err.Error())
	}

	return nil
}
