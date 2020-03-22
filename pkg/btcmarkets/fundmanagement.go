package btcmarkets

import (
	"net/http"
	"net/url"
	"path"
	"strconv"
	"time"
)

// WithdrawRequestCrypto is a generalized withdraw request type
type WithdrawRequestCrypto struct {
	Amount   float64 `json:"amount"`
	Currency string  `json:"currency"`
	Address  string  `json:"address"`
}

// WithdrawRequestFiat is a generalized withdraw request type
type WithdrawRequestFiat struct {
	Amount        int64  `json:"amount"`
	Currency      string `json:"currency"`
	AccountName   string `json:"accountName"`
	AccountNumber string `json:"accountNumber"`
	BankName      string `json:"bankName"`
	BSBNumber     string `json:"bsbNumber"`
}

// PaymentDetails stores payment address
type PaymentDetails struct {
	Address string `json:"address"`
}

// WithdrawData stores data from asset transfers
type WithdrawData struct {
	ID             string         `json:"id"`
	AssetName      string         `json:"assetName"`
	Amount         float64        `json:"amount,string"`
	RequestType    string         `json:"type"`
	CreationTime   time.Time      `json:"creationTime"`
	Status         string         `json:"status"`
	Description    string         `json:"description"`
	Fee            float64        `json:"fee,string"`
	LastUpdate     string         `json:"lastUpdate"`
	PaymentDetails PaymentDetails `json:"paymentDetail,omitempty"`
}

// FundManagementServiceOp performs Fundmanagement operations on BTCMarkets
type FundManagementServiceOp struct {
	client *BTCMClient
}

// WithdrawCrypto This API is used to request to withdraw of crypto assets
func (f *FundManagementServiceOp) WithdrawCrypto(assetName, toAddress string, amount float64) (WithdrawData, error) {
	var wd WithdrawData

	wdreq := WithdrawRequestCrypto{
		Amount:   amount,
		Address:  toAddress,
		Currency: assetName,
	}

	req, err := f.client.NewRequest(http.MethodPost, path.Join(btcMarketsWithdrawals), wdreq)
	if err != nil {
		return wd, err
	}

	_, err = f.client.DoAuthenticated(req, wdreq, &wd)
	if err != nil {
		return wd, err
	}
	return wd, nil
}

// WithdrawFiat This API is used to request to withdraw of crypto assets
func (f *FundManagementServiceOp) WithdrawFiat(assetName, toAddress string, amount float64) (WithdrawData, error) {
	var wdf WithdrawData

	wdfreq := WithdrawRequestCrypto{
		Amount:   amount,
		Address:  toAddress,
		Currency: assetName,
	}

	req, err := f.client.NewRequest(http.MethodPost, path.Join(btcMarketsWithdrawals), wdfreq)
	if err != nil {
		return wdf, err
	}

	_, err = f.client.DoAuthenticated(req, wdfreq, &wdf)
	if err != nil {
		return wdf, err
	}
	return wdf, nil
}

// ListWithdrawls Returns list of withdrawals. This API supports pagination
func (f *FundManagementServiceOp) ListWithdrawls(before, after int64, limit int32) ([]WithdrawData, error) {
	var wd []WithdrawData

	params := url.Values{}
	if before > 0 {
		params.Add("before", strconv.FormatInt(before, 10))
	}
	if limit > 0 {
		params.Add("limit", strconv.Itoa(int(limit)))
	}

	req, err := f.client.NewRequest(http.MethodGet, path.Join(btcMarketsWithdrawals, "?"+params.Encode()), nil)
	if err != nil {
		return wd, err
	}

	_, err = f.client.DoAuthenticated(req, nil, &wd)
	if err != nil {
		return wd, err
	}
	return wd, nil
}

// GetWithdrawal This API is used to request to get withdraw by id.
func (f *FundManagementServiceOp) GetWithdrawal(orderID string) (WithdrawData, error) {
	var wd WithdrawData

	req, err := f.client.NewRequest(http.MethodDelete, path.Join(btcMarketsWithdrawals, orderID), nil)
	if err != nil {
		return wd, err
	}

	_, err = f.client.DoAuthenticated(req, nil, &wd)
	if err != nil {
		return wd, err
	}

	return wd, nil
}

// ListDeposits Returns list of depoists. This API supports pagination
func (f *FundManagementServiceOp) ListDeposits(before, after int64, limit int32) ([]WithdrawData, error) {
	var wd []WithdrawData

	params := url.Values{}
	if before > 0 {
		params.Add("before", strconv.FormatInt(before, 10))
	}
	if limit > 0 {
		params.Add("limit", strconv.Itoa(int(limit)))
	}
	if limit > 0 {
		params.Add("limit", strconv.Itoa(int(limit)))
	}

	req, err := f.client.NewRequest(http.MethodGet, path.Join(btcMarketsDeposits, "?"+params.Encode()), nil)
	if err != nil {
		return wd, err
	}

	_, err = f.client.DoAuthenticated(req, nil, &wd)
	if err != nil {
		return wd, err
	}
	return wd, nil
}
