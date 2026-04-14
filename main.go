package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
)

var query, setVal, delKey, format, search string
var pretty, stream, update bool

func main() {
	flag.BoolVar(&pretty, "p", false, "Pretty print JSON")
	flag.StringVar(&query, "q", "", "JSONPath query (e.g. .user.name)")
	flag.StringVar(&format, "f", "", "Format output (json, table, csv)")
	flag.StringVar(&search, "s", "", "Search for value")
	flag.StringVar(&setVal, "set", "", "Set value: key=value")
	flag.StringVar(&delKey, "del", "", "Delete key")
	flag.BoolVar(&update, "u", false, "Update file in place")
	flag.Parse()

	data, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "❌ Read error: %v\n", err)
		os.Exit(1)
	}

	// Try to parse as JSON array or object
	var js interface{}
	if err := json.Unmarshal(data, &js); err != nil {
		// Not JSON, treat as raw text
		fmt.Print(string(data))
		return
	}

	args := flag.Args()
	if len(args) > 0 {
		filename := args[0]
		fileData, err := os.ReadFile(filename)
		if err == nil {
			json.Unmarshal(fileData, &js)
		}
	}

	output(os.Stdout, js)

	if update && len(args) > 0 {
		out, _ := json.MarshalIndent(js, "", "  ")
		os.WriteFile(args[0], out, 0644)
	}
}

func output(w io.Writer, js interface{}) {
	var out []byte
	var err error

	if pretty {
		out, err = json.MarshalIndent(js, "", "  ")
	} else {
		out, err = json.Marshal(js)
	}
	if err != nil {
		fmt.Fprintf(os.Stderr, "❌ Output error: %v\n", err)
		return
	}
	fmt.Fprint(w, string(out))
}
