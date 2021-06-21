package main

import (
	"io/ioutil"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/iot-sense/utility"
)

func setupRouter() *gin.Engine {
	router := gin.Default()
	router.Use(cors.Default())
	// router.StaticFS("/assets", http.Dir("./assets"))

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "root page")
	})

	router.GET("/tags/list", func(c *gin.Context) {
		data := []byte(`["tag1", "tag2", "tag3"]`)
		c.Data(http.StatusOK, "application/json", data)
	})

	router.GET("/metrics/list", func(c *gin.Context) {
		data := []byte(`[
			{
				"metrics": "patient_card",
				"group":  "nckuh"
			}
		]`)
		c.Data(http.StatusOK, "application/json", data)
	})

	router.GET("/metrics/table/:metricsName", func(c *gin.Context) {
		data := []byte(`[
			{
				"data": "name",
				"tag":  "tag_name"
			},
			{
				"data": "department",
				"tag":  "tag_department"
			}
		]`)
		c.Data(http.StatusOK, "application/json", data)
	})

	router.POST("/metrics/table/:metricsName", func(c *gin.Context) {
		data, _ := ioutil.ReadAll(c.Request.Body)
		c.Data(http.StatusOK, "application/json", data)
	})

	return router
}

func main() {
	utility.Logger.Info("<< START SERVER >>")

	// // connect to mongoDB
	// utility.G_MONGO_CLIENT = utility.NewMongoClient()
	// // connect to postgresDB
	// utility.G_PG_CLIENT = utility.NewPgClient()

	router := setupRouter()
	utility.Logger.Info("<< END SERVER >>")

	port := utility.G_CONFIGER.GetString("plugin.port")
	router.Run(port)
}
