// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package subhosting_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/denoland/subhosting-go"
	"github.com/denoland/subhosting-go/internal/testutil"
	"github.com/denoland/subhosting-go/option"
)

func TestDeploymentAppLogGetWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := subhosting.NewClient(
		option.WithBaseURL(baseURL),
		option.WithBearerToken("My Bearer Token"),
	)
	_, err := client.Deployments.AppLogs.Get(
		context.TODO(),
		"string",
		subhosting.DeploymentAppLogGetParams{
			Cursor: subhosting.F("string"),
			Level:  subhosting.F(subhosting.DeploymentAppLogGetParamsLevelError),
			Limit:  subhosting.F(int64(1)),
			Order:  subhosting.F("string"),
			Q:      subhosting.F("string"),
			Region: subhosting.F(subhosting.DeploymentAppLogGetParamsRegionGcpAsiaEast1),
			Since:  subhosting.F(time.Now()),
			Sort:   subhosting.F("string"),
			Until:  subhosting.F(time.Now()),
		},
	)
	if err != nil {
		var apierr *subhosting.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
