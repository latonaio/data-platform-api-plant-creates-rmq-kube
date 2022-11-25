package dpfm_api_input_reader

type EC_MC struct {
}

type SDC struct {
	ConnectionKey       string   `json:"connection_key"`
	Result              bool     `json:"result"`
	RedisKey            string   `json:"redis_key"`
	Filepath            string   `json:"filepath"`
	APIStatusCode       int      `json:"api_status_code"`
	RuntimeSessionID    string   `json:"runtime_session_id"`
	BusinessPartnerID   *int     `json:"business_partner"`
	ServiceLabel        string   `json:"service_label"`
	Plant               Plant    `json:"Plant"`
	APISchema           string   `json:"api_schema"`
	Accepter            []string `json:"accepter"`
	OrderID             *int     `json:"order_id"`
	Deleted             bool     `json:"deleted"`
	SQLUpdateResult     *bool    `json:"sql_update_result"`
	SQLUpdateError      string   `json:"sql_update_error"`
	SubfuncResult       *bool    `json:"subfunc_result"`
	SubfuncError        string   `json:"subfunc_error"`
	ExconfResult        *bool    `json:"exconf_result"`
	ExconfError         string   `json:"exconf_error"`
	APIProcessingResult *bool    `json:"api_processing_result"`
	APIProcessingError  string   `json:"api_processing_error"`
}

type Plant struct {
	BusinessPartner      *int    `json:"BusinessPartner"`
	Plant                *string `json:"Plant"`
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
