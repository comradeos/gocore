package main

import "testing"

func TestBuildGreetingReportUsesDefaultName(t *testing.T) {
	report := BuildGreetingReport("   ")

	if report.NormalizedName != "gopher" {
		t.Fatalf("expected default name %q, got %q", "gopher", report.NormalizedName)
	}

	if report.NameLength != 6 {
		t.Fatalf("expected default name length 6, got %d", report.NameLength)
	}

	if report.UppercaseLetters != 0 {
		t.Fatalf("expected 0 uppercase letters, got %d", report.UppercaseLetters)
	}
}

func TestBuildGreetingReportCountsCharactersAndUppercase(t *testing.T) {
	report := BuildGreetingReport(" GoLang ")

	if report.NormalizedName != "GoLang" {
		t.Fatalf("expected trimmed name %q, got %q", "GoLang", report.NormalizedName)
	}

	if report.NameLength != 6 {
		t.Fatalf("expected 6 characters, got %d", report.NameLength)
	}

	if report.UppercaseLetters != 2 {
		t.Fatalf("expected 2 uppercase letters, got %d", report.UppercaseLetters)
	}
}

func TestFormatGreeting(t *testing.T) {
	report := GreetingReport{
		NormalizedName:   "Ada",
		NameLength:       3,
		UppercaseLetters: 1,
	}

	got := FormatGreeting(report)
	want := "Hello, Ada! Your name has 3 characters and 1 uppercase letter(s)."

	if got != want {
		t.Fatalf("expected %q, got %q", want, got)
	}
}
