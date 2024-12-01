package main

import "net/http"

type groceryListItem struct {
	Item  string `json:"item"`
	Store string `json:"store"`
}

var groceries = []groceryListItem{
	{Item: "eggs", Store: "Sams"},
	{Item: "milk", Store: "Lowes"},
}

func main() {
	router := gin.Default()
	router.GET("/items", getGroceryListItems)
	router.Run("localhost:8080")
}

func getGroceryListItems(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, groceries)
}
