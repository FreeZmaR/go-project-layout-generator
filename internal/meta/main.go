package meta

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/BurntSushi/toml"
	"github.com/FreeZmaR/go-project-layout-generator/internal/utils"
	"github.com/joho/godotenv"
	"gopkg.in/yaml.v3"
	"io"
	"os"
	"strings"
)

const (
	GoVersion1_13 GoVersion = "1.13"
	GoVersion1_14 GoVersion = "1.14"
	GoVersion1_15 GoVersion = "1.15"
	GoVersion1_16 GoVersion = "1.16"
	GoVersion1_17 GoVersion = "1.17"
	GoVersion1_18 GoVersion = "1.18"
	GoVersion1_19 GoVersion = "1.19"
	GoVersion1_20 GoVersion = "1.20"
	GoVersion1_21 GoVersion = "1.21"
	GoVersion1_22 GoVersion = "1.22"
)

type Main struct {
	ProjectName    string
	GoVersion      GoVersion
	ConfigFileType ConfigFileType
	ConfigTypes    []Config[any]
	configFilePath string
}

var _ Module = (*Main)(nil)

type GoVersion string

func (m *Main) GetProjectName() string {
	return m.ProjectName
}

func (m *Main) GetGoVersionList() []GoVersion {
	return []GoVersion{
		GoVersion1_22,
		GoVersion1_21,
		GoVersion1_20,
		GoVersion1_19,
		GoVersion1_18,
		GoVersion1_17,
		GoVersion1_16,
		GoVersion1_15,
		GoVersion1_14,
		GoVersion1_13,
	}
}

func (m *Main) GetConfigFileTypeList() []ConfigFileType {
	return []ConfigFileType{
		ConfigFileTypeYAML,
		ConfigFileTypeJSON,
		ConfigFileTypeTOML,
		ConfigFileTypeENV,
		ConfigFileTypeNone,
	}
}

func (m *Main) GetConfigFileType() ConfigFileType {
	return m.ConfigFileType
}

func (m *Main) GetGoVersion() GoVersion {
	return m.GoVersion
}

func (m *Main) SetGoVersion(v GoVersion) error {
	if !utils.OneOf(v, m.GetGoVersionList()...) {
		return ErrGoVersionNotSupported
	}

	m.GoVersion = v

	return nil
}

func (m *Main) SetConfigFileType(v ConfigFileType) error {
	if !utils.OneOf(v, m.GetConfigFileTypeList()...) {
		return ErrConfigFileTypeNotSupported
	}

	m.ConfigFileType = v

	return nil
}

func (m *Main) Scan(ctx context.Context, dir string) error {
	entities, err := os.ReadDir(dir)
	if err != nil {
		return err
	}

	for _, entity := range entities {
		if entity.IsDir() {
			continue
		}

		if entity.Name() == "go.mod" {
			if errScan := m.scanGoVersionAndProjectName(ctx, dir+"/"+entity.Name()); errScan != nil {
				continue
			}
		}

		if strings.Contains(entity.Name(), "config.") || strings.Contains(entity.Name(), ".env") {
			m.configFilePath = dir + "/" + entity.Name()
			if errScan := m.scanConfigFileType(ctx); errScan != nil {
				break
			}
		}
	}

	if m.ConfigFileType == ConfigFileTypeNone {
		return m.findConfigFileInConfigLayer(ctx, dir+"/config")
	}

	return nil
}

func (m *Main) scanGoVersionAndProjectName(ctx context.Context, filePath string) error {
	if ctx.Err() != nil {
		return ctx.Err()
	}

	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return err
	}

	m.findProjectName(ctx, data)
	m.findGoVersion(ctx, data)

	return nil
}

func (m *Main) findProjectName(ctx context.Context, data []byte) {
	if ctx.Err() != nil {
		return
	}

	strData := string(data)

	position := strings.Index(strData, "module")
	if position == -1 {
		return
	}

	for i := position + 6; i < len(strData); i++ {
		if strData[i] == '\n' {
			m.ProjectName = strings.TrimSpace(strData[position+6 : i])

			return
		}
	}
}

func (m *Main) findGoVersion(ctx context.Context, data []byte) {
	if ctx.Err() != nil {
		return
	}

	strData := string(data)

	position := strings.Index(strData, "go ")
	if position == -1 {
		return
	}

	for i := position + 3; i < len(strData); i++ {
		if strData[i] == '\n' {
			m.GoVersion = GoVersion(strings.TrimSpace(strData[position+3 : i]))

			return
		}
	}
}

func (m *Main) scanConfigFileType(ctx context.Context) error {
	if ctx.Err() != nil {
		return ctx.Err()
	}

	file, err := os.Open(m.configFilePath)
	if err != nil {
		return err
	}
	defer file.Close()

	m.ConfigFileType = GetConfigFileTypeByFileName(m.configFilePath)

	data, err := io.ReadAll(file)
	if err != nil {
		return err
	}

	return m.scanConfigTypes(ctx, data)
}

func (m *Main) findConfigFileInConfigLayer(ctx context.Context, configDir string) error {
	entities, err := os.ReadDir(configDir)
	if err != nil {
		return err
	}

	for _, entity := range entities {
		if entity.IsDir() {
			continue
		}

		if strings.Contains(entity.Name(), "config.") || strings.Contains(entity.Name(), ".env") {
			m.configFilePath = configDir + "/" + entity.Name()
			if errScan := m.scanConfigFileType(ctx); errScan != nil {
				break
			}
		}
	}

	return nil
}

func (m *Main) scanConfigTypes(ctx context.Context, data []byte) error {
	if ctx.Err() != nil {
		return ctx.Err()
	}

	mapData := make(map[string]any)

	var err error

	switch m.ConfigFileType {
	case ConfigFileTypeYAML:
		err = yaml.Unmarshal(data, &mapData)
	case ConfigFileTypeJSON:
		err = json.Unmarshal(data, &mapData)
	case ConfigFileTypeTOML:
		_, err = toml.Decode(string(data), &mapData)
	case ConfigFileTypeENV:
		strMap, errParse := godotenv.Parse(bytes.NewReader(data))
		if errParse != nil {
			err = errParse
			break
		}

		mapData = make(map[string]any, len(strMap))
		for k, v := range strMap {
			mapData[k] = v
		}
	default:
		return nil
	}

	if err != nil {
		return ErrMarshalConfig
	}

	for k, v := range mapData {
		m.ConfigTypes = append(
			m.ConfigTypes,
			Config[any]{
				Name:  k,
				Value: v,
			},
		)
	}

	return nil
}
