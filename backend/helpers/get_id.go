package helpers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetParamID(c *gin.Context) (int32, error) {
	databaseID, err := strconv.Atoi(c.Param("id"))
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
