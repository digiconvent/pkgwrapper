package registry

import "testing"

func TestResolveDataRoot_ReturnsATempPathDuringTests(t *testing.T) {
	got, err := resolveDataRoot()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got == "" {
		t.Error("expected a non-empty data root")
	}
}
