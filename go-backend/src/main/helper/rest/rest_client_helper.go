package rest

func POST(req *Request) *Response {
	return NewRestClient().POST(req)
}
