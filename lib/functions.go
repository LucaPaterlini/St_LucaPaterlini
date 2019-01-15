package lib

import (
	"../schema"
	"crypto/sha256"
	"errors"
	"fmt"
	"time"
)



var bidDict = map[string][]schema.BidRecord{}
var userDict = map[string]*schema.User{}
var sema = make(chan int,1)
// Record a user's bid on an item, each new bid must be at a higher price than before
func Bid(userId, itemId string,  value int64)(result int64,err error){
	// checking the existence of the user
	if _, ok := userDict[userId]; !ok {err = errors.New("user not found");return }
	if _, ok := bidDict[itemId]; !ok {err = errors.New("item not found");return }

	lastVal := bidDict[itemId][len(bidDict[itemId])-1].Value
	// lock
	sema<- 1

	if lastVal>= value {err = errors.New("the offer is to low for "+itemId+ " require at least: "+string(lastVal));return}
	if bidDict[itemId][len(bidDict[itemId])-1].UserId==userId{
		err = errors.New("this bid overbid the bid of the same user")
		result = lastVal
		return
	}
	// add the bid
	t := time.Now().UnixNano()
	newBidtmp := schema.UserBidRecord{Value:value,Time:t,ItemId:itemId}
	item := schema.BidRecord{UserId:userId,Value:value,Time:t}
	bidDict[itemId]= append(bidDict[itemId],item)
	userDict[userId].Bids =append(userDict[userId].Bids,newBidtmp)
	result= bidDict[itemId][len(bidDict[itemId])-1].Value
	// unlock
	<-sema
	return
}

// Get the current winning bid for an item
func Winning(itemId string)(item schema.BidRecord,err error){
	if val,ok:=bidDict[itemId];ok{
		item = val[len(bidDict[itemId])-1]
	}else {
		err = errors.New("item not found")
	}
	return
}

// Get all the bids for an item
func AllBids(itemId string)(array []schema.BidRecord,err error){
	if val,ok:=bidDict[itemId];ok{
		array = val
	}else {
		err = errors.New("item not found")
	}
	return
}


//Get all the items on which a user has bid
func UserBids(userId string)(userItems[]schema.UserBidRecord,err error){
	if user, ok := userDict[userId]; ok {
		userItems = user.Bids
	}else {
	err = errors.New("user not found")
	}
	return
}

// create a user
func CreteUser(name,surname string,dob int64)(userId string, err error){
	h := sha256.New()
	if _,ok:=userDict[userId];ok{
		err = errors.New("user already inserted")
		return
	}
	h.Write([]byte(name+surname+string(dob)))
	userId =fmt.Sprintf("%x", h.Sum(nil))
	userDict[userId]=&schema.User{Name:name,Surname:surname,DateOfBirth:dob,Bids:nil}
	return
}
