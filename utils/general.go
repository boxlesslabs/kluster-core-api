//=============================================================================
// developer: boxlesslabsng@gmail.com
// General utility library
//=============================================================================
 
/**
 **
 * @struct GeneralUtil
 **
 * @GenerateRandomString() returns a random string of fixed lenght
 * @GenerateRandomInt() returns a random string of integers
 * @RemoveDuplicates() removes duplicate values from a slice
**/

package utils

import (
	"crypto/sha256"
	"fmt"
	"math/big"
	"math/rand"
	"path"
	"reflect"
	"runtime"
	"strings"
	"time"
)

const (
	charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	intCharset = "0123456789"
)

var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

type GeneralUtil struct {}

func (util *GeneralUtil) GenerateRandomString(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func (util *GeneralUtil) GenerateRandomInt(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = intCharset[seededRand.Intn(len(intCharset))]
	}
	return string(b)
}

func (util *GeneralUtil) RemoveDuplicates(elements []string) []string {
	encountered := map[string]bool{}
	for v := range elements {
		encountered[elements[v]] = true
	}

	result := []string{}
	for key, _ := range encountered {
		result = append(result, key)
	}
	return result
}

func (util *GeneralUtil) ConvertStructToMap(st interface{}) map[string]interface{} {

	reqRules := make(map[string]interface{})

	v := reflect.ValueOf(st)
	t := reflect.TypeOf(st)

	for i := 0; i < v.NumField(); i++ {
		key := strings.ToLower(t.Field(i).Name)
		typ := v.FieldByName(t.Field(i).Name).Kind().String()
		structTag := t.Field(i).Tag.Get("json")
		jsonName := strings.TrimSpace(strings.Split(structTag, ",")[0])
		value := v.FieldByName(t.Field(i).Name)

		if jsonName != ""  && jsonName != "-" {
			key = jsonName
		}

		if typ == "string" {
			if !(value.String() == "" && strings.Contains(structTag, "omitempty")) {
				reqRules[key] = value.String()
			}
		} else if typ == "int" {
			reqRules[key] = value.Int()
		} else {
			reqRules[key] = value.Interface()
		}

	}

	return reqRules
}

func (util *GeneralUtil) GetFileName(skip int) string {
	_, b, ln, _ := runtime.Caller(skip)
	return fmt.Sprintf("file:%s   line:%d",path.Base(b),ln)
}

func (util *GeneralUtil) HashPassword(password string) string {

	_hash_ := sha256.Sum256([]byte(password))
	bigIntRep := new(big.Int)
	bigIntRep = bigIntRep.SetBytes(_hash_[:])

	to16 := bigIntRep.Text(16)
	return to16
}

func (util *GeneralUtil) CleanRoute(root string, path string) string {
	return root + path
}