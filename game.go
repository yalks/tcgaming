package tcgaming

import (
	"encoding/json"
	"fmt"
)

const (
	GameModeTest = 0
	GameModeLive = 1
)

type LaunchGameResponse struct {
	GameURL string `json:"game_url"`
	Token   string `json:"token"`
}

func (c *Client) LaunchGameRNG(username string, productType, gameMode int, gameCode, platform string) (*LaunchGameResponse, error) {
	params := Request{
		Method:      "lg",
		Username:    username,
		ProductType: productType,
		GameMode:    gameMode,
		GameCode:    gameCode,
		Platform:    platform,
	}

	resp, err := c.sendRequest(params)
	if err != nil {
		return nil, err
	}

	if err := checkResponseError(resp); err != nil {
		return nil, err
	}

	// For launch game API, the data is directly in the response, not nested under "data" or "result"
	raw, err := c.SendRawRequest(params)
	if err != nil {
		return nil, err
	}

	var fullResponse struct {
		Status       int     `json:"status"`
		GameURL      string  `json:"game_url"`
		Token        string  `json:"token"`
		ErrorDesc    *string `json:"error_desc"`
		ErrorMessage string  `json:"error_message"`
	}

	if err := json.Unmarshal(raw.Body, &fullResponse); err != nil {
		return nil, fmt.Errorf("failed to unmarshal launch game response: %w", err)
	}

	// Check for API errors
	if fullResponse.Status != 0 {
		errMsg := "Unknown error"
		if fullResponse.ErrorMessage != "" {
			errMsg = fullResponse.ErrorMessage
		} else if fullResponse.ErrorDesc != nil {
			errMsg = *fullResponse.ErrorDesc
		}
		return nil, NewProcessException(errMsg, fullResponse.Status)
	}

	result := &LaunchGameResponse{
		GameURL: fullResponse.GameURL,
		Token:   fullResponse.Token,
	}

	return result, nil
}

func (c *Client) LaunchGameLottery(username string, productType, gameMode int, gameCode, platform, view string) (*LaunchGameResponse, error) {
	betMode := "Traditional"
	series := []SeriesConfig{
		{
			GameGroupCode: "SSC",
			PrizeModeID:   1,
			MaxSeries:     1956,
			MinSeries:     1700,
			MaxBetSeries:  1950,
			DefaultSeries: 1800,
		},
	}

	params := Request{
		Method:      "lg",
		Username:    username,
		ProductType: productType,
		GameCode:    gameCode,
		GameMode:    gameMode,
		Platform:    platform,
		BetMode:     betMode,
		View:        view,
		Series:      series,
	}

	resp, err := c.sendRequest(params)
	if err != nil {
		return nil, err
	}

	if err := checkResponseError(resp); err != nil {
		return nil, err
	}

	var result LaunchGameResponse
	if err := unmarshalResponseData(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

type GameInfo struct {
	DisplayStatus int    `json:"displayStatus"`
	GameType      string `json:"gameType"`
	GameName      string `json:"gameName"`
	TCGGameCode   string `json:"tcgGameCode"`
	ProductCode   string `json:"productCode"`
	ProductType   string `json:"productType"`
	Platform      string `json:"platform"`
	GameSubType   string `json:"gameSubType"`
	TrialSupport  bool   `json:"trialSupport"`
}

type PageInfo struct {
	TotalPage   int `json:"totalPage"`
	CurrentPage int `json:"currentPage"`
	TotalCount  int `json:"totalCount"`
}

type GameListResponse struct {
	Games    []GameInfo `json:"games"`
	PageInfo PageInfo   `json:"page_info"`
}

func (c *Client) GetGameList(productType int, platform, clientType, gameType string, page, pageSize int) (*GameListResponse, error) {
	params := Request{
		Method:      "tgl",
		ProductType: productType,
		Platform:    platform,
		ClientType:  clientType,
		GameType:    gameType,
		Page:        page,
		PageSize:    pageSize,
	}

	resp, err := c.sendRequest(params)
	if err != nil {
		return nil, err
	}

	if err := checkResponseError(resp); err != nil {
		return nil, err
	}

	// For game list API, the data is directly in the response, not nested under "data" or "result"
	// We need to unmarshal the entire response body to get the games and page_info
	raw, err := c.SendRawRequest(params)
	if err != nil {
		return nil, err
	}

	var fullResponse struct {
		Status    int        `json:"status"`
		Games     []GameInfo `json:"games"`
		PageInfo  PageInfo   `json:"page_info"`
		ErrorDesc *string    `json:"error_desc"`
	}

	if err := json.Unmarshal(raw.Body, &fullResponse); err != nil {
		return nil, fmt.Errorf("failed to unmarshal game list response: %w", err)
	}

	// Check for API errors
	if fullResponse.Status != 0 {
		errMsg := "Unknown error"
		if fullResponse.ErrorDesc != nil {
			errMsg = *fullResponse.ErrorDesc
		}
		return nil, NewProcessException(errMsg, fullResponse.Status)
	}

	result := &GameListResponse{
		Games:    fullResponse.Games,
		PageInfo: fullResponse.PageInfo,
	}

	return result, nil
}

type PlayerRank struct {
	Rank     int     `json:"rank"`
	Username string  `json:"username"`
	WinLoss  float64 `json:"win_loss"`
	Turnover float64 `json:"turnover"`
}

type GameRankResponse struct {
	Rankings []PlayerRank `json:"rankings"`
	GameCode string       `json:"game_code"`
}

func (c *Client) GetGameRank(productType int, gameCategory, gameCode, startDate, endDate string, count int) (*GameRankResponse, error) {
	params := Request{
		Method:       "pgr",
		ProductType:  productType,
		GameCategory: gameCategory,
		GameCode:     gameCode,
		StartDate:    startDate,
		EndDate:      endDate,
		Count:        count,
	}

	resp, err := c.sendRequest(params)
	if err != nil {
		return nil, err
	}

	if err := checkResponseError(resp); err != nil {
		return nil, err
	}

	var result GameRankResponse
	if err := unmarshalResponseData(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}
