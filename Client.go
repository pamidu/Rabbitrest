package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Person struct {
	Name string
	Age  int
}

func main() {

	for i := 0; i < 10000; i++ {
		person := &Person{"Jonh", i}
		buf, _ := json.Marshal(person)
		body := bytes.NewBuffer(buf)
		r, _ := http.Post("http://localhost:8080", "text/json", body)
		response, _ := ioutil.ReadAll(r.Body)
		fmt.Println(string(response))
	}

}
