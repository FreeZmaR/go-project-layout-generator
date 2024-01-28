package generator

const (
	DefaultProjectName = "default-app"
	DefaultModName     = "default-app"
	DefaultGoVersion   = "1.21"
)

type ProjectSetting struct {
	modName             string
	projectName         string
	goVersion           string
	withCodeExample     bool
	useFX               bool
	packages            []packageType
	needInstallPackages bool
	isDefault           bool
}

type packageType struct {
	name string
	link string
}

func newDefaultProjectSetting() *ProjectSetting {
	return &ProjectSetting{
		modName:             DefaultModName,
		projectName:         DefaultProjectName,
		goVersion:           DefaultGoVersion,
		withCodeExample:     false,
		useFX:               false,
		needInstallPackages: false,
		isDefault:           true,
	}
}

func (s *ProjectSetting) SetModName(name string) *ProjectSetting {
	s.modName = name

	return s
}

func (s *ProjectSetting) SetProjectName(name string) *ProjectSetting {
	s.projectName = name

	return s
}

func (s *ProjectSetting) SetGoVersion(version string) *ProjectSetting {
	s.goVersion = version

	return s
}

func (s *ProjectSetting) SetWithCodeExample(withExample bool) *ProjectSetting {
	s.withCodeExample = withExample

	return s
}

func (s *ProjectSetting) SetDefault() *ProjectSetting {
	s.isDefault = true

	return s
}

func (s *ProjectSetting) SetCustom() *ProjectSetting {
	s.isDefault = false

	return s
}

func (s *ProjectSetting) ModName() string {
	return s.modName
}

func (s *ProjectSetting) ProjectName() string {
	return s.projectName
}

func (s *ProjectSetting) GoVersion() string {
	return s.goVersion
}
