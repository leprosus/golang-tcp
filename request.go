package tcp

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

type Versioner interface {
	Version() (ver Version)
}

type Reader interface {
	Read() (object interface{}, err error)
}

type Request struct {
	fmt.Stringer
	Versioner
	Reader

	body []byte
}

func (req *Request) Version() (ver Version, err error) {
	version := struct {
		Version string `json:"ver"`
	}{}

	err = req.Read(&version)
	if err != nil {
		return
	}

	slices := strings.Split(version.Version, ".")
	if len(slices) != 2 {
		return ver, fmt.Errorf("bad version format; user `major.minor` pattern")
	}

	val, err := strconv.ParseUint(slices[0], 10, 64)
	if err != nil {
		return
	}

	ver.Major = uint(val)

	val, err = strconv.ParseUint(slices[1], 10, 64)
	if err != nil {
		return
	}

	ver.Minor = uint(val)

	return
}

func (req *Request) Command() (com Command, err error) {
	command := struct {
		Command string `json:"com"`
	}{}

	err = req.Read(&command)
	if err != nil {
		return
	}

	slices := strings.Split(command.Command, "/")
	if len(slices) != 2 {
		return com, fmt.Errorf("bad command format; user `object/action` pattern")
	}

	com.Object = strings.ToLower(slices[0])
	com.Action = strings.ToLower(slices[1])

	return
}

func (req *Request) Read(structure interface{}) (err error) {
	err = json.Unmarshal(req.body, &structure)

	return
}

func (req *Request) String() string {
	return string(req.body)
}
