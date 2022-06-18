package main

import (
	"flag"
	"os"
	"strings"
)

var (
	key   string
	value string
	push  bool
)

func init() {
	flag.StringVar(&key, "key", "", "environment variable key")
	flag.StringVar(&value, "value", "", "environment variable value")
	flag.BoolVar(&push, "push", false, "push value to the end of environment value")
}

func main() {
	flag.Parse()
	w := os.Stdout
	for _, e := range os.Environ() {
		currentKey, currentValue, ok := strings.Cut(e, "=")
		if !ok {
			continue
		}
		if currentKey != key {
			continue
		}
		if strings.Contains(currentValue, value) {
			currentValue = strings.ReplaceAll(currentValue, ":"+value, "")
		}
		if push {
			w.WriteString("export " + currentKey + "=" + currentValue + ":" + value + "\n")
		} else {
			w.WriteString("export " + currentKey + "=" + value + ":" + currentValue + "\n")
		}
		os.Exit(0)
	}
	w.WriteString("export " + key + "=" + value + "\n")
}
