package gobag

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// PrettyPrintJSON ...
func PrettyPrintJSON(b []byte) error {
	var prettyJSON bytes.Buffer
	err := json.Indent(&prettyJSON, b, "", "\t")
	if err != nil {
		return err
	}
	fmt.Println(prettyJSON.String())
	return nil
}

func MarshalPrettyPrintJSON(v interface{}) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}
	return PrettyPrintJSON(b)
}
