package meta

import "strings"

const (
	ConfigFileTypeNone ConfigFileType = 0
	ConfigFileTypeYAML ConfigFileType = 1
	ConfigFileTypeJSON ConfigFileType = 2
	ConfigFileTypeTOML ConfigFileType = 3
	ConfigFileTypeENV  ConfigFileType = 4
)

type Config[T any] struct {
	Name  string
	Value T
}

type ConfigFileType int

func GetConfigFileTypeByFileName(fileName string) ConfigFileType {
	if strings.Contains(fileName, ".yaml") {
		return ConfigFileTypeYAML
	}

	if strings.Contains(fileName, ".json") {
		return ConfigFileTypeJSON
	}

	if strings.Contains(fileName, ".toml") {
		return ConfigFileTypeTOML
	}

	if strings.Contains(fileName, ".env") {
		return ConfigFileTypeENV
	}

	return ConfigFileTypeNone
}

func (c ConfigFileType) Name() string {
	switch c {
	case ConfigFileTypeYAML:
		return "YAML"
	case ConfigFileTypeJSON:
		return "JSON"
	case ConfigFileTypeTOML:
		return "TOML"
	case ConfigFileTypeENV:
		return "ENV"
	default:
		return "None"
	}
}

func (c ConfigFileType) Description() string {
	switch c {
	case ConfigFileTypeYAML:
		return "YAML file (.yaml)"
	case ConfigFileTypeJSON:
		return "JSON file (.json)"
	case ConfigFileTypeTOML:
		return "TOML file (.toml)"
	case ConfigFileTypeENV:
		return "ENV file (.env)"
	default:
		return "Without config file"
	}
}

func (c ConfigFileType) getFileName() string {
	switch c {
	case ConfigFileTypeYAML:
		return "example.config.yaml"
	case ConfigFileTypeJSON:
		return "example.config.json"
	case ConfigFileTypeTOML:
		return "example.config.toml"
	case ConfigFileTypeENV:
		return "example.config.env"
	default:
		return ""
	}
}
