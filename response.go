package tcgaming

import (
	"encoding/json"
	"fmt"
)

// checkResponseError checks if the response contains an error
// Matches Java SDK's error handling pattern
func checkResponseError(resp *Response) error {
	if resp.Status != 0 {
		errMsg := "Unknown error"
		// Priority: error_message (Java SDK) > error_desc (PHP) 
		if resp.ErrorMessage != "" {
			errMsg = resp.ErrorMessage
		} else if resp.ErrorDesc != nil {
			errMsg = *resp.ErrorDesc
		}
		// Return ProcessException to match Java SDK
		return NewProcessException(errMsg, resp.Status)
	}
	return nil
}

// unmarshalResponseData attempts to unmarshal response data from various fields
func unmarshalResponseData(resp *Response, target interface{}) error {
	// Try resp.Data first
	if len(resp.Data) > 0 && string(resp.Data) != "null" {
		if err := json.Unmarshal(resp.Data, target); err == nil {
			return nil
		}
	}
	
	// Try resp.Result
	if len(resp.Result) > 0 && string(resp.Result) != "null" {
		if err := json.Unmarshal(resp.Result, target); err == nil {
			return nil
		}
	}
	
	// No data found in response
	return fmt.Errorf("no data found in response")
}

