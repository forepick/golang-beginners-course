package serialization

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"testing"
)

type (
	Person struct {
		FirstName string			`json:"first_name" xml:"Given-Name"`
		LastName string				`json:"last_name" xml:"Last-Name"`
		BirthYear int
	}

)

func TestMarshaling(t *testing.T){
	p := &Person{"Roy", "Pearl", 1980}

	buf, err := json.Marshal(p)
	if err != nil {
		panic("something went wrong")
	}
	fmt.Println(string(buf))

	buf, err = xml.Marshal(p)
	if err != nil {
		panic("something went wrong again")
	}
	fmt.Println(string(buf))

}

func TestUnMarshaling(t *testing.T){
	input := `{"first_name":"Roy", "last_name": "Pearl"}`
	p := &Person{}
	err := json.Unmarshal([]byte(input), p)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", p)

}
