// Code generated by github.com/actgardner/gogen-avro. DO NOT EDIT.
/*
 * SOURCE:
 *     schemas.asvc
 */

package ethvm

type UnionBlockTransactionLogTraceTransfer struct {
	Block       *Block
	Transaction *Transaction
	Log         *Log
	Trace       *Trace
	Transfer    *Transfer
	UnionType   UnionBlockTransactionLogTraceTransferTypeEnum
}

type UnionBlockTransactionLogTraceTransferTypeEnum int

const (
	UnionBlockTransactionLogTraceTransferTypeEnumBlock       UnionBlockTransactionLogTraceTransferTypeEnum = 0
	UnionBlockTransactionLogTraceTransferTypeEnumTransaction UnionBlockTransactionLogTraceTransferTypeEnum = 1
	UnionBlockTransactionLogTraceTransferTypeEnumLog         UnionBlockTransactionLogTraceTransferTypeEnum = 2
	UnionBlockTransactionLogTraceTransferTypeEnumTrace       UnionBlockTransactionLogTraceTransferTypeEnum = 3
	UnionBlockTransactionLogTraceTransferTypeEnumTransfer    UnionBlockTransactionLogTraceTransferTypeEnum = 4
)
