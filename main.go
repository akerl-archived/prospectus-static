package main

import (
	"fmt"
	"os"

	"github.com/akerl/prospectus/v3/plugin"
)

type config struct {
	Value string `json:"value"`
}

func run(cfg config) error {
	if cfg.Value == "" {
		return fmt.Errorf("required arg not provided: value")
	}

	fmt.Printf(cfg.Value)
	return nil
}

func main() {
	cfg := config{}
	err := plugin.ParseConfig(&cfg)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to load config: %s", err)
		os.Exit(1)
	}
	err = run(cfg)
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(1)
	}
}
