package dpfm_api_input_reader

import (
	"data-platform-api-plant-creates-rmq-kube/DPFM_API_Caller/requests"
)

func (sdc *SDC) ConvertToPlant() *requests.Plant {
	data := sdc.Plant
	return &requests.Plant{
		BusinessPartner:      data.BusinessPartner,
		Plant:                data.Plant,
		PlantFullName:        data.PlantFullName,
		PlantName:            data.PlantName,
		Language:             data.Language,
		CreationDate:         data.CreationDate,
		CreationTime:         data.CreationTime,
		LastChangeDate:       data.LastChangeDate,
		LastChangeTime:       data.LastChangeTime,
		PlantFoundationDate:  data.PlantFoundationDate,
		PlantLiquidationDate: data.PlantLiquidationDate,
		SearchTerm1:          data.SearchTerm1,
		SearchTerm2:          data.SearchTerm2,
		PlantDeathDate:       data.PlantDeathDate,
		PlantIsBlocked:       data.PlantIsBlocked,
		GroupPlantName1:      data.GroupPlantName1,
		GroupPlantName2:      data.GroupPlantName2,
		AddressID:            data.AddressID,
		Country:              data.Country,
		TimeZone:             data.TimeZone,
		PlantIDByExtSystem:   data.PlantIDByExtSystem,
		IsMarkedForDeletion:  data.IsMarkedForDeletion,
	}

}
