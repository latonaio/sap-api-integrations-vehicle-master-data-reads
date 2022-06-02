package sap_api_output_formatter

import (
	"encoding/json"
	"sap-api-integrations-vehicle-master-data-reads/SAP_API_Caller/responses"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
	"golang.org/x/xerrors"
)

func ConvertToVehicleCollection(raw []byte, l *logger.Logger) ([]VehicleCollection, error) {
	pm := &responses.VehicleCollection{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to VehicleCollection. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	vehicleCollection := make([]VehicleCollection, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		vehicleCollection = append(vehicleCollection, VehicleCollection{
			ObjectID:                  data.ObjectID,
			CustomerID:                data.CustomerID,
			ProductUUID:               data.ProductUUID,
			VIN:                       data.VIN,
			Description:               data.Description,
			LicensePlateNumber:        data.LicensePlateNumber,
			ModelYear:                 data.ModelYear,
			OdometerReading:           data.OdometerReading,
			TUVValidity:               data.TUVValidity,
			ID:                        data.ID,
			RegisteredProductCategory: data.RegisteredProductCategory,
			ProductID:                 data.ProductID,
			ModelCatalog:              data.ModelCatalog,
			ModelSalesCode:            data.ModelSalesCode,
			DeliveryDate:              data.DeliveryDate,
			PurchaseDate:              data.PurchaseDate,
			RegistrationDate:          data.RegistrationDate,
			IndividualProductUUID:     data.IndividualProductUUID,
			InstallationPointUUID:     data.InstallationPointUUID,
			ChassisNumber:             data.ChassisNumber,
			CSNumber:                  data.CSNumber,
			EngineCapacity:            data.EngineCapacity,
			EngineCapacityUnitCode:    data.EngineCapacityUnitCode,
			EngineNumber:              data.EngineNumber,
			FuelType:                  data.FuelType,
			EntityLastChangedOn:       data.EntityLastChangedOn,
			ETag:                      data.ETag,
		})
	}

	return vehicleCollection, nil
}
