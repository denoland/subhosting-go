// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package subhosting_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/denoland/subhosting-go"
	"github.com/denoland/subhosting-go/internal/testutil"
	"github.com/denoland/subhosting-go/option"
)

func TestDeploymentDelete(t *testing.T) {
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
	err := client.Deployments.Delete(context.TODO(), "abcde12vwxyz")
	if err != nil {
		var apierr *subhosting.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestDeploymentGet(t *testing.T) {
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
	_, err := client.Deployments.Get(context.TODO(), "abcde12vwxyz")
	if err != nil {
		var apierr *subhosting.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestDeploymentRedeployWithOptionalParams(t *testing.T) {
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
	_, err := client.Deployments.Redeploy(
		context.TODO(),
		"abcde12vwxyz",
		subhosting.DeploymentRedeployParams{
			Databases: subhosting.F(map[string]string{
				"default": "a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11",
			}),
			Description: subhosting.F("Updated description"),
			EnvVars: subhosting.F(map[string]string{
				"MY_ENV":            "hey",
				"ENV_TO_BE_DELETED": nil,
			}),
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
