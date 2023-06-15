package main

import (
	"personalprojects/Backend/Go-Programming/go-crud/controllers"
	"personalprojects/Backend/Go-Programming/go-crud/initializers"

	"github.com/gin-gonic/gin"
)


func init(){
	//calling initializers function
	initializers.LoadEnvVariables()
	initializers.ConnectionToDatabase()

}

func main(){
	r := gin.Default()
	r.POST("/posts", controllers.PostsCreate )
	r.GET("/all_posts", controllers.GetAllPosts)
	r.GET("/post/:id", controllers.GetSinglePost)
  r.PUT("/post/:id", controllers.PostUpdate)
  r.DELETE("/post/:id", controllers.PostDelete)

	r.Run()
 
}