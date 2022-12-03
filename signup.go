package main

import (
	"io"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func signup(c *gin.Context){
  bodyAsByteArray, _ := io.ReadAll(c.Request.Body)
  dataarr := strings.Split(string(bodyAsByteArray), "&")

  firstrow := strings.Split(dataarr[0], "=")
  secondrow := strings.Split(dataarr[1], "=")
  thirdrow := strings.Split(dataarr[2], "=")
  fourthrow := strings.Split(dataarr[3], "=")
  fifthrow := strings.Split(dataarr[4], "=")
  sixthrow := strings.Split(dataarr[5], "=")

  var _id string = uuid.New().String();
  age,_ := strconv.Atoi(sixthrow[1])

  usr := User{
    id: _id,
    username: firstrow[1],
    email: secondrow[1],
    first_name: thirdrow[1],
    last_name: fourthrow[1],
    password: fifthrow[1],
    age: age,
  }

  db = append(db, usr);
  c.JSON(201, gin.H{"data": gin.H{ "id": _id, firstrow[0]: firstrow[1], secondrow[0]: secondrow[1], thirdrow[0]: thirdrow[1], fourthrow[0]: fourthrow[1], fifthrow[0]: fifthrow[1], sixthrow[0]: sixthrow[1] }} )
}