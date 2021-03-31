package creator

import (
	"fmt"
	"io"
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
func Create(source, destination, checkout string) error {

	tempDir, err := ioutil.TempDir(os.TempDir(), "skaphos-*")
	if err != nil {
		return err
	}
	defer os.Remove(tempDir)

	if isGit(source) { // clone
		git.Clone(source, tempDir, git.CloneOptions{
			Branch: checkout,
		})
	} else if !pathExists(source) {
		return fmt.Errorf("source path does not exist")
	}

	transformFile := path.Join(tempDir, "transformations.yml")
	destination, err = filepath.Abs(destination)
	if err != nil {
		return fmt.Errorf("invalid destination")
	}

	exists := pathExists(destination)
	empty, _ := isEmpty(destination)
	if exists && !empty {
		return fmt.Errorf("destination is not empty")
	}

	if !pathExists(transformFile) {
		return fmt.Errorf("source does not have a `transformations.yml` file")
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

	if _, err = os.Stat(absPath); err != nil {
		return false
	}
	return true
}

func isEmpty(name string) (bool, error) {
	f, err := os.Open(name)
	if err != nil {
		return false, err
	}
	defer f.Close()

	_, err = f.Readdirnames(1)
	if err == io.EOF {
		return true, nil
	}
	return false, err // Either not empty or error, suits both cases
}
