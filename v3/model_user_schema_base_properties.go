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

type UserSchemaBaseProperties struct {
	Login *UserSchemaAttribute `json:"login,omitempty"`
	FirstName *UserSchemaAttribute `json:"firstName,omitempty"`
	LastName *UserSchemaAttribute `json:"lastName,omitempty"`
	MiddleName *UserSchemaAttribute `json:"middleName,omitempty"`
	HonorificPrefix *UserSchemaAttribute `json:"honorificPrefix,omitempty"`
	HonorificSuffix *UserSchemaAttribute `json:"honorificSuffix,omitempty"`
	Email *UserSchemaAttribute `json:"email,omitempty"`
	Title *UserSchemaAttribute `json:"title,omitempty"`
	DisplayName *UserSchemaAttribute `json:"displayName,omitempty"`
	NickName *UserSchemaAttribute `json:"nickName,omitempty"`
	ProfileUrl *UserSchemaAttribute `json:"profileUrl,omitempty"`
	SecondEmail *UserSchemaAttribute `json:"secondEmail,omitempty"`
	MobilePhone *UserSchemaAttribute `json:"mobilePhone,omitempty"`
	PrimaryPhone *UserSchemaAttribute `json:"primaryPhone,omitempty"`
	StreetAddress *UserSchemaAttribute `json:"streetAddress,omitempty"`
	City *UserSchemaAttribute `json:"city,omitempty"`
	State *UserSchemaAttribute `json:"state,omitempty"`
	ZipCode *UserSchemaAttribute `json:"zipCode,omitempty"`
	CountryCode *UserSchemaAttribute `json:"countryCode,omitempty"`
	PostalAddress *UserSchemaAttribute `json:"postalAddress,omitempty"`
	PreferredLanguage *UserSchemaAttribute `json:"preferredLanguage,omitempty"`
	Locale *UserSchemaAttribute `json:"locale,omitempty"`
	Timezone *UserSchemaAttribute `json:"timezone,omitempty"`
	UserType *UserSchemaAttribute `json:"userType,omitempty"`
	EmployeeNumber *UserSchemaAttribute `json:"employeeNumber,omitempty"`
	CostCenter *UserSchemaAttribute `json:"costCenter,omitempty"`
	Organization *UserSchemaAttribute `json:"organization,omitempty"`
	Division *UserSchemaAttribute `json:"division,omitempty"`
	Department *UserSchemaAttribute `json:"department,omitempty"`
	ManagerId *UserSchemaAttribute `json:"managerId,omitempty"`
	Manager *UserSchemaAttribute `json:"manager,omitempty"`
}
