// File generated from our OpenAPI spec by Stainless.

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

func TestOrganizationAnalyticsGet(t *testing.T) {
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
	_, err := client.Organizations.Analytics.Get(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		subhosting.OrganizationAnalyticsGetParams{
			Since: subhosting.F(time.Now()),
			Until: subhosting.F(time.Now()),
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
