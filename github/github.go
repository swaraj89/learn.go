package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	resp, err := http.Get("https://api.github.com/users/swaraj89")
	if err != nil {
		log.Fatalf("error: %s", err)
		// Equivalent code
		// log.printf("error: %s", err)
		// os.Exit(1)
	}

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("error: %s", resp.Status)
	}

	fmt.Printf("Content-Type: %s\n", resp.Header.Get("Content-Type"))
	/* if _, err := io.Copy(os.Stdout, resp.Body); err != nil {
		log.Fatalf("Error: Can't copy - %s", err)
	} */

	var r Reply
	dec := json.NewDecoder(resp.Body)
	//pointers in go are used for refrence
	if err := dec.Decode(&r); err != nil {
		log.Fatalf("error: can't decode - %s", err)
	}
	// fmt.Println(r)
	fmt.Printf("%#v\n", r)
}

/*  JSON <-> GO
true/false <-> true/false
string <-> string
null <-> nil
number <-> float64, float32, int8, int 16, int 32, int64, uint8 ...
array <-> []any ([] interface{})
object <-> map[string]any, Struct

Encoding/JSON API
____________________________

JSON -> io.Reader -> Go: json.Decoder
JSON -> []byte -> Go: json.Unmarshal
  Go -> io.writer -> JSON: Json.Encoder
  Go -> []byte -> JSON: Json.Marshal
*/

type Reply struct {
	Name         string
	Public_Repos int
}
