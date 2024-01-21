package generator

import "strings"

const (
	projectNamePlaceholder = "{{projectName}}"
	goVersionPlaceholder   = "{{goVersion}}"
)

func replaceProjectName(str string, projectName string) string {
	return strings.ReplaceAll(str, projectNamePlaceholder, projectName)
}

func replaceGoVersion(str string, goVersion string) string {
	return strings.ReplaceAll(str, goVersionPlaceholder, goVersion)
}

func replaceAll(str, projectName, goVersion string) string {
	return replaceProjectName(replaceGoVersion(str, goVersion), projectName)
}
