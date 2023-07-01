# PostLoginParametersRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | Pointer to **string** | username for your account | [optional] 
**AccountType** | Pointer to [**AccountType**](AccountType.md) |  | [optional] [default to _2]
**DeviceId** | Pointer to **string** |  | [optional] 
**DeviceName** | Pointer to **string** | device name | [optional] [default to "test"]
**ExtInfo** | Pointer to [**PostLoginParametersRequestExtInfo**](PostLoginParametersRequestExtInfo.md) |  | [optional] 
**Grade** | Pointer to **int32** |  | [optional] [default to 1]
**Pwd** | Pointer to **string** | pwd &#x3D; md5(passwordHash + password) | [optional] 
**RegionId** | Pointer to **int32** |  | [optional] [default to 1]

## Methods

### NewPostLoginParametersRequest

`func NewPostLoginParametersRequest() *PostLoginParametersRequest`

NewPostLoginParametersRequest instantiates a new PostLoginParametersRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostLoginParametersRequestWithDefaults

`func NewPostLoginParametersRequestWithDefaults() *PostLoginParametersRequest`

NewPostLoginParametersRequestWithDefaults instantiates a new PostLoginParametersRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *PostLoginParametersRequest) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *PostLoginParametersRequest) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *PostLoginParametersRequest) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *PostLoginParametersRequest) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetAccountType

`func (o *PostLoginParametersRequest) GetAccountType() AccountType`

GetAccountType returns the AccountType field if non-nil, zero value otherwise.

### GetAccountTypeOk

`func (o *PostLoginParametersRequest) GetAccountTypeOk() (*AccountType, bool)`

GetAccountTypeOk returns a tuple with the AccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountType

`func (o *PostLoginParametersRequest) SetAccountType(v AccountType)`

SetAccountType sets AccountType field to given value.

### HasAccountType

`func (o *PostLoginParametersRequest) HasAccountType() bool`

HasAccountType returns a boolean if a field has been set.

### GetDeviceId

`func (o *PostLoginParametersRequest) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *PostLoginParametersRequest) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *PostLoginParametersRequest) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *PostLoginParametersRequest) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### GetDeviceName

`func (o *PostLoginParametersRequest) GetDeviceName() string`

GetDeviceName returns the DeviceName field if non-nil, zero value otherwise.

### GetDeviceNameOk

`func (o *PostLoginParametersRequest) GetDeviceNameOk() (*string, bool)`

GetDeviceNameOk returns a tuple with the DeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceName

`func (o *PostLoginParametersRequest) SetDeviceName(v string)`

SetDeviceName sets DeviceName field to given value.

### HasDeviceName

`func (o *PostLoginParametersRequest) HasDeviceName() bool`

HasDeviceName returns a boolean if a field has been set.

### GetExtInfo

`func (o *PostLoginParametersRequest) GetExtInfo() PostLoginParametersRequestExtInfo`

GetExtInfo returns the ExtInfo field if non-nil, zero value otherwise.

### GetExtInfoOk

`func (o *PostLoginParametersRequest) GetExtInfoOk() (*PostLoginParametersRequestExtInfo, bool)`

GetExtInfoOk returns a tuple with the ExtInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtInfo

`func (o *PostLoginParametersRequest) SetExtInfo(v PostLoginParametersRequestExtInfo)`

SetExtInfo sets ExtInfo field to given value.

### HasExtInfo

`func (o *PostLoginParametersRequest) HasExtInfo() bool`

HasExtInfo returns a boolean if a field has been set.

### GetGrade

`func (o *PostLoginParametersRequest) GetGrade() int32`

GetGrade returns the Grade field if non-nil, zero value otherwise.

### GetGradeOk

`func (o *PostLoginParametersRequest) GetGradeOk() (*int32, bool)`

GetGradeOk returns a tuple with the Grade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrade

`func (o *PostLoginParametersRequest) SetGrade(v int32)`

SetGrade sets Grade field to given value.

### HasGrade

`func (o *PostLoginParametersRequest) HasGrade() bool`

HasGrade returns a boolean if a field has been set.

### GetPwd

`func (o *PostLoginParametersRequest) GetPwd() string`

GetPwd returns the Pwd field if non-nil, zero value otherwise.

### GetPwdOk

`func (o *PostLoginParametersRequest) GetPwdOk() (*string, bool)`

GetPwdOk returns a tuple with the Pwd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPwd

`func (o *PostLoginParametersRequest) SetPwd(v string)`

SetPwd sets Pwd field to given value.

### HasPwd

`func (o *PostLoginParametersRequest) HasPwd() bool`

HasPwd returns a boolean if a field has been set.

### GetRegionId

`func (o *PostLoginParametersRequest) GetRegionId() int32`

GetRegionId returns the RegionId field if non-nil, zero value otherwise.

### GetRegionIdOk

`func (o *PostLoginParametersRequest) GetRegionIdOk() (*int32, bool)`

GetRegionIdOk returns a tuple with the RegionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionId

`func (o *PostLoginParametersRequest) SetRegionId(v int32)`

SetRegionId sets RegionId field to given value.

### HasRegionId

`func (o *PostLoginParametersRequest) HasRegionId() bool`

HasRegionId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


