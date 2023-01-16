package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Printf(githubInfo("swaraj89"))
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

// githubinfo returns name and numbe rof public repos for login
func githubInfo(login string) (string, int, error) {
	url := "https://api.github.com/users/" + login
	resp, err := http.Get(url)
	if err != nil {
		return "", 0, err
	}

	if resp.StatusCode != http.StatusOK {
		return "", 0, fmt.Errorf("%#v - %s", url, resp.Status)
	}

	fmt.Printf("Content-Type: %s\n", resp.Header.Get("Content-Type"))
	/* if _, err := io.Copy(os.Stdout, resp.Body); err != nil {
		log.Fatalf("Error: Can't copy - %s", err)
	} */

	//annonymous struct
	var r struct {
		Name         string
		Public_Repos int
	}
	dec := json.NewDecoder(resp.Body)
	//pointers in go are used for refrence
	if err := dec.Decode(&r); err != nil {
		log.Fatalf("error: can't decode - %s", err)
	}
	// fmt.Println(r)
	// fmt.Printf("%#v\n", r)
	return r.Name, r.Public_Repos, nil
}

type Reply struct {
	Name         string
	Public_Repos int
}
