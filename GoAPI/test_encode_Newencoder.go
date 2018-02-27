package main
import (
	"os"
	"fmt"
	"reflect"
	"encoding/json"
//	"log"
)

type User struct {
	Username	string
	Pass		string
}

func main() {
//	u1 := User{"user1", "pass1"}

    enc := json.NewEncoder(os.Stdout)
    d := map[string]int{"apple": 5, "lettuce": 7}
    enc.Encode(d)

    fmt.Println(d)
    fmt.Println(reflect.TypeOf(enc))

}