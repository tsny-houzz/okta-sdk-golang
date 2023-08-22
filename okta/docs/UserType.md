# UserType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to **time.Time** | A timestamp from when the User Type was created | [optional] [readonly] 
**CreatedBy** | Pointer to **string** | The user ID of the account that created the User Type | [optional] [readonly] 
**Default** | Pointer to **bool** | A boolean value to indicate if this is the default User Type | [optional] [readonly] 
**Description** | Pointer to **string** | The human-readable description of the User Type | [optional] 
**DisplayName** | **string** | The human-readable name of the User Type | 
**Id** | Pointer to **string** | The unique key for the User Type | [optional] [readonly] 
**LastUpdated** | Pointer to **time.Time** | A timestamp from when the User Type was most recently updated | [optional] [readonly] 
**LastUpdatedBy** | Pointer to **string** | The user ID of the most recent account to edit the User Type | [optional] [readonly] 
**Name** | **string** | The name of the User Type. The name must start with A-Z or a-z and contain only A-Z, a-z, 0-9, or underscore (_) characters.   This value becomes read-only after creation and can&#39;t be updated. | 
**Links** | Pointer to [**UserTypeLinks**](UserTypeLinks.md) |  | [optional] 

## Methods

### NewUserType

`func NewUserType(displayName string, name string, ) *UserType`

NewUserType instantiates a new UserType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserTypeWithDefaults

`func NewUserTypeWithDefaults() *UserType`

NewUserTypeWithDefaults instantiates a new UserType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *UserType) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *UserType) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *UserType) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *UserType) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetCreatedBy

`func (o *UserType) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *UserType) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *UserType) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *UserType) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetDefault

`func (o *UserType) GetDefault() bool`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *UserType) GetDefaultOk() (*bool, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *UserType) SetDefault(v bool)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *UserType) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### GetDescription

`func (o *UserType) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UserType) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UserType) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UserType) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayName

`func (o *UserType) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *UserType) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *UserType) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetId

`func (o *UserType) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserType) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserType) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *UserType) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastUpdated

`func (o *UserType) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *UserType) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *UserType) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *UserType) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetLastUpdatedBy

`func (o *UserType) GetLastUpdatedBy() string`

GetLastUpdatedBy returns the LastUpdatedBy field if non-nil, zero value otherwise.

### GetLastUpdatedByOk

`func (o *UserType) GetLastUpdatedByOk() (*string, bool)`

GetLastUpdatedByOk returns a tuple with the LastUpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedBy

`func (o *UserType) SetLastUpdatedBy(v string)`

SetLastUpdatedBy sets LastUpdatedBy field to given value.

### HasLastUpdatedBy

`func (o *UserType) HasLastUpdatedBy() bool`

HasLastUpdatedBy returns a boolean if a field has been set.

### GetName

`func (o *UserType) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UserType) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UserType) SetName(v string)`

SetName sets Name field to given value.


### GetLinks

`func (o *UserType) GetLinks() UserTypeLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *UserType) GetLinksOk() (*UserTypeLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *UserType) SetLinks(v UserTypeLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *UserType) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

