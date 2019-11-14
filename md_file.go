package mdfile

import (
	"bytes"
	"errors"
	"github.com/hashicorp/hcl/v2"

	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"gopkg.in/russross/blackfriday.v2"
)


type MdFile interface {
	FmtHclCodeInMd() ([]byte, error)
}

type impleMdFile struct {
	md       *[]byte
	filename string
}

func NewMdFile(md *[]byte, filename string) MdFile {
	return &impleMdFile{md: md, filename: filename}
}

// FmtHclCodeInMd formats HCL code in Markdown.
// returns error if code has syntax error.
func (i impleMdFile) FmtHclCodeInMd() ([]byte, error) {
	var synerr error
	n := blackfriday.New(blackfriday.WithExtensions(blackfriday.FencedCode)).Parse(*i.md)

	n.Walk(i.hclFmtWalkerFunc(&synerr))
	if synerr != nil {
		return nil, synerr
	}

	return *i.md, nil
}

func (i impleMdFile) isHclCodeBlock(node *blackfriday.Node) bool {
	return node.Type == blackfriday.CodeBlock && (string(node.CodeBlockData.Info) == "hcl" || string(node.CodeBlockData.Info) == "hcl-terraform")
}

func (i impleMdFile) hclFmtWalkerFunc(synerr *error) blackfriday.NodeVisitor {
	return func(node *blackfriday.Node, entering bool) blackfriday.WalkStatus {
		if i.isHclCodeBlock(node) {
			_, syntaxDiags := hclsyntax.ParseConfig(node.Literal, i.filename, hcl.Pos{Line: 1, Column: 1})
			if syntaxDiags.HasErrors() {
				*synerr = errors.New("[tffmtmd] failed to format hcl source code. Please check syntax")
				return blackfriday.Terminate
			}
			result := hclwrite.Format(node.Literal)
			*i.md = bytes.ReplaceAll(*i.md, bytes.TrimRight(node.Literal, "\n"), bytes.TrimRight(result, "\n"))
		}
		return blackfriday.GoToNext
	}
}
