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
import (
	"time"
)

type TrustedOrigin struct {
	Links map[string]interface{} `json:"_links,omitempty"`
	Created time.Time `json:"created,omitempty"`
	CreatedBy string `json:"createdBy,omitempty"`
	Id string `json:"id,omitempty"`
	LastUpdated time.Time `json:"lastUpdated,omitempty"`
	LastUpdatedBy string `json:"lastUpdatedBy,omitempty"`
	Name string `json:"name,omitempty"`
	Origin string `json:"origin,omitempty"`
	Scopes []Scope `json:"scopes,omitempty"`
	Status string `json:"status,omitempty"`
}
