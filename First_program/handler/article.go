package handler

import (
	. "First_program/models"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func IndexAPI(c *gin.Context) {
	c.String(http.StatusOK, "It works")
}

func AddarticleAPI(c *gin.Context) {
	req := Article{}
	if err := c.ShouldBindWith(&req, binding.Form); err != nil {
		c.JSON(http.StatusOK, gin.H{"errno": "-1", "errmsg": "参数不匹配，请重试"})
		return
	}
	ra, err := req.Addarticle()
	if err != nil {
		log.Fatalln(err)
	}
	msg := fmt.Sprintf("新增成功!", ra)
	c.JSON(http.StatusOK, gin.H{
		"msg": msg,
	})
}

func GetarticlesAPI(c *gin.Context) {
	var a Article
	articles, err := a.Getarticles()
	if err != nil {
		log.Fatalln(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"msg":      "success",
		"articles": articles,
	})

}

func GetarticleAPI(c *gin.Context) {
	cid := c.Param("id")
	id, err := strconv.Atoi(cid)
	if err != nil {
		log.Fatalln(err)
	}
	a := Article{ID: id}
	article, err := a.Getarticle()
	if err != nil {
		log.Fatalln(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"msg":     "success",
		"article": article,
	})

}

func ModarticleAPI(c *gin.Context) {
	cid := c.Param("id")
	id, err := strconv.Atoi(cid)
	if err != nil {
		log.Fatalln(err)
	}
	a := Article{ID: id}
	err = c.Bind(&a)
	if err != nil {
		log.Fatalln(err)
	}
	ra, err := a.Modarticle()
	if err != nil {
		log.Fatalln(err)
	}
	msg := fmt.Sprintf("更新成功!", ra)
	c.JSON(http.StatusOK, gin.H{
		"msg": msg,
	})
}

func DelarticleAPI(c *gin.Context) {
	cid := c.Param("id")
	id, err := strconv.Atoi(cid)
	if err != nil {
		log.Fatalln(err)
	}
	a := Article{ID: id}
	ra, err := a.Delarticle()
	if err != nil {
		log.Fatalln(err)
	}
	msg := fmt.Sprintf("删除成功!", ra)
	c.JSON(http.StatusOK, gin.H{
		"msg": msg,
	})
}
