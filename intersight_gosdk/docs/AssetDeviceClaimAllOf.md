# AssetDeviceClaimAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceUpdates** | Pointer to [**[]AssetConnectionControlMessage**](asset.ConnectionControlMessage.md) |  | [optional] 
**SecurityToken** | Pointer to **string** | Obtained from the device connector management UI or API (REST endpoint &#39;/connector/SecurityTokens&#39;). | [optional] 
**SerialNumber** | Pointer to **string** | Obtained from the device connector management UI or API (REST endpoint &#39;/connector/DeviceIdentifiers&#39;). | [optional] 
**Account** | Pointer to [**IamAccountRelationship**](iam.Account.Relationship.md) |  | [optional] 
**Device** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 

## Methods

### NewAssetDeviceClaimAllOf

`func NewAssetDeviceClaimAllOf() *AssetDeviceClaimAllOf`

NewAssetDeviceClaimAllOf instantiates a new AssetDeviceClaimAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetDeviceClaimAllOfWithDefaults

`func NewAssetDeviceClaimAllOfWithDefaults() *AssetDeviceClaimAllOf`

NewAssetDeviceClaimAllOfWithDefaults instantiates a new AssetDeviceClaimAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceUpdates

`func (o *AssetDeviceClaimAllOf) GetDeviceUpdates() []AssetConnectionControlMessage`

GetDeviceUpdates returns the DeviceUpdates field if non-nil, zero value otherwise.

### GetDeviceUpdatesOk

`func (o *AssetDeviceClaimAllOf) GetDeviceUpdatesOk() (*[]AssetConnectionControlMessage, bool)`

GetDeviceUpdatesOk returns a tuple with the DeviceUpdates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceUpdates

`func (o *AssetDeviceClaimAllOf) SetDeviceUpdates(v []AssetConnectionControlMessage)`

SetDeviceUpdates sets DeviceUpdates field to given value.

### HasDeviceUpdates

`func (o *AssetDeviceClaimAllOf) HasDeviceUpdates() bool`

HasDeviceUpdates returns a boolean if a field has been set.

### GetSecurityToken

`func (o *AssetDeviceClaimAllOf) GetSecurityToken() string`

GetSecurityToken returns the SecurityToken field if non-nil, zero value otherwise.

### GetSecurityTokenOk

`func (o *AssetDeviceClaimAllOf) GetSecurityTokenOk() (*string, bool)`

GetSecurityTokenOk returns a tuple with the SecurityToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityToken

`func (o *AssetDeviceClaimAllOf) SetSecurityToken(v string)`

SetSecurityToken sets SecurityToken field to given value.

### HasSecurityToken

`func (o *AssetDeviceClaimAllOf) HasSecurityToken() bool`

HasSecurityToken returns a boolean if a field has been set.

### GetSerialNumber

`func (o *AssetDeviceClaimAllOf) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *AssetDeviceClaimAllOf) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *AssetDeviceClaimAllOf) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *AssetDeviceClaimAllOf) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### GetAccount

`func (o *AssetDeviceClaimAllOf) GetAccount() IamAccountRelationship`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *AssetDeviceClaimAllOf) GetAccountOk() (*IamAccountRelationship, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *AssetDeviceClaimAllOf) SetAccount(v IamAccountRelationship)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *AssetDeviceClaimAllOf) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetDevice

`func (o *AssetDeviceClaimAllOf) GetDevice() AssetDeviceRegistrationRelationship`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *AssetDeviceClaimAllOf) GetDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *AssetDeviceClaimAllOf) SetDevice(v AssetDeviceRegistrationRelationship)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *AssetDeviceClaimAllOf) HasDevice() bool`

HasDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


