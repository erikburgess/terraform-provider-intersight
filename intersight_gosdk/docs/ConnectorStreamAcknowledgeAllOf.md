# ConnectorStreamAcknowledgeAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AckSequence** | Pointer to **int64** | The latest message sequence processed in the cloud. Device connector will drop all messages up to this sequence from its cache. | [optional] 

## Methods

### NewConnectorStreamAcknowledgeAllOf

`func NewConnectorStreamAcknowledgeAllOf() *ConnectorStreamAcknowledgeAllOf`

NewConnectorStreamAcknowledgeAllOf instantiates a new ConnectorStreamAcknowledgeAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectorStreamAcknowledgeAllOfWithDefaults

`func NewConnectorStreamAcknowledgeAllOfWithDefaults() *ConnectorStreamAcknowledgeAllOf`

NewConnectorStreamAcknowledgeAllOfWithDefaults instantiates a new ConnectorStreamAcknowledgeAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAckSequence

`func (o *ConnectorStreamAcknowledgeAllOf) GetAckSequence() int64`

GetAckSequence returns the AckSequence field if non-nil, zero value otherwise.

### GetAckSequenceOk

`func (o *ConnectorStreamAcknowledgeAllOf) GetAckSequenceOk() (*int64, bool)`

GetAckSequenceOk returns a tuple with the AckSequence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAckSequence

`func (o *ConnectorStreamAcknowledgeAllOf) SetAckSequence(v int64)`

SetAckSequence sets AckSequence field to given value.

### HasAckSequence

`func (o *ConnectorStreamAcknowledgeAllOf) HasAckSequence() bool`

HasAckSequence returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


