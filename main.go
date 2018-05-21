package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	flag "github.com/spf13/pflag"
	"gopkg.in/src-d/go-git.v4"
)

var sourceFetcher SourceFetcher
var replacer TokenReplacer

func init() {
	sourceFetcher = githubSourceFetcher{}
	replacer = defaultReplacer{}
	flag.Usage = usage
}

func main() {
	var licenseName = flag.StringP("license", "l", "MIT", "the license to use for the project")

	flag.Parse()

	if flag.NArg() < 1 {
		flag.Usage()
		return
	}

	projectName := flag.Args()[0]
	projectPath := "./" + projectName + "/"

	git.PlainInit(projectPath+".git", true)

	settings := map[string]string{
		"fullname": getUserName(),
		"year":     time.Now().Local().Format("2006"),
		"project":  projectName,
	}

	license := sourceFetcher.fetchLicense(strings.ToLower(*licenseName))

	if license.Body == "" {
		fmt.Println("Could not find a license called", *licenseName)
		return
	}

	gitignore := sourceFetcher.fetchGitIgnore("Go")

	writeFile(projectPath+"LICENSE", replacer.replace(settings, license.Body))
	writeFile(projectPath+"README.md", replacer.replace(settings, README))
	writeFile(projectPath+".gitignore", gitignore.Source)
	writeFile(projectPath+"Gopkg.toml", GOPKG)
	writeFile(projectPath+".travis.yml", TRAVIS)

	writeFile(projectPath+replacer.replace(settings, "[project].go"), replacer.replace(settings, PACKAGEFILE))
	writeFile(projectPath+replacer.replace(settings, "[project]_test.go"), replacer.replace(settings, PACKAGEFILETEST))

	ensureFolder(projectPath + "cmd")
	writeFile(projectPath+"cmd/main.go", replacer.replace(settings, CMDFILE))
	writeFile(projectPath+"cmd/main_test.go", replacer.replace(settings, CMDFILETEST))
}

func usage() {
	fmt.Println("Usage of new:")
	fmt.Println("\n\tnew projectname [options]")
	fmt.Println("\nOptions are:")
	flag.PrintDefaults()
}

func writeFile(filename, content string) {
	file, err := os.Create(filename)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	fmt.Fprintf(file, content)
}

func ensureFolder(folderName string) {
	if _, err := os.Stat(folderName); os.IsNotExist(err) {
		os.Mkdir(folderName, os.ModePerm)
	}
}
