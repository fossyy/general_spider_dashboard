package config

import (
	"encoding/json"
	"fmt"
	"general_spider_controll_panel/types"
	configView "general_spider_controll_panel/view/config"
	"github.com/google/uuid"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

func GET(w http.ResponseWriter, r *http.Request) {
	configView.Main("config maker").Render(r.Context(), w)
}

func POST(w http.ResponseWriter, r *http.Request) {
	configDir := "configs"
	if _, err := os.Stat(configDir); os.IsNotExist(err) {
		err := os.Mkdir(configDir, os.ModePerm)
		if err != nil {
			panic(err)
		}
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Error when parsing data : %s", err.Error())
		return
	}
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println("unable to get current directory")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	basePath := filepath.Join(currentDir, configDir)
	cleanBasePath := filepath.Clean(basePath)
	configID := uuid.New()

	var config types.Config
	err = json.Unmarshal(body, &config)
	if err != nil {
		fmt.Fprintf(w, "Error when parsing data : %s", err.Error())
		return
	}

	parsedURL, err := url.Parse(config.BaseURL)
	if err != nil {
		fmt.Fprintf(w, "Error when parsing base url : %s", err.Error())
		return
	}
	if _, err := os.Stat(filepath.Join(cleanBasePath, parsedURL.Hostname())); os.IsNotExist(err) {
		err := os.Mkdir(filepath.Join(cleanBasePath, parsedURL.Hostname()), os.ModePerm)
		if err != nil {
			panic(err)
		}
	}
	configPath := filepath.Join(cleanBasePath, parsedURL.Hostname(), configID.String())
	cleanSaveFolder := filepath.Clean(configPath)

	if !strings.HasPrefix(cleanSaveFolder, cleanBasePath) {
		fmt.Println("invalid path")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	configFile, err := os.Create(configPath + ".json")
	if err != nil {
		fmt.Println("unable to create config file : ", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer configFile.Close()

	_, err = configFile.Write(body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Received : %s\n", configID)
}
