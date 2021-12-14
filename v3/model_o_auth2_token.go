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

type OAuth2Token struct {
	Embedded map[string]interface{} `json:"_embedded,omitempty"`
	Links map[string]interface{} `json:"_links,omitempty"`
	ClientId string `json:"clientId,omitempty"`
	Created time.Time `json:"created,omitempty"`
	ExpiresAt time.Time `json:"expiresAt,omitempty"`
	Id string `json:"id,omitempty"`
	Issuer string `json:"issuer,omitempty"`
	LastUpdated time.Time `json:"lastUpdated,omitempty"`
	Scopes []string `json:"scopes,omitempty"`
	Status string `json:"status,omitempty"`
	UserId string `json:"userId,omitempty"`
}
