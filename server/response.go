package server

import "bytes"

type Response struct {
	value []byte
}

func ok() Response {
	return Response{value: []byte("+OK\r\n")}
}

func badRequest(value []byte) Response {
	return Response{value: bytes.Join([][]byte{[]byte("-"), value, []byte("\r\n")}, []byte{})}
}

func (p *Response) getValue() []byte {
	return p.value
}

func (p *Response) isOk() bool {
	return len(p.value) > 0 && p.value[0] == '+'
}

func (p *Response) isError() bool {
	return len(p.value) > 0 && p.value[0] == '-'
}
