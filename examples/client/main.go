package main

import (
	"fmt"
	"log"

	"github.com/dgrr/http2"
	"github.com/valyala/fasthttp"
)

func main() {
	c := &fasthttp.HostClient{
		Addr:  "api.binance.com:443",
		IsTLS: true,
	}
	http2.ConfigureClient(c)

	req := fasthttp.AcquireRequest()
	res := fasthttp.AcquireResponse()

	req.Header.SetMethod("GET")
	// TODO: Use SetRequestURI
	req.URI().Update("https://api.binance.com/api/v3/exchangeInfo")

	for i := 0; i < 4; i++ {
		res.ResetBody()

		err := c.Do(req, res)
		if err != nil {
			log.Fatalln(err)
		}

		body := res.Body()

		fmt.Printf("%d: %d\n", res.Header.StatusCode(), len(body))
		res.Header.VisitAll(func(k, v []byte) {
			fmt.Printf("%s: %s\n", k, v)
		})
		fmt.Println("------------------------")
	}

	// fmt.Printf("%s\n", body)
}
