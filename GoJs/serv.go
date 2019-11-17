package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"net/http"
)

type Response struct {
	Status string `json:"status"`
}

type MessageFromJs struct {
	Message string `json:"message"`
}



func TestPost(c *gin.Context){
	var messageFromJs MessageFromJs
	var response Response
	response.Status = "ok"
	c.BindJSON(&messageFromJs)
	fmt.Println(messageFromJs)
	c.JSON(http.StatusOK, response)

}

func main() {
	//gin.SetMode(gin.ReleaseMode)
	gin.SetMode(gin.DebugMode)
	//gin.DefaultWriter = ioutil.Discard
	r := gin.Default()
	r.Use(cors.Default())
	r.POST("test/", TestPost)
	r.Run("127.0.0.1:8083")
}
