package http

import "github.com/gin-gonic/gin"

type RouterSchema interface {
	RegisterSchema(router *gin.Engine)
}
