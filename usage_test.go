// File generated from our OpenAPI spec by Stainless.

package subhosting_test

import (
	"context"
	"os"
	"testing"

	"github.com/denoland/subhosting-go"
	"github.com/denoland/subhosting-go/internal/testutil"
	"github.com/denoland/subhosting-go/option"
)

func TestUsage(t *testing.T) {
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
	organization, err := client.Organizations.Get(context.TODO(), "DEPLOY_ORG_ID")
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v\n", organization.ID)
}
