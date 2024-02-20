// File generated from our OpenAPI spec by Stainless.

package subhosting

import (
	"os"

	"github.com/denoland/subhosting-go/option"
)

// Client creates a struct with services and top level methods that help with
// interacting with the subhosting API. You should not instantiate this client
// directly, and instead use the [NewClient] method instead.
type Client struct {
	Options       []option.RequestOption
	Organizations *OrganizationService
	Databases     *DatabaseService
	Projects      *ProjectService
	Deployments   *DeploymentService
	Domains       *DomainService
}

// NewClient generates a new client with the default option read from the
// environment (DEPLOY_ACCESS_TOKEN). The option passed in as arguments are applied
// after these default arguments, and all option will be passed down to the
// services and requests that this client makes.
func NewClient(opts ...option.RequestOption) (r *Client) {
	defaults := []option.RequestOption{option.WithEnvironmentProduction()}
	if o, ok := os.LookupEnv("DEPLOY_ACCESS_TOKEN"); ok {
		defaults = append(defaults, option.WithBearerToken(o))
	}
	opts = append(defaults, opts...)

	r = &Client{Options: opts}

	r.Organizations = NewOrganizationService(opts...)
	r.Databases = NewDatabaseService(opts...)
	r.Projects = NewProjectService(opts...)
	r.Deployments = NewDeploymentService(opts...)
	r.Domains = NewDomainService(opts...)

	return
}
