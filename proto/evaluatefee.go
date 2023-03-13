package proto

import "google.golang.org/protobuf/runtime/protoimpl"

type EvaluateFeeTransaction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SourceAddress    string       `protobuf:"bytes,1,opt,name=source_address,json=sourceAddress,proto3" json:"source_address,omitempty"`
	Nonce            int64        `protobuf:"varint,2,opt,name=nonce,proto3" json:"nonce,omitempty"`
	ExprCondition    string       `protobuf:"bytes,3,opt,name=expr_condition,json=exprCondition,proto3" json:"expr_condition,omitempty"`
	Operations       []*Operation `protobuf:"bytes,4,rep,name=operations,proto3" json:"operations,omitempty"`
	Metadata         string       `protobuf:"bytes,5,opt,name=metadata,proto3" json:"metadata,omitempty"`
	FeeLimit         int64        `protobuf:"varint,6,opt,name=fee_limit,json=feeLimit,proto3" json:"fee_limit,omitempty"`
	GasPrice         int64        `protobuf:"varint,7,opt,name=gas_price,json=gasPrice,proto3" json:"gas_price,omitempty"`
	CeilLedgerSeq    int64        `protobuf:"varint,8,opt,name=ceil_ledger_seq,json=ceilLedgerSeq,proto3" json:"ceil_ledger_seq,omitempty"`
	ChainId          int64        `protobuf:"varint,9,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty"`
	AddressPrefix    string       `protobuf:"bytes,10,opt,name=address_prefix,json=addressPrefix,proto3" json:"address_prefix,omitempty"` //it represent the address is raw
	RawSourceAddress []byte       `protobuf:"bytes,11,opt,name=raw_source_address,json=rawSourceAddress,proto3" json:"raw_source_address,omitempty"`
}
