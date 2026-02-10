//  Licensed under the Apache License 2.0
//  @author Selman TUNÇ <selmantunc@gmail.com>
//  @link: https://github.com/stnc/go-mvc-blog-clean-code
//  @license: Apache License 2.0
package main

import (
	"io"
	"log"
	"net/http"
	"os"

	"stncCms/app/domain/repository"
	"stncCms/app/web/controller"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	//To load our environmental variables.

	if err := godotenv.Load(); err != nil {
		log.Println("no env gotten")
	}
}

func main() {


	db := repository.DbConnect()
	services, err := repository.RepositoriesInit(db)
	if err != nil {
		panic(err)
	}
	//defer services.Close()
	services.Automigrate()



	optionsHandle := controller.InitOptions(services.Options)



	r := gin.Default()

	r.Use(gin.Recovery())
	r.Use(gin.Logger())

	store := cookie.NewStore([]byte("SpeedyGonzales"))
	store.Options(sessions.Options{Path: "/", HttpOnly: true, MaxAge: 3600 * 8}) //Also set Secure: true if using SSL, you should though

	r.Use(sessions.Sessions("myCRM", store))



	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	r.Static("/assets", "./public/static")

	r.StaticFS("/upload", http.Dir("./public/upl"))

	r.GET("/", controller.Index)



	optionsGroup := r.Group("/admin/options")
	{
		optionsGroup.GET("/", optionsHandle.Index)
		// optionsGroup.POST("update", optionsHandle.Update)
		// optionsGroup.GET("makbuzNo", optionsHandle.MakbuzNo)
	}

	
	// Logging to a file.
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)
	//Starting the application
	appPort := os.Getenv("PORT")
	if appPort == "" {
		appPort = "8080" //localhost
	}
	log.Fatal(r.Run(":" + appPort))

}
