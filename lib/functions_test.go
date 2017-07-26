package lib

import (
	"testing"
	"encoding/json"
)

// Checks the effects of the new bid
func TestBidNew(t *testing.T) {
	Bid("1","pony",5)
	if len_bid_dict["pony"]!=1 ||
		bid_dict["pony"][0].Value!=5 ||
		bid_dict["pony"][0].User_id!="1" {
		t.Error("Error inserting a new Bid\n"+
			"  len:",len_bid_dict["pony"],
			", value:",bid_dict["pony"][0].Value,
			", user_id:",bid_dict["pony"][0].User_id)}
}

// Checks the effects of a lower bid
func TestBidLower(t *testing.T) {
	Bid("1","pony",4)
	if len_bid_dict["pony"]!=1 ||
		bid_dict["pony"][0].Value!=5 ||
		bid_dict["pony"][0].User_id!="1" {
		t.Error("Error inserting a new Bid\n"+
			"  len:",len_bid_dict["pony"],
			", value:",bid_dict["pony"][0].Value,
			", user_id:",bid_dict["pony"][0].User_id)}
}

// Check the effects of the second bid
func TestBidSecondUserSameItemSame(t *testing.T) {
	Bid("1","pony",10)
	if bid_dict["pony"][len_bid_dict["pony"]-1].Value!=10 {
		t.Error("Error inserting a second Bid same user same item\n"+
			"  len:",len_bid_dict["pony"],
			", value:",bid_dict["pony"][len_bid_dict["pony"]-1].Value,
			", user_id:",bid_dict["pony"][len_bid_dict["pony"]-1].User_id)}
}

// Check the effect of the third bid made by a different user
func TestBidThirdUserDifferentItemSame(t *testing.T) {
	Bid("2","pony",20)
	if bid_dict["pony"][len_bid_dict["pony"]-1].Value!=20 ||
		bid_dict["pony"][len_bid_dict["pony"]-1].User_id!="2"{
		t.Error("Error inserting a third Bid different user same item\n"+
			"  len:",len_bid_dict["pony"],
			", value:",bid_dict["pony"][len_bid_dict["pony"]-1].Value,
			", user_id:",bid_dict["pony"][len_bid_dict["pony"]-1].User_id)}
}

// Check the Insert of an offer for another item
func TestBidItemDifferent(t *testing.T) {
	Bid("1","dog",3)
	if  len_bid_dict["dog"]!=1 ||
	  bid_dict["dog"][len_bid_dict["dog"]-1].Value!=3 ||
		bid_dict["dog"][len_bid_dict["dog"]-1].User_id!="1"{
		t.Error("Error inserting a new Bid for a new Item\n"+
			"  len:",len_bid_dict["pony"],
			", value:",bid_dict["pony"][len_bid_dict["pony"]-1].Value,
			", user_id:",bid_dict["pony"][len_bid_dict["pony"]-1].User_id)}
}

// Check the best Bid for a not existing item
func TestWinningItemNotExists(t *testing.T){
	if Winning("Cthulhu")!=-1{
		t.Error("Error inserting an item that doesn't exists\n"+
			"  len:",len_bid_dict["Cthulhu"],
			", value:",bid_dict["Cthulhu"][len_bid_dict["Cthulhu"]-1].Value,
			", user_id:",bid_dict["Cthulhu"][len_bid_dict["Cthulhu"]-1].User_id)}
}

// Check the best Bid for an existing item
func TestWinningItemExists(t *testing.T){
	if Winning("pony")!=20{
		t.Error("Error inserting an item that have to exist \n"+
			"  len:",len_bid_dict["pony"],
			", value:",bid_dict["pony"][len_bid_dict["pony"]-1].Value,
			", user_id:",bid_dict["pony"][len_bid_dict["pony"]-1].User_id)}
}

// Check the the previus bids for an existing item
func TestAllBids(t *testing.T){

	s:=All_bids("pony")
	var jsonBlob = []byte(s)
	var bids_array []Bid_item

	err := json.Unmarshal(jsonBlob, &bids_array)

	if err != nil {t.Error("Error while getting the pony's bids")}
	if !(bids_array[0].Value < bids_array[1].Value ||
		bids_array[1].Value  < bids_array[2].Value) {t.Error("Wrong bid's order")}
}

// Check the previous bids for a not existing item
func TestAllBids_NoItem(t *testing.T){
	s:=All_bids("crocodile")
	if s!="[]"{t.Error("Unexpected element inside the test of an empty bids list")}

}

// Check the list of items wich a user have made a bid
func TestUserBids(t *testing.T){
	s:=User_bids("1")
	comparison := `["pony","dog"]`
	if s=="null"  {t.Error("no Bids available for the user wich id is 1")}
	if s != comparison{t.Error("wrong list of items")}
}

// Check the list of items wich a user have made a bid for a user that not made any bid
func TestUserBids_NoItem(t *testing.T){
	s:=User_bids("666")
	if s!="[]"{t.Error("Unexpected element inside the list of items")}
}