package dpfm_api_processing_formatter

import (
	dpfm_api_input_reader "data-platform-api-plant-creates-rmq-kube/DPFM_API_Input_Reader"
)

func ConvertToGeneralUpdates(general dpfm_api_input_reader.General) *GeneralUpdates {
	data := general

	return &GeneralUpdates{
			BusinessPartner:      data.BusinessPartner,
			Plant:                data.Plant,
			PlantFullName:        data.PlantFullName,
			PlantName:            data.PlantName,
			Language:             data.Language,
			PlantFoundationDate:  data.PlantFoundationDate,
			PlantLiquidationDate: data.PlantLiquidationDate,
			PlantDeathDate:       data.PlantDeathDate,
			AddressID:            data.AddressID,
			Country:              data.Country,
			TimeZone:             data.TimeZone,
			PlantIDByExtSystem:   data.PlantIDByExtSystem,
	}
}

func ConvertToStorageLoationUpdates(general dpfm_api_input_reader.General, accounting dpfm_api_input_reader.StorageLocation) *StorageLocationUpdates {
	dataGeneral := general
	data := storageLocation

	return &StorageLocationUpdates{
		    BusinessPartner:              dataGeneral.BusinessPartner,
		    Plant:                        dataGeneral.Plant,
		    StorageLocation:              data.StorageLocation,
		    StorageLocationFullName:      data.StorageLocationFullName,
		    StorageLocationName:          data.StorageLocationName,
		    StorageLocationIDByExtSystem: data.StorageLocationIDByExtSystem,
	}
}
