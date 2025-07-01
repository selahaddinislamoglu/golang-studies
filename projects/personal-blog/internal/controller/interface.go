package controller

import (
	"github.com/gin-gonic/gin"
)

type HTTPController interface {
	GuestHome(context *gin.Context)
	GuestReadArticle(context *gin.Context)
	AdminHome(context *gin.Context)
	AdminCreateArticle(context *gin.Context)
	AdminEditArticle(context *gin.Context)

	AdminCreateArticleHandler(context *gin.Context)
	AdminEditArticleHandler(context *gin.Context)
	AdminDeleteArticleHandler(context *gin.Context)
}
