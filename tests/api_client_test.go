package api_test

import (
	"testing"

	"github.com/aspose-words-cloud/aspose-words-cloud-go/api"
	"github.com/aspose-words-cloud/aspose-words-cloud-go/api/models"
)

func TestAppSidAndAppKey(t *testing.T) {
	p := []struct {
		appKey string
		appSid string
	}{
		{"", "x"},
		{"x", ""},
	}
	for _, cp := range p {
		config := models.Configuration{
			AppKey: cp.appKey,
			AppSid: cp.appSid,
		}
		_, err := api.NewAPIClient(&config)

		if err == nil {
			t.Error(err)
		}
	}
}

func TestBaseUrl(t *testing.T) {
	config := models.Configuration{
		AppKey:  "x",
		AppSid:  "x",
		BaseUrl: "http://localhost",
	}
	_, err := api.NewAPIClient(&config)

	if err != nil {
		t.Error(err)
	}
}
