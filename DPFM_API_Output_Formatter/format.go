package dpfm_api_output_formatter

import (
	dpfm_api_input_reader "data-platform-api-plant-creates-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_processing_formatter "data-platform-api-plant-creates-rmq-kube/DPFM_API_Processing_Formatter"
	"encoding/json"
	"time"

	"golang.org/x/xerrors"
)

func ConvertToGeneralCreates(sdc *dpfm_api_input_reader.SDC) (*General, error) {
	data := sdc.General

	general, err := TypeConverter[*General](data)
	if err != nil {
		return nil, err
	}
	// general.CreationDate = *getSystemDatePtr()
	// general.CreationTime = *getSystemTimePtr()
	// general.LastChangeDate = getSystemDatePtr()
	// general.LastChangeTime = getSystemTimePtr()

	return general, nil
}

func ConvertToStorageLocationCreates(sdc *dpfm_api_input_reader.SDC) (*StorageLocation, error) {
	data := sdc.StorageLocation

	storageLocation, err := TypeConverter[*StorageLocation](data)
	if err != nil {
		return nil, err
	}

	return storageLocation, nil
}

func ConvertToGeneralUpdates(generalData dpfm_api_input_reader.General) (*General, error) {
	data := generalData

	general, err := TypeConverter[*General](data)
	if err != nil {
		return nil, err
	}

	return general, nil
}

func ConvertToStorageLocationUpdates(generalData dpfm_api_input_reader.StorageLocation) (*StorageLocation, error) {
	data := storageLocationData

	storageLocation, err := TypeConverter[*StorageLocation](data)
	if err != nil {
		return nil, err
	}

	return storageLocation, nil
}

func TypeConverter[T any](data interface{}) (T, error) {
	var dist T
	b, err := json.Marshal(data)
	if err != nil {
		return dist, xerrors.Errorf("Marshal error: %w", err)
	}
	err = json.Unmarshal(b, &dist)
	if err != nil {
		return dist, xerrors.Errorf("Unmarshal error: %w", err)
	}
	return dist, nil
}

func getSystemDatePtr() *string {
	// jst, _ := time.LoadLocation("Asia/Tokyo")
	// day := time.Now().In(jst)

	day := time.Now()
	res := day.Format("2006-01-02")
	return &res
}

func getSystemTimePtr() *string {
	// jst, _ := time.LoadLocation("Asia/Tokyo")
	// day := time.Now().In(jst)

	day := time.Now()
	res := day.Format("15:04:05")
	return &res
}
