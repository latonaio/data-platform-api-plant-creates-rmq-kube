package dpfm_api_output_formatter

type General struct {
	BusinessPartner      int    `json:"BusinessPartner"`
	Plant                string `json:"Plant"`
	PlantFullName        string `json:"PlantFullName"`
	PlantName            string `json:"PlantName"`
	Language             string `json:"Language"`
	CreationDate         string `json:"CreationDate"`
	CreationTime         string `json:"CreationTime"`
	LastChangeDate       string `json:"LastChangeDate"`
	LastChangeTime       string `json:"LastChangeTime"`
	PlantFoundationDate  string `json:"PlantFoundationDate"`
	PlantLiquidationDate string `json:"PlantLiquidationDate"`
	SearchTerm1          string `json:"SearchTerm1"`
	SearchTerm2          string `json:"SearchTerm2"`
	PlantDeathDate       string `json:"PlantDeathDate"`
	PlantIsBlocked       bool   `json:"PlantIsBlocked"`
	GroupPlantName1      string `json:"GroupPlantName1"`
	GroupPlantName2      string `json:"GroupPlantName2"`
	AddressID            int    `json:"AddressID"`
	Country              string `json:"Country"`
	TimeZone             string `json:"TimeZone"`
	PlantIDByExtSystem   string `json:"PlantIDByExtSystem"`
	IsMarkedForDeletion  bool   `json:"IsMarkedForDeletion"`
}

type StorageLocation struct {
	BusinessPartner              int    `json:"BusinessPartner"`
	Plant                        string `json:"Plant"`
	StorageLocation              string `json:"StorageLocation"`
	StorageLocationFullName      string `json:"StorageLocationFullName"`
	StorageLocationName          string `json:"StorageLocationName"`
	CreationDate                 string `json:"CreationDate"`
	CreationTime                 string `json:"CreationTime"`
	LastChangeDate               string `json:"LastChangeDate"`
	LastChangeTime               string `json:"LastChangeTime"`
	SearchTerm1                  string `json:"SearchTerm1"`
	SearchTerm2                  string `json:"SearchTerm2"`
	StorageLocationIsBlocked     bool   `json:"StorageLocationIsBlocked"`
	GroupStorageLocationName1    string `json:"GroupStorageLocationName1"`
	GroupStorageLocationName2    string `json:"GroupStorageLocationName2"`
	StorageLocationIDByExtSystem string `json:"StorageLocationIDByExtSystem"`
	IsMarkedForDeletion          bool   `json:"IsMarkedForDeletion"`
}
