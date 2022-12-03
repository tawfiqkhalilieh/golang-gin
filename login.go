package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func login(c *gin.Context){
  for _, usr := range db {
    if usr.email == c.Param("email") && usr.password == c.Param("password"){
      c.JSON(http.StatusOK, gin.H{ "authorized": true, "user": gin.H{ "id": usr.id, "username": usr.username, "email": usr.email, "first_name": usr.first_name  }})
      return
    }
  }
  c.JSON(422 , gin.H{ "authorized": false })
}
