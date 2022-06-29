package vuln_structs

import (
	types "main/modg/types"
	"sync"
)

var DOM = []string{
	"document.cookie=",
	"document.referrer=",
	"window.name=",
	"history.pushstate(",
	"history.replacestate(",
	"localstorage.setitem(",
	"localstorage.getitem(",
	"document.url=",
	"document.documenturi=",
	"document.urlencoded=",
	"document.baseuri=",
	"sessionstorage=",
	"document.write(",
	"document.writeIn(",
	"innerHTML=",
	"outerHTML=",
	"eval(",
	"setTimeout(",
	"setInterval(",
	"{{__html",
	"location=",
	"location.href=",
	"location.search=",
	"location.hash=",
	"location.pathname=",
}

var (
	Content_results []types.Result
	Mutmod          = &sync.Mutex{}
	Request_limit   = make(chan string, 10)
	Wg              = sync.WaitGroup{}
)
