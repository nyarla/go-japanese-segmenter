package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/nyarla/go-japanese-segmenter/internal/tinydictgen"
)

var (
	pkg  string
	bias string
	dict string
)

func init() {
	flag.StringVar(&pkg, "pkg", "dictionary", "package name of generated code.")
	flag.StringVar(&bias, "bias", "-332", "string of initial bias")
	flag.StringVar(&dict, "json", "./dict.json", "dictionary json-data of TinySegmenter")
}

func main() {
	if !flag.Parsed() {
		flag.Parse()
	}

	r, err := os.Open(dict)
	if err != nil {
		r.Close()
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
	defer r.Close()

	decoder := json.NewDecoder(r)
	data := make(tinydictgen.JSONData)

	if err2 := decoder.Decode(&data); err2 != nil {
		fmt.Fprintln(os.Stderr, err2.Error())
		os.Exit(1)
	}

	wd, err3 := os.Getwd()
	if err3 != nil {
		fmt.Fprintln(os.Stderr, err3.Error())
		os.Exit(1)
	}

	w, err4 := os.OpenFile(filepath.Join(wd, pkg+"_generated.go"), os.O_WRONLY|os.O_CREATE, 0644)
	if err4 != nil {
		fmt.Fprintln(os.Stderr, err4.Error())
		os.Exit(1)
	}

	defer w.Close()

	if err5 := tinydictgen.Render(w, pkg, bias, data); err5 != nil {
		fmt.Fprintln(os.Stderr, err4.Error())
		os.Exit(1)
	}
}
