package models

import "gorm.io/gorm"

type article struct {
	gorm.Model `json:"-"`
	Title      string `json:"title"`
	Content    string `json:"content"`
	Author     string `json:"author"`
}

func CreateArticle(Title, Content, Author string) {
	db.Create(&article{Title: Title, Content: Content, Author: Author})
}

func GetArticleById(id int) article {
	var article article
	db.Where("id = ?", id).First(&article)
	return article
}

func GetArticleListByAuthor(author string) []article {
	var articles []article
	db.Where("author = ?", author).Find(&articles)
	return articles
}
func GetAllArticleList() []article {
	var articles []article
	db.Find(&articles)
	return articles
}

func GetArticleByTitle(title string) article {
	var article article
	db.Where("title = ?", title).First(&article)
	return article
}

func SreachArticle(keyword string) []article {
	var articles []article
	db.Where("title like ?", "%"+keyword+"%").Find(&articles)
	return articles
}
func SreachArticleByContent(keyword string) []article {
	var articles []article
	db.Where("content like ?", "%"+keyword+"%").Find(&articles)
	return articles
}
