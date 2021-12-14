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

type UserSchema struct {
	Id string `json:"id,omitempty"`
	Schema string `json:"$schema,omitempty"`
	Name string `json:"name,omitempty"`
	Title string `json:"title,omitempty"`
	LastUpdated string `json:"lastUpdated,omitempty"`
	Created string `json:"created,omitempty"`
	Definitions *UserSchemaDefinitions `json:"definitions,omitempty"`
	Type_ string `json:"type,omitempty"`
	Properties *UserSchemaProperties `json:"properties,omitempty"`
	Links map[string]interface{} `json:"_links,omitempty"`
}
