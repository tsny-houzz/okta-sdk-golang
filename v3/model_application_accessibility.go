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

type ApplicationAccessibility struct {
	ErrorRedirectUrl string `json:"errorRedirectUrl,omitempty"`
	LoginRedirectUrl string `json:"loginRedirectUrl,omitempty"`
	SelfService bool `json:"selfService,omitempty"`
}
