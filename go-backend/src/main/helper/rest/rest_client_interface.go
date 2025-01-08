package rest

type IRestClient interface {
	POST(req *Request) *Response
}
