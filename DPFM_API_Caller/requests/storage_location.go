package requests

type StorageLocation struct {
	BusinessPartner              *int    `json:"BusinessPartner"`
	Plant                        string  `json:"Plant"`
	StorageLocation              string  `json:"StorageLocation"`
	StorageLocationFullName      string  `json:"StorageLocationFullName"`
	StorageLocationName          string  `json:"StorageLocationName"`
	CreationDate                 *string `json:"CreationDate"`
	CreationTime                 *string `json:"CreationTime"`
	LastChangeDate               *string `json:"LastChangeDate"`
	LastChangeTime               *string `json:"LastChangeTime"`
	SearchTerm1                  string  `json:"SearchTerm1"`
	SearchTerm2                  string  `json:"SearchTerm2"`
	StorageLocationIsBlocked     *bool   `json:"StorageLocationIsBlocked"`
	GroupStorageLocationName1    string  `json:"GroupStorageLocationName1"`
	GroupStorageLocationName2    string  `json:"GroupStorageLocationName2"`
	StorageLocationIDByExtSystem string  `json:"StorageLocationIDByExtSystem"`
	IsMarkedForDeletion          *bool   `json:"IsMarkedForDeletion"`
}
