package cryptowallet

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

const (
	methodGet  = "GET"
	methodPost = "POST"
	methodPut  = "PUT"
)

const (
	OperationTypeIn           = "IN"
	OperationTypeOut          = "OUT"
	OperationTypeCreateWallet = "CREATE_WALLET"
)

type Config struct {
	ApiKey  string
	ApiUrl  string
	Test    bool
	Logging uint8
}

type Client struct {
	ApiKey string
	ApiUrl string
	Test   bool
	log    logger
}

type Response struct {
	StatusCode int
	Status     string      `json:"status"`
	Error      string      `json:"error"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
}

type Currency struct {
	Code           string `json:"currencyCode"`
	TokenID        string `json:"tokenId"`
	FractionNumber uint8  `json:"fractionNumber"`
	Enabled        bool   `json:"enabled"`
}

type Address struct {
	Hash string `json:"hash"`
}

type Balance struct {
	Hash         string  `json:"walletId"`
	Balance      uint64  `json:"balance"`
	BalanceFloat float64 `json:"-"`
}

type Transaction struct {
	ID                 uint32   `json:"id"`
	Timestamp          string   `json:"timestamp"`
	TimestampLong      uint64   `json:"timestampLong"`
	Currency           string   `json:"currency"`
	Amount             uint64   `json:"amount"`
	Type               string   `json:"type"`
	DestinationAddress string   `json:"destinationAddress"`
	SourceAddresses    []string `json:"sourceAddresses"`
	IpAddress          string   `json:"ipAddress"`
}

type TransactionsData struct {
	Transactions  []*Transaction `json:"content"`
	TotalPages    uint32         `json:"totalPages"`
	TotalElements uint32         `json:"totalElements"`
	Size          uint32         `json:"size"`
	Page          uint32         `json:"page"`
}

func (b *Balance) normalizeNumber(fractionNumber uint8) *Balance {
	b.BalanceFloat = 0
	return b
}

func NewClient(config *Config) *Client {
	if config == nil {
		config = &Config{}
	}
	client := &Client{
		ApiKey: config.ApiKey,
		ApiUrl: strings.Trim(config.ApiUrl, "/"),
		Test:   config.Test,
		log:    newWalletLogger(config.Logging),
	}

	return client
}

func (c *Client) GetCurrencies() ([]*Currency, error) {
	respData, err := c.makeRequest("/currencies", methodGet, nil, TestResponseGetCurrencies)
	if err != nil {
		return nil, err
	}

	respDataList, ok := respData.Data.([]interface{})
	if !ok {
		if c.log.Enable {
			c.log.Error.Printf("data format error: %s", ErrResponseFormat.Error())
		}
		return nil, ErrResponseFormat
	}

	var currenciesList []*Currency
	for _, item := range respDataList {
		currencyItem := &Currency{}
		b, err := json.Marshal(item)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(b, currencyItem)
		if err != nil {
			return nil, err
		}
		currenciesList = append(currenciesList, currencyItem)
	}

	return currenciesList, nil
}

func (c *Client) GetTokens(currencyCode string) ([]*Currency, error) {
	respData, err := c.makeRequest(fmt.Sprintf("/currencies/%s", currencyCode), methodGet, nil, TestResponseGetTokens)
	if err != nil {
		return nil, err
	}

	respDataList, ok := respData.Data.([]interface{})
	if !ok {
		if c.log.Enable {
			c.log.Error.Printf("data format error: %s", ErrResponseFormat.Error())
		}
		return nil, ErrResponseFormat
	}

	var tokensList []*Currency
	for _, item := range respDataList {
		tokenItem := &Currency{}
		b, err := json.Marshal(item)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(b, tokenItem)
		if err != nil {
			return nil, err
		}
		tokensList = append(tokensList, tokenItem)
	}

	return tokensList, nil
}

func (c *Client) AddCoinAddress(currencyCode string) (*Address, error) {
	respData, err := c.makeRequest(fmt.Sprintf("/wallet/%s", currencyCode), methodPut, nil, TestResponseAddAddress)
	if err != nil {
		return nil, err
	}

	respDataItem, ok := respData.Data.(interface{})
	if !ok {
		if c.log.Enable {
			c.log.Error.Printf("data format error: %s", ErrResponseFormat.Error())
		}
		return nil, ErrResponseFormat
	}

	addressItem := &Address{}
	b, err := json.Marshal(respDataItem)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(b, addressItem)
	if err != nil {
		return nil, err
	}

	return addressItem, nil
}

func (c *Client) AddTokenAddress(currencyCode string, tokenCode string, tokenID string) (*Address, error) {
	var params []string
	if tokenCode != "" {
		params = append(params, fmt.Sprintf("code=%s", tokenCode))
	}
	if tokenID != "" {
		params = append(params, fmt.Sprintf("tokenId=%s", tokenID))
	}

	respData, err := c.makeRequest(fmt.Sprintf("/wallet/%s?%s", currencyCode, strings.Join(params, "&")), methodPut, nil, TestResponseAddAddress)
	if err != nil {
		return nil, err
	}

	respDataItem, ok := respData.Data.(interface{})
	if !ok {
		if c.log.Enable {
			c.log.Error.Printf("data format error: %s", ErrResponseFormat.Error())
		}
		return nil, ErrResponseFormat
	}

	addressItem := &Address{}
	b, err := json.Marshal(respDataItem)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(b, addressItem)
	if err != nil {
		return nil, err
	}

	return addressItem, nil
}

func (c *Client) GetBalanceByCoinAddress(currencyCode string, address string) (*Balance, error) {
	respData, err := c.makeRequest(fmt.Sprintf("/wallet/status/%s/%s", currencyCode, address), methodGet, nil, TestResponseGetBalance)
	if err != nil {
		return nil, err
	}

	respDataItem, ok := respData.Data.(interface{})
	if !ok {
		if c.log.Enable {
			c.log.Error.Printf("data format error: %s", ErrResponseFormat.Error())
		}
		return nil, ErrResponseFormat
	}

	balanceItem := &Balance{}
	b, err := json.Marshal(respDataItem)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(b, balanceItem)
	if err != nil {
		return nil, err
	}

	balanceItem = balanceItem.normalizeNumber(0)

	return balanceItem, nil
}

func (c *Client) GetBalanceByTokenAddress(currencyCode string, address string, tokenCode string, tokenID string) (*Balance, error) {
	var params []string
	if tokenCode != "" {
		params = append(params, fmt.Sprintf("code=%s", tokenCode))
	}
	if tokenID != "" {
		params = append(params, fmt.Sprintf("tokenId=%s", tokenID))
	}

	respData, err := c.makeRequest(fmt.Sprintf("/wallet/status/%s/%s?%s", currencyCode, address, strings.Join(params, "&")), methodGet, nil, TestResponseGetBalance)
	if err != nil {
		return nil, err
	}

	respDataItem, ok := respData.Data.(interface{})
	if !ok {
		if c.log.Enable {
			c.log.Error.Printf("data format error: %s", ErrResponseFormat.Error())
		}
		return nil, ErrResponseFormat
	}

	balanceItem := &Balance{}
	b, err := json.Marshal(respDataItem)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(b, balanceItem)
	if err != nil {
		return nil, err
	}

	balanceItem = balanceItem.normalizeNumber(0)

	return balanceItem, nil
}

func (c *Client) GetTransactions(page uint8, fromTime uint64, operationType string) (*TransactionsData, error) {
	var params []string
	params = append(params, fmt.Sprintf("page=%d", page))
	params = append(params, "size=100")
	params = append(params, "sort=timestampLong,DESC")

	jsonBody := map[string]interface{}{"fromTime": fromTime, "types": []string{operationType}}
	jsonBodyValue, _ := json.Marshal(jsonBody)

	respData, err := c.makeRequest(fmt.Sprintf("/operation?%s", strings.Join(params, "&")), methodPost, jsonBodyValue, TestResponseGetTransactions)
	if err != nil {
		return nil, err
	}

	respDataItem, ok := respData.Data.(interface{})
	if !ok {
		if c.log.Enable {
			c.log.Error.Printf("data format error: %s", ErrResponseFormat.Error())
		}
		return nil, ErrResponseFormat
	}

	transactionsDataItem := &TransactionsData{}
	b, err := json.Marshal(respDataItem)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(b, transactionsDataItem)
	if err != nil {
		return nil, err
	}

	return transactionsDataItem, nil
}

func (c *Client) makeRequest(url string, method string, jsonBody []byte, testResponseString string) (*Response, error) {
	respData := &Response{}
	if c.Test {
		if c.log.Enable {
			c.log.Info.Printf("TEST REQUEST to: %s %s, loading: %s", strings.ToUpper(method), url, testResponseString)
		}
		data, err := c.sandboxRequest(testResponseString)
		if err != nil {
			return nil, err
		}
		respData = data
	} else {
		if c.log.Enable {
			c.log.Info.Printf("API REQUEST to: %s %s; body: %s", strings.ToUpper(method), url, string(jsonBody))
		}
		data, err := c.apiRequest(url, method, jsonBody)
		if err != nil {
			return nil, err
		}
		respData = data
	}

	if c.log.Enable {
		c.log.Info.Printf("api code: %d; api status field: '%s'; api error field: '%s'; api message field: '%s'", respData.StatusCode, respData.Status, respData.Error, respData.Message)
		c.log.Info.Println("api data field:", respData.Data)
	}

	return respData, nil
}

func (c *Client) apiRequest(url string, method string, jsonBody []byte) (*Response, error) {
	respData := &Response{}
	req, err := http.NewRequest(method, fmt.Sprintf("%s%s", c.ApiUrl, url), bytes.NewBuffer(jsonBody))
	if err != nil {
		if c.log.Enable {
			c.log.Error.Printf("creating request error: %s", err.Error())
		}
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("X-API-Key", c.ApiKey)

	client := &http.Client{
		Timeout: time.Duration(5 * time.Second),
	}
	resp, err := client.Do(req)
	if err != nil {
		if c.log.Enable {
			c.log.Error.Printf("sending request error: %s", err.Error())
		}
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		if c.log.Enable {
			c.log.Error.Printf("response body error: %s", err.Error())
		}
		return nil, err
	}
	if err = json.Unmarshal(body, &respData); err != nil {
		if c.log.Enable {
			c.log.Error.Printf("response json error: %s", err.Error())
		}
		return nil, err
	}

	if 200 != resp.StatusCode {
		if c.log.Enable {
			c.log.Error.Printf("response code error %d: %s", resp.StatusCode, body)
		}
		return nil, fmt.Errorf("error %d: %s", resp.StatusCode, respData.Status)
	}

	respData.StatusCode = resp.StatusCode

	return respData, nil
}

func (c *Client) sandboxRequest(testResponseString string) (*Response, error) {
	respData := &Response{}
	if err := json.Unmarshal([]byte(testResponseString), &respData); err != nil {
		respData.StatusCode = 500
		return nil, err
	}
	respData.StatusCode = 200
	return respData, nil
}
