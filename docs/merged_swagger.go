package docs

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type SwaggerDoc struct {
	Swagger     string                   `json:"swagger"`
	Info        map[string]interface{}   `json:"info"`
	Paths       map[string]interface{}   `json:"paths"`
	Definitions map[string]interface{}   `json:"definitions,omitempty"`
	Tags        []map[string]interface{} `json:"tags,omitempty"`
}

func GenerateDocs() error {
	mergedDoc := SwaggerDoc{
		Swagger:     "2.0",
		Info:        make(map[string]interface{}),
		Paths:       make(map[string]interface{}),
		Definitions: make(map[string]interface{}),
		Tags:        make([]map[string]interface{}, 0),
	}

	currentDir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("failed to get current directory: %w", err)
	}

	apiDir := filepath.Join(currentDir, "docs", "api")

	fmt.Printf("Searching for swagger files in: %s\n", apiDir)

	pattern := filepath.Join(apiDir, "*.swagger.json")
	files, err := filepath.Glob(pattern)
	if err != nil {
		return fmt.Errorf("failed to glob files in %s: %w", pattern, err)
	}

	if len(files) == 0 {
		return fmt.Errorf("no swagger files found in %s", pattern)
	}

	fmt.Printf("Found swagger files: %v\n", files)

	firstFile, err := os.ReadFile(files[0])
	if err != nil {
		return fmt.Errorf("failed to read first file: %w", err)
	}

	var firstDoc SwaggerDoc
	if err := json.Unmarshal(firstFile, &firstDoc); err != nil {
		return fmt.Errorf("failed to unmarshal first file: %w", err)
	}
	mergedDoc.Info = firstDoc.Info

	for _, file := range files {
		data, err := os.ReadFile(file)
		if err != nil {
			return fmt.Errorf("failed to read file %s: %w", file, err)
		}

		var doc SwaggerDoc
		if err := json.Unmarshal(data, &doc); err != nil {
			return fmt.Errorf("failed to unmarshal file %s: %w", file, err)
		}

		for path, pathData := range doc.Paths {
			mergedDoc.Paths[path] = pathData
		}

		if doc.Definitions != nil {
			for name, definition := range doc.Definitions {
				mergedDoc.Definitions[name] = definition
			}
		}

		if doc.Tags != nil {
			mergedDoc.Tags = append(mergedDoc.Tags, doc.Tags...)
		}
	}

	mergedJSON, err := json.MarshalIndent(mergedDoc, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal merged doc: %w", err)
	}

	outputPath := filepath.Join(apiDir, "swagger.json")

	if err := os.MkdirAll(filepath.Dir(outputPath), 0o755); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}

	if err := os.WriteFile(outputPath, mergedJSON, 0o644); err != nil {
		return fmt.Errorf("failed to write merged file to %s: %w", outputPath, err)
	}

	fmt.Printf("Successfully generated swagger.json at: %s\n", outputPath)
	return nil
}
