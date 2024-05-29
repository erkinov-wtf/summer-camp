package interfaces

import "github.com/gin-gonic/gin"

type AppInterface interface {
	Create(context *gin.Context)
	GetById(context *gin.Context)
	Patch(context *gin.Context)
}
