// +build scripts

// This program generates /internal/models/hooks.go.
// It can be invoked by running go run scripts/models/gen_user_audit_hooks.go

// Supported args:
// --print-only

package main

import (
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

// https://blog.carlmjohnson.net/post/2016-11-27-how-to-use-go-generate/

// ! define types to be generated as nullable
var TYPES = []string{"string", "int", "int16", "int32", "int64", "float32", "float64", "bool"}

var (
	PATH_PROJECT_ROOT = GetProjectRootDir()
	PATH_MOD_FILE     = PATH_PROJECT_ROOT + "/go.mod"
	TEMPLATE_PATH     = PATH_PROJECT_ROOT + "/scripts/nullable/nullable_types.gotmpl"
)

type Nullable struct {
	Type string
}

func (n Nullable) Short() string {
	if len(n.Type) > 0 {
		return n.Type[0:1]
	}

	return ""
}

func (n Nullable) NullableType() string {
	return strings.ToUpper(n.Short()) + n.Type[1:]
}

func GetProjectRootDir() string {

	if val, ok := os.LookupEnv("PROJECT_ROOT_DIR"); ok {
		return val
	}

	return "/app"
}

// get all structs in above methods packages
// that have a CreatedBy and UpdatedBy field
func main() {
	nullableTemplate := template.Must(template.New("nullable_types.gotmpl").ParseFiles(TEMPLATE_PATH))

	nullableTypes := make([]Nullable, 0)
	for _, t := range TYPES {
		nullableTypes = append(nullableTypes, Nullable{
			Type: t,
		})
	}

	for _, nt := range nullableTypes {
		f, err := os.Create(filepath.Join(PATH_PROJECT_ROOT, nt.Type+".go"))

		if err != nil {
			log.Fatal(err)
		}

		defer f.Close()

		if err := nullableTemplate.Execute(f, nt); err != nil {
			log.Fatal(err)
		}
	}
}
