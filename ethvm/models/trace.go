// Code generated by github.com/actgardner/gogen-avro. DO NOT EDIT.
/*
 * SOURCES:
 *     block.schema.v1.asvc
 *     pendingtx.schema.v1.asvc
 */

package models

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
	return "{\"fields\":[{\"desc\":\"Signals if an error happened during execution\",\"name\":\"isError\",\"type\":\"boolean\"},{\"desc\":\"Stores the error message\",\"name\":\"msg\",\"type\":\"string\"},{\"desc\":\"An array describing transfers\",\"name\":\"transfers\",\"type\":{\"items\":{\"fields\":[{\"doc\":\"Type of op executed inside the transaction\",\"name\":\"op\",\"type\":\"int\"},{\"doc\":\"Raw value of the transaction\",\"name\":\"value\",\"type\":\"bytes\"},{\"doc\":\"Address of the sender\",\"name\":\"from\",\"type\":\"bytes\"},{\"doc\":\"Balance of the sender\",\"name\":\"fromBalance\",\"type\":\"bytes\"},{\"doc\":\"Address of the receiver\",\"name\":\"to\",\"type\":\"bytes\"},{\"doc\":\"Balance of the receiver\",\"name\":\"toBalance\",\"type\":\"bytes\"},{\"doc\":\"Raw input data\",\"name\":\"input\",\"type\":\"bytes\"}],\"name\":\"Transfer\",\"namespace\":\"io.enkrypt.bolt.models.avro\",\"type\":\"record\"},\"type\":\"array\"}}],\"name\":\"Trace\",\"namespace\":\"io.enkrypt.bolt.models.avro\",\"type\":\"record\"}"
}

func (r *Trace) Serialize(w io.Writer) error {
	return writeTrace(r, w)
}
