package generator

import "strings"

const (
	projectNamePlaceholder      = "{{projectName}}"
	shortProjectNamePlaceholder = "{{shortProjectName}}"
	goVersionPlaceholder        = "{{goVersion}}"
)

func replaceAll(str, projectName, shortProjectName, goVersion string) string {
	str = replaceProjectName(str, projectName)
	str = replaceShortProjectName(str, shortProjectName)
	str = replaceGoVersion(str, goVersion)

	return str
}

func replaceProjectName(str string, projectName string) string {
	return strings.ReplaceAll(str, projectNamePlaceholder, projectName)
}

func replaceGoVersion(str string, goVersion string) string {
	return strings.ReplaceAll(str, goVersionPlaceholder, goVersion)
}

func replaceShortProjectName(str string, shortProjectName string) string {
	return strings.ReplaceAll(str, shortProjectNamePlaceholder, shortProjectName)
}
