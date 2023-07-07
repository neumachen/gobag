package gobag

import (
	"bytes"
)

var emptyJSON = []byte(`{}`)

// AppendJSON appends a the given bytesToAdd to the originalBytes. If
// originalBytes is emptyJSON the value for bytesToAdd is returned instead IF
// bytesToAdd is also not an emptyJSON. If it is, nil is returned.
// This is not meant to be used for JSON patching.
func AppendJSON(originalBytes, bytesToAdd []byte) []byte {
	compareA := bytes.Compare(originalBytes, emptyJSON)
	compareB := bytes.Compare(bytesToAdd, emptyJSON)

	// both are nil
	if AreAllNil(originalBytes, bytesToAdd) {
		return nil
	}

	// both are empty
	if compareA+compareB == 0 {
		return nil
	}

	// compare is
	if compareA != 0 && (compareB == 0 || IsNil(bytesToAdd)) {
		return originalBytes
	}

	if (compareA == 0 || IsNil(originalBytes)) && compareB != 0 {
		return bytesToAdd
	}

	var b bytes.Buffer

	b.Write([]byte(","))
	b.Write(bytesToAdd[1:])
	defer b.Reset()

	return append(originalBytes[:len(originalBytes)-1], b.Bytes()...)
}
