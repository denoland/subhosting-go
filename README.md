# Deno Deploy Subhosting REST API client for Go

<a href="https://pkg.go.dev/github.com/denoland/subhosting-go"><img src="https://pkg.go.dev/badge/github.com/denoland/subhosting-go.svg" alt="Go Reference"></a>

This library provides convenient access to the
[Deno Deploy Subhosting](https://deno.com/subhosting) REST API, which allows you
to programmatically deploy untrusted, third-party code into the cloud, from Go.

The REST API documentation can be found
[on apidocs.deno.com](https://apidocs.deno.com/). The full API of this library
can be found in [api.md](api.md).

To learn more about Subhosting,
[check out our documentation](https://docs.deno.com/subhosting/manual).

## Installation

<!-- x-release-please-start-version -->

```go
import (
	"github.com/denoland/subhosting-go" // imported as subhosting
)
```

<!-- x-release-please-end -->

Or to pin the version:

<!-- x-release-please-start-version -->

```sh
go get -u 'github.com/denoland/subhosting-go@v0.1.0-alpha.1'
```

<!-- x-release-please-end -->

## Requirements

This library requires Go 1.18+.

## Usage

Before you begin, you'll need to have a Deno Deploy access token and an ID for
the Deno Deploy organization you're using for Subhosting.

- You can find or create a personal access token
  [in the dashboard here](https://dash.deno.com/account#access-tokens)
- Your org ID can be found near the top of the page on your Deno Deploy
  dashboard [as described here](https://docs.deno.com/subhosting/api)

The code examples below assume your access token is stored in a
`DEPLOY_ACCESS_TOKEN` environment variable and your Deno Deploy org ID is stored
in a `DEPLOY_ORG_ID` environment variable.

```go
package main

import (
	"context"
	"fmt"

	"github.com/denoland/subhosting-go"
	"github.com/denoland/subhosting-go/option"
)

func main() {
	org_id := "<My Org Id>"
	client := subhosting.NewClient(
		option.WithBearerToken("My Bearer Token"), // defaults to os.LookupEnv("DEPLOY_ACCESS_TOKEN")
	)
	organization, err := client.Organizations.Get(context.TODO(), org_id)
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("%+v\n", organization.Name)
}
```

### Request Fields

All request parameters are wrapped in a generic `Field` type, which we use to
distinguish zero values from null or omitted fields.

This prevents accidentally sending a zero value if you forget a required
parameter, and enables explicitly sending `null`, `false`, `''`, or `0` on
optional parameters. Any field not specified is not sent.

To construct fields with values, use the helpers `String()`, `Int()`, `Float()`,
or most commonly, the generic `F[T]()`. To send a null, use `Null[T]()`, and to
send a nonconforming value, use `Raw[T](any)`. For example:

```go
params := FooParams{
	Name: subhosting.F("hello"),

	// Explicitly send `"description": null`
	Description: subhosting.Null[string](),

	Point: subhosting.F(subhosting.Point{
		X: subhosting.Int(0),
		Y: subhosting.Int(1),

		// In cases where the API specifies a given type,
		// but you want to send something else, use `Raw`:
		Z: subhosting.Raw[int64](0.01), // sends a float
	}),
}
```

### Response Objects

All fields in response structs are value types (not pointers or wrappers).

If a given field is `null`, not present, or invalid, the corresponding field
will simply be its zero value.

All response structs also include a special `JSON` field, containing more
detailed information about each property, which you can use like so:

```go
if res.Name == "" {
	// true if `"name"` is either not present or explicitly null
	res.JSON.Name.IsNull()

	// true if the `"name"` key was not present in the repsonse JSON at all
	res.JSON.Name.IsMissing()

	// When the API returns data that cannot be coerced to the expected type:
	if res.JSON.Name.IsInvalid() {
		raw := res.JSON.Name.Raw()

		legacyName := struct{
			First string `json:"first"`
			Last  string `json:"last"`
		}{}
		json.Unmarshal([]byte(raw), &legacyName)
		name = legacyName.First + " " + legacyName.Last
	}
}
```

These `.JSON` structs also include an `Extras` map containing any properties in
the json response that were not specified in the struct. This can be useful for
API features not yet present in the SDK.

```go
body := res.JSON.ExtraFields["my_unexpected_field"].Raw()
```

### RequestOptions

This library uses the functional options pattern. Functions defined in the
`option` package return a `RequestOption`, which is a closure that mutates a
`RequestConfig`. These options can be supplied to the client or at individual
requests. For example:

```go
client := subhosting.NewClient(
	// Adds a header to every request made by the client
	option.WithHeader("X-Some-Header", "custom_header_info"),
)

client.Organizations.Get(context.TODO(), ...,
	// Override the header
	option.WithHeader("X-Some-Header", "some_other_custom_header_info"),
	// Add an undocumented field to the request body, using sjson syntax
	option.WithJSONSet("some.json.path", map[string]string{"my": "object"}),
)
```

The full list of request options is
[here](https://pkg.go.dev/github.com/denoland/subhosting-go/option).

### Pagination

This library provides some conveniences for working with paginated list
endpoints.

You can use `.ListAutoPaging()` methods to iterate through items across all
pages:

```go
// TODO
```

Or you can use simple `.List()` methods to fetch a single page and receive a
standard response object with additional helper methods like `.GetNextPage()`,
e.g.:

```go
// TODO
```

### Errors

When the API returns a non-success status code, we return an error with type
`*subhosting.Error`. This contains the `StatusCode`, `*http.Request`, and
`*http.Response` values of the request, as well as the JSON of the error body
(much like other response objects in the SDK).

To handle errors, we recommend that you use the `errors.As` pattern:

```go
_, err := client.Organizations.Get(context.TODO(), "DEPLOY_ORG_ID")
if err != nil {
	var apierr *subhosting.Error
	if errors.As(err, &apierr) {
		println(string(apierr.DumpRequest(true)))  // Prints the serialized HTTP request
		println(string(apierr.DumpResponse(true))) // Prints the serialized HTTP response
	}
	panic(err.Error()) // GET "/organizations/{organizationId}": 400 Bad Request { ... }
}
```

When other errors occur, they are returned unwrapped; for example, if HTTP
transport fails, you might receive `*url.Error` wrapping `*net.OpError`.

### Timeouts

Requests do not time out by default; use context to configure a timeout for a
request lifecycle.

Note that if a request is [retried](#retries), the context timeout does not
start over. To set a per-retry timeout, use `option.WithRequestTimeout()`.

```go
// This sets the timeout for the request, including all the retries.
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
defer cancel()
client.Organizations.Get(
	ctx,
	"DEPLOY_ORG_ID",
	// This sets the per-retry timeout
	option.WithRequestTimeout(20*time.Second),
)
```

## Retries

Certain errors will be automatically retried 2 times by default, with a short
exponential backoff. We retry by default all connection errors, 408 Request
Timeout, 409 Conflict, 429 Rate Limit, and >=500 Internal errors.

You can use the `WithMaxRetries` option to configure or disable this:

```go
// Configure the default for all requests:
client := subhosting.NewClient(
	option.WithMaxRetries(0), // default is 2
)

// Override per-request:
client.Organizations.Get(
	context.TODO(),
	"DEPLOY_ORG_ID",
	option.WithMaxRetries(5),
)
```

### Middleware

We provide `option.WithMiddleware` which applies the given middleware to
requests.

```go
func Logger(req *http.Request, next option.MiddlewareNext) (res *http.Response, err error) {
	// Before the request
	start := time.Now()
	LogReq(req)

	// Forward the request to the next handler
	res, err = next(req)

	// Handle stuff after the request
	end := time.Now()
	LogRes(res, err, start - end)

    return res, err
}

client := subhosting.NewClient(
	option.WithMiddleware(Logger),
)
```

When multiple middlewares are provided as variadic arguments, the middlewares
are applied left to right. If `option.WithMiddleware` is given multiple times,
for example first in the client then the method, the middleware in the client
will run first and the middleware given in the method will run next.

You may also replace the default `http.Client` with
`option.WithHTTPClient(client)`. Only one http client is accepted (this
overwrites any previous client) and receives requests after any middleware has
been applied.

## Semantic Versioning

This package generally follows [SemVer](https://semver.org/spec/v2.0.0.html)
conventions, though certain backwards-incompatible changes may be released as
minor versions:

1. Changes to library internals which are technically public but not intended or
   documented for external use. _(Please open a GitHub issue to let us know if
   you are relying on such internals)_.
2. Changes that we do not expect to impact the vast majority of users in
   practice.

We take backwards-compatibility seriously and work hard to ensure you can rely
on a smooth upgrade experience.

We are keen for your feedback; please open an
[issue](https://www.github.com/denoland/subhosting-go/issues) with questions,
bugs, or suggestions.

## Contributing

This library is auto-generated with
[Stainless API](https://www.stainlessapi.com/) based on our
[OpenAPI spec](https://apidocs.deno.com/). If you’re interested in contributing
to the readme/documentation, feel free to submit a PR. However, since our
OpenAPI spec is generated from our private Deno Deploy repository, if you’re
interested in contributing code, please provide feedback in the
[issues](https://github.com/denoland/subhosting-go/issues) and we’ll review it.
