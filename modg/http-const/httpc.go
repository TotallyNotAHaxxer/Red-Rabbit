package constantshttp

import "golang.org/x/net/html"

var (
	Finder_func func(n *html.Node, tagged string) *html.Node

	DOCEXE = []string{
		"doc",
		"docx",
		"pdf",
		"csv",
		"xls",
		"xlsx",
		"zip",
		"tar",
		"gz",
	}
)
