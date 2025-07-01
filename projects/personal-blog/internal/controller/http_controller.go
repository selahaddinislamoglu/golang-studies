package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/selahaddinislamoglu/golang-studies/projects/personal-blog/internal/model"
	"github.com/selahaddinislamoglu/golang-studies/projects/personal-blog/internal/service"
)

type PersonalBlogHTTPController struct {
	service service.Service
}

func NewPersonalBlogHTTPController(service service.Service) HTTPController {
	return &PersonalBlogHTTPController{
		service: service,
	}
}

func (c *PersonalBlogHTTPController) GuestHome(context *gin.Context) {

	articles, err := c.service.GetAllArticlesSummary()
	if err != nil {
		fmt.Println(err)
		return
	}

	context.HTML(200, "home.html", gin.H{
		"Title":    "Personal Blog",
		"Articles": articles,
	})
}
func (c *PersonalBlogHTTPController) GuestReadArticle(context *gin.Context) {
	id := context.Param("id")
	if id == "" {
		context.String(400, "Article ID is required")
		return
	}

	var articleID int64
	_, err := fmt.Sscanf(id, "%d", &articleID)
	if err != nil {
		context.String(400, "Invalid Article ID format")
		return
	}

	article, err := c.service.GetArticleByID(articleID)
	if err != nil {
		context.String(500, "Error retrieving article: %v", err)
		return
	}

	context.HTML(200, "article.html", article)

}
func (c *PersonalBlogHTTPController) AdminHome(context *gin.Context) {

	articles, err := c.service.GetAllArticlesSummary()
	if err != nil {
		fmt.Println(err)
		return
	}

	context.HTML(200, "admin.html", gin.H{
		"Title":    "Admin Dashboard",
		"Articles": articles,
	})
}
func (c *PersonalBlogHTTPController) AdminCreateArticle(context *gin.Context) {
	context.HTML(200, "create.html", gin.H{
		"Title": "Create New Article",
	})
}
func (c *PersonalBlogHTTPController) AdminEditArticle(context *gin.Context) {
	id := context.Param("id")
	if id == "" {
		context.String(400, "Article ID is required")
		return
	}

	var articleID int64
	_, err := fmt.Sscanf(id, "%d", &articleID)
	if err != nil {
		context.String(400, "Invalid Article ID format")
		return
	}
	article, err := c.service.GetArticleByID(articleID)
	if err != nil {
		context.String(500, "Error retrieving article: %v", err)
		return
	}

	context.HTML(200, "edit.html", gin.H{
		"Title":   "Edit Article",
		"Article": article,
	})
}
func (c *PersonalBlogHTTPController) AdminCreateArticleHandler(context *gin.Context) {

	title := context.PostForm("title")
	content := context.PostForm("content")

	if title == "" || content == "" {
		context.String(400, "Title and content are required")
		return
	}

	article := &model.Article{
		ArticleSummary: model.ArticleSummary{
			Title: title,
		},
		Content: content,
	}

	id, err := c.service.CreateArticle(article)
	if err != nil {
		context.String(500, "Error creating article: %v", err)
		return
	}

	context.Redirect(302, fmt.Sprintf("/article/%d", id))
}
func (c *PersonalBlogHTTPController) AdminEditArticleHandler(context *gin.Context) {

	id := context.Param("id")
	if id == "" {
		context.String(400, "Article ID is required")
		return
	}

	var articleID int64
	_, err := fmt.Sscanf(id, "%d", &articleID)
	if err != nil {
		context.String(400, "Invalid Article ID format")
		return
	}

	title := context.PostForm("title")
	content := context.PostForm("content")

	if title == "" || content == "" {
		context.String(400, "Title and content are required")
		return
	}
	article := &model.Article{
		ArticleSummary: model.ArticleSummary{
			ID:    articleID,
			Title: title,
		},
		Content: content,
	}

	err = c.service.UpdateArticle(article)
	if err != nil {
		context.String(500, "Error updating article: %v", err)
		return
	}

	context.Redirect(302, fmt.Sprintf("/article/%d", articleID))

}
func (c *PersonalBlogHTTPController) AdminDeleteArticleHandler(context *gin.Context) {

	id := context.Param("id")
	if id == "" {
		context.String(400, "Article ID is required")
		return
	}

	var articleID int64
	_, err := fmt.Sscanf(id, "%d", &articleID)
	if err != nil {
		context.String(400, "Invalid Article ID format")
		return
	}

	err = c.service.DeleteArticle(articleID)
	if err != nil {
		context.String(500, "Error deleting article: %v", err)
		return
	}

	context.Redirect(302, "/admin")
}
