package main

import (
	"github.com/gin-gonic/gin"
)

type server struct {
	quoteController *quoteController
	engine          *gin.Engine
}

func NewServer(
	quoteController *quoteController,
) *server {
	return &server{
		quoteController: quoteController,
		engine:          gin.New(),
	}
}
