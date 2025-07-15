package tcgaming

type CreateUserResponse struct {
	Status    int     `json:"status"`
	ErrorDesc *string `json:"error_desc"`
}

func (c *Client) CreateUser(username, password string) (*CreateUserResponse, error) {
	params := Request{
		Method:   "cm",
		Username: username,
		Password: password,
		Currency: c.Currency,
	}

	resp, err := c.sendRequest(params)
	if err != nil {
		return nil, err
	}

	if err := checkResponseError(resp); err != nil {
		return nil, err
	}

	// Return the exact API response
	return &CreateUserResponse{
		Status:    resp.Status,
		ErrorDesc: resp.ErrorDesc,
	}, nil
}

type UpdatePasswordResponse struct {
	Status    int     `json:"status"`
	ErrorDesc *string `json:"error_desc"`
}

func (c *Client) UpdatePassword(username, password string) (*UpdatePasswordResponse, error) {
	params := Request{
		Method:   "up",
		Username: username,
		Password: password,
		Currency: c.Currency,
	}

	resp, err := c.sendRequest(params)
	if err != nil {
		return nil, err
	}

	if err := checkResponseError(resp); err != nil {
		return nil, err
	}

	// Return the exact API response
	return &UpdatePasswordResponse{
		Status:    resp.Status,
		ErrorDesc: resp.ErrorDesc,
	}, nil
}