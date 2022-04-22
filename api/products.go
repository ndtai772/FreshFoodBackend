package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (server *Server) getProduct(ctx *gin.Context) {
	id, err := parseIdUri(ctx)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	product, err := server.store.GetProduct(ctx, id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, product)
}

func (server *Server) listProducts(ctx *gin.Context) {
	products, err := server.store.ListProducts(ctx)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	

	ctx.JSON(http.StatusOK, products)
} 