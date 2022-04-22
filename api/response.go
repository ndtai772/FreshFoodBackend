package api

import "github.com/gin-gonic/gin"

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

func msgResponse(msg string) gin.H {
	return gin.H{"msg": msg}
}
