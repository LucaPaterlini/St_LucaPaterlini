package main

import (
	"testing"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// Checks the connection to the API
func TestHello(t *testing.T) {
	// Setup
	_ , err := http.Get("http://"+ADDRESS+TEST_PORT)
	if err != nil {t.Error(fmt.Sprintf("No communication with the API err:%s ",err))}
}
// Checks the effects of a request with wrong parameters
func TestBidWrongParameters(t *testing.T){

	resp , err := http.Get("http://"+ADDRESS+TEST_PORT+"/bid?user_id=6&item_id=66&value=\"rhino\")")
	if err != nil {t.Error(fmt.Sprintf("No communication with the API err:%s ",err))}

	defer resp.Body.Close()
	body_bytearray,err := ioutil.ReadAll(resp.Body)
	if err == nil {
		if string(body_bytearray)!= "{error: true, message: \"Wrong Parameters\"}"{
			t.Error("There are wrong parameters in the bid request that have not triggered an errror")}
	}else{
		t.Error("Error while parsing the response")
	}
}
// Checks the effect of a lower than needed offer
func TestBidLowerOffer(t *testing.T){
	_ , err := http.Get("http://"+ADDRESS+TEST_PORT+"/bid?user_id=1&item_id=\"duck\"&value=20")
	if err != nil {t.Error(fmt.Sprintf("No communication with the API err:%s ",err))}

	resp , err := http.Get("http://"+ADDRESS+TEST_PORT+"/bid?user_id=1&item_id=\"duck\"&value=10")
	if err != nil {t.Error(fmt.Sprintf("No communication with the API err:%s ",err))}

	defer resp.Body.Close()

	body_bytearray,err := ioutil.ReadAll(resp.Body)

	body_string :=string(body_bytearray)
	if body_string!="{error: true, message: \"Your bid of 10 is too lower for \"duck\"\"}"{
		t.Error(fmt.Sprintf("Wrong response while inserting %s",body_string))
	}
}
// Checks the effect of a lower than needed offer
func TestBidBestOffer(t *testing.T){
	resp , err := http.Get("http://"+ADDRESS+TEST_PORT+"/bid?user_id=1&item_id=\"onix\"&value=10")
	if err != nil {t.Error(fmt.Sprintf("No communication with the API err:%s ",err))}

	defer resp.Body.Close()
	body_bytearray,err := ioutil.ReadAll(resp.Body)

	body_string :=string(body_bytearray)
	if body_string!="{error: false, message: \"Your bid of 10 is the best offer for \"onix\"\"}"{
		t.Error(fmt.Sprintf("Wrong response while inserting %s",body_string))
	}
}
// Checks the response of a request for the best bid of af an item that have more then one bid
func TestWinningAvailable(t *testing.T){
	resp , err := http.Get("http://"+ADDRESS+TEST_PORT+"/winning?item_id=\"duck\"")
	if err != nil {t.Error(fmt.Sprintf("No communication with the API err:%s ",err))}
	defer resp.Body.Close()
	body_bytearray,err := ioutil.ReadAll(resp.Body)

	body_string :=string(body_bytearray)
	if body_string!="{error: false, message: \"The current best bid for \"duck\" is 20\"}"{
		t.Error(fmt.Sprintf("Got the wrong value for the best offer err:%s",body_string))
	}
}
// Checks the response of a request for the best bid that have not received offers before
func TestWinningNotAvailable(t *testing.T){
	resp , err := http.Get("http://"+ADDRESS+TEST_PORT+"/winning?item_id=\"magickarp\"")
	if err != nil {t.Error(fmt.Sprintf("No communication with the API err:%s ",err))}
	defer resp.Body.Close()
	body_bytearray,err := ioutil.ReadAll(resp.Body)

	body_string :=string(body_bytearray)
	if body_string!="{error: false, message: \"There is no offer for \"magickarp\"\"}"{
		t.Error(fmt.Sprintf("Got a value instead of getting no offer err:%s",body_string))
	}
}


// Checks the response of for the request af all bets of a specific item
func TestListAvailable(t *testing.T){
	resp , err := http.Get("http://"+ADDRESS+TEST_PORT+"/list?item_id=\"duck\"")
	if err != nil {t.Error(fmt.Sprintf("No communication with the API err:%s ",err))}
	defer resp.Body.Close()

	body_bytearray,err := ioutil.ReadAll(resp.Body)
	body_string :=string(body_bytearray)

	if !strings.HasPrefix(body_string,"{error: false, message:\"[{\"User_id\":\"1\",\"Value\":20,"){
		t.Error(fmt.Sprintf("Not got the list err:%s",body_string))
	}
}
// Checks the response of for the request af all bets of a specific item that have received no bid
func TestListNotAvailable(t *testing.T){
	resp , err := http.Get("http://"+ADDRESS+TEST_PORT+"/list?item_id=\"caterpie\"")
	if err != nil {t.Error(fmt.Sprintf("No communication with the API err:%s ",err))}
	defer resp.Body.Close()

	body_bytearray,err := ioutil.ReadAll(resp.Body)
	body_string :=string(body_bytearray)

	if !strings.HasPrefix(body_string,"{error: false, message:\"[]\"}"){
		t.Error(fmt.Sprintf("Got a list instead of not getting it err:%s",body_string))
	}
}

// Checks the response to the request of the list of item to whom the user have made a bid
func TestUserBidsAvailable(t *testing.T){
	resp , err := http.Get("http://"+ADDRESS+TEST_PORT+"/user_bids?user_id=1")
	if err != nil {t.Error(fmt.Sprintf("No communication with the API err:%s ",err))}
	defer resp.Body.Close()

	body_bytearray,err := ioutil.ReadAll(resp.Body)
	body_string :=string(body_bytearray)

	if !strings.HasPrefix(body_string,`{error: false, message:"["\"duck\"","\"onix\""]"}`){
		t.Error(fmt.Sprintf("Not got the right list of item to whom the user have made a bid err:%s",body_string))
	}
}
// // Checks the response to the request of the list of item to whom the user have made a bib for a user that have made no bid
func TestUserBidsNotAvailable(t *testing.T){
	resp , err := http.Get("http://"+ADDRESS+TEST_PORT+"/user_bids?user_id=2")
	if err != nil {t.Error(fmt.Sprintf("No communication with the API err:%s ",err))}
	defer resp.Body.Close()

	body_bytearray,err := ioutil.ReadAll(resp.Body)
	body_string :=string(body_bytearray)

	if !strings.HasPrefix(body_string,"{error: false, message:\"[]\"}"){
		t.Error(fmt.Sprintf("Got a list instead of not getting it err:%s",body_string))
	}
}

