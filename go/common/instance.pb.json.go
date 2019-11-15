// Code generated by protoc-gen-jsonify. DO NOT EDIT.
// source: common/instance.proto

package common
import (
	"bytes"
	"encoding/json"
	"github.com/gogo/protobuf/jsonpb"
)

// InstanceJSONMarshaler describes the default jsonpb.Marshaler used by all
// instances of Instance. This struct is safe to replace or modify but
// should not be done so concurrently.
var InstanceJSONMarshaler = new(jsonpb.Marshaler)
// MarshalJSON satisfies the encoding/json Marshaler interface. This method
// uses the more correct jsonpb package to correctly marshal the message.
func (m *Instance) MarshalJSON() ([]byte, error) {
	if m == nil {
		return json.Marshal(nil)
	}
	buf := &bytes.Buffer{}
	if err := InstanceJSONMarshaler.Marshal(buf, m); err != nil {
	  return nil, err
	}
	return buf.Bytes(), nil
}
var _ json.Marshaler = (*Instance)(nil)
// InstanceJSONUnmarshaler describes the default jsonpb.Unmarshaler used by all
// instances of Instance. This struct is safe to replace or modify but
// should not be done so concurrently.
var InstanceJSONUnmarshaler = new(jsonpb.Unmarshaler)
// UnmarshalJSON satisfies the encoding/json Unmarshaler interface. This method
// uses the more correct jsonpb package to correctly unmarshal the message.
func (m *Instance) UnmarshalJSON(b []byte) error {
	return InstanceJSONUnmarshaler.Unmarshal(bytes.NewReader(b), m)
}
var _ json.Unmarshaler = (*Instance)(nil)

