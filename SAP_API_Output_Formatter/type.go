package sap_api_output_formatter

type VehicleCollection struct {
	ObjectID                  string `json:"ObjectID"`
	CustomerID                string `json:"CustomerID"`
	ProductUUID               string `json:"ProductUUID"`
	VIN                       string `json:"VIN"`
	Description               string `json:"Description"`
	LicensePlateNumber        string `json:"LicensePlateNumber"`
	ModelYear                 string `json:"ModelYear"`
	OdometerReading           string `json:"OdometerReading"`
	TUVValidity               string `json:"TUVValidity"`
	ID                        string `json:"ID"`
	RegisteredProductCategory string `json:"RegisteredProductCategory"`
	ProductID                 string `json:"ProductID"`
	ModelCatalog              string `json:"ModelCatalog"`
	ModelSalesCode            string `json:"ModelSalesCode"`
	DeliveryDate              string `json:"DeliveryDate"`
	PurchaseDate              string `json:"PurchaseDate"`
	RegistrationDate          string `json:"RegistrationDate"`
	IndividualProductUUID     string `json:"IndividualProductUUID"`
	InstallationPointUUID     string `json:"InstallationPointUUID"`
	ChassisNumber             string `json:"ChassisNumber"`
	CSNumber                  string `json:"CSNumber"`
	EngineCapacity            string `json:"EngineCapacity"`
	EngineCapacityUnitCode    string `json:"EngineCapacityUnitCode"`
	EngineNumber              string `json:"EngineNumber"`
	FuelType                  string `json:"FuelType"`
	EntityLastChangedOn       string `json:"EntityLastChangedOn"`
	ETag                      string `json:"ETag"`
}
