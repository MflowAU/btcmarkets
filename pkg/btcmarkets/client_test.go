package btcmarkets

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"testing"
)

func setup() (*BTCMClient, *http.ServeMux, *httptest.Server, func(), error) {
	mux := http.NewServeMux()

	ts := httptest.NewServer(mux)

	b, _ := url.Parse(ts.URL)
	w, _ := url.Parse(ts.URL)

	conf := ClientConfig{
		BaseURL:    b,
		WsURL:      w,
		APIKey:     "25d55ef7-f33e-49e8",
		APISecret:  "TXlTdXBlclNlY3JldEtleQ==",
		Httpclient: ts.Client(),
	}

	client, err := NewBTCMClient(conf)
	if err != nil {
		return nil, nil, nil, nil, err
	}

	return client, mux, ts, ts.Close, nil
}
func TestSignMessage(t *testing.T) {
	client, _, _, teardown, err := setup()
	defer teardown()
	if err != nil {
		t.Error(err.Error())
	}

	tm := strconv.FormatInt(1257894000000, 10)
	// GET/v3/accounts/me/trading-fees1257894000000
	smsg := client.signMessage("GET/v3/accounts/me/trading-fees" + tm)
	if smsg != "MbgRsg2A+uHLWoT1gRkknReAooeeAdonNBmbVaQ3fo0DrB2lJ6k7xMIHZ/kJa1H94A+ghbbF39zXIhUnYO5T8g==" {
		t.Errorf("Expecting MbgRsg2A+uHLWoT1gRkknReAooeeAdonNBmbVaQ3fo0DrB2lJ6k7xMIHZ/kJa1H94A+ghbbF39zXIhUnYO5T8g== got %s", smsg)
	}
}

// func TestGetServerTime(t *testing.T) {
// 	mocktimeserver := func(w http.ResponseWriter, r *http.Request) {
// 		w.WriteHeader(http.StatusOK)
// 		w.Write([]byte(`
// 		{
// 			"timestamp": "2019-09-01T18:34:27.045000Z"
// 		  }
// 		`))
// 	}
// 	client, mux, _, teardown, err := setup()
// 	defer teardown()
// 	if err != nil {
// 		t.Error(err.Error())
// 	}

// 	mux.HandleFunc("/v3/time", mocktimeserver)
// 	st, err := client.GetServerTime()
// 	ty := reflect.TypeOf(st).String()
// 	if ty != "btcmarkets.ServerTime" {
// 		t.Errorf("Expected btcmarkets.ServerTime got %s", ty)
// 	}

// }
