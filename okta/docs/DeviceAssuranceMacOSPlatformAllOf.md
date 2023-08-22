# DeviceAssuranceMacOSPlatformAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DiskEncryptionType** | Pointer to [**DeviceAssuranceMacOSPlatformAllOfDiskEncryptionType**](DeviceAssuranceMacOSPlatformAllOfDiskEncryptionType.md) |  | [optional] 
**OsVersion** | Pointer to [**OSVersionThreeComponents**](OSVersionThreeComponents.md) |  | [optional] 
**ScreenLockType** | Pointer to [**DeviceAssuranceAndroidPlatformAllOfScreenLockType**](DeviceAssuranceAndroidPlatformAllOfScreenLockType.md) |  | [optional] 
**SecureHardwarePresent** | Pointer to **bool** |  | [optional] 
**ThirdPartySignalProviders** | Pointer to [**DeviceAssuranceMacOSPlatformAllOfThirdPartySignalProviders**](DeviceAssuranceMacOSPlatformAllOfThirdPartySignalProviders.md) |  | [optional] 

## Methods

### NewDeviceAssuranceMacOSPlatformAllOf

`func NewDeviceAssuranceMacOSPlatformAllOf() *DeviceAssuranceMacOSPlatformAllOf`

NewDeviceAssuranceMacOSPlatformAllOf instantiates a new DeviceAssuranceMacOSPlatformAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceAssuranceMacOSPlatformAllOfWithDefaults

`func NewDeviceAssuranceMacOSPlatformAllOfWithDefaults() *DeviceAssuranceMacOSPlatformAllOf`

NewDeviceAssuranceMacOSPlatformAllOfWithDefaults instantiates a new DeviceAssuranceMacOSPlatformAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDiskEncryptionType

`func (o *DeviceAssuranceMacOSPlatformAllOf) GetDiskEncryptionType() DeviceAssuranceMacOSPlatformAllOfDiskEncryptionType`

GetDiskEncryptionType returns the DiskEncryptionType field if non-nil, zero value otherwise.

### GetDiskEncryptionTypeOk

`func (o *DeviceAssuranceMacOSPlatformAllOf) GetDiskEncryptionTypeOk() (*DeviceAssuranceMacOSPlatformAllOfDiskEncryptionType, bool)`

GetDiskEncryptionTypeOk returns a tuple with the DiskEncryptionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskEncryptionType

`func (o *DeviceAssuranceMacOSPlatformAllOf) SetDiskEncryptionType(v DeviceAssuranceMacOSPlatformAllOfDiskEncryptionType)`

SetDiskEncryptionType sets DiskEncryptionType field to given value.

### HasDiskEncryptionType

`func (o *DeviceAssuranceMacOSPlatformAllOf) HasDiskEncryptionType() bool`

HasDiskEncryptionType returns a boolean if a field has been set.

### GetOsVersion

`func (o *DeviceAssuranceMacOSPlatformAllOf) GetOsVersion() OSVersionThreeComponents`

GetOsVersion returns the OsVersion field if non-nil, zero value otherwise.

### GetOsVersionOk

`func (o *DeviceAssuranceMacOSPlatformAllOf) GetOsVersionOk() (*OSVersionThreeComponents, bool)`

GetOsVersionOk returns a tuple with the OsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsVersion

`func (o *DeviceAssuranceMacOSPlatformAllOf) SetOsVersion(v OSVersionThreeComponents)`

SetOsVersion sets OsVersion field to given value.

### HasOsVersion

`func (o *DeviceAssuranceMacOSPlatformAllOf) HasOsVersion() bool`

HasOsVersion returns a boolean if a field has been set.

### GetScreenLockType

`func (o *DeviceAssuranceMacOSPlatformAllOf) GetScreenLockType() DeviceAssuranceAndroidPlatformAllOfScreenLockType`

GetScreenLockType returns the ScreenLockType field if non-nil, zero value otherwise.

### GetScreenLockTypeOk

`func (o *DeviceAssuranceMacOSPlatformAllOf) GetScreenLockTypeOk() (*DeviceAssuranceAndroidPlatformAllOfScreenLockType, bool)`

GetScreenLockTypeOk returns a tuple with the ScreenLockType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScreenLockType

`func (o *DeviceAssuranceMacOSPlatformAllOf) SetScreenLockType(v DeviceAssuranceAndroidPlatformAllOfScreenLockType)`

SetScreenLockType sets ScreenLockType field to given value.

### HasScreenLockType

`func (o *DeviceAssuranceMacOSPlatformAllOf) HasScreenLockType() bool`

HasScreenLockType returns a boolean if a field has been set.

### GetSecureHardwarePresent

`func (o *DeviceAssuranceMacOSPlatformAllOf) GetSecureHardwarePresent() bool`

GetSecureHardwarePresent returns the SecureHardwarePresent field if non-nil, zero value otherwise.

### GetSecureHardwarePresentOk

`func (o *DeviceAssuranceMacOSPlatformAllOf) GetSecureHardwarePresentOk() (*bool, bool)`

GetSecureHardwarePresentOk returns a tuple with the SecureHardwarePresent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureHardwarePresent

`func (o *DeviceAssuranceMacOSPlatformAllOf) SetSecureHardwarePresent(v bool)`

SetSecureHardwarePresent sets SecureHardwarePresent field to given value.

### HasSecureHardwarePresent

`func (o *DeviceAssuranceMacOSPlatformAllOf) HasSecureHardwarePresent() bool`

HasSecureHardwarePresent returns a boolean if a field has been set.

### GetThirdPartySignalProviders

`func (o *DeviceAssuranceMacOSPlatformAllOf) GetThirdPartySignalProviders() DeviceAssuranceMacOSPlatformAllOfThirdPartySignalProviders`

GetThirdPartySignalProviders returns the ThirdPartySignalProviders field if non-nil, zero value otherwise.

### GetThirdPartySignalProvidersOk

`func (o *DeviceAssuranceMacOSPlatformAllOf) GetThirdPartySignalProvidersOk() (*DeviceAssuranceMacOSPlatformAllOfThirdPartySignalProviders, bool)`

GetThirdPartySignalProvidersOk returns a tuple with the ThirdPartySignalProviders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThirdPartySignalProviders

`func (o *DeviceAssuranceMacOSPlatformAllOf) SetThirdPartySignalProviders(v DeviceAssuranceMacOSPlatformAllOfThirdPartySignalProviders)`

SetThirdPartySignalProviders sets ThirdPartySignalProviders field to given value.

### HasThirdPartySignalProviders

`func (o *DeviceAssuranceMacOSPlatformAllOf) HasThirdPartySignalProviders() bool`

HasThirdPartySignalProviders returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

