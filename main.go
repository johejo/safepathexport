package main

import (
	"flag"
	"os"
	"os/exec"
	"strings"
)

var (
	key   string
	value string
	push  bool
	shell bool

	w = os.Stdout
)

func init() {
	flag.StringVar(&key, "key", "", "environment variable key")
	flag.StringVar(&value, "value", "", "environment variable value")
	flag.BoolVar(&push, "push", false, "push value to the end of environment value")
	flag.BoolVar(&shell, "shell", false, "invoke a shell process to get environment variables")
}

func main() {
	flag.Parse()
	if shell {
		shell := os.Getenv("SHELL")
		if shell == "" {
			shell = "sh"
		}
		cmd := exec.Command(shell, "-c", "echo $"+key)
		cmd.Stderr = os.Stderr
		out, err := cmd.Output()
		if err != nil {
			os.Exit(1)
		}
		merge(strings.TrimSpace(string(out)))
		os.Exit(0)
	}
	for _, e := range os.Environ() {
		currentKey, currentValue, ok := strings.Cut(e, "=")
		if !ok {
			continue
		}
		if currentKey != key {
			continue
		}
		merge(currentValue)
		os.Exit(0)
	}
	w.WriteString("export " + key + "=" + value + "\n")
}

func merge(currentValue string) {
	if strings.Contains(currentValue, value) {
		currentValue = strings.ReplaceAll(currentValue, ":"+value, "")
	}
	if push {
		w.WriteString("export " + key + `="` + currentValue + ":" + value + `"` + "\n")
	} else {
		w.WriteString("export " + key + `="` + value + ":" + currentValue + `"` + "\n")
	}
}
