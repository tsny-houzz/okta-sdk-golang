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

type FactorStatus string

// List of FactorStatus
const (
	PENDING_ACTIVATION_FactorStatus FactorStatus = "PENDING_ACTIVATION"
	ACTIVE_FactorStatus FactorStatus = "ACTIVE"
	INACTIVE_FactorStatus FactorStatus = "INACTIVE"
	NOT_SETUP_FactorStatus FactorStatus = "NOT_SETUP"
	ENROLLED_FactorStatus FactorStatus = "ENROLLED"
	DISABLED_FactorStatus FactorStatus = "DISABLED"
	EXPIRED_FactorStatus FactorStatus = "EXPIRED"
)
