package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/RAI015/runbook-cli/internal/generator"
)

func main() {
	if err := run(os.Args[1:], os.Stdout, os.Stderr); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(args []string, stdout, stderr io.Writer) error {
	if len(args) == 0 {
		printUsage(stderr)
		return errors.New("missing command")
	}

	switch args[0] {
	case "generate":
		return runGenerate(args[1:], stdout, stderr)
	default:
		printUsage(stderr)
		return fmt.Errorf("unknown command: %s", args[0])
	}
}

func runGenerate(args []string, stdout, stderr io.Writer) error {
	fs := flag.NewFlagSet("generate", flag.ContinueOnError)
	fs.SetOutput(io.Discard)

	inputPath := fs.String("i", "", "input YAML file path")
	outputPath := fs.String("o", "", "output Markdown file path")

	if err := fs.Parse(args); err != nil {
		printGenerateUsage(stderr)
		return fmt.Errorf("invalid arguments: %w", err)
	}

	if *inputPath == "" {
		printGenerateUsage(stderr)
		return errors.New("missing required flag: -i")
	}
	if *outputPath == "" {
		printGenerateUsage(stderr)
		return errors.New("missing required flag: -o")
	}

	input, err := os.ReadFile(*inputPath)
	if err != nil {
		return fmt.Errorf("read input file: %w", err)
	}

	markdown, err := generator.GenerateFromYAML(input)
	if err != nil {
		return fmt.Errorf("generate markdown: %w", err)
	}

	if err := os.WriteFile(*outputPath, []byte(markdown), 0o644); err != nil {
		return fmt.Errorf("write output file: %w", err)
	}

	fmt.Fprintf(stdout, "generated %s\n", *outputPath)
	return nil
}

func printUsage(w io.Writer) {
	fmt.Fprintln(w, "Usage:")
	fmt.Fprintln(w, "  runbook generate -i runbook.yaml -o runbook.md")
}

func printGenerateUsage(w io.Writer) {
	fmt.Fprintln(w, "Usage:")
	fmt.Fprintln(w, "  runbook generate -i runbook.yaml -o runbook.md")
}
