package OSINT_Engine

import (
	"strings"
)

func Spider() {
	for _, n := range TAB0.FindAllStringSubmatch(Pages_Crawled, -1) {
		cond1 := strings.Contains(strings.ToLower(n[1]), REGEX_URL0)
		if !cond1 {
			Links_Crawled = append(Links_Crawled, n[1])
		}
	}
	if len(Links_Crawled) < 1 {
		for _, link := range TAB1.FindAllStringSubmatch(Pages_Crawled, -1) {
			cond1 := strings.Contains(strings.ToLower(link[1]), REGEX_URL0)
			if !cond1 {
				Links_Crawled = append(Links_Crawled, link[1])
			}
		}
	}
}
