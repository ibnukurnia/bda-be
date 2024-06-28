//go:build exclude

package main

import (
	"os"
	"regexp"
	"strings"

	"github.com/thanhpk/randstr"
)

func checkFile() (bool, string) {
	basePath, _ := os.Getwd()
	envPath := basePath + "/.env"

	_, err := os.Stat(envPath)

	return os.IsExist(err), envPath
}

func main() {
	ok, envPath := checkFile()
	if !ok {
		_, err := os.Create(envPath)
		if err != nil {
			panic(err.Error())
		}
	}

	c, err := os.ReadFile(envPath)
	if err != nil {
		panic(err.Error())
	}

	key, err := regexp.Compile("APP_KEY=.*")
	if err != nil {
		panic(err.Error())
	}

	m := key.FindString(string(c))
	nk := randstr.Base64(32)

	n := strings.Replace(string(c), m, "APP_KEY="+nk, 1)

	if err := os.WriteFile(envPath, []byte(n), 0644); err != nil {
		panic(err.Error())
	}
}
