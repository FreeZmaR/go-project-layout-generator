package generator

type ProjectSetting struct {
	projectName         string
	goVersion           string
	withCodeExample     bool
	useFX               bool
	packages            []packageType
	needInstallPackages bool
}

type packageType struct {
	name string
	link string
}

func newDefaultProjectSetting() *ProjectSetting {
	return &ProjectSetting{
		projectName:         "",
		goVersion:           "1.21",
		withCodeExample:     false,
		useFX:               false,
		needInstallPackages: false,
	}
}

func (s *ProjectSetting) SetProjectName(name string) *ProjectSetting {
	s.projectName = name

	return s
}

func (s *ProjectSetting) SetGoVersion(version string) *ProjectSetting {
	s.goVersion = version

	return s
}
