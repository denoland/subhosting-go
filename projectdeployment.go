// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package subhosting

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/denoland/subhosting-go/internal/apijson"
	"github.com/denoland/subhosting-go/internal/apiquery"
	"github.com/denoland/subhosting-go/internal/param"
	"github.com/denoland/subhosting-go/internal/requestconfig"
	"github.com/denoland/subhosting-go/internal/shared"
	"github.com/denoland/subhosting-go/option"
)

// ProjectDeploymentService contains methods and other services that help with
// interacting with the subhosting API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewProjectDeploymentService] method
// instead.
type ProjectDeploymentService struct {
	Options []option.RequestOption
}

// NewProjectDeploymentService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewProjectDeploymentService(opts ...option.RequestOption) (r *ProjectDeploymentService) {
	r = &ProjectDeploymentService{}
	r.Options = opts
	return
}

// Create a deployment
//
// This API initiates a build process for a new deployment.
//
// Note that this process is asynchronous; the completion of this API doesn't mean
// the deployment is ready. In order to keep track of the progress of the build,
// call the "Get build logs of a deployment" API or the "Get deployment details"
// API.
func (r *ProjectDeploymentService) New(ctx context.Context, projectID string, body ProjectDeploymentNewParams, opts ...option.RequestOption) (res *shared.Deployment, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("projects/%s/deployments", projectID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// List deployments of a project
//
// This API returns a list of deployments belonging to the specified project in a
// pagenated manner.
//
// The URLs for the next, previous, first, and last page are returned in the `Link`
// header of the response, if any.
func (r *ProjectDeploymentService) List(ctx context.Context, projectID string, query ProjectDeploymentListParams, opts ...option.RequestOption) (res *[]shared.Deployment, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("projects/%s/deployments", projectID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type ProjectDeploymentNewParams struct {
	// A map whose key represents a file path, and the value is an asset that composes
	// the deployment.
	//
	// Each asset is one of the following three kinds:
	//
	// 1. A file with content data (which is UTF-8 for text, or base64 for binary)
	// 2. A file with a hash
	// 3. A symbolic link to another asset
	//
	// Assets that were uploaded in some of the previous deployments don't need to be
	// uploaded again. In this case, in order to identify the asset, just provide the
	// SHA-1 hash of the content.
	Assets param.Field[map[string]ProjectDeploymentNewParamsAssetsUnion] `json:"assets,required"`
	// An URL of the entry point of the application. This is the file that will be
	// executed when the deployment is invoked.
	EntryPointURL param.Field[string] `json:"entryPointUrl,required"`
	// A dictionary of environment variables to be set in the runtime environment of
	// the deployment.
	EnvVars param.Field[map[string]string] `json:"envVars,required"`
	// Compiler options to be used when building the deployment.
	//
	// If `null` is given, Deno's config file (i.e. `deno.json` or `deno.jsonc`) will
	// be auto-discovered, which may contain a `compilerOptions` field. If found, that
	// compiler options will be applied.
	//
	// If an empty object `{}` is given,
	// [the default compiler options](https://docs.deno.com/runtime/manual/advanced/typescript/configuration#how-deno-uses-a-configuration-file)
	// will be applied.
	CompilerOptions param.Field[ProjectDeploymentNewParamsCompilerOptions] `json:"compilerOptions"`
	// KV database ID mappings to associate with the deployment.
	//
	// A key represents a KV database name (e.g. `"default"`), and a value is a KV
	// database ID.
	//
	// Currently, only `"default"` database is supported. If any other database name is
	// specified, that will be rejected.
	//
	// If not provided, the deployment will be created with no KV database attached.
	Databases param.Field[map[string]string] `json:"databases" format:"uuid"`
	// A description of the created deployment. If not provided, an empty string will
	// be set.
	Description param.Field[string] `json:"description"`
	// An URL of the import map file.
	//
	// If `null` is given, import map auto-discovery logic will be performed, where it
	// looks for Deno's config file (i.e. `deno.json` or `deno.jsonc`) which may
	// contain an embedded import map or a path to an import map file. If found, that
	// import map will be used.
	//
	// If an empty string is given, no import map will be used.
	ImportMapURL param.Field[string] `json:"importMapUrl"`
	// An URL of the lock file.
	//
	// If `null` is given, lock file auto-discovery logic will be performed, where it
	// looks for Deno's config file (i.e. `deno.json` or `deno.jsonc`) which may
	// contain a path to a lock file or boolean value, such as `"lock": false` or
	// `"lock": "my-lock.lock"`. If a config file is found, the semantics of the lock
	// field is the same as the Deno CLI, so refer to
	// [the CLI doc page](https://docs.deno.com/runtime/manual/basics/modules/integrity_checking#auto-generated-lockfile).
	//
	// If an empty string is given, no lock file will be used.
	LockFileURL param.Field[string] `json:"lockFileUrl"`
}

func (r ProjectDeploymentNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ProjectDeploymentNewParamsAssets struct {
	Content  param.Field[string]                                   `json:"content"`
	Encoding param.Field[ProjectDeploymentNewParamsAssetsEncoding] `json:"encoding"`
	GitSha1  param.Field[string]                                   `json:"gitSha1"`
	Kind     param.Field[ProjectDeploymentNewParamsAssetsKind]     `json:"kind,required"`
	Target   param.Field[string]                                   `json:"target"`
}

func (r ProjectDeploymentNewParamsAssets) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ProjectDeploymentNewParamsAssets) implementsProjectDeploymentNewParamsAssetsUnion() {}

// Satisfied by [ProjectDeploymentNewParamsAssetsObject],
// [ProjectDeploymentNewParamsAssetsObject], [ProjectDeploymentNewParamsAssets].
type ProjectDeploymentNewParamsAssetsUnion interface {
	implementsProjectDeploymentNewParamsAssetsUnion()
}

type ProjectDeploymentNewParamsAssetsObject struct {
	Kind     param.Field[ProjectDeploymentNewParamsAssetsObjectKind]     `json:"kind,required"`
	Content  param.Field[string]                                         `json:"content"`
	Encoding param.Field[ProjectDeploymentNewParamsAssetsObjectEncoding] `json:"encoding"`
	GitSha1  param.Field[string]                                         `json:"gitSha1"`
}

func (r ProjectDeploymentNewParamsAssetsObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ProjectDeploymentNewParamsAssetsObject) implementsProjectDeploymentNewParamsAssetsUnion() {}

type ProjectDeploymentNewParamsAssetsObjectKind string

const (
	ProjectDeploymentNewParamsAssetsObjectKindFile ProjectDeploymentNewParamsAssetsObjectKind = "file"
)

func (r ProjectDeploymentNewParamsAssetsObjectKind) IsKnown() bool {
	switch r {
	case ProjectDeploymentNewParamsAssetsObjectKindFile:
		return true
	}
	return false
}

type ProjectDeploymentNewParamsAssetsObjectEncoding string

const (
	ProjectDeploymentNewParamsAssetsObjectEncodingUtf8   ProjectDeploymentNewParamsAssetsObjectEncoding = "utf-8"
	ProjectDeploymentNewParamsAssetsObjectEncodingBase64 ProjectDeploymentNewParamsAssetsObjectEncoding = "base64"
)

func (r ProjectDeploymentNewParamsAssetsObjectEncoding) IsKnown() bool {
	switch r {
	case ProjectDeploymentNewParamsAssetsObjectEncodingUtf8, ProjectDeploymentNewParamsAssetsObjectEncodingBase64:
		return true
	}
	return false
}

type ProjectDeploymentNewParamsAssetsEncoding string

const (
	ProjectDeploymentNewParamsAssetsEncodingUtf8   ProjectDeploymentNewParamsAssetsEncoding = "utf-8"
	ProjectDeploymentNewParamsAssetsEncodingBase64 ProjectDeploymentNewParamsAssetsEncoding = "base64"
)

func (r ProjectDeploymentNewParamsAssetsEncoding) IsKnown() bool {
	switch r {
	case ProjectDeploymentNewParamsAssetsEncodingUtf8, ProjectDeploymentNewParamsAssetsEncodingBase64:
		return true
	}
	return false
}

type ProjectDeploymentNewParamsAssetsKind string

const (
	ProjectDeploymentNewParamsAssetsKindFile    ProjectDeploymentNewParamsAssetsKind = "file"
	ProjectDeploymentNewParamsAssetsKindSymlink ProjectDeploymentNewParamsAssetsKind = "symlink"
)

func (r ProjectDeploymentNewParamsAssetsKind) IsKnown() bool {
	switch r {
	case ProjectDeploymentNewParamsAssetsKindFile, ProjectDeploymentNewParamsAssetsKindSymlink:
		return true
	}
	return false
}

// Compiler options to be used when building the deployment.
//
// If `null` is given, Deno's config file (i.e. `deno.json` or `deno.jsonc`) will
// be auto-discovered, which may contain a `compilerOptions` field. If found, that
// compiler options will be applied.
//
// If an empty object `{}` is given,
// [the default compiler options](https://docs.deno.com/runtime/manual/advanced/typescript/configuration#how-deno-uses-a-configuration-file)
// will be applied.
type ProjectDeploymentNewParamsCompilerOptions struct {
	Jsx                param.Field[string] `json:"jsx"`
	JsxFactory         param.Field[string] `json:"jsxFactory"`
	JsxFragmentFactory param.Field[string] `json:"jsxFragmentFactory"`
	JsxImportSource    param.Field[string] `json:"jsxImportSource"`
}

func (r ProjectDeploymentNewParamsCompilerOptions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ProjectDeploymentListParams struct {
	// The maximum number of items to return per page.
	Limit param.Field[int64] `query:"limit"`
	// Sort order, either `asc` or `desc`. Defaults to `asc`.
	Order param.Field[string] `query:"order"`
	// The page number to return.
	Page param.Field[int64] `query:"page"`
	// Query by deployment ID
	Q param.Field[string] `query:"q"`
	// The field to sort by, either `id` or `created_at`. Defaults to `created_at`.
	Sort param.Field[string] `query:"sort"`
}

// URLQuery serializes [ProjectDeploymentListParams]'s query parameters as
// `url.Values`.
func (r ProjectDeploymentListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
