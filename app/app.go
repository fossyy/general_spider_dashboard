package app

import (
	"bytes"
	"fmt"
	"general_spider_controll_panel/types"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

var scrapydURL = "http://localhost:6800"
var project = "general"
var spider = "general_engine"
var version = "1.0"
var eggPath = "general.egg"

type App struct {
	http.Server
	Database types.Database
}

func uploadEgg(scrapydURL, project, version, eggPath string) error {
	eggFile, err := os.Open(eggPath)
	if err != nil {
		return fmt.Errorf("failed to open egg file: %w", err)
	}
	defer eggFile.Close()

	var requestBody bytes.Buffer
	writer := multipart.NewWriter(&requestBody)

	if err := writer.WriteField("project", project); err != nil {
		return fmt.Errorf("failed to add project field: %w", err)
	}

	if err := writer.WriteField("version", version); err != nil {
		return fmt.Errorf("failed to add version field: %w", err)
	}

	part, err := writer.CreateFormFile("egg", eggPath)
	if err != nil {
		return fmt.Errorf("failed to add egg file field: %w", err)
	}
	if _, err := io.Copy(part, eggFile); err != nil {
		return fmt.Errorf("failed to copy egg file data: %w", err)
	}

	if err := writer.Close(); err != nil {
		return fmt.Errorf("failed to close multipart writer: %w", err)
	}

	url := fmt.Sprintf("%s/addversion.json", scrapydURL)
	req, err := http.NewRequest("POST", url, &requestBody)
	if err != nil {
		return fmt.Errorf("failed to create HTTP request: %w", err)
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send HTTP request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("upload failed with status: %s", resp.Status)
	}

	fmt.Println("Egg uploaded successfully!")
	return nil
}

func NewApp(addr string, handler http.Handler, database types.Database) *App {
	if err := uploadEgg(scrapydURL, project, version, eggPath); err != nil {
		panic("Error uploading egg: " + err.Error())
	}
	return &App{
		Server: http.Server{
			Addr:    addr,
			Handler: handler,
		},
		Database: database,
	}
}
