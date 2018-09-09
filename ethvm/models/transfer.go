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
	Op          string
	Value       string
	From        string
	FromBalance string
	To          string
	ToBalance   string
	Input       string
}

func DeserializeTransfer(r io.Reader) (*Transfer, error) {
	return readTransfer(r)
}

func NewTransfer() *Transfer {
	v := &Transfer{}

	return v
}

func (r *Transfer) Schema() string {
	return "{\"fields\":[{\"doc\":\"Type of op executed inside the transaction\",\"name\":\"op\",\"type\":\"string\"},{\"doc\":\"Raw value of the transaction\",\"name\":\"value\",\"type\":\"string\"},{\"doc\":\"Address of the sender\",\"name\":\"from\",\"type\":\"string\"},{\"doc\":\"Balance of the sender\",\"name\":\"fromBalance\",\"type\":\"string\"},{\"doc\":\"Address of the receiver\",\"name\":\"to\",\"type\":\"string\"},{\"doc\":\"Balance of the receiver\",\"name\":\"toBalance\",\"type\":\"string\"},{\"doc\":\"Raw input data\",\"name\":\"input\",\"type\":\"string\"}],\"name\":\"Transfer\",\"namespace\":\"io.enkrypt.bolt.models\",\"type\":\"record\"}"
}

func (r *Transfer) Serialize(w io.Writer) error {
	return writeTransfer(r, w)
}
