package meta

type Lib struct {
	Name    string
	Package LibPackage
	Configs []Config[any]
}

type LibPackage struct {
	Ref     string
	Version string
}
