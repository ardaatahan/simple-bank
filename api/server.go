package api

import (
	"fmt"

	db "github.com/ardaatahan/simplebank/db/sqlc"
	"github.com/ardaatahan/simplebank/token"
	"github.com/ardaatahan/simplebank/util"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type Server struct {
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
	router     *gin.Engine
}

func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}
	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}
	router := gin.Default()

	if validator, ok := binding.Validator.Engine().(*validator.Validate); ok {
		validator.RegisterValidation("currency", isValidCurrency)
	}

	router.POST("/users", server.createUser)

	router.GET("/accounts", server.listAccounts)
	router.POST("/accounts", server.createAccount)

	router.GET("/accounts/:id", server.getAccount)
	router.PUT("/accounts/:id", server.updateAccount)
	router.DELETE("/accounts/:id", server.deleteAccount)

	router.POST("/transfers", server.createTransfer)

	server.router = router
	return server, nil
}

func (server *Server) RunServer(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
