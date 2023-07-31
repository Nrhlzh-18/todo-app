package helpers

import (
	"os"
	"regexp"
)

func GetRootDirectory() string { // func untuk mendapatkan direktori utama (root directory) dari aplikasi
	projectDir := os.Getenv("PROJECT_DIR") // mengambil set PROJECT_DIR
	projectName := regexp.MustCompile(`^(.*` + projectDir + `)`) // mencocokkan dan mendapatkan direktori root dari direktori kerja saat ini
	currentWorkDirectory, _ := os.Getwd() // untuk mendapatkan direktori kerja saat ini (current working directory) dari aplikasi.
	rootPath := projectName.Find([]byte(currentWorkDirectory))

	return string(rootPath)
}
