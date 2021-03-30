package creator

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"regexp"

	"github.com/gogs/git-module"
	"github.com/rantav/go-archetype/generator"
)

// Create a new project based on the source template.
// If source is a git repository, it will be fetched.
func Create(source, destination string) error {

	tempDir, err := ioutil.TempDir(os.TempDir(), "skaphos-*")
	if err != nil {
		return err
	}
	defer os.Remove(tempDir)

	if isGit(source) {
		// clone
		git.Clone(source, tempDir)
	} else if !pathExists(source) {
		return fmt.Errorf("source path is not valid")
	}

	transformFile := path.Join(tempDir, "transformations.yml")
	destination, err = filepath.Abs(destination)
	if err != nil {
		return fmt.Errorf("invalid destination")
	}

	if !pathExists(transformFile) {
		return fmt.Errorf("project does not have a transformation.yml")
	}

	return generator.Generate(transformFile, tempDir, destination, []string{})
}

func isGit(source string) bool {
	re := regexp.MustCompile(`^(https?://|ssh://).*`)
	return re.Match([]byte(source))
}

func pathExists(source string) bool {
	absPath, err := filepath.Abs(source)
	if err != nil {
		return false
	}

	if _, err := os.Stat(absPath); os.IsNotExist(err) {
		return false
	}
	return true
}
