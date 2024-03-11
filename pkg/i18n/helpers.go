package i18n

import (
	"os"
	"path"
)

func getLocalePath() (string, error) {
	rootPath, err := getPwdDirPath()
	if err != nil {
		return "", err
	}
	return path.Join(rootPath, "locales"), nil
}

func getPwdDirPath() (string, error) {
	rootPath, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return rootPath, nil
}
