package main

import (
	"fmt"

	httptypes "github.com/rajatjindal/wasip2-golang/fermyon/spin/http-types"
	inboundhttp "github.com/rajatjindal/wasip2-golang/fermyon/spin/inbound-http"
	keyvalue "github.com/rajatjindal/wasip2-golang/fermyon/spin/key-value"
	"github.com/ydnar/wasm-tools-go/cm"
)

func init() {
	_ = inboundhttp.HandleRequest(httptypes.Request{})
	opt := keyvalue.StoreOpen("default")
	if opt.IsErr() {
		panic(opt.Err())
	}

	store := opt.OK()
	d := []byte("aa")
	store.Set("hello", cm.ToList(d))
	res := store.Get("hello")
	if res.IsErr() {
		panic(res.Err())
	}

	fmt.Println(string(res.OK().Some().Slice()))
}

func main() {

}
