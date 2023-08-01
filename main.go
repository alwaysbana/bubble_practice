package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/http"
)

var DB *gorm.DB

type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

func init() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/yuanshen?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn))
	DB = db

	if err != nil {
		panic("连接数据库失败,err%" + err.Error())
	}
}
func main() {
	//链接数据库
	r := gin.Default()
	//模型绑定
	DB.Debug().AutoMigrate(&Todo{})
	r.LoadHTMLGlob("templates/*")
	r.Static("/static", "static")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)

	})
	v1Group := r.Group("v1")
	{
		//添加
		v1Group.POST("/todo", func(c *gin.Context) {
			var todo Todo
			c.BindJSON(&todo)
			if err := DB.Create(&todo).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{"error": err.Error()})
			} else {
				c.JSON(http.StatusOK, todo)
			}

		})
		//修改
		v1Group.PUT("/todo", func(c *gin.Context) {

		})
		//删除
		v1Group.DELETE("/todo:id", func(c *gin.Context) {

		})
		//查看一个事项
		v1Group.GET("/todo:id", func(c *gin.Context) {

		})
		//查看所有事项
		v1Group.GET("/todo", func(c *gin.Context) {
			var todoList []Todo
			if err := DB.Find(&todoList).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{"error": err.Error()})
			} else {
				c.JSON(http.StatusOK, todoList)
			}

		})

	}
	r.Run()

}
