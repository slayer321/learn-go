package routes

import (
	"finalassignment/database"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {

	router := gin.Default()

	router.GET("/allbooks", database.GetAllBooks)
	router.POST("/addbooks", database.InsertBookData)

	return router

}
