package OSINT_Engine

import (
	"net/http"
	"regexp"
)

type Information struct {
	Search_Query     string
	Run_Twitter      bool
	Run_Facebook     bool
	Run_Linkedin     bool
	Results_per_page int
	Pages_to_Crawl   int
}

var (
	HTTP_ = http.Client{}

	OPTIONS                = Information{}
	URL                    = "https://google.com/search"
	REGEX_URL0             = "https://accounts.google.com/"
	Connection_Attempt     = 0
	Connection_Attempt_max = 3
	SET_PAGE               = func(p, n int) int {
		return (p - 1) * n
	}
	Pages_Crawled = ""
	Links_Crawled = []string{}
	TAB0          = regexp.MustCompile(`<a href="([^"]*)" onmousedown`)
	TAB1          = regexp.MustCompile(`<a href="/url\?q=([^"]+)&amp;sa=U[-&;=\d\w]*`)
)
