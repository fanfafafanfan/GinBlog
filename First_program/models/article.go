package models

import (
	"database/sql"
	"log"
)

type Article struct {
	ID             int    `form:"Id"`
	ArticleName    string `form:"Article_Name" binding:"required"`
	ArticleContent string `form:"Article_Content" binding:"required"`
	Author         string `form:"Author" binding:"required"`
}

var db *sql.DB

func (a *Article) Addarticle() (id int64, err error) {
	rs, err := db.Exec("INSERT INTO article(Article_Name, Article_Content, Author) VALUES (?, ?, ?)", a.ArticleName, a.ArticleContent, a.Author)
	if err != nil {
		return
	}
	id, err = rs.LastInsertId()
	return
}

func (a *Article) Getarticles() (articles []Article, err error) {
	articles = make([]Article, 0)
	rows, err := db.Query("SELECT Id, Article_Name, Article_Content, Author FROM article")
	defer rows.Close()

	if err != nil {
		return
	}

	for rows.Next() {
		var article Article
		rows.Scan(&article.ID, &article.ArticleName, &article.ArticleContent, &article.Author)
		articles = append(articles, article)
	}
	if err = rows.Err(); err != nil {
		return
	}
	return
}

func (a *Article) Getarticle() (article Article, err error) {
	err = db.QueryRow("SELECT Id, Article_Name, Article_Content, Author FROM article WHERE Id=?", a.ID).Scan(
		&article.ID, &article.ArticleName, &article.ArticleContent, &article.Author,
	)
	return
}

func (a *Article) Modarticle() (ra int64, err error) {
	stmt, err := db.Prepare("UPDATE article SET Article_Name=?, Article_Content=?, Author=? WHERE Id=?")
	defer stmt.Close()
	if err != nil {
		return
	}
	rs, err := stmt.Exec(a.ArticleName, a.ArticleContent, a.Author, a.ID)
	if err != nil {
		return
	}
	ra, err = rs.RowsAffected()
	return
}

func (a *Article) Delarticle() (ra int64, err error) {
	rs, err := db.Exec("DELETE FROM article WHERE Id=?", a.ID)
	if err != nil {
		log.Fatalln(err)
	}
	ra, err = rs.RowsAffected()
	return
}
