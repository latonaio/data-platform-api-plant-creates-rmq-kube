package requests

type General struct {
	BusinessPartner      *int    `json:"BusinessPartner"`
	Plant                string  `json:"Plant"`
	PlantFullName        string  `json:"PlantFullName"`
	PlantName            *string `json:"PlantName"`
	Language             *string `json:"Language"`
	CreationDate         *string `json:"CreationDate"`
	CreationTime         *string `json:"CreationTime"`
	LastChangeDate       *string `json:"LastChangeDate"`
	LastChangeTime       *string `json:"LastChangeTime"`
	PlantFoundationDate  string  `json:"PlantFoundationDate"`
	PlantLiquidationDate string  `json:"PlantLiquidationDate"`
	SearchTerm1          string  `json:"SearchTerm1"`
	SearchTerm2          string  `json:"SearchTerm2"`
	PlantDeathDate       string  `json:"PlantDeathDate"`
	PlantIsBlocked       bool    `json:"PlantIsBlocked"`
	GroupPlantName1      string  `json:"GroupPlantName1"`
	GroupPlantName2      string  `json:"GroupPlantName2"`
	AddressID            int     `json:"AddressID"`
	Country              string  `json:"Country"`
	TimeZone             string  `json:"TimeZone"`
	PlantIDByExtSystem   string  `json:"PlantIDByExtSystem"`
	IsMarkedForDeletion  bool    `json:"IsMarkedForDeletion"`
}
