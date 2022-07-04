package database

import (
	"context"
	"finalassignment/model"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const dbName = "mylibrary"
const colName = "books"

var collection *mongo.Collection

func Init() {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("MongoDB connection Success")

	collection = client.Database(dbName).Collection(colName)

	fmt.Println("Collection instance is ready")
}

func InsertBookData(c *gin.Context) {
	var books model.Library

	err := c.BindJSON(&books)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "JSON binding failed.",
		})
	}

	inserted, err := collection.InsertOne(context.Background(), books)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted book in library: ", inserted.InsertedID)

	c.JSON(http.StatusOK, gin.H{
		"message": "Added book to database",
	})

}

func GetAllBooks(c *gin.Context) {

	cursor, err := collection.Find(context.Background(), bson.M{})

	if err != nil {
		log.Fatal(err)
	}

	var results []bson.M

	if err = cursor.All(context.TODO(), &results); err != nil {
		log.Fatal(err)
	}

	for _, books := range results {

		fmt.Println(books)

	}

	c.JSON(http.StatusOK, results)
}

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}
