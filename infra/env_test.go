package infra_test

import (
	"os"
	"testing"

	"github.com/mkaiho/go-oauth-sandbox/infra"
)

func TestSimpleEnvLoader(t *testing.T) {
	t.Run("NewSimpleEnvLoader", func(t *testing.T) {
		t.Run("Returned SimpleEnvLoader", func(t *testing.T) {
			actual := infra.NewSimpleEnvLoader()
			if actual == nil {
				t.Fatal("expected to return value that is not null, but was not")
			}
		})
	})

	t.Run("Get", func(t *testing.T) {
		t.Run("Returned value from key", func(t *testing.T) {
			key := "testkey"
			expected := "testvalue"
			os.Setenv(key, expected)
			t.Cleanup(func() {
				os.Setenv(key, "")
			})

			sut := infra.NewSimpleEnvLoader()

			actual := sut.Get(key)

			if actual != expected {
				t.Fatalf("expected value is \"%s\", but was \"%s\"", expected, actual)
			}
		})

		t.Run("Returned empty string when value is not exits", func(t *testing.T) {
			key := "testkey"
			expected := ""

			sut := infra.NewSimpleEnvLoader()

			actual := sut.Get(key)

			if actual != expected {
				t.Fatalf("expected value is \"%s\", but was \"%s\"", expected, actual)
			}
		})
	})
}
