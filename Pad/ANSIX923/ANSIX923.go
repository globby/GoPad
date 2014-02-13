package ANSIX923
import (
	"errors"
	"strings"
)
/*
An implementation of the ANSI X.923 byte padding algorithm
*/
const zero = string(int(0))
func Pad(data []byte, length int) ([]byte, error){
	if data == nil{return data,errors.New("Data can not be empty")}
	if length < 1 {return data,errors.New("Length can not be < 1")}
	padl := length-(len(data)%length)
	return []byte(string(data)+strings.Repeat(zero,padl-1)+string(padl)), error(nil)
}
func Unpad(data []byte) ([]byte, error){
	if data == nil{return data,errors.New("Data can not be empty")}
	l := len(data);s := int(data[l-1])
	if l<s{return data,errors.New("Data not padded")}
	if e := string(data[l-s:l-1]);e!=strings.Repeat(zero,s-1){return data,errors.New("Data not padded")}
	return data[:l-s], error(nil)
}