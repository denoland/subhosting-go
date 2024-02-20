// File generated from our OpenAPI spec by Stainless.

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

func TestProjectDeploymentNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Projects.Deployments.New(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		subhosting.ProjectDeploymentNewParams{
			Assets: subhosting.F(map[string]subhosting.ProjectDeploymentNewParamsAssets{
				"main.ts": subhosting.ProjectDeploymentNewParamsAssetsFile(subhosting.ProjectDeploymentNewParamsAssetsFile{
					Content:  subhosting.F("Deno.serve((req: Request) => new Response(\"Hello World\"));\n"),
					Encoding: subhosting.F(subhosting.ProjectDeploymentNewParamsAssetsFileEncodingUtf8),
					GitSha1:  subhosting.F("string"),
					Kind:     subhosting.F(subhosting.ProjectDeploymentNewParamsAssetsFileKindFile),
				}),
				"images/cat1.png": subhosting.ProjectDeploymentNewParamsAssetsFile(subhosting.ProjectDeploymentNewParamsAssetsFile{
					Content:  subhosting.F("iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAQAAAC1HAwCAAAAC0lEQVR42mNk"),
					Encoding: subhosting.F(subhosting.ProjectDeploymentNewParamsAssetsFileEncodingBase64),
					GitSha1:  subhosting.F("string"),
					Kind:     subhosting.F(subhosting.ProjectDeploymentNewParamsAssetsFileKindFile),
				}),
				"images/cat2.png": subhosting.ProjectDeploymentNewParamsAssetsFile(subhosting.ProjectDeploymentNewParamsAssetsFile{
					Content:  subhosting.F("string"),
					Encoding: subhosting.F(subhosting.ProjectDeploymentNewParamsAssetsFileEncodingUtf8),
					GitSha1:  subhosting.F("5c4f8729e5c30a91a890e24d7285e89f418c637b"),
					Kind:     subhosting.F(subhosting.ProjectDeploymentNewParamsAssetsFileKindFile),
				}),
				"symlink.png": subhosting.ProjectDeploymentNewParamsAssetsSymlink(subhosting.ProjectDeploymentNewParamsAssetsSymlink{
					Target: subhosting.F("images/cat1.png"),
					Kind:   subhosting.F(subhosting.ProjectDeploymentNewParamsAssetsSymlinkKindSymlink),
				}),
			}),
			EntryPointURL: subhosting.F("main.ts"),
			EnvVars: subhosting.F(map[string]string{
				"MY_ENV": "hey",
			}),
			CompilerOptions: subhosting.F(subhosting.ProjectDeploymentNewParamsCompilerOptions{
				Jsx:                subhosting.F("string"),
				JsxFactory:         subhosting.F("string"),
				JsxFragmentFactory: subhosting.F("string"),
				JsxImportSource:    subhosting.F("string"),
			}),
			Databases: subhosting.F(map[string]string{
				"default": "5b484959-cba2-482d-95ab-ba592784af80",
			}),
			Description:  subhosting.F("My first deployment"),
			ImportMapURL: subhosting.Null[string](),
			LockFileURL:  subhosting.Null[string](),
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

func TestProjectDeploymentListWithOptionalParams(t *testing.T) {
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
	_, err := client.Projects.Deployments.List(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		subhosting.ProjectDeploymentListParams{
			Limit: subhosting.F(int64(1)),
			Order: subhosting.F("string"),
			Page:  subhosting.F(int64(1)),
			Q:     subhosting.F("string"),
			Sort:  subhosting.F("string"),
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
