package main

import (

    "net/http"    // this is to create methodes and requests
    "github.com/gin-gonic/gin" // this is to import gin and work with it

)

func main(){

    r := gin.Default() // default gin router
    
    // a simple get methode :
    r.GET("/ping", func(c *gin.Context) {
    // Return JSON response
    c.JSON(http.StatusOK, gin.H{
      "message": "pong",
    })
  })

    r.run() // gin router start
    // to start and run gin router we need to have atleast one methode in running in it.
    }
