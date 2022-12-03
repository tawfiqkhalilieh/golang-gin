package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// should be moved to routes/change_password
func change_password(c *gin.Context){
  for index, usr := range db {
    if usr.id == c.Param("id") {
      if usr.password == c.Param("password") {
        usr.password = c.Param("newpassword")
        db[index] = usr;
        c.JSON(http.StatusOK, gin.H{ "status": "succsess" })
      } else {
        c.JSON(422, gin.H{ "status": "failed" })
      }
      return
    }
  }
  c.JSON(404, gin.H{ "status": "failed" })
}

// should be moved to routes/change_username
func change_username(c *gin.Context) {
  for index, usr := range db {
    if usr.id == c.Param("id") {
      if usr.password == c.Param("password") {
        usr.username = c.Param("newusername")
        db[index] = usr;
        c.JSON(http.StatusOK, gin.H{ "status": "succsess" })
      } else {
        c.JSON(422, gin.H{ "status": "failed" })
      }
      return
    }
  }
  c.JSON(404, gin.H{ "status": "failed" })
}

// should be moved to routes/change_email
func change_email(c *gin.Context){
  for index, usr := range db {
    if usr.id == c.Param("id") {
      if usr.password == c.Param("password") {
        usr.email = c.Param("newemail")
        db[index] = usr;
        c.JSON(http.StatusOK, gin.H{ "status": "succsess" })
      } else {
        c.JSON(422, gin.H{ "status": "failed" })
      }
      return
    }
  }
  c.JSON(404, gin.H{ "status": "failed" })
}