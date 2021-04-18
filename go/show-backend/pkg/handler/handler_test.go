package function

import (
	"encoding/json"
	"fmt"
	handler "github.com/openfaas/templates-sdk/go-http"
	"net/url"
	"testing"
)

func TestHandle(t *testing.T) {
	testcase := []handler.Request{
		{QueryString: "format=json"},
		{QueryString: "format=yaml"},
		{QueryString: ""},
		{},
	}
	for _, req := range testcase {
		res,err := Handle(req)
		if err!=nil{
			t.Fail()
		} else {
			fmt.Println(string(res.Body))
		}
	}
}

func TestQueryString(t *testing.T) {
	var message []byte

	vList,err := url.ParseQuery("format=json&format=yaml")
	if err != nil {
		fmt.Printf("Query string parse error!\n")
	}
	formatList, exists := vList["format"]
	if exists {
		format := formatList[0]
		if format=="json"{
			message, _ = json.Marshal(IPRes{Ip: "127.0.0.1"})
		}
	} else {
		message = []byte("127.0.0.1")
	}

	fmt.Printf("response: %s\n", string(message))
}
