package btcmarkets

import (
	"errors"
	"net/http"
	"net/url"
	"path"
	"strconv"
	"strings"
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

// TransferData convinience datatype to make code more readable
type TransferData WithdrawData

// DepositAddress stores deposit address data
type DepositAddress struct {
	Address   string `json:"address"`
	AssetName string `json:"assetName"`
}

// WithdrawalFeeData stores data for fees
type WithdrawalFee struct {
	AssetName string  `json:"assetName"`
	Fee       float64 `json:"fee,string"`
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

// GetDeposit This API returns a deposit by id.
func (f *FundManagementServiceOp) GetDeposit(orderID string) (WithdrawData, error) {
	// TODO: WithdrawData may be incompatible for this reponse data received from this endpoint
	var wd WithdrawData

	req, err := f.client.NewRequest(http.MethodGet, path.Join(btcMarketsDeposits, orderID), nil)
	if err != nil {
		return wd, err
	}

	_, err = f.client.DoAuthenticated(req, nil, &wd)
	if err != nil {
		return wd, err
	}

	return wd, nil
}

// ListTransfers A transfer record refers either to a deposit or withdraw and this API returns list of transfers covering both depoists and withdrawals. This API supports pagination
func (f *FundManagementServiceOp) ListTransfers(before, after int64, limit int32) ([]TransferData, error) {
	var t []TransferData

	if (before > 0) && (after >= 0) {
		return nil, errors.New("BTCMarkets only supports either before or after, not both")
	}

	params := url.Values{}
	if before > 0 {
		params.Set("before", strconv.FormatInt(before, 10))
	}
	if after >= 0 {
		params.Set("after", strconv.FormatInt(after, 10))
	}
	if limit > 0 {
		params.Add("limit", strconv.Itoa(int(limit)))
	}

	req, err := f.client.NewRequest(http.MethodGet, path.Join(btcMarketsTransfers, "?"+params.Encode()), nil)
	if err != nil {
		return t, err
	}

	_, err = f.client.DoAuthenticated(req, nil, &t)
	if err != nil {
		return t, err
	}

	return t, nil
}

// GetTransfers This API retruns either deposit or withdrawal by id
func (f *FundManagementServiceOp) GetTransfers(orderID string) (WithdrawData, error) {
	var wd WithdrawData

	req, err := f.client.NewRequest(http.MethodGet, path.Join(btcMarketsTransfers, orderID), nil)
	if err != nil {
		return wd, err
	}

	_, err = f.client.DoAuthenticated(req, nil, &wd)
	if err != nil {
		return wd, err
	}

	return wd, nil
}

// GetDepositeAddress returns deposit address for the given asset
// Note: The documentation at https://api.btcmarkets.net/doc/v3#tag/Fund-Management-APIs/paths/~1v3~1addresses/get is wrong
func (f *FundManagementServiceOp) GetDepositeAddress(assetName string, before, after, limit int64) (DepositAddress, error) {
	var da DepositAddress

	if (before > 0) && (after > 0) {
		return da, errors.New("BTCMarkets only supports either before or after, not both")
	}
	params := url.Values{}
	assetName = strings.ToUpper(assetName) //API expects assetName has to be uppercase
	params.Set("assetName", assetName)
	if before > 0 {
		params.Set("before", strconv.FormatInt(before, 10))
	}
	if after >= 0 {
		params.Set("after", strconv.FormatInt(after, 10))
	}
	if limit > 0 {
		params.Set("limit", strconv.FormatInt(limit, 10))
	}

	req, err := f.client.NewRequest(http.MethodGet, path.Join(btcMarketsAddresses, "?"+params.Encode()), nil)
	if err != nil {
		return da, err
	}

	_, err = f.client.DoAuthenticated(req, nil, &da)
	if err != nil {
		return da, err
	}

	return da, nil
}

// GetWithdrawalFees Returns fees associated with withdrawals.
// This API is public and does not require authentication as the fees as system wide and published on the website
func (f *FundManagementServiceOp) GetWithdrawalFees() ([]WithdrawalFee, error) {
	var wf []WithdrawalFee

	req, err := f.client.NewRequest(http.MethodGet, btcMarketsWithdrawalFees, nil)
	if err != nil {
		return wf, err
	}

	_, err = f.client.Do(req, &wf)
	if err != nil {
		return wf, err
	}

	return wf, nil
}
