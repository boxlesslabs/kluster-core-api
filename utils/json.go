//=============================================================================
// developer: boxlesslabsng@gmail.com
// JSON Utility library
//=============================================================================
 
/**
 **
 * @struct UnmarshallObjectId
 **
 * @Contains() returns true if item is contained in slice
 * @CleanRoute() returns string of path for routing
 * @EncodeJSON() encodes json to byte
 * @EncodeJSON() decodes json
 * @EncodeResponseAsJSON() encodes object to json
 * @TransformMAP() marshall and unmarshall json
**/

package utils

import (
	"encoding/json"
	"io"
)

type UnmarshallObjectId struct {
	ID string `bson:"ObjectId"`
}

func Contains(item string, items []string) bool {
	for _, this := range items {
		if this == item {
			return true
		}
	}
	return false
}

func CleanRoute(root string, path string) string {
	return root + path
}

func EncodeJSON(source interface{}) ([]byte, error) {
	target, err := json.Marshal(&source)
	if err != nil {
		return nil, err
	}
	return target, nil
}

func DecodeJSON(source []byte, target interface{}) error {
	// convert json to struct

	err := json.Unmarshal(source, &target)
	return err
}

func EncodeResponseAsJSON(data interface{}, w io.Writer) {
	enc := json.NewEncoder(w)
	enc.Encode(data)
}

func TransformMAP(source interface{}, target interface{}) error {
	src, err := json.Marshal(&source)
	if err != nil {
		return err
	}

	// convert json to struct
	_ = json.Unmarshal(src, &target)
	return err
}
