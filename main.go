package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"reflect"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	_ "github.com/proullon/ramsql/driver"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Menu_Product struct {
	_Id          primitive.ObjectID
	Id           primitive.ObjectID
	Name         string
	CampingTent  int
	CampingChair int
	Shirt        int
	Bag          int
	Ongcoffee    int
}

func Update_Product(c echo.Context) error {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	fmt.Println("clientOptions type:", reflect.TypeOf(clientOptions), "\n")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		fmt.Println("mongo.Connect() ERROR:", err)
		os.Exit(1)
	}
	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)
	col := client.Database("SSB-Store").Collection("Product")
	fmt.Println("Collection type:", reflect.TypeOf(col), "\n")
	fmt.Println("ctx", ctx)

	// Update data
	id, _ := primitive.ObjectIDFromHex("63ccf25668b2fee9d0d70b12")
	filter := bson.D{{"_id", id}}
	update := bson.D{{"$set", bson.D{{"campingtent", 100}}}}
	result, err := col.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println("result campingtent", result)

	update = bson.D{{"$set", bson.D{{"campingchair", 100}}}}
	result, err = col.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println("result campingchair", result)

	update = bson.D{{"$set", bson.D{{"shirt", 100}}}}
	result, err = col.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println("result campingchair", result)

	update = bson.D{{"$set", bson.D{{"bag", 100}}}}
	result, err = col.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println("result campingchair", result)

	update = bson.D{{"$set", bson.D{{"ongcoffee", 100}}}}
	result, err = col.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println("result campingchair", result)

	return c.JSON(http.StatusOK, "Update Success")
}

func Buy_Product_from_Web(c echo.Context) error {
	json_map := make(map[string]interface{})
	err := json.NewDecoder(c.Request().Body).Decode(&json_map)
	if err != nil {
		return err
	}
	fmt.Println("data", json_map)
	var buy_bag int = int(json_map["Bag"].(float64))
	var buy_campingchair int = int(json_map["Campingchair"].(float64))
	var buy_campingtent int = int(json_map["Campingtent"].(float64))
	var buy_ongcoffee int = int(json_map["Ongcoffee"].(float64))
	var buy_shirt int = int(json_map["Shirt"].(float64))

	fmt.Printf("Type %T %v\n", buy_bag, buy_bag)
	fmt.Printf("Type %T %v\n", buy_campingchair, buy_campingchair)
	fmt.Printf("Type %T %v\n", buy_campingtent, buy_campingtent)
	fmt.Printf("Type %T %v\n", buy_ongcoffee, buy_ongcoffee)
	fmt.Printf("Type %T %v\n", buy_shirt, buy_shirt)

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	fmt.Println("clientOptions type:", reflect.TypeOf(clientOptions), "\n")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		fmt.Println("mongo.Connect() ERROR:", err)
		os.Exit(1)
	}
	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)
	col := client.Database("SSB-Store").Collection("Product")
	fmt.Println("Collection type:", reflect.TypeOf(col), "\n")
	fmt.Println("ctx", ctx)

	// Get data
	var L Menu_Product
	output := []Menu_Product{}
	cursor, err := col.Find(context.TODO(), bson.D{})
	if err != nil {
		fmt.Println("Finding all documents ERROR:", err)
		defer cursor.Close(ctx)
	} else {
		for cursor.Next(ctx) {
			err := cursor.Decode(&L)
			fmt.Println("L", L)
			output = append(output, L)
			if err != nil {
				fmt.Println("cursor.Next() error:", err)
				os.Exit(1)
			}
		}
		fmt.Println("output", output)
	}

	// Update data
	id, _ := primitive.ObjectIDFromHex("63ccf25668b2fee9d0d70b12")
	filter := bson.D{{"_id", id}}

	update_bag := bson.D{{"$set", bson.D{{"bag", output[0].Bag - buy_bag}}}}
	result_bag, err := col.UpdateOne(context.TODO(), filter, update_bag)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println("result Buy Product bag", result_bag)

	update_campingchair := bson.D{{"$set", bson.D{{"campingchair", output[0].CampingChair - buy_campingchair}}}}
	result_campingchair, err := col.UpdateOne(context.TODO(), filter, update_campingchair)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println("result Buy Product campingchair", result_campingchair)

	update_campingtent := bson.D{{"$set", bson.D{{"campingtent", output[0].CampingTent - buy_campingtent}}}}
	result_campingtent, err := col.UpdateOne(context.TODO(), filter, update_campingtent)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println("result Buy Product campingtent", result_campingtent)

	update_ongcoffee := bson.D{{"$set", bson.D{{"ongcoffee", output[0].Ongcoffee - buy_ongcoffee}}}}
	result_ongcoffee, err := col.UpdateOne(context.TODO(), filter, update_ongcoffee)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println("result Buy Product ongcoffee", result_ongcoffee)

	update_shirt := bson.D{{"$set", bson.D{{"shirt", output[0].Shirt - buy_shirt}}}}
	result_shirt, err := col.UpdateOne(context.TODO(), filter, update_shirt)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println("result Buy Product ongcoffee", result_shirt)

	return c.JSON(http.StatusOK, result_campingtent)
}

func Post_Product(c echo.Context) error {
	d := &Menu_Product{}
	d.Name = "Big-Shop"
	d.CampingTent = 10
	d.CampingChair = 10
	d.Shirt = 10
	d.Bag = 10
	d.Ongcoffee = 10

	if err := c.Bind(d); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	fmt.Println("data", c.Bind(d))
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		fmt.Println("error:", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		fmt.Println("error:", err)
	}
	db := client.Database("SSB-Store")
	collection := db.Collection("Product")
	res, err := collection.InsertOne(context.Background(), d)

	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println("res", res.InsertedID.(primitive.ObjectID).Timestamp())
	return c.JSON(http.StatusOK, d)
}

func main() {

	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	e.Use(middleware.Logger())
	e.PATCH("update_product", Update_Product)
	e.PATCH("buy_product_from_web", Buy_Product_from_Web)
	e.POST("post_product", Post_Product)

	port := 2565
	log.Println("Starting ... port", port)
	log.Fatal(e.Start("localhost:2565"))
}
