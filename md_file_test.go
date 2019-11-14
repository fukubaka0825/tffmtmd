package mdfile_test

import (
	"bytes"
	"github.com/po3rin/mdfile"
	"io/ioutil"
	"testing"
)

func TestFmtHclCodeInMd(t *testing.T) {
	normalCases := []struct {
		outline    string
		inputfile  string
		goldenfile string
	}{
		{
			outline:    "normal(formated file match goldenfile)",
			inputfile:  "./testdata/normal.md",
			goldenfile: "./testdata/golden.md",
		},
		{
			outline:    "no code should be formacaed",
			inputfile:  "./testdata/golden.md",
			goldenfile: "./testdata/golden.md",
		},
	}

	abnormalCases := []struct {
		outline    string
		inputfile  string
		errSentence string
	}{
		{
			outline:    "abnormal(fail to format file) ",
			inputfile:  "./testdata/abnormal.md",
			errSentence: "[tffmtmd] failed to format hcl source code. Please check syntax",
		},
	}

	for _, ca := range normalCases {
		ca := ca
		t.Run(ca.outline, func(t *testing.T) {
			t.Parallel()
			md, err := ioutil.ReadFile(ca.inputfile)
			if err != nil {
				t.Fatalf("failed to read bytes from %v: ", ca.inputfile)
			}

			file := mdfile.NewMdFile(&md,ca.inputfile)
			got,err := file.FmtHclCodeInMd()
			if err != nil {
				t.Fatalf("got unexpected error: %v", err)
			}

			want, err := ioutil.ReadFile(ca.goldenfile)
			if err != nil {
				t.Fatalf("failed to read bytes from %v: ", ca.goldenfile)
			}

			if !bytes.Equal(want, got) {
				t.Errorf("unexpected result \n  want : %v \n got : %v",want,got)
			}
		})
	}

	for _, ca := range abnormalCases {
		ca := ca
		t.Run(ca.outline, func(t *testing.T) {
			t.Parallel()
			md, err := ioutil.ReadFile(ca.inputfile)
			if err != nil {
				t.Fatalf("failed to read bytes from %v: ", ca.inputfile)
			}

			file := mdfile.NewMdFile(&md,ca.inputfile)
			_,err = file.FmtHclCodeInMd()
			if err.Error() != ca.errSentence {
				t.Fatalf("unexpected errSentence \n want : %v \n got : %v",ca.errSentence,err)
			}
		})
	}
}
