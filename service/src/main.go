package main

import (
    "encoding/json"
    "log"
    "net/http"
    "myapp/src/parser"
    "github.com/gin-gonic/gin"
)

type ParseRequest struct {
    SourceCode string `json:"source_code"`
}

type ParseResponse struct {
    Variables []parser.Variable `json:"variables"`
}

func main() {
    router := gin.Default()

    router.POST("/parse", func(c *gin.Context) {
        var request ParseRequest
        if err := c.ShouldBindJSON(&request); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        variables := parser.Parse(request.SourceCode)

        response := ParseResponse{
            Variables: variables,
        }

        c.JSON(http.StatusOK, response)
    })

    log.Fatal(router.Run(":8080"))
}
