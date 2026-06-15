package main

import (
	"testing"
)

// ==========================================
// parseArgs tests
// ==========================================

func TestParseArgs_defaultBuild(t *testing.T) {
	cfg, err := parseArgs([]string{"prog.minipar"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if cfg.src != "prog.minipar" {
		t.Errorf("src: want prog.minipar, got %q", cfg.src)
	}
	if cfg.binName != "" {
		t.Errorf("binName: want empty (derive from src), got %q", cfg.binName)
	}
	// No emit flags should be set.
	if cfg.tokens.enabled || cfg.ast.enabled || cfg.tac.enabled || cfg.c.enabled {
		t.Error("no emit flags should be enabled for bare src invocation")
	}
}

func TestParseArgs_namedBinary(t *testing.T) {
	cfg, err := parseArgs([]string{"prog.minipar", "app"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if cfg.src != "prog.minipar" {
		t.Errorf("src: want prog.minipar, got %q", cfg.src)
	}
	if cfg.binName != "app" {
		t.Errorf("binName: want app, got %q", cfg.binName)
	}
}

func TestParseArgs_astStdout(t *testing.T) {
	// -ast without a path → AST to stdout.
	cfg, err := parseArgs([]string{"-ast", "prog.minipar"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !cfg.ast.enabled {
		t.Error("ast.enabled: want true")
	}
	if cfg.ast.path != "" {
		t.Errorf("ast.path: want empty (stdout), got %q", cfg.ast.path)
	}
	if cfg.src != "prog.minipar" {
		t.Errorf("src: want prog.minipar, got %q", cfg.src)
	}
}

func TestParseArgs_astToFile(t *testing.T) {
	// -ast=ast.json → AST to file.
	cfg, err := parseArgs([]string{"-ast=ast.json", "prog.minipar"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !cfg.ast.enabled {
		t.Error("ast.enabled: want true")
	}
	if cfg.ast.path != "ast.json" {
		t.Errorf("ast.path: want ast.json, got %q", cfg.ast.path)
	}
}

func TestParseArgs_tokensFlag(t *testing.T) {
	cfg, err := parseArgs([]string{"-tokens", "prog.minipar"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !cfg.tokens.enabled {
		t.Error("tokens.enabled: want true")
	}
	if cfg.tokens.path != "" {
		t.Errorf("tokens.path: want empty, got %q", cfg.tokens.path)
	}
}

func TestParseArgs_tokensToFile(t *testing.T) {
	cfg, err := parseArgs([]string{"-tokens=tok.txt", "prog.minipar"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !cfg.tokens.enabled {
		t.Error("tokens.enabled: want true")
	}
	if cfg.tokens.path != "tok.txt" {
		t.Errorf("tokens.path: want tok.txt, got %q", cfg.tokens.path)
	}
}

func TestParseArgs_tacFlag(t *testing.T) {
	cfg, err := parseArgs([]string{"-tac", "prog.minipar"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !cfg.tac.enabled {
		t.Error("tac.enabled: want true")
	}
}

func TestParseArgs_cToFile(t *testing.T) {
	cfg, err := parseArgs([]string{"-c=out.c", "prog.minipar"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !cfg.c.enabled {
		t.Error("c.enabled: want true")
	}
	if cfg.c.path != "out.c" {
		t.Errorf("c.path: want out.c, got %q", cfg.c.path)
	}
}

func TestParseArgs_multipleFlagsCombine(t *testing.T) {
	// All flags may be combined; binary_name is the second positional.
	cfg, err := parseArgs([]string{"-tokens", "-tac", "prog.minipar", "app"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !cfg.tokens.enabled {
		t.Error("tokens.enabled: want true")
	}
	if !cfg.tac.enabled {
		t.Error("tac.enabled: want true")
	}
	if cfg.ast.enabled {
		t.Error("ast.enabled: want false")
	}
	if cfg.src != "prog.minipar" {
		t.Errorf("src: want prog.minipar, got %q", cfg.src)
	}
	if cfg.binName != "app" {
		t.Errorf("binName: want app, got %q", cfg.binName)
	}
}

func TestParseArgs_allFlagsWithPaths(t *testing.T) {
	args := []string{
		"-tokens=tok.txt",
		"-ast=ast.json",
		"-tac=out.tac",
		"-c=out.c",
		"prog.minipar",
		"app",
	}
	cfg, err := parseArgs(args)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	checks := []struct {
		name    string
		enabled bool
		path    string
	}{
		{"tokens", cfg.tokens.enabled, cfg.tokens.path},
		{"ast", cfg.ast.enabled, cfg.ast.path},
		{"tac", cfg.tac.enabled, cfg.tac.path},
		{"c", cfg.c.enabled, cfg.c.path},
	}
	paths := []string{"tok.txt", "ast.json", "out.tac", "out.c"}
	for i, ch := range checks {
		if !ch.enabled {
			t.Errorf("%s.enabled: want true", ch.name)
		}
		if ch.path != paths[i] {
			t.Errorf("%s.path: want %q, got %q", ch.name, paths[i], ch.path)
		}
	}
	if cfg.src != "prog.minipar" {
		t.Errorf("src: want prog.minipar, got %q", cfg.src)
	}
	if cfg.binName != "app" {
		t.Errorf("binName: want app, got %q", cfg.binName)
	}
}

func TestParseArgs_noSrcError(t *testing.T) {
	_, err := parseArgs([]string{})
	if err == nil {
		t.Fatal("want error when no source file is given")
	}
}

func TestParseArgs_noSrcWithFlagError(t *testing.T) {
	_, err := parseArgs([]string{"-ast"})
	if err == nil {
		t.Fatal("want error when flag is given but source file is missing")
	}
}

func TestParseArgs_helpShort(t *testing.T) {
	cfg, err := parseArgs([]string{"-h"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !cfg.help {
		t.Error("help: want true for -h")
	}
}

func TestParseArgs_helpLong(t *testing.T) {
	cfg, err := parseArgs([]string{"--help"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !cfg.help {
		t.Error("help: want true for --help")
	}
}

func TestParseArgs_helpDoesNotRequireSrc(t *testing.T) {
	// -h/-help should not error even without a source file.
	_, err := parseArgs([]string{"-h"})
	if err != nil {
		t.Fatalf("help should not error when source is absent, got: %v", err)
	}
}
