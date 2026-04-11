package main

import (
	"fmt"
	"strings"
)

type GreetingReport struct {
	Name             string
	NormalizedName   string
	NameLength       int
	UppercaseLetters int
}

func BuildGreetingReport(name string) GreetingReport {
	normalized := strings.TrimSpace(name)
	if normalized == "" {
		normalized = "gopher"
	}

	return GreetingReport{
		Name:             name,
		NormalizedName:   normalized,
		NameLength:       len([]rune(normalized)),
		UppercaseLetters: countUppercaseLetters(normalized),
	}
}

func countUppercaseLetters(value string) int {
	count := 0
	for _, r := range value {
		if 'A' <= r && r <= 'Z' {
			count++
		}
	}
	return count
}

func FormatGreeting(report GreetingReport) string {
	return fmt.Sprintf(
		"Hello, %s! Your name has %d characters and %d uppercase letter(s).",
		report.NormalizedName,
		report.NameLength,
		report.UppercaseLetters,
	)
}

func main() {
	report := BuildGreetingReport("gopher")
	fmt.Println(FormatGreeting(report))
}
