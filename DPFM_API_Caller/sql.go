package dpfm_api_caller

import (
	"context"
	dpfm_api_input_reader "data-platform-api-plant-creates-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_output_formatter "data-platform-api-plant-creates-rmq-kube/DPFM_API_Output_Formatter"
	dpfm_api_processing_formatter "data-platform-api-plant-creates-rmq-kube/DPFM_API_Processing_Formatter"
	"sync"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
	"golang.org/x/xerrors"
)

func (c *DPFMAPICaller) createSqlProcess(
	ctx context.Context,
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	accepter []string,
	errs *[]error,
	log *logger.Logger,
) interface{} {
	var general            *dpfm_api_output_formatter.General
	var storageLocation    *[]dpfm_api_output_formatter.StorageLocation
	for _, fn := range accepter {
		switch fn {
		case "General":
			general = c.generalCreateSql(nil, mtx, input, output, errs, log)
		case "StorageLocation":
			storageLocation = c.storageLocationCreateSql(nil, mtx, input, output, errs, log)
		default:

		}
	}

	data := &dpfm_api_output_formatter.Message{
		General:            general,
		StorageLocation:    storageLocation,
	}

	return data
}

func (c *DPFMAPICaller) updateSqlProcess(
	ctx context.Context,
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	accepter []string,
	errs *[]error,
	log *logger.Logger,
) interface{} {
	var general *dpfm_api_output_formatter.General
	var storageLocation *[]dpfm_api_output_formatter.StorageLocation
	for _, fn := range accepter {
		switch fn {
		case "General":
			general = c.generalUpdateSql(mtx, input, output, errs, log)
		case "StorageLocation":
			storageLocation = c.storageLocationUpdateSql(mtx, input, output, errs, log)
		default:

		}
	}

	data := &dpfm_api_output_formatter.Message{
		General:            general,
		StorageLocation:    storageLocation,
	}

	return data
}

func (c *DPFMAPICaller) generalCreateSql(
	ctx context.Context,
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *dpfm_api_output_formatter.General {
	if ctx == nil {
		ctx = context.Background()
	}
	sessionID := input.RuntimeSessionID
	generalData := input.General
	res, err := c.rmq.SessionKeepRequest(nil, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": generalData, "function": "PlantGeneral", "runtime_session_id": sessionID})
	if err != nil {
		err = xerrors.Errorf("rmq error: %w", err)
		return nil
	}
	res.Success()
	if !checkResult(res) {
		output.SQLUpdateResult = getBoolPtr(false)
		output.SQLUpdateError = "General Data cannot insert"
		return nil
	}

	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data, err := dpfm_api_output_formatter.ConvertToGeneralCreates(input)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) storageLocationCreateSql(
	ctx context.Context,
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.StorageLocation {
	if ctx == nil {
		ctx = context.Background()
	}
	sessionID := input.RuntimeSessionID
	for i := range input.General.StorageLocation {
		input.General.StorageLocation[i].Plant = input.General.Plant
		storageLocationData := input.General.StorageLocation[i]

		res, err := c.rmq.SessionKeepRequest(ctx, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": accountingData, "function": "PlantStorageLocation", "runtime_session_id": sessionID})
		if err != nil {
			err = xerrors.Errorf("rmq error: %w", err)
			return nil
		}
		res.Success()
		if !checkResult(res) {
			output.SQLUpdateResult = getBoolPtr(false)
			output.SQLUpdateError = "StorageLocation Data cannot insert"
			return nil
		}
	}

	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data, err := dpfm_api_output_formatter.ConvertToStorageLocationCreates(input)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) generalUpdateSql(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *dpfm_api_output_formatter.General {
	general := input.General
	generalData := dpfm_api_processing_formatter.ConvertToGeneralUpdates(general)

	sessionID := input.RuntimeSessionID
	if generalIsUpdate(generalData) {
		res, err := c.rmq.SessionKeepRequest(nil, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": generalData, "function": "PlantGeneral", "runtime_session_id": sessionID})
		if err != nil {
			err = xerrors.Errorf("rmq error: %w", err)
			*errs = append(*errs, err)
			return nil
		}
		res.Success()
		if !checkResult(res) {
			output.SQLUpdateResult = getBoolPtr(false)
			output.SQLUpdateError = "General Data cannot insert"
			return nil
		}
	}

	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data, err := dpfm_api_output_formatter.ConvertToGeneralUpdates(general)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) storageLocationUpdateSql(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.StorageLocation {
	req := make([]dpfm_api_processing_formatter.StorageLocationUpdates, 0)
	sessionID := input.RuntimeSessionID

	general := input.General
	for _, storageLocation := range general.StorageLocation {
		storageLocationData := *dpfm_api_processing_formatter.ConvertToStorageLocationUpdates(general, storageLocation)

		if storageLocationIsUpdate(&storageLocationData) {
			res, err := c.rmq.SessionKeepRequest(nil, c.conf.RMQ.QueueToSQL()[0], map[string]interface{}{"message": storageLocationData, "function": "PlantStorageLocation", "runtime_session_id": sessionID})
			if err != nil {
				err = xerrors.Errorf("rmq error: %w", err)
				*errs = append(*errs, err)
				return nil
			}
			res.Success()
			if !checkResult(res) {
				output.SQLUpdateResult = getBoolPtr(false)
				output.SQLUpdateError = "StorageLocation Data cannot update"
				return nil
			}
		}
		req = append(req, storageLocationData)
	}

	if output.SQLUpdateResult == nil {
		output.SQLUpdateResult = getBoolPtr(true)
	}

	data, err := dpfm_api_output_formatter.ConvertToStorageLocationUpdates(&req)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func generalIsUpdate(general *dpfm_api_processing_formatter.GeneralUpdates) bool {
	businessPartner := general.BusinessPartner
	plant := general.Plant

	return !(businessPartner == 0 || plant == "")
}

func storageLocationIsUpdate(storageLocation *dpfm_api_processing_formatter.StorageLocationUpdates) bool {
	businessPartner := storageLocation.BusinessPartner
	plant := storageLocation.Plant
	storageLocation := storageLocation.StorageLocation

	return !(businessPartner == 0 || plant == "" || storageLocation == "")
}
