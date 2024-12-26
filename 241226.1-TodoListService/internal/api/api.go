package api

import (
	"net/http"
	"todolist/pkg/e"

	"github.com/gin-gonic/gin"
)

func Success(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": 1,
	})
}

func SuccessData(c *gin.Context, Data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"status": 1,
		"data":   Data,
	})
}

func FailError(c *gin.Context, err error) {
	if err2, ok := err.(e.Err); ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": 0,
			"code":   err2.Code,
			"errmsg": err2.Msg,
		})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": 0,
			"errmsg": err.Error(),
		})

	}
}
