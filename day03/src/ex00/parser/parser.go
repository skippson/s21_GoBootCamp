package parser

import (
	"day03/types"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/grailbio/base/tsv"
	"github.com/olivere/elastic"
)
type Pars types.Place

func (p Pars) readDataSet(filePath string) ([][]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := tsv.NewReader(file)
	data, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}
	fmt.Println("data", len(data))
	return data, nil
}

func (p Pars) createPlace(data []string) (types.Place, error) {
	var place types.Place
	id, err := strconv.Atoi(data[0])
	if err != nil {
		if data[0] == "" {
			log.Println(data)
		}
		return place, err
	}
	lon, err := strconv.ParseFloat(data[4], 64)
	if err != nil {
		return place, err
	}
	lat, err := strconv.ParseFloat(data[5], 64)
	if err != nil {
		return place, err
	}
	place = types.Place{
		ID:       id,
		Name:     data[1],
		Address:   data[2],
		Phone:    data[3],
		Location: elastic.GeoPoint{Lat: lat, Lon: lon},
	}
	return place, err

}
func (p Pars) ParseCSV(places []*types.Place) ([]*types.Place, error) {

	str_places, err := p.readDataSet("../../materials/data.csv")
	if err != nil {
		return nil, err
	}
	for i := 1; i < len(str_places); i++ {
		new_place, err := p.createPlace(str_places[i])
		if err != nil {
			return nil, err
		}
		places = append(places, &new_place)
	}
	fmt.Println(len(places))
	return places, nil
}
