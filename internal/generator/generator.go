package generator

import (
	"fmt"
	"strings"

	"gopkg.in/yaml.v3"
)

type Runbook struct {
	Title     string       `yaml:"title"`
	Owner     string       `yaml:"owner"`
	Severity  string       `yaml:"severity"`
	Purpose   string       `yaml:"purpose"`
	Prechecks []string     `yaml:"prechecks"`
	Steps     []Step       `yaml:"steps"`
	Rollback  RollbackPlan `yaml:"rollback"`
	Notes     []string     `yaml:"notes"`
}

type Step struct {
	Title string   `yaml:"title"`
	Items []string `yaml:"items"`
}

type RollbackPlan struct {
	Criteria []string `yaml:"criteria"`
	Actions  []string `yaml:"actions"`
}

func GenerateFromYAML(input []byte) (string, error) {
	var rb Runbook
	if err := yaml.Unmarshal(input, &rb); err != nil {
		return "", fmt.Errorf("parse yaml: %w", err)
	}

	if err := validate(rb); err != nil {
		return "", err
	}

	return renderMarkdown(rb), nil
}

func validate(rb Runbook) error {
	if strings.TrimSpace(rb.Title) == "" {
		return fmt.Errorf("missing required field: title")
	}
	if strings.TrimSpace(rb.Purpose) == "" {
		return fmt.Errorf("missing required field: purpose")
	}
	if len(rb.Steps) == 0 {
		return fmt.Errorf("missing required field: steps")
	}

	for i, step := range rb.Steps {
		if strings.TrimSpace(step.Title) == "" {
			return fmt.Errorf("missing required field: steps[%d].title", i)
		}
		if len(step.Items) == 0 {
			return fmt.Errorf("missing required field: steps[%d].items", i)
		}
	}

	return nil
}

func renderMarkdown(rb Runbook) string {
	var b strings.Builder

	b.WriteString("# ")
	b.WriteString(rb.Title)
	b.WriteString("\n\n")

	b.WriteString("## Purpose\n")
	b.WriteString(rb.Purpose)
	b.WriteString("\n\n")

	if rb.Owner != "" {
		b.WriteString("## Owner\n")
		b.WriteString(rb.Owner)
		b.WriteString("\n\n")
	}

	if rb.Severity != "" {
		b.WriteString("## Severity\n")
		b.WriteString(rb.Severity)
		b.WriteString("\n\n")
	}

	if len(rb.Prechecks) > 0 {
		b.WriteString("## Prechecks\n")
		for _, p := range rb.Prechecks {
			b.WriteString("- ")
			b.WriteString(p)
			b.WriteString("\n")
		}
		b.WriteString("\n")
	}

	b.WriteString("## Steps\n")
	for i, step := range rb.Steps {
		fmt.Fprintf(&b, "%d. %s\n", i+1, step.Title)
		for _, item := range step.Items {
			b.WriteString("   - ")
			b.WriteString(item)
			b.WriteString("\n")
		}
		b.WriteString("\n")
	}

	if len(rb.Rollback.Criteria) > 0 || len(rb.Rollback.Actions) > 0 {
		b.WriteString("## Rollback\n")
		if len(rb.Rollback.Criteria) > 0 {
			b.WriteString("### Criteria\n")
		}
		for _, v := range rb.Rollback.Criteria {
			b.WriteString("- ")
			b.WriteString(v)
			b.WriteString("\n")
		}
		if len(rb.Rollback.Criteria) > 0 && len(rb.Rollback.Actions) > 0 {
			b.WriteString("\n")
		}
		if len(rb.Rollback.Actions) > 0 {
			b.WriteString("### Actions\n")
		}
		for _, r := range rb.Rollback.Actions {
			b.WriteString("- ")
			b.WriteString(r)
			b.WriteString("\n")
		}
		b.WriteString("\n")
	}

	if len(rb.Notes) > 0 {
		b.WriteString("## Notes\n")
		for _, note := range rb.Notes {
			b.WriteString("- ")
			b.WriteString(note)
			b.WriteString("\n")
		}
		b.WriteString("\n")
	}

	return strings.TrimSpace(b.String()) + "\n"
}
