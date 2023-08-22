# \PolicyApi

All URIs are relative to *https://subdomain.okta.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActivatePolicy**](PolicyApi.md#ActivatePolicy) | **Post** /api/v1/policies/{policyId}/lifecycle/activate | Activate a Policy
[**ActivatePolicyRule**](PolicyApi.md#ActivatePolicyRule) | **Post** /api/v1/policies/{policyId}/rules/{ruleId}/lifecycle/activate | Activate a Policy Rule
[**ClonePolicy**](PolicyApi.md#ClonePolicy) | **Post** /api/v1/policies/{policyId}/clone | Clone an existing Policy
[**CreatePolicy**](PolicyApi.md#CreatePolicy) | **Post** /api/v1/policies | Create a Policy
[**CreatePolicyRule**](PolicyApi.md#CreatePolicyRule) | **Post** /api/v1/policies/{policyId}/rules | Create a Policy Rule
[**CreatePolicySimulation**](PolicyApi.md#CreatePolicySimulation) | **Post** /api/v1/policies/simulate | Create a Policy Simulation
[**DeactivatePolicy**](PolicyApi.md#DeactivatePolicy) | **Post** /api/v1/policies/{policyId}/lifecycle/deactivate | Deactivate a Policy
[**DeactivatePolicyRule**](PolicyApi.md#DeactivatePolicyRule) | **Post** /api/v1/policies/{policyId}/rules/{ruleId}/lifecycle/deactivate | Deactivate a Policy Rule
[**DeletePolicy**](PolicyApi.md#DeletePolicy) | **Delete** /api/v1/policies/{policyId} | Delete a Policy
[**DeletePolicyResourceMapping**](PolicyApi.md#DeletePolicyResourceMapping) | **Delete** /api/v1/policies/{policyId}/mappings/{mappingId} | Delete a policy resource Mapping
[**DeletePolicyRule**](PolicyApi.md#DeletePolicyRule) | **Delete** /api/v1/policies/{policyId}/rules/{ruleId} | Delete a Policy Rule
[**GetPolicy**](PolicyApi.md#GetPolicy) | **Get** /api/v1/policies/{policyId} | Retrieve a Policy
[**GetPolicyMapping**](PolicyApi.md#GetPolicyMapping) | **Get** /api/v1/policies/{policyId}/mappings/{mappingId} | Retrieve a policy resource Mapping
[**GetPolicyRule**](PolicyApi.md#GetPolicyRule) | **Get** /api/v1/policies/{policyId}/rules/{ruleId} | Retrieve a Policy Rule
[**ListPolicies**](PolicyApi.md#ListPolicies) | **Get** /api/v1/policies | List all Policies
[**ListPolicyApps**](PolicyApi.md#ListPolicyApps) | **Get** /api/v1/policies/{policyId}/app | List all Applications mapped to a Policy
[**ListPolicyMappings**](PolicyApi.md#ListPolicyMappings) | **Get** /api/v1/policies/{policyId}/mappings | List all resources mapped to a Policy
[**ListPolicyRules**](PolicyApi.md#ListPolicyRules) | **Get** /api/v1/policies/{policyId}/rules | List all Policy Rules
[**MapResourceToPolicy**](PolicyApi.md#MapResourceToPolicy) | **Post** /api/v1/policies/{policyId}/mappings | Map a resource to a Policy
[**ReplacePolicy**](PolicyApi.md#ReplacePolicy) | **Put** /api/v1/policies/{policyId} | Replace a Policy
[**ReplacePolicyRule**](PolicyApi.md#ReplacePolicyRule) | **Put** /api/v1/policies/{policyId}/rules/{ruleId} | Replace a Policy Rule



## ActivatePolicy

> ActivatePolicy(ctx, policyId).Execute()

Activate a Policy



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {
    policyId := "00plrilJ7jZ66Gn0X0g3" // string | `id` of the Policy

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PolicyApi.ActivatePolicy(context.Background(), policyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PolicyApi.ActivatePolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyId** | **string** | &#x60;id&#x60; of the Policy | 

### Other Parameters

Other parameters are passed through a pointer to a apiActivatePolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActivatePolicyRule

> ActivatePolicyRule(ctx, policyId, ruleId).Execute()

Activate a Policy Rule



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {
    policyId := "00plrilJ7jZ66Gn0X0g3" // string | `id` of the Policy
    ruleId := "ruld3hJ7jZh4fn0st0g3" // string | `id` of the Policy Rule

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PolicyApi.ActivatePolicyRule(context.Background(), policyId, ruleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PolicyApi.ActivatePolicyRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyId** | **string** | &#x60;id&#x60; of the Policy | 
**ruleId** | **string** | &#x60;id&#x60; of the Policy Rule | 

### Other Parameters

Other parameters are passed through a pointer to a apiActivatePolicyRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ClonePolicy

> ListPolicies200ResponseInner ClonePolicy(ctx, policyId).Execute()

Clone an existing Policy



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {
    policyId := "00plrilJ7jZ66Gn0X0g3" // string | `id` of the Policy

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PolicyApi.ClonePolicy(context.Background(), policyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PolicyApi.ClonePolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ClonePolicy`: ListPolicies200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `PolicyApi.ClonePolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyId** | **string** | &#x60;id&#x60; of the Policy | 

### Other Parameters

Other parameters are passed through a pointer to a apiClonePolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ListPolicies200ResponseInner**](ListPolicies200ResponseInner.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePolicy

> ListPolicies200ResponseInner CreatePolicy(ctx).Policy(policy).Activate(activate).Execute()

Create a Policy



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {
    policy := openapiclient.listPolicies_200_response_inner{AccessPolicy: openapiclient.NewAccessPolicy()} // ListPolicies200ResponseInner | 
    activate := true // bool |  (optional) (default to true)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PolicyApi.CreatePolicy(context.Background()).Policy(policy).Activate(activate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PolicyApi.CreatePolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreatePolicy`: ListPolicies200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `PolicyApi.CreatePolicy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **policy** | [**ListPolicies200ResponseInner**](ListPolicies200ResponseInner.md) |  | 
 **activate** | **bool** |  | [default to true]

### Return type

[**ListPolicies200ResponseInner**](ListPolicies200ResponseInner.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePolicyRule

> ListPolicyRules200ResponseInner CreatePolicyRule(ctx, policyId).PolicyRule(policyRule).Execute()

Create a Policy Rule



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {
    policyId := "00plrilJ7jZ66Gn0X0g3" // string | `id` of the Policy
    policyRule := openapiclient.listPolicyRules_200_response_inner{AccessPolicyRule: openapiclient.NewAccessPolicyRule()} // ListPolicyRules200ResponseInner | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PolicyApi.CreatePolicyRule(context.Background(), policyId).PolicyRule(policyRule).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PolicyApi.CreatePolicyRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreatePolicyRule`: ListPolicyRules200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `PolicyApi.CreatePolicyRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyId** | **string** | &#x60;id&#x60; of the Policy | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreatePolicyRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **policyRule** | [**ListPolicyRules200ResponseInner**](ListPolicyRules200ResponseInner.md) |  | 

### Return type

[**ListPolicyRules200ResponseInner**](ListPolicyRules200ResponseInner.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePolicySimulation

> []SimulatePolicyEvaluations CreatePolicySimulation(ctx).SimulatePolicy(simulatePolicy).Expand(expand).Execute()

Create a Policy Simulation



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {
    simulatePolicy := []openapiclient.SimulatePolicyBody{*openapiclient.NewSimulatePolicyBody("AppInstance_example")} // []SimulatePolicyBody | 
    expand := "expand=EVALUATED&expand=RULE" // string | Use `expand=EVALUATED` to include a list of evaluated but not matched policies and policy rules. Use `expand=RULE` to include details about why a rule condition was (not) matched. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PolicyApi.CreatePolicySimulation(context.Background()).SimulatePolicy(simulatePolicy).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PolicyApi.CreatePolicySimulation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreatePolicySimulation`: []SimulatePolicyEvaluations
    fmt.Fprintf(os.Stdout, "Response from `PolicyApi.CreatePolicySimulation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePolicySimulationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **simulatePolicy** | [**[]SimulatePolicyBody**](SimulatePolicyBody.md) |  | 
 **expand** | **string** | Use &#x60;expand&#x3D;EVALUATED&#x60; to include a list of evaluated but not matched policies and policy rules. Use &#x60;expand&#x3D;RULE&#x60; to include details about why a rule condition was (not) matched. | 

### Return type

[**[]SimulatePolicyEvaluations**](SimulatePolicyEvaluations.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeactivatePolicy

> DeactivatePolicy(ctx, policyId).Execute()

Deactivate a Policy



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {
    policyId := "00plrilJ7jZ66Gn0X0g3" // string | `id` of the Policy

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PolicyApi.DeactivatePolicy(context.Background(), policyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PolicyApi.DeactivatePolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyId** | **string** | &#x60;id&#x60; of the Policy | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeactivatePolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeactivatePolicyRule

> DeactivatePolicyRule(ctx, policyId, ruleId).Execute()

Deactivate a Policy Rule



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {
    policyId := "00plrilJ7jZ66Gn0X0g3" // string | `id` of the Policy
    ruleId := "ruld3hJ7jZh4fn0st0g3" // string | `id` of the Policy Rule

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PolicyApi.DeactivatePolicyRule(context.Background(), policyId, ruleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PolicyApi.DeactivatePolicyRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyId** | **string** | &#x60;id&#x60; of the Policy | 
**ruleId** | **string** | &#x60;id&#x60; of the Policy Rule | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeactivatePolicyRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePolicy

> DeletePolicy(ctx, policyId).Execute()

Delete a Policy



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {
    policyId := "00plrilJ7jZ66Gn0X0g3" // string | `id` of the Policy

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PolicyApi.DeletePolicy(context.Background(), policyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PolicyApi.DeletePolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyId** | **string** | &#x60;id&#x60; of the Policy | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePolicyResourceMapping

> DeletePolicyResourceMapping(ctx, policyId, mappingId).Execute()

Delete a policy resource Mapping



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {
    policyId := "00plrilJ7jZ66Gn0X0g3" // string | `id` of the Policy
    mappingId := "maplr2rLjZ6NsGn1P0g3" // string | `id` of the policy resource Mapping

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PolicyApi.DeletePolicyResourceMapping(context.Background(), policyId, mappingId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PolicyApi.DeletePolicyResourceMapping``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyId** | **string** | &#x60;id&#x60; of the Policy | 
**mappingId** | **string** | &#x60;id&#x60; of the policy resource Mapping | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePolicyResourceMappingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePolicyRule

> DeletePolicyRule(ctx, policyId, ruleId).Execute()

Delete a Policy Rule



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {
    policyId := "00plrilJ7jZ66Gn0X0g3" // string | `id` of the Policy
    ruleId := "ruld3hJ7jZh4fn0st0g3" // string | `id` of the Policy Rule

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PolicyApi.DeletePolicyRule(context.Background(), policyId, ruleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PolicyApi.DeletePolicyRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyId** | **string** | &#x60;id&#x60; of the Policy | 
**ruleId** | **string** | &#x60;id&#x60; of the Policy Rule | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePolicyRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPolicy

> ListPolicies200ResponseInner GetPolicy(ctx, policyId).Expand(expand).Execute()

Retrieve a Policy



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {
    policyId := "00plrilJ7jZ66Gn0X0g3" // string | `id` of the Policy
    expand := "expand_example" // string |  (optional) (default to "")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PolicyApi.GetPolicy(context.Background(), policyId).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PolicyApi.GetPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPolicy`: ListPolicies200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `PolicyApi.GetPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyId** | **string** | &#x60;id&#x60; of the Policy | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **expand** | **string** |  | [default to &quot;&quot;]

### Return type

[**ListPolicies200ResponseInner**](ListPolicies200ResponseInner.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPolicyMapping

> PolicyMapping GetPolicyMapping(ctx, policyId, mappingId).Execute()

Retrieve a policy resource Mapping



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {
    policyId := "00plrilJ7jZ66Gn0X0g3" // string | `id` of the Policy
    mappingId := "maplr2rLjZ6NsGn1P0g3" // string | `id` of the policy resource Mapping

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PolicyApi.GetPolicyMapping(context.Background(), policyId, mappingId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PolicyApi.GetPolicyMapping``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPolicyMapping`: PolicyMapping
    fmt.Fprintf(os.Stdout, "Response from `PolicyApi.GetPolicyMapping`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyId** | **string** | &#x60;id&#x60; of the Policy | 
**mappingId** | **string** | &#x60;id&#x60; of the policy resource Mapping | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPolicyMappingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**PolicyMapping**](PolicyMapping.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPolicyRule

> ListPolicyRules200ResponseInner GetPolicyRule(ctx, policyId, ruleId).Execute()

Retrieve a Policy Rule



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {
    policyId := "00plrilJ7jZ66Gn0X0g3" // string | `id` of the Policy
    ruleId := "ruld3hJ7jZh4fn0st0g3" // string | `id` of the Policy Rule

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PolicyApi.GetPolicyRule(context.Background(), policyId, ruleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PolicyApi.GetPolicyRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPolicyRule`: ListPolicyRules200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `PolicyApi.GetPolicyRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyId** | **string** | &#x60;id&#x60; of the Policy | 
**ruleId** | **string** | &#x60;id&#x60; of the Policy Rule | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPolicyRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ListPolicyRules200ResponseInner**](ListPolicyRules200ResponseInner.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPolicies

> []ListPolicies200ResponseInner ListPolicies(ctx).Type_(type_).Status(status).Expand(expand).Execute()

List all Policies



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {
    type_ := "type__example" // string | 
    status := "status_example" // string |  (optional)
    expand := "expand_example" // string |  (optional) (default to "")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PolicyApi.ListPolicies(context.Background()).Type_(type_).Status(status).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PolicyApi.ListPolicies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListPolicies`: []ListPolicies200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `PolicyApi.ListPolicies`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListPoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type_** | **string** |  | 
 **status** | **string** |  | 
 **expand** | **string** |  | [default to &quot;&quot;]

### Return type

[**[]ListPolicies200ResponseInner**](ListPolicies200ResponseInner.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPolicyApps

> []ListApplications200ResponseInner ListPolicyApps(ctx, policyId).Execute()

List all Applications mapped to a Policy



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {
    policyId := "00plrilJ7jZ66Gn0X0g3" // string | `id` of the Policy

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PolicyApi.ListPolicyApps(context.Background(), policyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PolicyApi.ListPolicyApps``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListPolicyApps`: []ListApplications200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `PolicyApi.ListPolicyApps`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyId** | **string** | &#x60;id&#x60; of the Policy | 

### Other Parameters

Other parameters are passed through a pointer to a apiListPolicyAppsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]ListApplications200ResponseInner**](ListApplications200ResponseInner.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPolicyMappings

> []PolicyMapping ListPolicyMappings(ctx, policyId).Execute()

List all resources mapped to a Policy



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {
    policyId := "00plrilJ7jZ66Gn0X0g3" // string | `id` of the Policy

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PolicyApi.ListPolicyMappings(context.Background(), policyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PolicyApi.ListPolicyMappings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListPolicyMappings`: []PolicyMapping
    fmt.Fprintf(os.Stdout, "Response from `PolicyApi.ListPolicyMappings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyId** | **string** | &#x60;id&#x60; of the Policy | 

### Other Parameters

Other parameters are passed through a pointer to a apiListPolicyMappingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]PolicyMapping**](PolicyMapping.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPolicyRules

> []ListPolicyRules200ResponseInner ListPolicyRules(ctx, policyId).Execute()

List all Policy Rules



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {
    policyId := "00plrilJ7jZ66Gn0X0g3" // string | `id` of the Policy

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PolicyApi.ListPolicyRules(context.Background(), policyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PolicyApi.ListPolicyRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListPolicyRules`: []ListPolicyRules200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `PolicyApi.ListPolicyRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyId** | **string** | &#x60;id&#x60; of the Policy | 

### Other Parameters

Other parameters are passed through a pointer to a apiListPolicyRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]ListPolicyRules200ResponseInner**](ListPolicyRules200ResponseInner.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MapResourceToPolicy

> PolicyMapping MapResourceToPolicy(ctx, policyId).PolicyMappingRequest(policyMappingRequest).Execute()

Map a resource to a Policy



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {
    policyId := "00plrilJ7jZ66Gn0X0g3" // string | `id` of the Policy
    policyMappingRequest := *openapiclient.NewPolicyMappingRequest() // PolicyMappingRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PolicyApi.MapResourceToPolicy(context.Background(), policyId).PolicyMappingRequest(policyMappingRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PolicyApi.MapResourceToPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MapResourceToPolicy`: PolicyMapping
    fmt.Fprintf(os.Stdout, "Response from `PolicyApi.MapResourceToPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyId** | **string** | &#x60;id&#x60; of the Policy | 

### Other Parameters

Other parameters are passed through a pointer to a apiMapResourceToPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **policyMappingRequest** | [**PolicyMappingRequest**](PolicyMappingRequest.md) |  | 

### Return type

[**PolicyMapping**](PolicyMapping.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplacePolicy

> ListPolicies200ResponseInner ReplacePolicy(ctx, policyId).Policy(policy).Execute()

Replace a Policy



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {
    policyId := "00plrilJ7jZ66Gn0X0g3" // string | `id` of the Policy
    policy := openapiclient.listPolicies_200_response_inner{AccessPolicy: openapiclient.NewAccessPolicy()} // ListPolicies200ResponseInner | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PolicyApi.ReplacePolicy(context.Background(), policyId).Policy(policy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PolicyApi.ReplacePolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReplacePolicy`: ListPolicies200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `PolicyApi.ReplacePolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyId** | **string** | &#x60;id&#x60; of the Policy | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplacePolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **policy** | [**ListPolicies200ResponseInner**](ListPolicies200ResponseInner.md) |  | 

### Return type

[**ListPolicies200ResponseInner**](ListPolicies200ResponseInner.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplacePolicyRule

> ListPolicyRules200ResponseInner ReplacePolicyRule(ctx, policyId, ruleId).PolicyRule(policyRule).Execute()

Replace a Policy Rule



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/okta/okta-sdk-golang"
)

func main() {
    policyId := "00plrilJ7jZ66Gn0X0g3" // string | `id` of the Policy
    ruleId := "ruld3hJ7jZh4fn0st0g3" // string | `id` of the Policy Rule
    policyRule := openapiclient.listPolicyRules_200_response_inner{AccessPolicyRule: openapiclient.NewAccessPolicyRule()} // ListPolicyRules200ResponseInner | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PolicyApi.ReplacePolicyRule(context.Background(), policyId, ruleId).PolicyRule(policyRule).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PolicyApi.ReplacePolicyRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReplacePolicyRule`: ListPolicyRules200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `PolicyApi.ReplacePolicyRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyId** | **string** | &#x60;id&#x60; of the Policy | 
**ruleId** | **string** | &#x60;id&#x60; of the Policy Rule | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplacePolicyRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **policyRule** | [**ListPolicyRules200ResponseInner**](ListPolicyRules200ResponseInner.md) |  | 

### Return type

[**ListPolicyRules200ResponseInner**](ListPolicyRules200ResponseInner.md)

### Authorization

[apiToken](../README.md#apiToken), [oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
