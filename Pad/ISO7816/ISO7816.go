package ISO7816
import (
	"errors"
	"strings"
)
/*
An implementation of the ISO/IEC 7816-4 padding algorithm
*/
const fb, zero = string(80), string(0)
func Pad(data []byte, length int) ([]byte, error){
	if data == nil{return data,errors.New("Data can not be empty")}
	if length < 1 {return data,errors.New("Length can not be < 1")}
	padl := length-(len(data)%length)
	return []byte(string(data)+fb+strings.Repeat(zero,padl-1)), error(nil)
}
func Unpad(data []byte) ([]byte, error){
	if data == nil{return data,errors.New("Data can not be empty")}
	i, l := strings.Index(string(data),fb), len(data)
	if i == -1{return data,errors.New("Data not padded")}
	if e := string(data)[i:];e != fb+strings.Repeat(zero,l-i-1){return data,errors.New("Data not padded")}
	return data[:i], error(nil)
}