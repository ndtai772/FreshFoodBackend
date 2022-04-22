package api

import (
	"github.com/bwmarrin/snowflake"
	"github.com/gin-gonic/gin"
	db "github.com/ndtai772/hackathon/db/sqlc"
)

type Server struct {
	store       *db.Store
	router      *gin.Engine
	idGenerator *snowflake.Node
}

func NewServer(store *db.Store) *Server {
	idGenerator, err := snowflake.NewNode(1)
	if err != nil {
		panic(err)
	}
	server := &Server{store: store, idGenerator: idGenerator}
	server.setupRouter()
	return server
}

func (server *Server) setupRouter() {
	router := gin.Default()

	publicRoutes := router.Group("/")
	authRoutes := router.Group("/")

	// Upload
	authRoutes.POST("upload", server.uploadImage)

	// Products
	publicRoutes.GET("products/:id", server.getProduct)
	publicRoutes.GET("products/", server.listProducts)

	// Order
	publicRoutes.GET("orders", server.listOrders)
	publicRoutes.GET("orders/:id", server.getOrder)

	// Users
	publicRoutes.GET("users/:id", server.getUser)

	// Static files
	router.Static("/images", "./resources")

	server.router = router
}

// func unimplemented(msg string) func(ctx *gin.Context) {
// 	log.Println("Unimplemented handler function: " + msg)
// 	return func(ctx *gin.Context) {
// 		ctx.JSON(http.StatusNotImplemented, gin.H{
// 			"message": "unimplemented",
// 		})
// 	}
// }

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
