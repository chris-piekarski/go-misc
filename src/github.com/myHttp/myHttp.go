package myHttp

import (
	"fmt"
	"io/ioutil"
	"io"
	"net/http"
)

type data struct {
	b []byte
	filename string
}

var responseData *data
func init() {
	responseData = new(data)
	responseData.filename = ""
}

func SetFileToServe(f string) {
	responseData.filename = f
	if responseData.getData() {
		fmt.Printf("data to send is:\n %s", string(responseData.b))
	} else {
		fmt.Printf("FAILED to read data from %s\n", responseData.filename)
	}	
}

func MyHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Println(req.Method, req.URL, req.Header)
	io.WriteString(w, string(responseData.transform()))
}

func (d *data ) getData() bool {
	rd, err := ioutil.ReadFile(d.filename)
	d.b = rd
	if err != nil {
		return false
	}
	return true
}

func (d *data) transform() []byte {
	Rot13Mod(d.b)
	return d.b
}
