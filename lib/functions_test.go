package lib

import (
	"testing"
)


func TestAll(t *testing.T) {
	// check add for no user available
	val, err :=Bid("random1","banana",12)
	if err!=nil {t.Error(err.Error())}
	if val!=0{t.Error("Wrong highest offer value, expected 0, got:",val)}
	// create users
	idLuca,err :=CreteUser("luca","paterlini",111111111)
	if err!=nil {t.Error(err.Error())}
	idMario,err :=CreteUser("mario","paganino",666)
	if err!=nil {t.Error(err.Error())}
	// start the bidding process
	val, err =Bid(idLuca,"banana",1)
	if err!=nil {t.Error(err.Error())}
	val, err =Bid(idLuca,"banana",11)
	if err==nil {t.Error("Expected and Overbid error")}
	val, err =Bid(idMario,"banana",1)
	if err==nil {t.Error("Expeted a low value bid error")}
	val, err =Bid(idMario,"banana",12)

	if err!=nil {t.Error(err.Error())}
	val, err =Bid(idLuca,"apple",666)

	if err!=nil {t.Error(err.Error())}
	//retrieving the winning
	item,err := Winning("banana")
	if err!=nil || item.Value!=12 {t.Error(err.Error())}

	// compare the list of the bid
	listValue := []int64{1,11,12}
	listBiders := []string{idLuca,idMario}
	arrayBids,err := AllBids("banana")
	if err!=nil {t.Error(err.Error())}
	for i,v := range arrayBids {
		if v.Value!=listValue[i] || v.UserId!=listBiders[i]{
			t.Error("ArrayBids: wrong item values",v)
		}
	}
	// checks the bids of a specific user
	listUserBidsItems :=[]string{"banana","apple"}
	listUserBidsValue :=[]int64{1,666}
	arrayBidsUser,err := UserBids(idLuca)
	if err!=nil {t.Error(err.Error())}
	for i,v := range arrayBidsUser {
		if v.Value!=listUserBidsValue[i] || v.ItemId!=listUserBidsItems[i]{
			t.Error("arrayBidsUser: wrong item values",v)
		}
	}
}

