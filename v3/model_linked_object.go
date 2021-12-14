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

type LinkedObject struct {
	Links map[string]interface{} `json:"_links,omitempty"`
	Associated *LinkedObjectDetails `json:"associated,omitempty"`
	Primary *LinkedObjectDetails `json:"primary,omitempty"`
}
