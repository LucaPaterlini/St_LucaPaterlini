package schema

type BidRecord struct{
	UserId string
	Value int64
	Time int64
}

type User struct{
	Name string
	Surname string
	DateOfBirth int64
	Bids  []UserBidRecord
}

type UserBidRecord struct {
	Value int64
	Time int64
	ItemId string
}

