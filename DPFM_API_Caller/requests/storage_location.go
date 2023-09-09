package requests

type StorageLocation struct {
	BusinessPartner              int     `json:"BusinessPartner"`
	Plant                        string  `json:"Plant"`
	StorageLocation              string  `json:"StorageLocation"`
	StorageLocationFullName      *string `json:"StorageLocationFullName"`
	StorageLocationName          string  `json:"StorageLocationName"`
	StorageLocationIDByExtSystem *string `json:"StorageLocationIDByExtSystem"`
	CreationDate                 string  `json:"CreationDate"`
	LastChangeDate               string  `json:"LastChangeDate"`
	IsMarkedForDeletion          *bool   `json:"IsMarkedForDeletion"`
}
