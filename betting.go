package tcgaming

import (
	"encoding/json"
	"fmt"

	"github.com/goforj/godump"
)

type BetDetail struct {
	Username          string      `json:"username"`
	BetAmount         float64     `json:"betAmount"`
	ValidBetAmount    float64     `json:"validBetAmount"`
	WinAmount         float64     `json:"winAmount"`
	NetPnl            float64     `json:"netPnl"`
	Currency          string      `json:"currency"`
	TransactionTime   string      `json:"transactionTime"`
	GameCode          string      `json:"gameCode"`
	GameName          string      `json:"gameName"` // Live游戏也返回gameName
	BetOrderNo        string      `json:"betOrderNo"`
	BetTime           string      `json:"betTime"`
	ProductType       int         `json:"productType"`
	GameCategory      string      `json:"gameCategory"`
	SessionId         string      `json:"sessionId"`
	AdditionalDetails interface{} `json:"additionalDetails"`
}

// RNGBetDetail represents bet details for RNG and FISH games
// RNG和捕鱼游戏的投注详情结构
type RNGBetDetail struct {
	Username          string      `json:"username"`
	BetAmount         float64     `json:"betAmount"`
	ValidBetAmount    float64     `json:"validBetAmount"`
	WinAmount         float64     `json:"winAmount"`
	NetPnl            float64     `json:"netPnl"`
	Currency          string      `json:"currency"`
	TransactionTime   string      `json:"transactionTime"`
	GameCode          string      `json:"gameCode"`
	GameName          string      `json:"gameName"` // RNG/FISH特有字段
	BetOrderNo        string      `json:"betOrderNo"`
	BetTime           string      `json:"betTime"`
	ProductType       int         `json:"productType"`
	GameCategory      string      `json:"gameCategory"`
	SessionId         string      `json:"sessionId"`
	AdditionalDetails interface{} `json:"additionalDetails"`
}

// RNGBetDetailsResponse represents the response structure for RNG/FISH bet details
// RNG/FISH投注详情响应结构
type RNGBetDetailsResponse struct {
	Status    int            `json:"status"`
	Details   []RNGBetDetail `json:"details"`
	ErrorDesc *string        `json:"error_desc"`
	PageInfo  struct {
		TotalPage   int `json:"totalPage"`
		CurrentPage int `json:"currentPage"`
		TotalCount  int `json:"totalCount"`
	} `json:"page_info"`
}

type BetDetailsResponse struct {
	BatchName   string      `json:"batch_name"`
	BetDetails  []BetDetail `json:"bet_details"`
	TotalRecord int         `json:"total_record"`
	Page        int         `json:"page"`
}

func (c *Client) GetBetDetails(batchName string, page int) (*BetDetailsResponse, error) {
	params := Request{
		Method:    "bd",
		BatchName: batchName,
		Page:      page,
	}

	resp, err := c.sendRequest(params)
	if err != nil {
		return nil, err
	}

	if err := checkResponseError(resp); err != nil {
		return nil, err
	}

	var result BetDetailsResponse
	if err := unmarshalResponseData(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

type LiveBetDetailsResponse struct {
	Status    int         `json:"status"`
	Details   []BetDetail `json:"details"`
	ErrorDesc *string     `json:"error_desc"`
	PageInfo  struct {
		TotalPage   int `json:"totalPage"`
		CurrentPage int `json:"currentPage"`
		TotalCount  int `json:"totalCount"`
	} `json:"page_info"`
}

func (c *Client) GetLiveBetDetailsByMember(username, startDate, endDate string, page int) (*LiveBetDetailsResponse, error) {
	params := Request{
		Method:    "lbdm",
		Username:  username,
		StartDate: startDate,
		EndDate:   endDate,
		Page:      page,
	}
	godump.Dump(params)

	// For Live bet details API, the data is directly in the response, not nested under "data" or "result"
	raw, err := c.SendRawRequest(params)
	if err != nil {
		return nil, err
	}
	godump.Dump(string(raw.Body))

	var result LiveBetDetailsResponse
	if err := json.Unmarshal(raw.Body, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal Live bet details response: %w", err)
	}

	// Check for API errors
	if result.Status != 0 {
		errMsg := "Unknown error"
		if result.ErrorDesc != nil {
			errMsg = *result.ErrorDesc
		}
		return nil, NewProcessException(errMsg, result.Status)
	}

	return &result, nil
}

// GetRNGBetDetailsByMember gets RNG/FISH bet details for a specific member
// 获取指定玩家的RNG/FISH投注详情
// Note: According to documentation, RNG/FISH data is accessed via FTP, but this method tries API approach
func (c *Client) GetRNGBetDetailsByMember(username, startDate, endDate string, page int) (*RNGBetDetailsResponse, error) {
	params := Request{
		Method:    "bdm", // Use same method as Live but for RNG data
		Username:  username,
		StartDate: startDate,
		EndDate:   endDate,
		Page:      page,
	}
	godump.Dump(params)

	// For RNG bet details API, the data is directly in the response
	raw, err := c.SendRawRequest(params)
	if err != nil {
		return nil, err
	}
	godump.Dump(string(raw.Body))

	var result RNGBetDetailsResponse
	if err := json.Unmarshal(raw.Body, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal RNG bet details response: %w", err)
	}

	// Check for API errors
	if result.Status != 0 {
		errMsg := "Unknown error"
		if result.ErrorDesc != nil {
			errMsg = *result.ErrorDesc
		}
		return nil, NewProcessException(errMsg, result.Status)
	}

	return &result, nil
}

type LottoTransaction struct {
	TransactionID string  `json:"transaction_id"`
	Username      string  `json:"username"`
	GameCode      string  `json:"game_code"`
	IssueNumber   string  `json:"issue_number"`
	BetAmount     float64 `json:"bet_amount"`
	WinAmount     float64 `json:"win_amount"`
	NetAmount     float64 `json:"net_amount"`
	BetTime       string  `json:"bet_time"`
	DrawTime      string  `json:"draw_time"`
	SettleTime    string  `json:"settle_time"`
	Status        string  `json:"status"`
}

type LottoTransactionsResponse struct {
	Username     string             `json:"username"`
	Transactions []LottoTransaction `json:"transactions"`
	TotalRecord  int                `json:"total_record"`
	Page         int                `json:"page"`
	StartDate    string             `json:"start_date"`
	EndDate      string             `json:"end_date"`
}

func (c *Client) GetLottoTransactionsByMember(username, startDate, endDate string, page int) (*LottoTransactionsResponse, error) {
	params := Request{
		Method:    "lmb",
		Username:  username,
		StartDate: startDate,
		EndDate:   endDate,
		Page:      page,
	}

	resp, err := c.sendRequest(params)
	if err != nil {
		return nil, err
	}

	if err := checkResponseError(resp); err != nil {
		return nil, err
	}

	var result LottoTransactionsResponse
	if err := unmarshalResponseData(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

type LottoCode struct {
	GameCode string `json:"game_code"`
	GameName string `json:"game_name"`
}

type LottoCodesResponse struct {
	LottoCodes []LottoCode `json:"lotto_codes"`
}

func (c *Client) GetLottoCodes() (*LottoCodesResponse, error) {
	params := Request{
		Method: "glgl",
	}

	resp, err := c.sendRequest(params)
	if err != nil {
		return nil, err
	}

	if err := checkResponseError(resp); err != nil {
		return nil, err
	}

	var result LottoCodesResponse
	if err := unmarshalResponseData(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}
