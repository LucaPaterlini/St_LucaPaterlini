package main

import (
	"fmt"
	"github.com/asaskevich/govalidator"
	"github.com/valyala/fasthttp"
	"log"
	"strconv"
	"strings"
)

func logPanics(function fasthttp.RequestHandler)fasthttp.RequestHandler{
	return func(ctx *fasthttp.RequestCtx) {
		defer func() {
			if x := recover(); x!=nil{
				log.Printf("[%v] caught panic: %v",ctx.RemoteAddr(),x)
			}
		}()
		function(ctx)
	}
}

func middlewareEndpoint(ctx *fasthttp.RequestCtx,f func(map[string]interface{})string){
	d := parseQueryStringToDict(ctx.QueryArgs().String())
	response:= f(d)
	_,err :=fmt.Fprint(ctx, fmt.Sprintf("%v", response))
	if err== nil {ctx.SetStatusCode(fasthttp.StatusOK)
	} else {ctx.SetStatusCode(fasthttp.StatusInternalServerError)}
}

// Parsing the query to a dictionary
func parseQueryStringToDict(a string) map[string]interface{}{
	d := make(map[string]interface{})
	if len(a)<1 {return d}
	for _,t := range  strings.Split(a,"&"){
		g:= strings.Split(t,"=")
		if len(g)<2{continue}
		k,v :=g[0],g[1]
		if govalidator.IsInt(v){
			d[k], _ =strconv.ParseInt(v,10,64)
		}else if govalidator.IsFloat(v){
			d[k], _ =strconv.ParseFloat(v,64)
		}else {
			d[k]=v
		}
	}
	return d
}