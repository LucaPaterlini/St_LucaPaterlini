package handlerEndpoints

import (
	"encoding/json"
	"errors"
	"lib"
	"schema"
)


func HandlerBid(dInput map[string]interface{})string{
	dict :=make(map[string]interface{})
	var err error
	if checkKeys(dInput,[]string{"userId","itemId","value"})==true{
		val,e :=lib.Bid(dict["userId"].(string),
			dict["itemId"].(string),dict["value"].(int64))
		err =  e
		dict =map[string]interface{}{"lastVal":val}
	}else{err=errors.New(schema.WPARERRMSG)}
	return composeJson(dict,err)
}

func HandlerWinning(dInput map[string]interface{})string{
	dict :=make(map[string]interface{})
	var err error
	if checkKeys(dInput,[]string{"userId"})==true{
		item,e :=lib.Winning(dict["userId"].(string))
		err =  e
		dict =map[string]interface{}{"item":item}
	}else{err=errors.New(schema.WPARERRMSG)}
	return composeJson(dict,err)
}

func HandlerAllBids(dInput map[string]interface{})string{
	dict :=make(map[string]interface{})
	var err error
	if checkKeys(dInput,[]string{"userId"})==true{
		items,e :=lib.AllBids(dict["userId"].(string))
		err =  e
		st,_:= json.Marshal(items)
		dict =map[string]interface{}{"item":st}
	}else{err=errors.New(schema.WPARERRMSG)}
	return composeJson(dict,err)
}

func HandlerUserBids(dInput map[string]interface{})string{
	dict :=make(map[string]interface{})
	var err error
	if checkKeys(dInput,[]string{"userId"})==true{
		items,e :=lib.UserBids(dict["userId"].(string))
		err =  e
		st,_ := json.Marshal(items)
		dict =map[string]interface{}{"items":st}
	}else{err=errors.New(schema.WPARERRMSG)}
	return composeJson(dict,err)
}

func HandlerCreateUser(dInput map[string]interface{})string{
	dict :=make(map[string]interface{})
	var err error
	if checkKeys(dInput,[]string{"name","surname","dob"})==true{
		s,e :=lib.CreteUser(dInput["name"].(string),
			dInput["surname"].(string),
			dInput["dob"].(int64))
		err =  e
		dict =map[string]interface{}{"userId":s}
	}else{err=errors.New(schema.WPARERRMSG)}
	return composeJson(dict,err)
}