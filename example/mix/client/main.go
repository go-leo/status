package main

import (
	"bytes"
	"flag"
	"io"
	"log"
	"net/http"

	"github.com/go-leo/status"
	statuspb "github.com/go-leo/status/example/status"
)

var (
	addr = flag.String("addr", "localhost:60062", "the address to connect to")
	name = flag.String("name", "", "Name to greet")
)

func main() {
	flag.Parse()
	resp, err := http.Post("http://"+*addr+"/hello", "text/plain", bytes.NewBuffer([]byte(*name)))
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	err, ok := status.From(resp)
	if ok {
		var st status.Status
		var ok bool
		if st, ok = statuspb.IsDefault(err); ok {
			jsonData, _ := st.MarshalJSON()
			log.Fatalf("default error: %v, json: %s, header: %v", st, jsonData, st.Headers())
		} else if st, ok = statuspb.IsJustRpcStatus(err); ok {
			jsonData, _ := st.MarshalJSON()
			log.Fatalf("just rpc status error: %v, json: %s, header: %v", st, jsonData, st.Headers())
		} else if st, ok = statuspb.IsJustHttpStatus(err); ok {
			jsonData, _ := st.MarshalJSON()
			log.Fatalf("just http status error: %v, json: %s, header: %v", st, jsonData, st.Headers())
		} else if st, ok = statuspb.IsJustMessage(err); ok {
			jsonData, _ := st.MarshalJSON()
			log.Fatalf("just message error: %v, json: %s, header: %v", st, jsonData, st.Headers())
		} else if st, ok = statuspb.IsAllHave(err); ok {
			jsonData, _ := st.MarshalJSON()
			log.Fatalf("all have error: %v, json: %s, header: %v", st, jsonData, st.Headers())
		} else {
			jsonData, _ := st.MarshalJSON()
			log.Fatalf("custom error: %v, json: %s, header: %v", st, jsonData, st.Headers())
		}
	}

	message, _ := io.ReadAll(resp.Body)
	log.Printf("Greeting: %s", message)
}
