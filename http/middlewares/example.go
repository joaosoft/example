package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func CheckExample(ctx *gin.Context) {
	// do something
	fmt.Println("check middleware")
	ctx.Next()
}
