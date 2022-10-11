package nano

type Request struct {
	Method Method
	Url string
	Headers []Header
	Params []Param
	Body []byte
}

type Method uint8

type Header [2]string

type Param map[string][]string

const (
	GET Method = iota
	POST
	PUT
	DELETE
)

type Response struct {
	Status int
	Body   []byte
}

//go:wasm-module jig
//go:export httpRequest
func ImportedRequest(method Method, url string, reqBody []byte) *Response

//go:export allocateResponse
func allocateResponse(status, length int) *Response {
	return &Response{status, make([]byte, length)}
}

//go:export performRequest
func performRequest(url string) *Response {
	// GETS DATA, AND WRITES RES TO MEM
	res := ImportedRequest(POST, url, []byte("something~"))
	return allocateResponse(res.Status, len(string(res.Body)))
}
