package dpfm_api_processing_formatter

type GeneralUpdates struct {
	BusinessPartner      int    `json:"BusinessPartner"`
	Plant                string `json:"Plant"`
	PlantName            *string `json:"PlantName"`
	PlantFullName        *string `json:"PlantFullName"`
	Language             *string `json:"Language"`
	PlantFoundationDate  *string `json:"PlantFoundationDate"`
	PlantLiquidationDate *string `json:"PlantLiquidationDate"`
	PlantDeathDate       *string `json:"PlantDeathDate"`
	AddressID            *int    `json:"AddressID"`
	Country              *string `json:"Country"`
	TimeZone             *string `json:"TimeZone"`
	PlantIDByExtSystem   *string `json:"PlantIDByExtSystem"`
}

type StorageLocationUpdates struct {
	BusinessPartner              int    `json:"BusinessPartner"`
	Plant                        string `json:"Plant"`
	StorageLocation              string  `json:"StorageLocation"`
	StorageLocationFullName      *string `json:"StorageLocationFullName"`
	StorageLocationName          *string `json:"StorageLocationName"`
	StorageLocationIDByExtSystem *string `json:"StorageLocationIDByExtSystem"`
}
