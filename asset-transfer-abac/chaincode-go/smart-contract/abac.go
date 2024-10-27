package abac

package chaincode

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/v2/contractapi"
)

// SmartContract provides functions for managing an Asset
type SmartContract struct {
	contractapi.Contract
}

// Asset describes basic details of what makes up a simple asset
// Insert struct field in alphabetic order => to achieve determinism across languages
type Asset struct {
	BALANCE     float64 `json:"balance"`
	DEALERID    string  `json:"dealerid"`
	ID          string  `json:"ID"`
	MPIN        string  `json:"mpin"`
	MSISDN      string  `json:"msisdn"`
	REMARKS     string  `json:"remarks"`
	STATUS      string  `json:"status"`
	TRANSAMOUNT float64 `json:"transamount"`
	TRANSTYPE   string  `json:"transtype"`
}

// InitLedger adds a base set of assets to the ledger
func (s *SmartContract) InitLedger(ctx contractapi.TransactionContextInterface) error {
	assets := []Asset{
		{ID: "asset1", DEALERID: "DEALER101", MSISDN: "9877890123", MPIN: "1598", BALANCE: 100000.00, STATUS: "ACTIVE", TRANSAMOUNT: 100000.00, TRANSTYPE: "CREDIT", REMARKS: "Personal loan disbursement"},
		{ID: "asset2", DEALERID: "DEALER102", MSISDN: "9811234567", MPIN: "4321", BALANCE: 500.00, STATUS: "ACTIVE", TRANSAMOUNT: 500.00, TRANSTYPE: "INIT", REMARKS: "New account creation"},
		{ID: "asset3", DEALERID: "DEALER103", MSISDN: "9876543212", MPIN: "9012", BALANCE: 1500.00, STATUS: "ACTIVE", TRANSAMOUNT: 200.00, TRANSTYPE: "DEBIT", REMARKS: "Purchase transaction"},
		{ID: "asset4", DEALERID: "DEALER104", MSISDN: "9822345678", MPIN: "8765", BALANCE: 25000.00, STATUS: "ACTIVE", TRANSAMOUNT: 25000.00, TRANSTYPE: "CREDIT", REMARKS: "Business investment deposit"},
		{ID: "asset5", DEALERID: "DEALER105", MSISDN: "9844567890", MPIN: "1357", BALANCE: 0.00, STATUS: "INACTIVE", TRANSAMOUNT: 0.00, TRANSTYPE: "SUSPEND", REMARKS: "Account dormant - no activity for 6 months"},
		{ID: "asset6", DEALERID: "DEALER106", MSISDN: "9866789012", MPIN: "3579", BALANCE: 12000.00, STATUS: "ACTIVE", TRANSAMOUNT: 3000.00, TRANSTYPE: "DEBIT", REMARKS: "Electricity bill payment"},
		{ID: "asset7", DEALERID: "DEALER107", MSISDN: "9877890123", MPIN: "1598", BALANCE: 100000.00, STATUS: "ACTIVE", TRANSAMOUNT: 100000.00, TRANSTYPE: "CREDIT", REMARKS: "Personal loan disbursement"},
	}

	for _, asset := range assets {
		assetJSON, err := json.Marshal(asset)
		if err != nil {
			return err
		}

		err = ctx.GetStub().PutState(asset.ID, assetJSON)
		if err != nil {
			return fmt.Errorf("failed to put to world state: %v", err)
		}
	}

	return nil
}

// CreateAsset issues a new asset to the world state with given details.
func (s *SmartContract) CreateAsset(ctx contractapi.TransactionContextInterface, id string, dealerID string, msisdn string, mpin string, balance float64, status string, transAmount float64, transType string, remarks string) error {
	exists, err := s.AssetExists(ctx, id)
	if err != nil {
		return err
	}
	if exists {
		return fmt.Errorf("the asset %s already exists", id)
	}

	asset := Asset{
		ID:          id,
		DEALERID:    dealerID,
		MSISDN:      msisdn,
		MPIN:        mpin,
		BALANCE:     balance,
		STATUS:      status,
		TRANSAMOUNT: transAmount,
		TRANSTYPE:   transType,
		REMARKS:     remarks,
	}
	assetJSON, err := json.Marshal(asset)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(id, assetJSON)
}

// ReadAsset returns the asset stored in the world state with given id.
func (s *SmartContract) ReadAsset(ctx contractapi.TransactionContextInterface, id string) (*Asset, error) {
	assetJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return nil, fmt.Errorf("failed to read from world state: %v", err)
	}
	if assetJSON == nil {
		return nil, fmt.Errorf("the asset %s does not exist", id)
	}

	var asset Asset
	err = json.Unmarshal(assetJSON, &asset)
	if err != nil {
		return nil, err
	}

	return &asset, nil
}

// UpdateAsset updates an existing asset in the world state with provided parameters.
func (s *SmartContract) UpdateAsset(ctx contractapi.TransactionContextInterface, id string, dealerID string, msisdn string, mpin string, balance float64, status string, transAmount float64, transType string, remarks string) error {
	exists, err := s.AssetExists(ctx, id)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("the asset %s does not exist", id)
	}

	// overwriting original asset with new asset
	asset := Asset{
		ID:          id,
		DEALERID:    dealerID,
		MSISDN:      msisdn,
		MPIN:        mpin,
		BALANCE:     balance,
		STATUS:      status,
		TRANSAMOUNT: transAmount,
		TRANSTYPE:   transType,
		REMARKS:     remarks,
	}
	assetJSON, err := json.Marshal(asset)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(id, assetJSON)
}

// DeleteAsset deletes a given asset from the world state.
func (s *SmartContract) DeleteAsset(ctx contractapi.TransactionContextInterface, id string) error {
	exists, err := s.AssetExists(ctx, id)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("the asset %s does not exist", id)
	}

	return ctx.GetStub().DelState(id)
}

// AssetExists returns true when asset with given ID exists in world state
func (s *SmartContract) AssetExists(ctx contractapi.TransactionContextInterface, id string) (bool, error) {
	assetJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return false, fmt.Errorf("failed to read from world state: %v", err)
	}

	return assetJSON != nil, nil
}

// TransferAsset updates the DEALERID field of the asset with the given id in the world state.
func (s *SmartContract) TransferAsset(ctx contractapi.TransactionContextInterface, id string, newDealerID string) (string, error) {
	asset, err := s.ReadAsset(ctx, id)
	if err != nil {
		return "", err
	}

	oldDealerID := asset.DEALERID
	asset.DEALERID = newDealerID

	assetJSON, err := json.Marshal(asset)
	if err != nil {
		return "", err
	}

	err = ctx.GetStub().PutState(id, assetJSON)
	if err != nil {
		return "", err
	}

	return oldDealerID, nil
}

// GetAllAssets returns all assets found in world state
func (s *SmartContract) GetAllAssets(ctx contractapi.TransactionContextInterface) ([]*Asset, error) {
	resultsIterator, err := ctx.GetStub().GetStateByRange("", "")
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	var assets []*Asset
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		var asset Asset
		err = json.Unmarshal(queryResponse.Value, &asset)
		if err != nil {
			return nil, err
		}
		assets = append(assets, &asset)
	}

	return assets, nil
}
