package stitcher_on_pills_test

import (
	stitcher_on_pills "github.com/riyadattani/stitcher-on-pills"
	"strings"
	"testing"
)

func TestStitcher(t *testing.T) {
	t.Run("stitch more than two readers together", func(t *testing.T) {
		s1 := strings.NewReader("Salt")
		s2 := strings.NewReader("Pay")
		s3 := strings.NewReader("Rocks")
		got := stitcher_on_pills.Stitch(s1, s2, s3)

		want := "SaltPayRocks"
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}

func TestStitchURLS(t *testing.T) {
	t.Run("given two urls, when I stitch them, then I get the bodies from those urls concatenated together", func(t *testing.T) {

		url2 := "https://www.riyadattani.com"
		url1 := "https://www.quii.dev"
		got := stitcher_on_pills.StitchersURL(url1, url2)

		want := "Deep dive into pair programming"
		if !strings.Contains(got, want) {
			t.Errorf("did not have %q in %q", want, got)
		}

		want2 := "Speaking at GopherconUK"
		if !strings.Contains(got, want) {
			t.Errorf("did not have %q in %q", want2, got)
		}
	})
}
