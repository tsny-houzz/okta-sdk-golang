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

type ApplicationSettingsNotificationsVpnNetwork struct {
	Connection string `json:"connection,omitempty"`
	Exclude []string `json:"exclude,omitempty"`
	Include []string `json:"include,omitempty"`
}
