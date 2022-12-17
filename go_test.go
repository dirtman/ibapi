// Thanks to https://bitfieldconsulting.com/golang/test-scripts

package main

import (
	"testing"

	"github.com/rogpeppe/go-internal/testscript"
)

func Test(t *testing.T) {
	testscript.Run(t, testscript.Params{
		Dir: "testdata/script",
	})
}
