package PKCS7
import (
	"strings"
	"errors"
)
/*
An implementation of the PKCS7 padding algorithm as described in RFC 5652
*/
func Pad(data []byte, length int) ([]byte, error){
	if data == nil{return data,errors.New("Data can not be empty")}
	if length < 1 {return data,errors.New("Length can not be < 1")}
	padl := length-(len(data)%length)
	padb := string(padl)
	return []byte(string(data)+strings.Repeat(padb,padl)),error(nil)
}
func Unpad(data []byte) ([]byte, error){
	if data == nil{return data,errors.New("Data can not be empty")}
	if e:=int(data[len(data)-1]);e > len(data){
		return data,errors.New("Data not padded")
	}else if string(data[len(data)-e:]) != strings.Repeat(string(e),e){
		return data,errors.New("Data not padded")
	}else{
		return data[:len(data)-e],error(nil)
	}
}