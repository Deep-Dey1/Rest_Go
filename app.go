package main

import (

    "net/http"
//    "encoding/json"
    "github.com/gin-gonic/gin"

)

func main() {
    r := gin.Default()

    r.GET("/ping", func(c *gin.Context) {
    
    c.JSON(http.StatusOK, gin.H{

        "Web": "Its a challange You cant beat me : )",

    })
  
  })

    r.Run()
}

