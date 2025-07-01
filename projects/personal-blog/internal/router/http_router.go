package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/selahaddinislamoglu/golang-studies/projects/personal-blog/internal/controller"
)

type PersonalBlogRouter struct {
	httpController controller.HTTPController
}

func NewPersonalBlogHTTPRouter(httpController controller.HTTPController) HTTPRouter {
	return &PersonalBlogRouter{
		httpController: httpController,
	}
}

func (r *PersonalBlogRouter) SetupRoutes() (http.Handler, error) {

	router := gin.Default()
	router.LoadHTMLGlob("internal/templates/*.html")
	router.GET("/", r.httpController.GuestHome)
	router.GET("/article/:id", r.httpController.GuestReadArticle)

	admin := router.Group("/admin", gin.BasicAuth(gin.Accounts{
		"admin": "password",
	}))
	admin.GET("/", r.httpController.AdminHome)
	admin.GET("/create", r.httpController.AdminCreateArticle)
	admin.GET("/edit/:id", r.httpController.AdminEditArticle)
	admin.POST("/create", r.httpController.AdminCreateArticleHandler)
	admin.POST("/edit/:id", r.httpController.AdminEditArticleHandler)
	admin.POST("/delete/:id", r.httpController.AdminDeleteArticleHandler)

	return router, nil
}
