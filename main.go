package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	getEnv := func(key string) {
		val, ok := os.LookupEnv(key)
		if !ok { // yell if there's no env var set `unset FILEPATH`
			fmt.Fprintf(w, "The %s variable is unset.\nEnsure variable is set in runtime (export %[1]s=\"/etc/hosts\").", key)
		} else {
			content, err := ioutil.ReadFile(val) // read the file "content" set by the env var called in the function
			if err != nil {                      // yell if there's no file present
				fmt.Fprintf(w, "Is \"%s\" a file and NOT a directory? Check your config and environment.", val)
			}
			fmt.Fprintf(w, "%s", string(content)) // dump file content otherwise
		}
	}
	getEnv("FILEPATH") // read a file set by the env var (example: `export FILEPATH="/etc/hosts"`)
}

func main() {
	http.HandleFunc("/", handler)                // send root requests to the http file reader function
	log.Fatal(http.ListenAndServe(":8080", nil)) // serve on port 8080 for CloudRun
}
