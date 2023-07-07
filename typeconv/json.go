package typeconv

import "encoding/json"

func ToJSONRawMessagePtr(b []byte) *json.RawMessage {
	if b == nil {
		return nil
	}
	jRawMessage := json.RawMessage(b)
	return &jRawMessage
}

func JSONRawMessagePtrToBytes(j *json.RawMessage) []byte {
	if j == nil {
		return nil
	}
	return []byte(*j)
}
