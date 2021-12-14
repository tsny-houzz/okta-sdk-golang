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

type OAuth2ScopeConsentGrant struct {
	Embedded map[string]interface{} `json:"_embedded,omitempty"`
	Links map[string]interface{} `json:"_links,omitempty"`
	ClientId string `json:"clientId,omitempty"`
	UserId string `json:"userId,omitempty"`
	Created time.Time `json:"created,omitempty"`
	CreatedBy *OAuth2Actor `json:"createdBy,omitempty"`
	Id string `json:"id,omitempty"`
	Issuer string `json:"issuer,omitempty"`
	LastUpdated time.Time `json:"lastUpdated,omitempty"`
	ScopeId string `json:"scopeId,omitempty"`
	Source *OAuth2ScopeConsentGrantSource `json:"source,omitempty"`
	Status *OAuth2ScopeConsentGrantStatus `json:"status,omitempty"`
}
