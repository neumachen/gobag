package byteconv

import "encoding/json"

func ToJSONRawMessage(b []byte) json.RawMessage {
	if b == nil {
		return nil
	}
	jRawMessage := json.RawMessage(b)
	return jRawMessage
}

func ToJSONRawMessagePtr(b []byte) *json.RawMessage {
	if b == nil {
		return nil
	}
	jRawMessage := ToJSONRawMessage(b)
	return &jRawMessage
}
