// Code generated by github.com/actgardner/gogen-avro. DO NOT EDIT.
/*
 * SOURCE:
 *     schemas.asvc
 */

package ethvm

import (
	"io"
)

type Trace struct {
	IsError   bool
	Msg       string
	Transfers []*Transfer
}

func DeserializeTrace(r io.Reader) (*Trace, error) {
	return readTrace(r)
}

func NewTrace() *Trace {
	v := &Trace{
		Transfers: make([]*Transfer, 0),
	}

	return v
}

func (r *Trace) Schema() string {
	return "{\"fields\":[{\"desc\":\"\",\"name\":\"isError\",\"type\":\"boolean\"},{\"desc\":\"\",\"name\":\"msg\",\"type\":\"string\"},{\"desc\":\"\",\"name\":\"transfers\",\"type\":{\"items\":{\"fields\":[{\"doc\":\"Type of op executed inside the transaction\",\"name\":\"op\",\"type\":\"string\"},{\"doc\":\"\",\"name\":\"from\",\"type\":\"string\"},{\"doc\":\"\",\"name\":\"to\",\"type\":\"string\"},{\"doc\":\"\",\"name\":\"input\",\"type\":\"bytes\"}],\"name\":\"Transfer\",\"namespace\":\"io.enkrypt.bolt.avro\",\"type\":\"record\"},\"type\":\"array\"}}],\"name\":\"Trace\",\"namespace\":\"io.enkrypt.bolt.avro\",\"type\":\"record\"}"
}

func (r *Trace) Serialize(w io.Writer) error {
	return writeTrace(r, w)
}
