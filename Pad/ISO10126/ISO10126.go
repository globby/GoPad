package ISO10126
import(
	"errors"
	"crypto/rand"
	"io"
)
/*
An implementation of the ISO 10126 padding algorithm
*/
func Pad(data []byte, length int) ([]byte, error){
	if data == nil{return data,errors.New("Data can not be empty")}
	if length < 1 {return data,errors.New("Length can not be < 1")}
	padl := length-(len(data)%length)
	pad := string(padl)
	if padl > 1{
		p := make([]byte,padl-1)
		io.ReadFull(rand.Reader,p)
		pad = string(p)+pad
	}
	return []byte(string(data)+pad),error(nil)
}
func Unpad(data []byte) ([]byte, error){
	if data == nil{return data,errors.New("Data can not be empty")}
	if e:=int(data[len(data)-1]);e > len(data){
		return data,errors.New("Data not padded")
	}else{
		return data[:len(data)-e],error(nil)
	}
}