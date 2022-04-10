package http

import (
	"go-backend-template/internal/util/errors"

	"github.com/gin-gonic/gin"
)

func bindBody(payload interface{}, c *gin.Context) error {
	err := c.BindJSON(payload)

	if err != nil {
		return errors.New(errors.BadRequestError, err.Error())
	}

	return nil
}
