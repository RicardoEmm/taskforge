package httputil

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func ParseUUID(c *gin.Context, value, fieldName string) (uuid.UUID, bool) {
	parsedID, err := uuid.Parse(value)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid " + fieldName})
		return uuid.Nil, false
	}

	return parsedID, true
}
