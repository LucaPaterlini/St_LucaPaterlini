package lib

import ("time";"sync"
	"encoding/json"
)


type Bid_item struct{
	User_id string
	Value int64
	Time int64
}

var bid_dict = map[string][]Bid_item{}
var user_dict = map[string]map[string]bool{}

var len_bid_dict= map[string]int64{}
var len_user_dict= map[string]int64{}
var critical_section sync.Mutex



// Record a user's bid on an item, each new bid must be at a higher price than before
func Bid(user_id, item_id string,  value int64)int64{

	if bid_dict[item_id]!= nil && bid_dict[item_id][len_bid_dict[item_id]-1].Value>=value{ return -1}

	tmp_item :=Bid_item{user_id,value,int64(time.Now().UnixNano())}

	critical_section.Lock() // critical section start

	bid_dict[item_id] = append(bid_dict[item_id], tmp_item)
	len_bid_dict[item_id]+=1

	if len_user_dict[user_id]>0{
		user_dict[user_id][item_id] =  true
	} else{
		user_dict[user_id] = map[string]bool{item_id:true}
	}
	len_user_dict[user_id]+=1

	critical_section.Unlock() // critical section ends
	return len_bid_dict[item_id]
}

// Get the current winning bid for an item
func Winning(item_id string)int64{
	if len_bid_dict[item_id] !=0 {
		return bid_dict[item_id][len_bid_dict[item_id]-1].Value
	}
	return -1
}
// Get all the bids for an item
func All_bids(item_id string)string{
	if len_bid_dict[item_id] !=0 {
		json_return, _ := json.Marshal(bid_dict[item_id])
		return string(json_return)
	}
	return "[]"
}

//Get all the items on which a user has bid
func User_bids(user_id string)string{
	if len_user_dict[user_id] == 0 {return "[]"}
	keys := []string{}
	for k := range user_dict[user_id] {keys = append(keys, k)}
	json_return, _ := json.Marshal(keys)
	return string(json_return)
}
