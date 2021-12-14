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

type FactorResultType string

// List of FactorResultType
const (
	SUCCESS_FactorResultType FactorResultType = "SUCCESS"
	CHALLENGE_FactorResultType FactorResultType = "CHALLENGE"
	WAITING_FactorResultType FactorResultType = "WAITING"
	FAILED_FactorResultType FactorResultType = "FAILED"
	REJECTED_FactorResultType FactorResultType = "REJECTED"
	TIMEOUT_FactorResultType FactorResultType = "TIMEOUT"
	TIME_WINDOW_EXCEEDED_FactorResultType FactorResultType = "TIME_WINDOW_EXCEEDED"
	PASSCODE_REPLAYED_FactorResultType FactorResultType = "PASSCODE_REPLAYED"
	ERROR__FactorResultType FactorResultType = "ERROR"
	CANCELLED_FactorResultType FactorResultType = "CANCELLED"
)
