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

type OAuthResponseType string

// List of OAuthResponseType
const (
	CODE_OAuthResponseType OAuthResponseType = "code"
	TOKEN_OAuthResponseType OAuthResponseType = "token"
	ID_TOKEN_OAuthResponseType OAuthResponseType = "id_token"
)
