package function

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/url"
	"strings"

	handler "github.com/openfaas/templates-sdk/go-http"
)

type IPRes struct {
	Ip string `json:"ip"`
}

// Handle a function invocation
func Handle(req handler.Request) (handler.Response, error) {
	var err error
	var message string

	ip := GetPodIP()
	if ip == "" {
		message = "GetPodIP Error"
	} else {
		message = fmt.Sprintf("Content from backend: %s", ip)
	}

	// json output when query string = "json"
	queryString := req.QueryString
	if queryString != "" {
		valueMap, err := url.ParseQuery(queryString)
		if err != nil {
			fmt.Printf("Query string parse error!\n")
			goto OUT
		}
		formatList, exists := valueMap["format"]
		if exists {
			format := formatList[0]
			if format == "json" {
				messageBytes, err := json.Marshal(IPRes{Ip: ip})
				if err != nil {
					fmt.Printf("Json Marshal error!\n")
				}
				message = string(messageBytes)
			}
		}
	}


OUT:
	return handler.Response{
		Body:       []byte(message),
		StatusCode: http.StatusOK,
	}, err
}

func GetPodIP() string {
	ipList := GetIps()
	for _, ip := range ipList {
		// filter loopback and ipv6
		if ip.IsLoopback() || ip.To4() == nil {
			continue
		}
		ipStr := ip.String()
		if strings.HasPrefix(ipStr, "10.") {
			return ipStr
		}
	}
	return ""
}

func GetIps() []net.IP {

	var ifaces []net.Interface
	var err error
	ifaces, err = net.Interfaces()
	if err != nil {
		log.Printf("Get Interfaces Error\n")
		return nil
	}

	var ipList []net.IP
	for _, i := range ifaces {
		addrs, err := i.Addrs()
		if err != nil {
			log.Printf("Get Addrs Error!\n")
			continue
		}

		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			ipList = append(ipList, ip)
		}
	}

	return ipList
}
