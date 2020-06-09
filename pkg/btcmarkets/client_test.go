package btcmarkets

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"strconv"
	"testing"
	"time"

	"golang.org/x/time/rate"
)

func setup(r *rate.Limiter) (*BTCMClient, *http.ServeMux, *httptest.Server, func(), error) {
	mux := http.NewServeMux()

	ts := httptest.NewServer(mux)
	rl := rate.NewLimiter(rate.Every(10*time.Second), 50)
	if rl != nil {
		rl = r
	}

	b, _ := url.Parse(ts.URL)
	w, _ := url.Parse(ts.URL)

	conf := ClientConfig{
		BaseURL:     b,
		WsURL:       w,
		APIKey:      "25d55ef7-f33e-49e8",
		APISecret:   "TXlTdXBlclNlY3JldEtleQ==",
		Httpclient:  ts.Client(),
		RateLimiter: rl,
	}

	client, err := NewBTCMClient(conf)
	if err != nil {
		return nil, nil, nil, nil, err
	}

	return client, mux, ts, ts.Close, nil
}

func TestSignMessage(t *testing.T) {
	client, _, _, teardown, err := setup(nil)
	defer teardown()
	if err != nil {
		t.Error(err.Error())
	}

	tests := []struct {
		name string
		msg  string
		want string
	}{
		{
			name: "Positive test: GET/v3/accounts/me/trading-fees",
			msg:  "GET/v3/accounts/me/trading-fees",
			want: "MbgRsg2A+uHLWoT1gRkknReAooeeAdonNBmbVaQ3fo0DrB2lJ6k7xMIHZ/kJa1H94A+ghbbF39zXIhUnYO5T8g==",
		},
		// {
		// 	name: "Negative test: GET/v3/accounts/me/trading-fees",
		// 	msg:  "GET/v3/accounts/me/trading-fees",
		// 	want: "saf",
		// },
	}

	tm := strconv.FormatInt(1257894000000, 10)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := client.signMessage(tt.msg + tm)
			if tt.want != got {
				t.Errorf("client.singMessage() wanted = '%v' got = '%v' ", tt.want, got)
			}
		})

	}
}

func TestGetServerTime(t *testing.T) {
	mocktimeserver := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`
		{
			"timestamp": "2019-09-01T18:34:27.045000Z"
		  }
		`))
	}
	client, mux, _, teardown, err := setup(nil)
	defer teardown()
	if err != nil {
		t.Error(err.Error())
	}

	mux.HandleFunc("/v3/time", mocktimeserver)
	st, err := client.GetServerTime()
	ty := reflect.TypeOf(st).String()
	if ty != "btcmarkets.ServerTime" {
		t.Errorf("Expected btcmarkets.ServerTime got %s", ty)
	}

}
