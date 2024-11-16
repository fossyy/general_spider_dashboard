package handlerConfigByID

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func GET(w http.ResponseWriter, r *http.Request) {
	fileID := r.PathValue("id")
	configDir := "configs"
	if _, err := os.Stat(configDir); os.IsNotExist(err) {
		err := os.Mkdir(configDir, os.ModePerm)
		if err != nil {
			panic(err)
		}
	}

	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println("unable to get current directory")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	basePath := filepath.Join(currentDir, configDir)
	cleanBasePath := filepath.Clean(basePath)
	configPath := filepath.Join(cleanBasePath, fileID)

	cleanSaveFolder := filepath.Clean(configPath)
	if !strings.HasPrefix(cleanSaveFolder, cleanBasePath) {
		fmt.Println("invalid path")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var foundFile string
	err = filepath.Walk(cleanBasePath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && strings.HasSuffix(info.Name(), ".json") && info.Name() == fileID+".json" {
			foundFile = path
			return filepath.SkipDir
		}
		return nil
	})

	if err != nil {
		fmt.Println("Error walking through the directory:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if foundFile == "" {
		http.Error(w, "Config file not found", http.StatusNotFound)
		return
	}

	fileContent, err := os.ReadFile(foundFile)
	if err != nil {
		fmt.Println("unable to read config file: ", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(fileContent)
}
