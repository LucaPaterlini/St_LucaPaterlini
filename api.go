package main

import (
	"./lib"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
	"strconv"
	"reflect"
	"fmt"
)

func hello(c echo.Context)error{
	return c.String(http.StatusOK,
		"{error: false, message: \"Welcome to the Stratagem Demo Bidding System made by Luca Paterlini\"}")
}

const ADDRESS = "127.0.0.1"
const PORT =":8180"
const TEST_PORT = ":8081"

func bid_handler(c echo.Context)error{
	user_id := c.QueryParam("user_id")
	item_id := c.QueryParam("item_id")
	var value int64
	if reflect.TypeOf(user_id).Kind()==reflect.String {
		value, _ =strconv.ParseInt(c.QueryParam("value"),10,64)
	}

	if reflect.TypeOf(user_id).Kind()!=reflect.String ||
		reflect.TypeOf(item_id).Kind()!=reflect.String ||
		value==0 {
			return c.String(http.StatusBadRequest, "{error: true, message: \"Wrong Parameters\"}")
	}


	response := lib.Bid(user_id,item_id,value)

	if response==-1{
		return c.String(http.StatusBadRequest,
			fmt.Sprintf(
				"{error: true, message: \"Your bid of %v is too lower for %s\"}",
				value,item_id))}

	return c.String(http.StatusOK, fmt.Sprintf(
			"{error: false, message: \"Your bid of %v is the best offer for %s\"}",
			value,item_id))
}

func winning_handler(c echo.Context)error{
	item_id := c.QueryParam("item_id")
	best_bid:=lib.Winning(item_id)
	if best_bid>0{
		return c.String(http.StatusOK,
			fmt.Sprintf(
				"{error: false, message: \"The current best bid for %s is %v\"}",item_id,best_bid))
	}
	return c.String(http.StatusOK,
		fmt.Sprintf("{error: false, message: \"There is no offer for %s\"}",item_id))
}

func all_bids_handler(c echo.Context)error{
	response:= "{error: false, message:\""+lib.All_bids(c.QueryParam("item_id"))+"\"}"
	return c.String(http.StatusOK,response)
}

func user_bids_handler(c echo.Context)error{
	response:= "{error: false, message:\""+lib.User_bids(c.QueryParam("user_id"))+"\"}"
	return c.String(http.StatusOK,response)
}



func main(){
	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(middleware.CORS()) /* This allow CORS */

	e.GET("/",hello)
	e.GET("/bid",bid_handler)
	e.GET("/winning",winning_handler)
	e.GET("/list",all_bids_handler)
	e.GET( "/user_bids", user_bids_handler)


	e.POST("/",hello)
	e.POST("/bid",bid_handler)
	e.POST("/winning",winning_handler)
	e.POST("/list",all_bids_handler)
	e.POST( "/user_bids", user_bids_handler)

	e.Logger.Fatal(e.Start(ADDRESS+TEST_PORT))
}
