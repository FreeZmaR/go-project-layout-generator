package generator

type Structure struct {
	ProjectName string    `json:"project_name"`
	Dirs        []Dir     `json:"dirs"`
	GoVersion   string    `json:"go_version"`
	Packages    []Package `json:"packages"`
}

type Dir struct {
	Name  string `json:"name"`
	Dirs  []Dir  `json:"dirs"`
	Files []File `json:"files"`
}

type File struct {
	Name        string `json:"name"`
	ContentFile string `json:"content_file"`
}

type Package struct {
	Link string `json:"link"`
}
