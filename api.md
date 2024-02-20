# Shared Response Types

- <a href="https://pkg.go.dev/github.com/denoland/subhosting-go/internal/shared">shared</a>.<a href="https://pkg.go.dev/github.com/denoland/subhosting-go/internal/shared#Analytics">Analytics</a>
- <a href="https://pkg.go.dev/github.com/denoland/subhosting-go/internal/shared">shared</a>.<a href="https://pkg.go.dev/github.com/denoland/subhosting-go/internal/shared#Deployment">Deployment</a>
- <a href="https://pkg.go.dev/github.com/denoland/subhosting-go/internal/shared">shared</a>.<a href="https://pkg.go.dev/github.com/denoland/subhosting-go/internal/shared#Domain">Domain</a>
- <a href="https://pkg.go.dev/github.com/denoland/subhosting-go/internal/shared">shared</a>.<a href="https://pkg.go.dev/github.com/denoland/subhosting-go/internal/shared#KvDatabase">KvDatabase</a>
- <a href="https://pkg.go.dev/github.com/denoland/subhosting-go/internal/shared">shared</a>.<a href="https://pkg.go.dev/github.com/denoland/subhosting-go/internal/shared#Project">Project</a>

# Organizations

Response Types:

- <a href="https://pkg.go.dev/github.com/denoland/subhosting-go">subhosting</a>.<a href="https://pkg.go.dev/github.com/denoland/subhosting-go#Organization">Organization</a>

Methods:

- <code title="get /organizations/{organizationId}">client.Organizations.<a href="https://pkg.go.dev/github.com/denoland/subhosting-go#OrganizationService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, organizationID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/denoland/subhosting-go">subhosting</a>.<a href="https://pkg.go.dev/github.com/denoland/subhosting-go#Organization">Organization</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Analytics

Methods:

- <code title="get /organizations/{organizationId}/analytics">client.Organizations.Analytics.<a href="https://pkg.go.dev/github.com/denoland/subhosting-go#OrganizationAnalyticsService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, organizationID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/denoland/subhosting-go">subhosting</a>.<a href="https://pkg.go.dev/github.com/denoland/subhosting-go#OrganizationAnalyticsGetParams">OrganizationAnalyticsGetParams</a>) (<a href="https://pkg.go.dev/github.com/denoland/subhosting-go/internal/shared">shared</a>.<a href="https://pkg.go.dev/github.com/denoland/subhosting-go/internal/shared#Analytics">Analytics</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Projects

Methods:

- <code title="post /organizations/{organizationId}/projects">client.Organizations.Projects.<a href="https://pkg.go.dev/github.com/denoland/subhosting-go#OrganizationProjectService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, organizationID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/denoland/subhosting-go">subhosting</a>.<a href="https://pkg.go.dev/github.com/denoland/subhosting-go#OrganizationProjectNewParams">OrganizationProjectNewParams</a>) (<a href="https://pkg.go.dev/github.com/denoland/subhosting-go/internal/shared">shared</a>.<a href="https://pkg.go.dev/github.com/denoland/subhosting-go/internal/shared#Project">Project</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /organizations/{organizationId}/projects">client.Organizations.Projects.<a href="https://pkg.go.dev/github.com/denoland/subhosting-go#OrganizationProjectService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, organizationID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/denoland/subhosting-go">subhosting</a>.<a href="https://pkg.go.dev/github.com/denoland/subhosting-go#OrganizationProjectListParams">OrganizationProjectListParams</a>) ([]<a href="https://pkg.go.dev/github.com/denoland/subhosting-go/internal/shared">shared</a>.<a href="https://pkg.go.dev/github.com/denoland/subhosting-go/internal/shared#Project">Project</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Databases

Methods:

- <code title="post /organizations/{organizationId}/databases">client.Organizations.Databases.<a href="https://pkg.go.dev/github.com/denoland/subhosting-go#OrganizationDatabaseService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, organizationID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/denoland/subhosting-go">subhosting</a>.<a href="https://pkg.go.dev/github.com/denoland/subhosting-go#OrganizationDatabaseNewParams">OrganizationDatabaseNewParams</a>) (<a href="https://pkg.go.dev/github.com/denoland/subhosting-go/internal/shared">shared</a>.<a href="https://pkg.go.dev/github.com/denoland/subhosting-go/internal/shared#KvDatabase">KvDatabase</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /organizations/{organizationId}/databases">client.Organizations.Databases.<a href="https://pkg.go.dev/github.com/denoland/subhosting-go#OrganizationDatabaseService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, organizationID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/denoland/subhosting-go">subhosting</a>.<a href="https://pkg.go.dev/github.com/denoland/subhosting-go#OrganizationDatabaseListParams">OrganizationDatabaseListParams</a>) ([]<a href="https://pkg.go.dev/github.com/denoland/subhosting-go/internal/shared">shared</a>.<a href="https://pkg.go.dev/github.com/denoland/subhosting-go/internal/shared#KvDatabase">KvDatabase</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Domains

Methods:

- <code title="post /organizations/{organizationId}/domains">client.Organizations.Domains.<a href="https://pkg.go.dev/github.com/denoland/subhosting-go#OrganizationDomainService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, organizationID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/denoland/subhosting-go">subhosting</a>.<a href="https://pkg.go.dev/github.com/denoland/subhosting-go#OrganizationDomainNewParams">OrganizationDomainNewParams</a>) (<a href="https://pkg.go.dev/github.com/denoland/subhosting-go/internal/shared">shared</a>.<a href="https://pkg.go.dev/github.com/denoland/subhosting-go/internal/shared#Domain">Domain</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /organizations/{organizationId}/domains">client.Organizations.Domains.<a href="https://pkg.go.dev/github.com/denoland/subhosting-go#OrganizationDomainService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, organizationID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/denoland/subhosting-go">subhosting</a>.<a href="https://pkg.go.dev/github.com/denoland/subhosting-go#OrganizationDomainListParams">OrganizationDomainListParams</a>) ([]<a href="https://pkg.go.dev/github.com/denoland/subhosting-go/internal/shared">shared</a>.<a href="https://pkg.go.dev/github.com/denoland/subhosting-go/internal/shared#Domain">Domain</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Databases

Methods:

- <code title="patch /databases/{databaseId}">client.Databases.<a href="https://pkg.go.dev/github.com/denoland/subhosting-go#DatabaseService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, databaseID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/denoland/subhosting-go">subhosting</a>.<a href="https://pkg.go.dev/github.com/denoland/subhosting-go#DatabaseUpdateParams">DatabaseUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/denoland/subhosting-go/internal/shared">shared</a>.<a href="https://pkg.go.dev/github.com/denoland/subhosting-go/internal/shared#KvDatabase">KvDatabase</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Projects

Methods:

- <code title="patch /projects/{projectId}">client.Projects.<a href="https://pkg.go.dev/github.com/denoland/subhosting-go#ProjectService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, projectID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/denoland/subhosting-go">subhosting</a>.<a href="https://pkg.go.dev/github.com/denoland/subhosting-go#ProjectUpdateParams">ProjectUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/denoland/subhosting-go/internal/shared">shared</a>.<a href="https://pkg.go.dev/github.com/denoland/subhosting-go/internal/shared#Project">Project</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /projects/{projectId}">client.Projects.<a href="https://pkg.go.dev/github.com/denoland/subhosting-go#ProjectService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, projectID <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="get /projects/{projectId}">client.Projects.<a href="https://pkg.go.dev/github.com/denoland/subhosting-go#ProjectService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, projectID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/denoland/subhosting-go/internal/shared">shared</a>.<a href="https://pkg.go.dev/github.com/denoland/subhosting-go/internal/shared#Project">Project</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Analytics

Methods:

- <code title="get /projects/{projectId}/analytics">client.Projects.Analytics.<a href="https://pkg.go.dev/github.com/denoland/subhosting-go#ProjectAnalyticsService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, projectID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/denoland/subhosting-go">subhosting</a>.<a href="https://pkg.go.dev/github.com/denoland/subhosting-go#ProjectAnalyticsGetParams">ProjectAnalyticsGetParams</a>) (<a href="https://pkg.go.dev/github.com/denoland/subhosting-go/internal/shared">shared</a>.<a href="https://pkg.go.dev/github.com/denoland/subhosting-go/internal/shared#Analytics">Analytics</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Deployments

Methods:

- <code title="post /projects/{projectId}/deployments">client.Projects.Deployments.<a href="https://pkg.go.dev/github.com/denoland/subhosting-go#ProjectDeploymentService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, projectID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/denoland/subhosting-go">subhosting</a>.<a href="https://pkg.go.dev/github.com/denoland/subhosting-go#ProjectDeploymentNewParams">ProjectDeploymentNewParams</a>) (<a href="https://pkg.go.dev/github.com/denoland/subhosting-go/internal/shared">shared</a>.<a href="https://pkg.go.dev/github.com/denoland/subhosting-go/internal/shared#Deployment">Deployment</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /projects/{projectId}/deployments">client.Projects.Deployments.<a href="https://pkg.go.dev/github.com/denoland/subhosting-go#ProjectDeploymentService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, projectID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/denoland/subhosting-go">subhosting</a>.<a href="https://pkg.go.dev/github.com/denoland/subhosting-go#ProjectDeploymentListParams">ProjectDeploymentListParams</a>) ([]<a href="https://pkg.go.dev/github.com/denoland/subhosting-go/internal/shared">shared</a>.<a href="https://pkg.go.dev/github.com/denoland/subhosting-go/internal/shared#Deployment">Deployment</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Deployments

Methods:

- <code title="delete /deployments/{deploymentId}">client.Deployments.<a href="https://pkg.go.dev/github.com/denoland/subhosting-go#DeploymentService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, deploymentID <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="get /deployments/{deploymentId}">client.Deployments.<a href="https://pkg.go.dev/github.com/denoland/subhosting-go#DeploymentService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, deploymentID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/denoland/subhosting-go/internal/shared">shared</a>.<a href="https://pkg.go.dev/github.com/denoland/subhosting-go/internal/shared#Deployment">Deployment</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /deployments/{deploymentId}/redeploy">client.Deployments.<a href="https://pkg.go.dev/github.com/denoland/subhosting-go#DeploymentService.Redeploy">Redeploy</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, deploymentID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/denoland/subhosting-go">subhosting</a>.<a href="https://pkg.go.dev/github.com/denoland/subhosting-go#DeploymentRedeployParams">DeploymentRedeployParams</a>) (<a href="https://pkg.go.dev/github.com/denoland/subhosting-go/internal/shared">shared</a>.<a href="https://pkg.go.dev/github.com/denoland/subhosting-go/internal/shared#Deployment">Deployment</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## BuildLogs

Response Types:

- <a href="https://pkg.go.dev/github.com/denoland/subhosting-go">subhosting</a>.<a href="https://pkg.go.dev/github.com/denoland/subhosting-go#DeploymentBuildLogGetResponse">DeploymentBuildLogGetResponse</a>

Methods:

- <code title="get /deployments/{deploymentId}/build_logs">client.Deployments.BuildLogs.<a href="https://pkg.go.dev/github.com/denoland/subhosting-go#DeploymentBuildLogService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, deploymentID <a href="https://pkg.go.dev/builtin#string">string</a>) ([]<a href="https://pkg.go.dev/github.com/denoland/subhosting-go">subhosting</a>.<a href="https://pkg.go.dev/github.com/denoland/subhosting-go#DeploymentBuildLogGetResponse">DeploymentBuildLogGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## AppLogs

Response Types:

- <a href="https://pkg.go.dev/github.com/denoland/subhosting-go">subhosting</a>.<a href="https://pkg.go.dev/github.com/denoland/subhosting-go#DeploymentAppLogGetResponse">DeploymentAppLogGetResponse</a>

Methods:

- <code title="get /deployments/{deploymentId}/app_logs">client.Deployments.AppLogs.<a href="https://pkg.go.dev/github.com/denoland/subhosting-go#DeploymentAppLogService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, deploymentID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/denoland/subhosting-go">subhosting</a>.<a href="https://pkg.go.dev/github.com/denoland/subhosting-go#DeploymentAppLogGetParams">DeploymentAppLogGetParams</a>) ([]<a href="https://pkg.go.dev/github.com/denoland/subhosting-go">subhosting</a>.<a href="https://pkg.go.dev/github.com/denoland/subhosting-go#DeploymentAppLogGetResponse">DeploymentAppLogGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Domains

Methods:

- <code title="patch /domains/{domainId}">client.Domains.<a href="https://pkg.go.dev/github.com/denoland/subhosting-go#DomainService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, domainID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/denoland/subhosting-go">subhosting</a>.<a href="https://pkg.go.dev/github.com/denoland/subhosting-go#DomainUpdateParams">DomainUpdateParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="delete /domains/{domainId}">client.Domains.<a href="https://pkg.go.dev/github.com/denoland/subhosting-go#DomainService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, domainID <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="get /domains/{domainId}">client.Domains.<a href="https://pkg.go.dev/github.com/denoland/subhosting-go#DomainService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, domainID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/denoland/subhosting-go/internal/shared">shared</a>.<a href="https://pkg.go.dev/github.com/denoland/subhosting-go/internal/shared#Domain">Domain</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /domains/{domainId}/verify">client.Domains.<a href="https://pkg.go.dev/github.com/denoland/subhosting-go#DomainService.Verify">Verify</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, domainID <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

## Certificates

Methods:

- <code title="post /domains/{domainId}/certificates">client.Domains.Certificates.<a href="https://pkg.go.dev/github.com/denoland/subhosting-go#DomainCertificateService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, domainID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/denoland/subhosting-go">subhosting</a>.<a href="https://pkg.go.dev/github.com/denoland/subhosting-go#DomainCertificateNewParams">DomainCertificateNewParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="post /domains/{domainId}/certificates/provision">client.Domains.Certificates.<a href="https://pkg.go.dev/github.com/denoland/subhosting-go#DomainCertificateService.Provision">Provision</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, domainID <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
