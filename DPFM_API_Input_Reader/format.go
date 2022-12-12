package dpfm_api_input_reader

import (
	"data-platform-api-plant-creates-rmq-kube/DPFM_API_Caller/requests"
)

func (sdc *SDC) ConvertToGeneral() *requests.General {
	data := sdc.General
	return &requests.General{
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

func (sdc *SDC) ConvertToStorageLocation() *requests.StorageLocation {
	dataGeneral := sdc.General
	data := sdc.General.StorageLocation
	return &requests.StorageLocation{
		BusinessPartner:              dataGeneral.BusinessPartner,
		Plant:                        dataGeneral.Plant,
		StorageLocation:              data.StorageLocation,
		StorageLocationFullName:      data.StorageLocationFullName,
		StorageLocationName:          data.StorageLocationName,
		CreationDate:                 data.CreationDate,
		CreationTime:                 data.CreationTime,
		LastChangeDate:               data.LastChangeDate,
		LastChangeTime:               data.LastChangeTime,
		SearchTerm1:                  data.SearchTerm1,
		SearchTerm2:                  data.SearchTerm2,
		StorageLocationIsBlocked:     data.StorageLocationIsBlocked,
		GroupStorageLocationName1:    data.GroupStorageLocationName1,
		GroupStorageLocationName2:    data.GroupStorageLocationName2,
		StorageLocationIDByExtSystem: data.StorageLocationIDByExtSystem,
		IsMarkedForDeletion:          data.IsMarkedForDeletion,
	}
}
