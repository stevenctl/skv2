package codegen

import (
	log "github.com/sirupsen/logrus"
	"github.com/solo-io/autopilot/codegen/util"
	"golang.org/x/tools/imports"
	"io/ioutil"

	"os"
	"path/filepath"
	"strings"
)

func Run(dir string, forceOverwrite bool) error {
	config := filepath.Join(dir, "autopilot.yaml")

	project, err := Load(config)
	if err != nil {
		return err
	}

	if err := project.Validate(); err != nil {
		return err
	}

	files, err := Generate(project)
	if err != nil {
		return err
	}

	for _, file := range files {
		name := filepath.Join(os.Getenv("GOPATH"), "src", file.OutPath)
		content := file.Content

		if !forceOverwrite && file.SkipOverwrite {
			if _, err := os.Stat(name); err == nil {
				log.Printf("skippinng file %v because it exists", name)
				continue
			}
		}

		if err := os.MkdirAll(filepath.Dir(name), 0777); err != nil {
			return err
		}

		perm := file.Permission
		if perm == 0 {
			perm = 0644
		}

		log.Printf("Writing %v", name)

		if err := ioutil.WriteFile(name, []byte(content), perm); err != nil {
			return err
		}

		if !strings.HasSuffix(name, ".go") {
			continue
		}

		formatted, err := imports.Process(name, []byte(content), nil)
		if err != nil {
			return err
		}

		if err := ioutil.WriteFile(name, []byte(formatted), 0644); err != nil {
			return err
		}
	}

	log.Printf("Generating Deepcopy types for %v", project.TypesImportPath)
	if err := util.DeepcopyGen(project.TypesImportPath); err != nil {
		return err
	}

	log.Printf("Finished generating %v", project.ApiVersion+"."+project.Kind)

	return nil
}
