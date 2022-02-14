package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

func main() {
    r := gin.Default()
    r.GET("/ping", func(ctx *gin.Context) {
        ctx.JSON(http.StatusOK, gin.H{
            "message": "pong",
        })
    })
    panic(r.Run(":20001"))
}
