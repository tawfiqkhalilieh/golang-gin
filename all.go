package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// should be moved to routes/get_all_users
func get_all_users(c *gin.Context){
  data := []Dictionary{}
  for _, usr := range db {
    data = append(data, Dictionary{ "id": usr.id, "username": usr.username, "email": usr.email, "first_name": usr.first_name, "last_name": usr.last_name, "age": usr.age, "password": usr.password } )
  }
  c.JSON(http.StatusOK , gin.H{ "response": data });
}


// should be moved to routes/delete_all_users
func delete_all_users(c *gin.Context) {
  db = []User {}
  c.JSON(204, gin.H{ "status": "succsess" })
  return
}