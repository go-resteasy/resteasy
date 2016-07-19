package client

/*
=============================================================
Client HTTP
=============================================================
*/

//Client : Resteasy Client
type Client struct {
}

//NewClient for Client
func NewClient() Client {
	return Client{}
}

/*
=============================================================
WebTaget HTTP
=============================================================
*/

//WebTaget for client
type WebTaget struct {
}

//Target for client
func (client Client) Target(url string) (targe WebTaget) {
	return
}

/*
=============================================================
Response HTTP
=============================================================
*/

//Response Http Response
type Response struct {
}

/*
=============================================================
Request HTTP
=============================================================
*/

//Request HTTP request
type Request struct{}

//Request Http
func (target WebTaget) Request() (req Request) {
	return
}

//Get http method get
func (req Request) Get() (res Response) {
	return
}
