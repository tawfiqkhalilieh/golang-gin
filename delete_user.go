package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// should be moved to routes/delete_user
func delete_user(c *gin.Context){
  for index, usr := range db {
    if usr.id == c.Param("id") {
      db = append(db[:index], db[index+1:]...)
      c.JSON(http.StatusOK, gin.H{ "status": "succsess" })
      return
    }
  }
  c.JSON(404, gin.H{ "status": "failed" })
}