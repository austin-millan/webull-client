# PostLoginParametersRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | **string** | username for your account | [optional] 
**AccountType** | [**AccountType**](AccountType.md) |  | [optional] 
**DeviceId** | **string** |  | [optional] 
**DeviceName** | **string** | device name | [optional] [default to test]
**ExtInfo** | [**PostLoginParametersRequestExtInfo**](PostLoginParametersRequest_extInfo.md) |  | [optional] 
**Grade** | **int32** |  | [optional] [default to 1]
**Pwd** | **string** | pwd &#x3D; md5(passwordHash + password) | [optional] 
**RegionId** | **int32** |  | [optional] [default to 1]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


