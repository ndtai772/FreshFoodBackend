package api

import (
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func (server *Server) uploadImage(ctx *gin.Context) {
	file, _ := ctx.FormFile("file")
	log.Println(file.Filename)

	dst := "./resources/"

	validSuffixes := [...]string{".png", ".jpg", ".jpeg"}

	for i := 0; i < len(validSuffixes); i++ {
		if strings.HasSuffix(file.Filename, validSuffixes[i]) {
			dst += server.idGenerator.Generate().String() + validSuffixes[i]

			ctx.SaveUploadedFile(file, dst)

			ctx.JSON(http.StatusOK, gin.H {
				"image_path": dst,
			})
			return
		}
	}

	ctx.JSON(http.StatusBadRequest, msgResponse("invalid image path, only support png, jpg, jpeg"))
}
