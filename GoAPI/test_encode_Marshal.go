package main
import (
	"fmt"
	"reflect"
	"encoding/json"
	"log"
)

type User struct {
	Username	string
	Pass		string
}

func main() {
	u1 := User{"user1", "pass1"}


	v1, err := json.Marshal(u1)
	if err != nil {
		log.Fatal(err)
	}
	
	var v2 User

	fmt.Println(reflect.TypeOf(v2))
	
	err = json.Unmarshal(v1, &v2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(v2)
	fmt.Println(reflect.TypeOf(v2))

}