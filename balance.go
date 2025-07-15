package tcgaming

import (
	"encoding/json"
)

type BalanceResponse struct {
	Status    int      `json:"status"`
	Balance   float64  `json:"balance"`
	ErrorDesc *string  `json:"error_desc"`
}

func (c *Client) GetBalance(username string, productType int) (*BalanceResponse, error) {
	params := Request{
		Method:      "gb",
		Username:    username,
		ProductType: productType,
	}

	resp, err := c.sendRequest(params)
	if err != nil {
		return nil, err
	}

	if err := checkResponseError(resp); err != nil {
		return nil, err
	}

	// Return the exact API response
	return &BalanceResponse{
		Status:    resp.Status,
		Balance:   resp.Balance,
		ErrorDesc: resp.ErrorDesc,
	}, nil
}

const (
	FundTypeIn  = 1
	FundTypeOut = 2
)

type TransferResponse struct {
	Status            int     `json:"status"`
	ErrorDesc         *string `json:"error_desc"`
	TransactionStatus string  `json:"transaction_status"` // API returns "SUCCESS" or "FAILED"
}

func (c *Client) UserTransfer(username string, productType, fundType int, amount float64, referenceNo string) (*TransferResponse, error) {
	params := Request{
		Method:      "ft",
		Username:    username,
		ProductType: productType,
		FundType:    fundType,
		Amount:      amount,
		ReferenceNo: referenceNo,
	}

	resp, err := c.sendRequest(params)
	if err != nil {
		return nil, err
	}

	if err := checkResponseError(resp); err != nil {
		return nil, err
	}

	// Return the exact API response
	return &TransferResponse{
		Status:            resp.Status,
		ErrorDesc:         resp.ErrorDesc,
		TransactionStatus: resp.TransactionStatus,
	}, nil
}

type TransactionStatusResponse struct {
	Status             int             `json:"status"`
	ErrorDesc          *string         `json:"error_desc"`
	TransactionStatus  string          `json:"transaction_status"`
	TransactionDetails json.RawMessage `json:"transaction_details"`
}

func (c *Client) CheckTransfer(productType int, referenceNo string) (*TransactionStatusResponse, error) {
	params := Request{
		Method:      "cs",
		ProductType: productType,
		RefNo:       referenceNo,
	}

	resp, err := c.sendRequest(params)
	if err != nil {
		return nil, err
	}

	if err := checkResponseError(resp); err != nil {
		return nil, err
	}

	// Return the exact API response
	return &TransactionStatusResponse{
		Status:             resp.Status,
		ErrorDesc:          resp.ErrorDesc,
		TransactionStatus:  resp.TransactionStatus,
		TransactionDetails: resp.TransactionDetails,
	}, nil
}