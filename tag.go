package main

import (
	"fmt"
	"github.com/blang/semver"
	"github.com/charmbracelet/huh"
	"os"
	"sort"
	"strings"
)

func main() {

	if err := checkGitInPath(); err != nil {
		fail("Error: %s", err)
	}

	if err := updateGitTags(); err != nil {
		fail("Error: %s", err)
	}

	var prefix string

	huh.NewInput().
		Title("What's the prefix of your tags?").
		Description("ie: if your tag is v1.0.0, prefix is `v`").
		Value(&prefix).
		Run()

	tags, err := getGitTags()
	if err != nil {
		fail("Error: %s", err)
	}

	tagsString := string(tags)

	lines := strings.Split(tagsString, "\n")
	tagMap := make(map[string]semver.Version)

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" || !strings.HasPrefix(line, prefix) {
			continue
		}
		versionString := line[len(prefix):]
		version, err := semver.Parse(versionString)
		if err != nil {
			continue
		}
		majorMinorKey := fmt.Sprintf("%d.%d", version.Major, version.Minor)

		if oldVersion, exists := tagMap[majorMinorKey]; exists {
			if version.GT(oldVersion) {
				tagMap[majorMinorKey] = version
			}
		} else {
			tagMap[majorMinorKey] = version
		}
	}

	var keys []string
	for k := range tagMap {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		fmt.Println(tagMap[k].String())
	}
}

// fail prints an error message and exits with a non-zero exit code
func fail(format string, args ...interface{}) {
	_, _ = fmt.Fprintf(os.Stderr, format+"\n", args...)
	os.Exit(1)
}
