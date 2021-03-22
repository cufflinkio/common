// Code generated by protoc-gen-go-drpc. DO NOT EDIT.
// protoc-gen-go-drpc version: v0.0.18
// source: scope.proto

package pb

import (
	bytes "bytes"

	jsonpb "github.com/gogo/protobuf/jsonpb"
	proto "github.com/gogo/protobuf/proto"

	drpc "storj.io/drpc"
)

type drpcEncoding_File_scope_proto struct{}

func (drpcEncoding_File_scope_proto) Marshal(msg drpc.Message) ([]byte, error) {
	return proto.Marshal(msg.(proto.Message))
}

func (drpcEncoding_File_scope_proto) Unmarshal(buf []byte, msg drpc.Message) error {
	return proto.Unmarshal(buf, msg.(proto.Message))
}

func (drpcEncoding_File_scope_proto) JSONMarshal(msg drpc.Message) ([]byte, error) {
	var buf bytes.Buffer
	err := new(jsonpb.Marshaler).Marshal(&buf, msg.(proto.Message))
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (drpcEncoding_File_scope_proto) JSONUnmarshal(buf []byte, msg drpc.Message) error {
	return jsonpb.Unmarshal(bytes.NewReader(buf), msg.(proto.Message))
}
