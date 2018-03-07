package tcp

import "encoding/json"

type Versioner interface {
	Version() (ver Version)
}

type Reader interface {
	Read() (object interface{}, err error)
}

type Request struct {
	Versioner
	Reader

	body []byte
}

type Version struct {
	Major uint
	Minor uint
}

func (req *Request) Version() (ver Version) {
	return
}

func (req *Request) Read(structure interface{}) (object interface{}, err error) {
	err = json.Unmarshal(req.body, &structure)

	return
}
