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

type Transfer struct {
	Op          int32
	Value       []byte
	From        []byte
	FromBalance []byte
	To          []byte
	ToBalance   []byte
	Input       []byte
}

func DeserializeTransfer(r io.Reader) (*Transfer, error) {
	return readTransfer(r)
}

func NewTransfer() *Transfer {
	v := &Transfer{}

	return v
}

func (r *Transfer) Schema() string {
	return "{\"fields\":[{\"doc\":\"Type of op executed inside the transaction\",\"name\":\"op\",\"type\":\"int\"},{\"doc\":\"Raw value of the transaction\",\"name\":\"value\",\"type\":\"bytes\"},{\"doc\":\"Address of the sender\",\"name\":\"from\",\"type\":\"bytes\"},{\"doc\":\"Balance of the sender\",\"name\":\"fromBalance\",\"type\":\"bytes\"},{\"doc\":\"Address of the receiver\",\"name\":\"to\",\"type\":\"bytes\"},{\"doc\":\"Balance of the receiver\",\"name\":\"toBalance\",\"type\":\"bytes\"},{\"doc\":\"Raw input data\",\"name\":\"input\",\"type\":\"bytes\"}],\"name\":\"Transfer\",\"namespace\":\"io.enkrypt.bolt.models.avro\",\"type\":\"record\"}"
}

func (r *Transfer) Serialize(w io.Writer) error {
	return writeTransfer(r, w)
}
