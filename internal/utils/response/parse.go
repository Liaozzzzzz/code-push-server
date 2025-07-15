package response

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/liaozzzzzz/code-push-server/internal/utils/errors"
)

// Parse body json data to struct
func ParseJSON(c *gin.Context, obj interface{}) *errors.BusinessError {
	if err := c.ShouldBindJSON(obj); err != nil {
		return errors.NewBusinessErrorf(errors.CodeInvalidParams, "Failed to parse json: %s", err.Error())
	}
	return nil
}

// Parse query parameter to struct
func ParseQuery(c *gin.Context, obj interface{}) *errors.BusinessError {
	if err := c.ShouldBindQuery(obj); err != nil {
		return errors.NewBusinessErrorf(errors.CodeInvalidParams, "Failed to parse query: %s", err.Error())
	}
	return nil
}

// Parse body form data to struct
func ParseForm(c *gin.Context, obj interface{}) *errors.BusinessError {
	if err := c.ShouldBindWith(obj, binding.Form); err != nil {
		return errors.NewBusinessErrorf(errors.CodeInvalidParams, "Failed to parse form: %s", err.Error())
	}
	return nil
}
