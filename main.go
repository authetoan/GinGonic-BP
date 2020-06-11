package main

import "github.com/gin-gonic/gin"
import "GoBP/routers"
func main() {
    r := gin.Default()
    routers.ApiRoutes(r)
    r.Run(":8080")
}
