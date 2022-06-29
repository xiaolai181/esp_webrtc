package article

import (
	"esp_webrtc/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Article struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

func CreateArticle(c *gin.Context) {
	var article Article
	c.BindJSON(&article)
	if article.Title == "" || article.Content == "" || article.Author == "" {
		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{"error": "title, content, author is empty"})
		return
	}
	models.CreateArticle(article.Title, article.Content, article.Author)
}

func GetArticleListByAuthor(c *gin.Context) {
	author := c.Param("author")
	if author == "" {
		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{"error": "author is empty"})
		return
	}
	articles := models.GetArticleListByAuthor(author)
	c.JSON(http.StatusOK, gin.H{"articles": articles})
}

func GetAllArticleList(c *gin.Context) {
	articles := models.GetAllArticleList()
	c.JSON(http.StatusOK, gin.H{"articles": articles})
}
func GetArticleByTitle(c *gin.Context) {
	title := c.Param("title")
	if title == "" {
		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{"error": "title is empty"})
		return
	}
	article := models.GetArticleByTitle(title)
	c.JSON(http.StatusOK, gin.H{"article": article})
}

func SreachArticle(c *gin.Context) {
	keyword := c.Param("keyword")
	if keyword == "" {
		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{"error": "keyword is empty"})
		return
	}
	articles := models.SreachArticle(keyword)
	c.JSON(http.StatusOK, gin.H{"articles": articles})
}
func SreachArticleByContent(c *gin.Context) {
	keyword := c.Param("keyword")
	if keyword == "" {
		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{"error": "keyword is empty"})
		return
	}
	articles := models.SreachArticleByContent(keyword)
	c.JSON(http.StatusOK, gin.H{"articles": articles})
}

func GetArticleById(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{"error": "id is empty"})
		return
	}
	ArticleId, err := strconv.Atoi(id)
	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{"error": err.Error()})
		return
	}

	article := models.GetArticleById(ArticleId)
	c.JSON(http.StatusOK, gin.H{"article": article})
}
