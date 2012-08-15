package view

import (
	"github.com/hoisie/mustache"
	"os"
	"testing"
)

const templateRoot string = "./templates"

func findFiles(dir string) []string {
	f, err := os.Open(dir)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	contents, err := f.Readdir(0)
	if err != nil {
		panic(err)
	}

	ret := make([]string, 0)
	for _, info := range contents {
		fn := dir + "/" + info.Name()
		if info.IsDir() {
			ret = append(ret, findFiles(fn)...)
		} else {
			ret = append(ret, fn)
		}
	}

	return ret
}

// Make sure all templates are good
func TestTemplates(t *testing.T) {
	files := findFiles(templateRoot)

	for _, f := range files {
		_, err := template.Parse(f)

		// Assure there are no errors
		if err != nil {
			t.Error("Error parsing template!", f, err)
		}
	}
}
