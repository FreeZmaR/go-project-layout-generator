package meta

type App struct {
	Name string
	Type *AppType
}

type AppType struct {
	Name    string
	Modules []AppModules
}

type AppModules struct {
	Name        string
	ConfigTypes []Config[any]
	Lib         string
	LibVersion  string
}
