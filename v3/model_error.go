/*
 * Okta API
 *
 * Allows customers to easily access the Okta API
 *
 * API version: 2.8.1
 * Contact: devex-public@okta.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type ModelError struct {
	// An Okta code for this type of error
	ErrorCode string `json:"errorCode,omitempty"`
	// A short description of what caused this error. Sometimes this contains dynamically-generated information about your specific error.
	ErrorSummary string `json:"errorSummary,omitempty"`
	// An Okta code for this type of error
	ErrorLink string `json:"errorLink,omitempty"`
	// A unique identifier for this error. This can be used by Okta Support to help with troubleshooting.
	ErrorId string `json:"errorId,omitempty"`
	ErrorCauses []ErrorErrorCauses `json:"errorCauses,omitempty"`
}
