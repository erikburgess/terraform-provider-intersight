# IamCertificateAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Certificate** | Pointer to [**X509Certificate**](x509.Certificate.md) |  | [optional] 
**Status** | Pointer to **string** | Status of the certificate. | [optional] [readonly] [default to "PendingValidation"]
**CertificateRequest** | Pointer to [**IamCertificateRequestRelationship**](iam.CertificateRequest.Relationship.md) |  | [optional] 

## Methods

### NewIamCertificateAllOf

`func NewIamCertificateAllOf() *IamCertificateAllOf`

NewIamCertificateAllOf instantiates a new IamCertificateAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamCertificateAllOfWithDefaults

`func NewIamCertificateAllOfWithDefaults() *IamCertificateAllOf`

NewIamCertificateAllOfWithDefaults instantiates a new IamCertificateAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertificate

`func (o *IamCertificateAllOf) GetCertificate() X509Certificate`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *IamCertificateAllOf) GetCertificateOk() (*X509Certificate, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *IamCertificateAllOf) SetCertificate(v X509Certificate)`

SetCertificate sets Certificate field to given value.

### HasCertificate

`func (o *IamCertificateAllOf) HasCertificate() bool`

HasCertificate returns a boolean if a field has been set.

### GetStatus

`func (o *IamCertificateAllOf) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *IamCertificateAllOf) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *IamCertificateAllOf) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *IamCertificateAllOf) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCertificateRequest

`func (o *IamCertificateAllOf) GetCertificateRequest() IamCertificateRequestRelationship`

GetCertificateRequest returns the CertificateRequest field if non-nil, zero value otherwise.

### GetCertificateRequestOk

`func (o *IamCertificateAllOf) GetCertificateRequestOk() (*IamCertificateRequestRelationship, bool)`

GetCertificateRequestOk returns a tuple with the CertificateRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateRequest

`func (o *IamCertificateAllOf) SetCertificateRequest(v IamCertificateRequestRelationship)`

SetCertificateRequest sets CertificateRequest field to given value.

### HasCertificateRequest

`func (o *IamCertificateAllOf) HasCertificateRequest() bool`

HasCertificateRequest returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


