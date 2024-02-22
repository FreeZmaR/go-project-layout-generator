package modules

import (
	"context"
	"os"
)

type AppNetHTTPServer struct {
	name            string
	mainDirName     string
	description     string
	versions        []string
	selectedVersion string
}

var _ Module = (*AppNetHTTPServer)(nil)

func NewApp() *AppNetHTTPServer {
	return &AppNetHTTPServer{
		name:            "AppNetHTTPServer",
		mainDirName:     "httpsrv",
		description:     "Sub application with http server",
		selectedVersion: "net/http",
	}
}

func (m *AppNetHTTPServer) Name() string {
	return m.name
}

func (m *AppNetHTTPServer) Description() string {
	return m.description
}

func (m *AppNetHTTPServer) SelectedVersion() string {
	return m.selectedVersion
}

func (m *AppNetHTTPServer) Versions() []string {
	return nil
}

func (m *AppNetHTTPServer) SelectVersion(_ string) error {
	return nil
}

func (m *AppNetHTTPServer) Generate(ctx context.Context, dir string, subModules []Module) error {
	//TODO implement me
	panic("implement me")
}

func (m *AppNetHTTPServer) generateImports(ctx context.Context, file *os.File) error {
	if ctx.Err() != nil {
		return ctx.Err()
	}

	imports := []string{
		"context",
		"net",
		"net/http",
	}

	importTMP := "import (\n"

	for _, imp := range imports {
		importTMP += "\t\"" + imp + "\"\n"
	}

	importTMP += ")\n"

	_, err := file.WriteString(importTMP)

	return err
}

func (m *AppNetHTTPServer) generateAppStruct(ctx context.Context, file *os.File) error {
	if ctx.Err() != nil {
		return ctx.Err()
	}

	types := map[string]string{
		"srv":      "*http.Server",
		"serverCH": "chan error",
	}

	typeTMP := "type App struct {"

	for name, typ := range types {
		typeTMP += "\n\t" + name + " " + typ
	}

	typeTMP += "\n}\n"

	_, err := file.WriteString(typeTMP)

	return err
}

func (m *AppNetHTTPServer) generateDefaultAppConstructor(ctx context.Context, file *os.File) error {
	if ctx.Err() != nil {
		return ctx.Err()
	}

	args := map[string]string{
		"cfg": "*types.HTTPServer",
	}

	var argsTMP string
	for name, typ := range args {
		argsTMP += name + " " + typ + ", "
	}

	argsTMP = argsTMP[:len(argsTMP)-2]

	if len(argsTMP) > 0 {
		argsTMP = ", " + argsTMP
	}

	constructorTMP := "func NewApp("

	_, err := file.WriteString(constructorTMP)

	return err
}
