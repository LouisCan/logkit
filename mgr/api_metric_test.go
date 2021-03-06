package mgr

import (
	"net/http"

	"github.com/json-iterator/go"
	"github.com/qiniu/logkit/metric"
	. "github.com/qiniu/logkit/utils/models"
	"github.com/stretchr/testify/assert"
)

type respMetricUsage struct {
	Code string   `json:"code"`
	Data []Option `json:"data"`
}

type respMetricKeys struct {
	Code string                   `json:"code"`
	Data map[string]KeyValueSlice `json:"data"`
}

type respMetricOptions struct {
	Code string              `json:"code"`
	Data map[string][]Option `json:"data"`
}

func metricAPITest(p *testParam) {
	t := p.t
	rs := p.rs
	var got1 respMetricUsage

	url := "http://127.0.0.1" + rs.address + "/logkit/metric/usages"
	respCode, respBody, err := makeRequest(url, http.MethodGet, []byte{})
	assert.NoError(t, err, string(respBody))
	assert.Equal(t, http.StatusOK, respCode)
	if err = jsoniter.Unmarshal(respBody, &got1); err != nil {
		t.Fatalf("respBody %v unmarshal failed, error is %v", respBody, err)
	}
	assert.Equal(t, metric.GetMetricUsages(), got1.Data)

	var got2 respMetricOptions
	url = "http://127.0.0.1" + rs.address + "/logkit/metric/options"
	respCode, respBody, err = makeRequest(url, http.MethodGet, []byte{})
	assert.NoError(t, err, string(respBody))
	assert.Equal(t, http.StatusOK, respCode)
	if err = jsoniter.Unmarshal(respBody, &got2); err != nil {
		t.Fatalf("respBody %v unmarshal failed, error is %v", respBody, err)
	}
	metricOpts := make(map[string][]Option)
	for key, option := range metric.GetMetricOptions() {
		if opt, ok := option.([]Option); ok {
			metricOpts[key] = opt
		}
	}
	assert.Equal(t, metricOpts, got2.Data)

	var got3 respMetricKeys
	url = "http://127.0.0.1" + rs.address + "/logkit/metric/keys"
	respCode, respBody, err = makeRequest(url, http.MethodGet, []byte{})
	assert.NoError(t, err, string(respBody))
	assert.Equal(t, http.StatusOK, respCode)
	if err = jsoniter.Unmarshal(respBody, &got3); err != nil {
		t.Fatalf("respBody %v unmarshal failed, error is %v", respBody, err)
	}
	metricKeys := make(map[string]KeyValueSlice)
	for key, MKey := range metric.GetMetricTypeKey() {
		if mkey, ok := MKey.(KeyValueSlice); ok {
			metricKeys[key] = mkey
		}
	}
	assert.Equal(t, metricKeys, got3.Data)
}
