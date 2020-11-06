package zip_test

import (
	zip "awesomeProject"
	"testing"
)

func TestZip(t *testing.T) {
	err := zip.Zip("/Users/guochenguang/go/src/zip/destt1", "/Users/guochenguang/go/src/zip/t2.zip")
	if err != nil {
		t.Fatal(err)
	}
}

func TestUnzip(t *testing.T) {
	err := zip.Unzip("/Users/guochenguang/go/src/zip/t2.zip", "/Users/guochenguang/go/src/zip/t2/")
	if err != nil {
		t.Fatal(err)
	}
}