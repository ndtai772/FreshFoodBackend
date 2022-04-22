package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (server *Server) getUser(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H {
		"id": 1,
		"username": "Nguyễn Văn A",
		"phone_number": "0123456789",
		"email": "nguyenvana@gmail.com",
		"address": "144 Xuân Thủy, Dịch Vọng, Cầu giấy, Hà Nội",
	})
}
