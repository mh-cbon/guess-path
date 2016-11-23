package guesspath

import (
	"os"
	"path/filepath"
	"runtime"
)

var gpath = ""

func init() {
	gpath = os.Getenv("GOPATH")
}

func Glob(
	staticValue string,
	PkgName string,
	RootPkgResPath string,
	Glob string,
) string {
	p := Path(
		staticValue,
		PkgName,
		RootPkgResPath,
	)
	if p != "" {
		return filepath.Join(p, Glob)
	}
	return ""
}

func Path(
	staticValue string,
	PkgName string,
	RootPkgResPath string,
) string {

	if staticValue != "" {
		return staticValue

	} else if k := runtimePath(); exists(k) {
		return filepath.Join(k, RootPkgResPath)

	} else if k := filepath.Join("vendor", PkgName); PkgName!="" && exists(k) {
		return filepath.Join(k, RootPkgResPath)

	}
	return ""
}

func runtimePath() string {
	_, p, _, _ := runtime.Caller(3)
	return filepath.Dir(p)
}

func exists(p string) bool {
	_, err := os.Stat(p)
	return !os.IsNotExist(err)
}
