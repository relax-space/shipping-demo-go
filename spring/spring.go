package spring

import (
	"net/http"

	"github.com/pangpanglabs/goutils/httpreq"
)

func Create(url string, reqDto *ReqCreateDto) (statusCode int, respDto *RespCreateDto, err error) {
	req := httpreq.New(http.MethodPost, url, reqDto)
	respDto = &RespCreateDto{}
	statusCode, err = req.Call(&respDto)
	return
}

func Query(url string, reqDto *ReqQueryDto) (statusCode int, respDto *RespQueryDto, err error) {
	req := httpreq.New(http.MethodPost, url, reqDto)
	respDto = &RespQueryDto{}
	statusCode, err = req.Call(&respDto)
	return
}

func Cancel(url string, reqDto *ReqCancelDto) (statusCode int, respDto *RespCancelDto, err error) {
	req := httpreq.New(http.MethodPost, url, reqDto)
	respDto = &RespCancelDto{}
	statusCode, err = req.Call(&respDto)
	return
}

func PushQuery(url string) (statusCode int, respDto *RespPushQueryDto, err error) {
	req := httpreq.New(http.MethodGet, url, nil)
	respDto = &RespPushQueryDto{}
	statusCode, err = req.Call(&respDto)
	return
}
