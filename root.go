package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func root(c *gin.Context) { 
	c.JSON(http.StatusOK, gin.H{ "response": "hello world" })
}