package utils

import (
	"testing"
)

func TestOptions_ParseValidFlags(t *testing.T) {
	opts := NewOptions()
	args := []string{"-l", "-r", "-R", "-a", "-t"}
	err := opts.Parse(args)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	for _, flag := range "lrRat" {
		if !opts.IsFlagSet(rune(flag)) {
			t.Errorf("expected flag %c to be set", flag)
		}
	}
}

func Test2Options_ParseValidFlags(t *testing.T) {
	opts := NewOptions()
	args := []string{"--l"}
	err := opts.Parse(args)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	for _, flag := range "l" {
		if !opts.IsFlagSet(rune(flag)) {
			t.Errorf("expected flag %c to be set", flag)
		}
	}
}

func TestMixedOptions_ParseValidFlags(t *testing.T) {
	opts := NewOptions()
	args := []string{"-lrRat"}
	err := opts.Parse(args)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	for _, flag := range "lrRat" {
		if !opts.IsFlagSet(rune(flag)) {
			t.Errorf("expected flag %c to be set", flag)
		}
	}
}

func Test2MixedOptions_ParseValidFlags(t *testing.T) {
	opts := NewOptions()
	args := []string{"-lr", "-Rat"}
	err := opts.Parse(args)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	for _, flag := range "lrRat" {
		if !opts.IsFlagSet(rune(flag)) {
			t.Errorf("expected flag %c to be set", flag)
		}
	}
}

func Test3MixedOptions_ParseValidFlags(t *testing.T) {
	opts := NewOptions()
	args := []string{"-lr", "-R", "-at"}
	err := opts.Parse(args)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	for _, flag := range "lrRat" {
		if !opts.IsFlagSet(rune(flag)) {
			t.Errorf("expected flag %c to be set", flag)
		}
	}
}

//#################################################################################

func TestOptions_ParseInvalidFlags(t *testing.T) {
	opts := NewOptions()
	args := []string{"-z", "--invalid", "-123", "---l", "--laR", "-l=rta"}
	err := opts.Parse(args)
	if err == nil {
		t.Fatal("expected an error for invalid flags, got nil")
	}

	expectedErrorMsg := "invalid flag: z"
	if err.Error() != expectedErrorMsg {
		t.Errorf("expected error %q, got %q", expectedErrorMsg, err.Error())
	}
}

func TestOptions_MixedFlags(t *testing.T) {
	opts := NewOptions()
	args := []string{"--l", "-Rta", "-r", "--invalid"}
	err := opts.Parse(args)
	if err == nil {
		t.Fatal("expected an error for invalid flags, got nil")
	}

	// Valid flags check
	if !opts.IsFlagSet('l') || !opts.IsFlagSet('R') || !opts.IsFlagSet('t') || !opts.IsFlagSet('a') || !opts.IsFlagSet('r') {
		t.Errorf("expected valid flags 'l', 'R', 't', 'a' to be set")
	}

	// Invalid flags check
	expectedErrorMsg := "invalid argument: --invalid"
	if err.Error() != expectedErrorMsg {
		t.Errorf("expected error %q, got %q", expectedErrorMsg, err.Error())
	}
}

func TestOptions_EmptyArgs(t *testing.T) {
	opts := NewOptions()
	args := []string{}
	err := opts.Parse(args)
	if err != nil {
		t.Fatal("expected an nil for empty arguments, got error")
	}
}

func TestOptions_CheckUnsetFlag(t *testing.T) {
	opts := NewOptions()
	args := []string{"-l"}
	err := opts.Parse(args)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if opts.IsFlagSet('r') {
		t.Errorf("did not expect flag 'r' to be set")
	}
}
