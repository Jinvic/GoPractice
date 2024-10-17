package shared

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetArticleID(c *gin.Context) (uint, error) {
	articleID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return 0, err
	}
	return uint(articleID), nil
}
