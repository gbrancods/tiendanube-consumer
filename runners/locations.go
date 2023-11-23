package runners

import (
	"github.com/gbrancods/tiendanube/locations"
)

func LocationsRunner() {
	r := CreateLocation()
	GetAllLocations()
	UpdateLocation(r.ID)
	GetLocationsByID(r.ID)
	DeleteLocation(r.ID)
	GetInventoryLevels(r.ID)
}

func CreateLocation() (r locations.Location) {
	r, err := locations.New(
		locations.SetName(
			locations.Name{
				EnUS: "Name",
				EsAR: "Nombre",
				PtBR: "Nome",
			},
		),
		locations.SetAddress(
			locations.Address{
				Zipcode:        "15503205",
				Street:         "César Waldemar Caldorin",
				Number:         "8888",
				Floor:          "AP3",
				Locality:       "Locality",
				Reference:      "REF12",
				BetweenStreets: "Between the streets",
				City:           "Votuporanga",
				Province: locations.Province{
					Code: "SP",
					Name: "São Paulo",
				},
				Region: locations.Region{
					Code: "SE",
					Name: "Sudeste",
				},
				Country: locations.Country{
					Code: "BR",
					Name: "Brasil",
				},
			},
		),
	).Create()

	if err != nil {
		panicErr("Error creating location", err)
	}

	prettyPrint("Location created with success!", r)
	return
}

func UpdateLocation(locationID string) {

	c, err := locations.New(
		locations.SetName(
			locations.Name{
				EnUS: "Name",
				EsAR: "Nombre",
				PtBR: "Nome",
			},
		),
		locations.SetAddress(
			locations.Address{
				Zipcode:        "15503205",
				Street:         "César Waldemar Caldorin",
				Number:         "7777",
				Floor:          "AP3",
				Locality:       "Locality",
				Reference:      "REF12",
				BetweenStreets: "Between the streets",
				City:           "Votuporanga",
				Province: locations.Province{
					Code: "SP",
					Name: "São Paulo",
				},
				Region: locations.Region{
					Code: "SE",
					Name: "Sudeste",
				},
				Country: locations.Country{
					Code: "BR",
					Name: "Brasil",
				},
			},
		),
	).Update(locationID)

	if err != nil {
		panicErr("Error updating location", err)
	}

	prettyPrint("Location updated!", c)
}

func GetAllLocations() {
	r, err := locations.GetAll()
	if err != nil {
		panicErr("Error getting all locations ", err)
	}

	prettyPrint("Locations:", r)
}

func GetLocationsByID(locationID string) {
	r, err := locations.GetByID(locationID)
	if err != nil {
		panicErr("Error getting all locations ", err)
	}

	prettyPrint("Location:", r)
}

func DeleteLocation(locationID string) {
	err := locations.Delete(locationID)
	if err != nil {
		panicErr("Error getting all locations ", err)
	}

	prettyPrint("Location deleted!", locationID)
}

func GetInventoryLevels(locationID string) {
	r, err := locations.GetInventoryLevels(locationID)
	if err != nil {
		panicErr("Error getting inventory levels ", err)
	}

	prettyPrint("Inventory levels:", r)
}
