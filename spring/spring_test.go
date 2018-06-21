package spring

import (
	"fmt"
	"os"
	"testing"
)

var (
	url = os.Getenv("SHIPPING_URL")
)

func Test_Create(t *testing.T) {
	s := &ReqCreateDto{
		Eid:    "10001",
		Method: "spring.create",
		Content: &CreateContentDto{
			Quantity: 1,
			Sender: &SenderDto{
				Code:    "CR00",
				Name:    "rose3c",
				Company: "Eland",
				City:    "北京市,北京市,通州区",
				Address: "恒通商务园B37座",
				Phone:   "13051447878",
			},
			Receiver: &ReceiverDto{
				Code:    "CR02",
				Name:    "hh",
				Company: "Eland",
				City:    "北京市,北京市,通州区",
				Address: "恒通商务园B51座",
				Phone:   "13051447877",
			},
		},
	}
	statusCode, respDto, err := Create(url, s)
	fmt.Println(statusCode, respDto, err)
}

func Test_Query(t *testing.T) {
	s := &ReqQueryDto{
		Eid:    "10001",
		Method: "spring.query",
		Content: &QueryContentDto{
			OrderNo: "1803021273351",
		},
	}
	statusCode, respDto, err := Query(url, s)
	fmt.Println(statusCode, respDto, err)
}
func Test_Cancel(t *testing.T) {
	s := &ReqCancelDto{
		Eid:    "10001",
		Method: "spring.cancel",
		Content: &CancelContentDto{
			OrderNo: "1803021273351",
		},
	}
	statusCode, respDto, err := Cancel(url, s)
	fmt.Println(statusCode, respDto, err)
}

func Test_PushQuery(t *testing.T) {
	url = fmt.Sprintf("%v?state=spring.push.query&bill_no=%v", url, "2017849728")
	statusCode, respDto, err := PushQuery(url)
	fmt.Println(statusCode, respDto, err)
}
