package main

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type server struct {
	quoteController *quoteController
	logger          *zap.SugaredLogger
	engine          *gin.Engine
}

func NewServer(
	quoteController *quoteController,
	logger *zap.SugaredLogger,
) *server {
	return &server{
		quoteController: quoteController,
		logger:          logger,
		engine:          gin.New(),
	}
}
