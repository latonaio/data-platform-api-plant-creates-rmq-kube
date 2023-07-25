package requests

type General struct {
	BusinessPartner      int     `json:"BusinessPartner"`
	Plant                string  `json:"Plant"`
	PlantName            string  `json:"PlantName"`
	PlantFullName        *string `json:"PlantFullName"`
	Language             string  `json:"Language"`
	PlantFoundationDate  *string `json:"PlantFoundationDate"`
	PlantLiquidationDate *string `json:"PlantLiquidationDate"`
	PlantDeathDate       *string `json:"PlantDeathDate"`
	AddressID            *int    `json:"AddressID"`
	Country              string  `json:"Country"`
	TimeZone             *string `json:"TimeZone"`
	PlantIDByExtSystem   *string `json:"PlantIDByExtSystem"`
	CreationDate         string  `json:"CreationDate"`
	LastChangeDate       string  `json:"LastChangeDate"`
	IsMarkedForDeletion  *bool   `json:"IsMarkedForDeletion"`
}
