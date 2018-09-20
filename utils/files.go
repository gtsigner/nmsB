package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v2"
)

var (
	FILE_EXT_DOT    = "."
	JSON_EXTENSIONS = []string{"json"}
	YAML_EXTENSIONS = []string{"yaml", "yml"}
)

func FileExists(name string) (bool, error) {
	_, err := os.Stat(name)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func FileExtension(fpath string) string {
	ext := filepath.Ext(fpath)
	withoutDot := strings.TrimLeft(ext, FILE_EXT_DOT)
	return strings.ToLower(withoutDot)
}

func IsFileExtension(fpath string, extensions []string) bool {
	ext := FileExtension(fpath)
	for _, ext2 := range extensions {
		if ext == ext2 {
			return true
		}
	}
	return false
}

func ReadObject(fpath string, v interface{}) error {
	if IsJSON(fpath) {
		err := ReadJSON(fpath, v)
		return err
	}

	if IsYAML(fpath) {
		err := ReadYAML(fpath, v)
		return err
	}

	return fmt.Errorf("unable to read object from file [ %s ], because extention unknown", fpath)
}

func ReadJSON(fpath string, v interface{}) error {
	bytes, err := ioutil.ReadFile(fpath)
	if err != nil {
		return err
	}
	err = json.Unmarshal(bytes, v)
	return err
}

func ReadYAML(fpath string, v interface{}) error {
	bytes, err := ioutil.ReadFile(fpath)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(bytes, v)
	return err
}

func IsYAML(fpath string) bool {
	match := IsFileExtension(fpath, YAML_EXTENSIONS)
	return match
}

func IsJSON(fpath string) bool {
	match := IsFileExtension(fpath, JSON_EXTENSIONS)
	return match
}
