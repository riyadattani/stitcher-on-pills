package stitcher_on_pills_test

import (
	"io"
	"strings"
	"testing"
)

func TestStitcher(t *testing.T) {
	t.Run("stitch two readers together", func(t *testing.T) {
		s1 := strings.NewReader("Salt")
		s2 := strings.NewReader("Pay")
		got := stitch(s1, s2)

		want := "SaltPay"
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

}

func stitch(s1 io.Reader, s2 io.Reader) string {
	byteA1, _ := io.ReadAll(s1)
	byteA2, _ := io.ReadAll(s2)

	return string(byteA1) + string(byteA2)
}
