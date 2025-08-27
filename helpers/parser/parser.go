package parser

import (
	"encoding/json"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Resource struct {
	Type  string `yaml:"type"`
	Count int    `yaml:"count"`
}

type Application struct {
	Resources []Resource `yaml:"resources"`
}

type Inventory struct {
	Applications map[string]Application `yaml:"applications"`
}

// AppResource holds an application name and its resource
type AppResource struct {
	ApplicationName string
	ResourceType    string
	ResourceCount   int
}

// Convert to the format Terraform expects
type TfResource struct {
	ApplicationName string `json:"application_name"`
	ResourceType    string `json:"resource_type"`
	ResourceCount   int    `json:"resource_count"`
}

// ParseInventoryAll parses inventory.yaml and returns all applications and their resources
func ParseInventoryAll(filePath string) ([]AppResource, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open inventory file: %w", err)
	}
	defer file.Close()

	var inv Inventory
	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&inv); err != nil {
		return nil, fmt.Errorf("failed to decode yaml: %w", err)
	}

	var results []AppResource
	for appName, app := range inv.Applications {
		for _, res := range app.Resources {
			results = append(results, AppResource{
				ApplicationName: appName,
				ResourceType:    res.Type,
				ResourceCount:   res.Count,
			})
		}
	}
	if len(results) == 0 {
		return nil, fmt.Errorf("no resources found in inventory")
	}
	return results, nil
}

func ToTerraform(filePath string) string {
	resources, err := ParseInventoryAll(filePath)
	if err != nil {
		return ""
	}

	var tfResources []TfResource

	for _, r := range resources {
		tfResources = append(tfResources, TfResource{
			ApplicationName: r.ApplicationName,
			ResourceType:    r.ResourceType,
			ResourceCount:   r.ResourceCount,
		})
	}
	jsonBytes, _ := json.Marshal(tfResources)
	jsonStr := string(jsonBytes)
	return jsonStr
}
