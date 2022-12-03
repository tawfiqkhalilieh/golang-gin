package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// should be moved to routes/get_user_username
func get_user_username(c *gin.Context) {
  if c.Param("username") != "" {
    for _, usr := range db {
    if usr.username == c.Param("username") {
        c.JSON(http.StatusOK, gin.H{ "user": gin.H{ "id": usr.id, "username": usr.username, "email": usr.email, "first_name": usr.first_name, "last_name": usr.last_name, "age": usr.age, "password": usr.password  }})
        return
      }
    }
  } else {
    c.JSON(http.StatusNotFound, "")
  }
}

// should be moved to routes/get_user_id
func get_user_id(c *gin.Context) {
  if c.Param("id") != "" {
    for _, usr := range db {
      if usr.id == c.Param("id") {
          c.JSON(http.StatusOK, gin.H{ "user": Dictionary { "id": usr.id, "username": usr.username, "email": usr.email, "first_name": usr.first_name, "last_name": usr.last_name, "age": usr.age, "password": usr.password } })
          return
      }
    }
  } 
  c.JSON(http.StatusNotFound, "User not found")
}