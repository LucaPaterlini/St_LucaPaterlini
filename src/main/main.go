package main

import (
	"flag"
	cors "github.com/AdhityaRamadhanus/fasthttpcors"
	"github.com/valyala/fasthttp"
	"handlerEndpoints"
	"log"
	"schema"
	"strconv"
)

var (
	addr = flag.String("addr",schema.IPADDR+":"+strconv.Itoa(schema.PORT),"TCP address to listen to")
	compress = flag.Bool("compress",false, "Whether to enable transparent response compression ")
)



func main(){

	// Corse Handeler
	withCors := cors.DefaultHandler()

	flag.Parse()
	h := withCors.CorsMiddleware(logPanics(routingHandler))
	if *compress {
		h = fasthttp.CompressHandler(h)
	}
	if err := fasthttp.ListenAndServe(*addr,h);err != nil{
		log.Fatalf("Errror in ListenAndServer: %s",err)
	}

}



func routingHandler (ctx *fasthttp.RequestCtx) {
	ctx.SetContentType("text/json; charset=utf-8")

	switch string(ctx.Path()) {

	//Call n°0
	case "/bid":
		middlewareEndpoint(ctx, handlerEndpoints.HandlerBid)
	// Call n°1
	case "/winning":
		middlewareEndpoint(ctx, handlerEndpoints.HandlerWinning)
	// Call n°2
	case "/allbids":
		middlewareEndpoint(ctx, handlerEndpoints.HandlerAllBids)
	// Call n°3
	case "/userbids":
		middlewareEndpoint(ctx, handlerEndpoints.HandlerUserBids)
	// Call n°4
	case "/createuser":
		middlewareEndpoint(ctx, handlerEndpoints.HandlerCreateUser)

	default:
		ctx.Error("{\"error\":\"Unsupported Path\"}",fasthttp.StatusNotFound)
	}
}