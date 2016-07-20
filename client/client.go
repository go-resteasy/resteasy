package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"strings"
)

/*
=============================================================
Client HTTP
=============================================================
*/

//Client : Resteasy Client
type Client struct {
}

//NewClient for Client
func NewClient() *Client {
	return &Client{}
}

var RestClient = &Client{}

/*
=============================================================
WebTaget HTTP
=============================================================
*/

//WebTaget for client
type WebTaget struct {
	request *Request
	url     string
}

//NewTarget for client
func (client Client) Target(url string) (target *WebTaget) {
	target = &WebTaget{url: url}
	return
}

/*
=============================================================
Response HTTP
=============================================================
*/

//Response Http Response
type Response struct {
	*http.Response
	// data string
	data []byte
}

//ReadEntity convert json/xml to struct
//convert json/xml to struct
func (resp Response) ReadEntity(t interface{}) {
	fmt.Printf("type:%v:", reflect.TypeOf(t))
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		//todo
	}
	fmt.Printf("response: %v\n", string(data))
	fmt.Printf("header: %v\n", string(resp.Header.Get("Content-Type")))
	if strings.Contains(resp.Header.Get("Content-Type"), "application/json") {
		json.Unmarshal(data, t)
	}
}

/*
=============================================================
Entity HTTP
=============================================================
*/

/*Entity for post
* Entity
 */
type Entity struct {
	Data      interface{}
	MediaType string
}

// NewEntity for post or put
func NewEntity(entity interface{}, mediaTyep string) Entity {
	return Entity{Data: entity, MediaType: mediaTyep}
}

/*
=============================================================
Request HTTP
=============================================================
*/

//Request HTTP request
type Request struct {
	url string
}

//NewRequest Http
func (target WebTaget) NewRequest() (req *Request) {
	// target.request = &Request{}
	req = &Request{url: target.url}
	return
}

//Get http method get
func (req Request) Get() (res *Response, err error) {
	resp, err := http.Get(req.url)
	if err != nil {
		//todo
	}
	res = &Response{resp, nil}
	return
}

//Post method
func (req Request) Post(entity Entity) (res *Response, err error) {

	if strings.HasPrefix(entity.MediaType, "application/json") {
		data, err := json.Marshal(entity.Data)
		resp, err := http.Post(req.url, entity.MediaType, bytes.NewReader(data))
		if err != nil {
			//todo
		}
		res = &Response{resp, nil}
	}

	return
}

//Post method
func (req Request) Put(entity Entity) (res *Response, err error) {

	if strings.HasPrefix(entity.MediaType, "application/json") {
		data, err := json.Marshal(entity.Data)
		client := http.DefaultClient
		req, err := http.NewRequest(http.MethodPut, req.url, bytes.NewReader(data))
		req.Header.Add("Content-Type", entity.MediaType)
		resp, err := client.Do(req)
		if err != nil {
			//todo
		}
		res = &Response{resp, nil}
	}

	return
}
