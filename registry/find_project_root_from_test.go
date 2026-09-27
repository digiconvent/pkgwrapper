package registry

import (
	"os"
	"path/filepath"
	"testing"
)

func TestFindProjectRootFrom_WalksUpUntilItFindsGoMod(t *testing.T) {
	root := t.TempDir()
	if err := os.WriteFile(filepath.Join(root, "go.mod"), []byte("module fixture\n"), 0644); err != nil {
		t.Fatal(err)
	}
	nested := filepath.Join(root, "a", "b", "c")
	if err := os.MkdirAll(nested, 0755); err != nil {
		t.Fatal(err)
	}

	got, err := findProjectRootFrom(nested)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got != root {
		t.Errorf("expected %q, got %q", root, got)
	}
}

func TestFindProjectRootFrom_ErrorsWhenNoGoModExists(t *testing.T) {
	_, err := findProjectRootFrom(t.TempDir())
	if err == nil {
		t.Fatal("expected an error when no go.mod is found")
	}
}
