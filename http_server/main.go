package main

import (
	"fmt"
	"net/http"

	"github.com/shh26b/trying-go/http_server/utils"
)

func main() {
	var env = utils.Env()

	handler := &RootHandler{}
	addr := fmt.Sprintf(":%v", env.Port)

	err := http.ListenAndServe(addr, handler)
	utils.Fatal(err)
}
