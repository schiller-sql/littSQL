package helpers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetParamIDInt32(c *gin.Context) (int32, error) {
	return GetParamInt32(c, "id")
}

func GetParamInt32(c *gin.Context, paramName string) (int32, error) {
	databaseID, err := strconv.Atoi(c.Param(paramName))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "an id is required to access the correct resource"})
		return 0, err
	}
	return int32(databaseID), nil
}

func GetJwtID(c *gin.Context) int32 {
	id, _ := c.Get("id")
	return id.(int32)
}
