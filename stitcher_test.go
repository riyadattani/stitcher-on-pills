package stitcher_on_pills_test

import (
	"bytes"
	"io"
	"strings"
	"testing"
)

func TestStitcher(t *testing.T) {
	t.Run("stitch more than two readers together", func(t *testing.T) {
		s1 := strings.NewReader("Salt")
		s2 := strings.NewReader("Pay")
		s3 := strings.NewReader("Rocks")
		got := stitch(s1, s2, s3)

		want := "SaltPayRocks"
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}

func stitch(readers ...io.Reader) string {
	var buf bytes.Buffer
	for _, rds := range readers {
		_, _ = io.Copy(&buf, rds)
	}

	return buf.String()
}
