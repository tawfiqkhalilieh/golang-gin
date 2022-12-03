package main

import (
	"github.com/gin-gonic/gin"
)

type Dictionary map[string]interface{}

// should be moved to models/user
type User struct {
  id string
  username string
  email string
  first_name string
  last_name string
  age  int
  password string
}

// fake users database
var db = []User { }

func main() {
  r := gin.Default()
  r.GET("/", root);
  r.GET("/get_user/username/:username", get_user_username)
  r.GET("/get_user/id/:id", get_user_id)
  r.GET("/all", get_all_users)
  r.POST("/login/:email/:password", login)
  r.POST("/signup/:id/:password", signup)
  r.PATCH("/change/password/:id/:password/:newpassword", change_password)
  r.PATCH("/change/username/:id/:password/:newusername", change_username)
  r.PATCH("/change/email/:id/:password/:newemail", change_email)
  r.DELETE("/delete/user/:id", delete_user)
  r.DELETE("/all/", delete_all_users)
  r.Run()
}