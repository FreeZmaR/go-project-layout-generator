package generator

import "embed"

const (
	defaultStructureFileName = "default_struct.json"
	defaultContentDirName    = "default_content"
)

var (
	//go:embed  default_struct.json
	defaultStructureFS embed.FS
)

type Structure struct {
	ProjectName string    `json:"projectName"`
	Dirs        []Dir     `json:"dirs"`
	GoVersion   string    `json:"goVersion"`
	Packages    []Package `json:"packages"`
	Files       []File    `json:"files"`
}

type Dir struct {
	Name      string `json:"name"`
	Dirs      []Dir  `json:"dirs"`
	Files     []File `json:"files"`
	IsDefault bool   `json:"isDefault,omitempty"`
	IsExample bool   `json:"isExample,omitempty"`
}

type File struct {
	Name        string `json:"name"`
	ContentFile string `json:"contentFile,omitempty"`
	Content     string `json:"content,omitempty"`
	IsDefault   bool   `json:"isDefault,omitempty"`
	IsExample   bool   `json:"isExample,omitempty"`
}

type Package struct {
	Link string `json:"link"`
}
