package tcgaming

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/goforj/godump"
)

type Client struct {
	config     *Config
	Currency   string
	httpClient *http.Client
}

type Option func(*Client)

func WithHTTPClient(client *http.Client) Option {
	return func(c *Client) {
		c.httpClient = client
	}
}

func WithTimeout(timeout time.Duration) Option {
	return func(c *Client) {
		c.httpClient.Timeout = timeout
	}
}

// NewClient creates a new client with the given configuration
// This matches the Java SDK pattern
func NewClient(config *Config, currency string, opts ...Option) (*Client, error) {
	if err := config.Validate(); err != nil {
		return nil, err
	}

	client := &Client{
		config:   config,
		Currency: currency,
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}

	for _, opt := range opts {
		opt(client)
	}

	return client, nil
}

// NewClientLegacy creates a client with the old parameter style (for backward compatibility)
func NewClientLegacy(merchantCode, desKey, signKey, currency string, opts ...Option) *Client {
	config := &Config{
		APIUrl:       "http://www.url.com/doBusiness.do",
		MerchantCode: merchantCode,
		DESKey:       desKey,
		SHA256Key:    signKey,
	}

	client, _ := NewClient(config, currency, opts...)
	return client
}

func (c *Client) SetURL(url string) {
	c.config.APIUrl = url
}

// SendRawRequest sends a request and returns the raw response for debugging
func (c *Client) SendRawRequest(params Request) (*RawResponse, error) {
	jsonData, err := json.Marshal(params)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	encrypted, err := c.encryptText(string(jsonData))
	if err != nil {
		return nil, fmt.Errorf("failed to encrypt request: %w", err)
	}

	sign := c.generateSign(encrypted)

	formData := url.Values{}
	formData.Set("merchant_code", c.config.MerchantCode)
	formData.Set("params", encrypted)
	formData.Set("sign", sign)

	req, err := http.NewRequest("POST", c.config.APIUrl, bytes.NewBufferString(formData.Encode()))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	return &RawResponse{
		Body:       body,
		StatusCode: resp.StatusCode,
	}, nil
}

type Request struct {
	Method       string         `json:"method"`
	Username     string         `json:"username,omitempty"`
	Password     string         `json:"password,omitempty"`
	Currency     string         `json:"currency,omitempty"`
	ProductType  int            `json:"product_type,omitempty"`
	FundType     int            `json:"fund_type,omitempty"`
	Amount       float64        `json:"amount,omitempty"`
	ReferenceNo  string         `json:"reference_no,omitempty"`
	RefNo        string         `json:"ref_no,omitempty"`
	GameMode     int            `json:"game_mode,omitempty"`
	GameCode     string         `json:"game_code,omitempty"`
	Platform     string         `json:"platform,omitempty"`
	BatchName    string         `json:"batch_name,omitempty"`
	Page         int            `json:"page,omitempty"`
	PageSize     int            `json:"page_size,omitempty"`
	StartDate    string         `json:"start_date,omitempty"`
	EndDate      string         `json:"end_date,omitempty"`
	ClientType   string         `json:"client_type,omitempty"`
	GameType     string         `json:"game_type,omitempty"`
	View         string         `json:"view,omitempty"`
	BetMode      string         `json:"lottery_bet_mode,omitempty"`
	Series       []SeriesConfig `json:"series,omitempty"`
	GameCategory string         `json:"game_category,omitempty"`
	Count        int            `json:"count,omitempty"`
}

type SeriesConfig struct {
	GameGroupCode string `json:"game_group_code"`
	PrizeModeID   int    `json:"prize_mode_id"`
	MaxSeries     int    `json:"max_series"`
	MinSeries     int    `json:"min_series"`
	MaxBetSeries  int    `json:"max_bet_series"`
	DefaultSeries int    `json:"default_series"`
}

type Response struct {
	Status       int     `json:"status"`
	ErrorMessage string  `json:"error_message,omitempty"` // Java SDK uses error_message
	ErrorDesc    *string `json:"error_desc,omitempty"`    // PHP version compatibility

	// Direct fields from API responses
	Balance            float64         `json:"balance,omitempty"`             // GetBalance API
	TransactionStatus  string          `json:"transaction_status,omitempty"`  // Transfer APIs
	TransactionDetails json.RawMessage `json:"transaction_details,omitempty"` // Check Transfer API

	// Generic fields for other responses
	Data   json.RawMessage `json:"data,omitempty"`   // Additional response data
	Result json.RawMessage `json:"result,omitempty"` // Some APIs might use result
}

// RawResponse represents the raw response from the API
type RawResponse struct {
	Body       []byte
	StatusCode int
}

func (c *Client) sendRequest(params Request) (*Response, error) {
	jsonData, err := json.Marshal(params)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	encrypted, err := c.encryptText(string(jsonData))
	if err != nil {
		return nil, fmt.Errorf("failed to encrypt request: %w", err)
	}

	sign := c.generateSign(encrypted)

	formData := url.Values{}
	formData.Set("merchant_code", c.config.MerchantCode)
	formData.Set("params", encrypted)
	formData.Set("sign", sign)

	req, err := http.NewRequest("POST", c.config.APIUrl, bytes.NewBufferString(formData.Encode()))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}
	godump.Dump(body)

	// Check if response is HTML (error page)
	if len(body) > 0 && body[0] == '<' {
		// Try to extract meaningful error from HTML if possible
		bodyStr := string(body)
		if len(bodyStr) > 500 {
			bodyStr = bodyStr[:500] + "..."
		}
		return nil, fmt.Errorf("received HTML response instead of JSON (status: %d): %s", resp.StatusCode, bodyStr)
	}

	// Check for empty response
	if len(body) == 0 {
		return nil, fmt.Errorf("received empty response from server (status: %d)", resp.StatusCode)
	}

	var response Response
	if err := json.Unmarshal(body, &response); err != nil {
		// Show the actual response for debugging
		bodyStr := string(body)
		if len(bodyStr) > 200 {
			bodyStr = bodyStr[:200] + "..."
		}
		return nil, fmt.Errorf("failed to unmarshal response: %w (response: %s)", err, bodyStr)
	}

	return &response, nil
}

func (c *Client) encryptText(plainText string) (string, error) {
	encryptor := NewDESEncrypt(c.config.DESKey)
	return encryptor.Encrypt(plainText)
}

func (c *Client) decryptText(encryptedText string) (string, error) {
	encryptor := NewDESEncrypt(c.config.DESKey)
	return encryptor.Decrypt(encryptedText)
}

func (c *Client) generateSign(params string) string {
	h := sha256.New()
	h.Write([]byte(params + c.config.SHA256Key))
	return hex.EncodeToString(h.Sum(nil))
}
