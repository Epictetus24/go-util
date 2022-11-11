package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

var regexes = []*regexp.Regexp{
	regexp.MustCompile(`(?i)login`),
	regexp.MustCompile(`(?i)pfx`),
	regexp.MustCompile(`(?i).cer`),
	regexp.MustCompile(`(?i)publishsettings`),
	regexp.MustCompile(`(?i)cspkg`),
	regexp.MustCompile(`(?i).config`),
}

func walkFn(path string, f os.FileInfo, err error) error {
	for _, r := range regexes {
		if r.MatchString(path) {
			fmt.Printf("[+] HIT %s : %s\n", r, path)
		}
	}
	return nil
}

func stringHunt(path2 string, f2 os.FileInfo, err error) error {
	if f2.IsDir() {
		return nil
	}

	fn, err := os.Open(path2)
	if err != nil {
		log.Fatalln(err)
	}
	defer fn.Close()

	scanner := bufio.NewScanner(fn)
	azurestrings := []string{"TokenCache", "Tenant", "PublishSettings", "FileURL", "ManagementPortalURL"}
	azurecnt := len(azurestrings)

	for scanner.Scan() {
		file := scanner.Text()
		for x := 0; x < azurecnt; x++ {
			str := azurestrings[x]
			if strings.Contains(file, str) {
				fmt.Printf("[+] Possible JSON profile %s : %s\n", str, path2)
			}
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatalln(err)
	}
	return nil
}

func main() {
	root := os.Args[1]
	if err := filepath.Walk(root, walkFn); err != nil {
		log.Panicln(err)
	}
	if err := filepath.Walk(root, stringHunt); err != nil {
		log.Panicln(err)
	}
}
