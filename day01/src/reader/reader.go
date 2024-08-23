package reader

import (
	"encoding/json"
	"encoding/xml"
	"os"
	models "secondDay/models"
)

type DB models.Recipes

func (f DB) DBReader(filename string, format string) ([]byte, *models.Recipes, error) {
	var recipe models.Recipes
	fileCont, err := os.ReadFile(filename)
	if err != nil {
		return nil, &recipe, err
	}
	var data []byte
	if format == "json" {
		if err = json.Unmarshal(fileCont, &recipe); err != nil {
			return nil, &recipe, err
		}
		data, err = xml.MarshalIndent(recipe, "", "    ")
		if err != nil {
			return nil, &recipe, err
		}
	} else {
		if err = xml.Unmarshal(fileCont, &recipe); err != nil {
			return nil, &recipe, err
		}
		data, err = json.MarshalIndent(recipe, "", "    ")
		if err != nil {
			return nil, &recipe, err
		}
	}
	return data, &recipe, nil
}