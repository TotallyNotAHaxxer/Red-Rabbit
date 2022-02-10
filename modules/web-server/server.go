package main

import (
	"embed"
	"fmt"
	"os"
	"text/template"

	"github.com/manifoldco/promptui"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
	"github.com/sirupsen/logrus"
)

var (
	//go:embed templates/*.tmpl
	rootFs embed.FS
)

// type set structure for system information
type SysInfo struct {
	Hostname string `bson:hostname`
	Platform string `bson:platform`
	CPU      string `bson:cpu`
	RAM      uint64 `bson:ram`
	Disk     uint64 `bson:disk`
}

// application values
type appValues struct {
	AppName  string
	YourName string
}

func main() {

	var (
		err       error
		fp        *os.File
		templates *template.Template
		subdirs   []string
	)

	values := SysInfo{}
	value2name := appValues{}
	info := new(SysInfo)
	hostStat, _ := host.Info()
	cpuStat, _ := cpu.Info()
	vmStat, _ := mem.VirtualMemory()
	diskStat, _ := disk.Usage("/")
	info.Platform = hostStat.Platform
	info.CPU = cpuStat[0].ModelName
	info.RAM = vmStat.Total / 1024 / 1024
	info.Disk = diskStat.Total / 1024 / 1024
	info.Hostname = hostStat.Hostname
	value2name.AppName = "RR5_host"
	// setting variables for code generatioon and static files
	values.Hostname = info.Hostname
	values.CPU = info.CPU
	values.Platform = info.Platform
	values.RAM = info.RAM
	values.Disk = info.Disk

	rootFsMapping := map[string]string{
		"index.html.tmpl": "static/index.html",
		"main.go.tmpl":    "main.go",
	}

	if err = os.Mkdir(value2name.AppName, 0755); err != nil {
		logrus.WithError(err).Errorf("error attempting to create application directory '%s'", values.Hostname)
	}

	if err = os.Chdir(value2name.AppName); err != nil {
		logrus.WithError(err).Errorf("error changing to new directory '%s'", values.Hostname)
	}

	subdirs = []string{
		"static",
	}

	for _, dirname := range subdirs {
		if err = os.MkdirAll(dirname, 0755); err != nil {
			logrus.WithError(err).Fatalf("unable to create subdirectory %s", dirname)
		}
	}

	/*
	 * Process templates
	 */
	if templates, err = template.ParseFS(rootFs, "templates/*.tmpl"); err != nil {
		logrus.WithError(err).Fatal("error parsing root templates files")
	}

	for templateName, outputPath := range rootFsMapping {
		if fp, err = os.Create(outputPath); err != nil {
			logrus.WithError(err).Fatalf("unable to create file %s for writing", outputPath)
		}

		defer fp.Close()

		if err = templates.ExecuteTemplate(fp, templateName, values); err != nil {
			logrus.WithError(err).Fatalf("unable to exeucte template %s", templateName)
		}
	}

	fmt.Printf("   cd %s\n", value2name.AppName)
	fmt.Printf("   go run .\n")
}

func stringPrompt(label, defaultValue string) string {
	var (
		err    error
		result string
	)

	prompt := promptui.Prompt{
		Label:   label,
		Default: defaultValue,
	}

	if result, err = prompt.Run(); err != nil {
		logrus.WithError(err).Fatalf("error asking for '%s'", label)
	}

	return result
}
