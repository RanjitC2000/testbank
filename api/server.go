package api

import (
	db "github.com/RanjitC2000/testbank/db/sqlc"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type Server struct {
	store  *db.Store
	router *gin.Engine
}

func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, "Hello World!")
	})
	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts", server.listAccounts)
	// router.POST("/transfers", server.createTransfer)
	server.router = router
	return server
}
