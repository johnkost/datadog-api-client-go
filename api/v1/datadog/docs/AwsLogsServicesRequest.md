# AwsLogsServicesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | Pointer to **string** | Your AWS Account ID without dashes. | 
**Services** | Pointer to **[]string** | Array of services IDs set to enable automatic log collection. Discover the list of available services with the Get list of AWS log ready services API endpoint | 

## Methods

### GetAccountId

`func (o *AwsLogsServicesRequest) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *AwsLogsServicesRequest) GetAccountIdOk() (string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAccountId

`func (o *AwsLogsServicesRequest) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### SetAccountId

`func (o *AwsLogsServicesRequest) SetAccountId(v string)`

SetAccountId gets a reference to the given string and assigns it to the AccountId field.

### GetServices

`func (o *AwsLogsServicesRequest) GetServices() []string`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *AwsLogsServicesRequest) GetServicesOk() ([]string, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasServices

`func (o *AwsLogsServicesRequest) HasServices() bool`

HasServices returns a boolean if a field has been set.

### SetServices

`func (o *AwsLogsServicesRequest) SetServices(v []string)`

SetServices gets a reference to the given []string and assigns it to the Services field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

