package myhttp

import (
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
)

const (
	plain    = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	rot13ked = "NOPQRSTUVWXYZABCDEFGHIJKLMnopqrstuvwxyzabcdefghijklm"
)

func TestRot13Mod(t *testing.T) {
	got := []byte(plain)
	Rot13Mod(got)
	if string(got) != rot13ked {
		t.Errorf("Rot13Mod(%q) = %q, want %q", plain, got, rot13ked)
	}
}

func TestRot13Table(t *testing.T) {
	got := []byte(plain)
	Rot13Table(got)
	if string(got) != rot13ked {
		t.Errorf("Rot13Table(%q) = %q, want %q", plain, got, rot13ked)
	}
}

func TestRot13IsItsOwnInverse(t *testing.T) {
	b := []byte("Hello, World! 123")
	want := string(b)
	Rot13Mod(b)
	Rot13Mod(b)
	if string(b) != want {
		t.Errorf("Rot13Mod twice = %q, want %q", b, want)
	}
}

func TestServeFileThroughHandler(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "data.txt")
	if err := os.WriteFile(path, []byte("Hello, World"), 0600); err != nil {
		t.Fatal(err)
	}

	SetFileToServe(path)

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	MyHandler(rec, req)

	if got, want := rec.Body.String(), "Uryyb, Jbeyq"; got != want {
		t.Errorf("MyHandler served %q, want %q", got, want)
	}
}

func TestSetFileToServeMissingFile(t *testing.T) {
	// Exercises the failure path; should not panic.
	SetFileToServe(filepath.Join(t.TempDir(), "does-not-exist"))
}

func BenchmarkRot13Mod(b *testing.B) {
	buf := []byte(plain)
	for i := 0; i < b.N; i++ {
		Rot13Mod(buf)
	}
}

func BenchmarkRot13Table(b *testing.B) {
	buf := []byte(plain)
	for i := 0; i < b.N; i++ {
		Rot13Table(buf)
	}
}
