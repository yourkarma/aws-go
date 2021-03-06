// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

// Package sdb provides a client for Amazon SimpleDB.
package sdb

import (
	"net/http"
	"time"

	"github.com/stripe/aws-go/aws"
	"github.com/stripe/aws-go/gen/endpoints"
)

import (
	"encoding/xml"
	"io"
)

// SDB is a client for Amazon SimpleDB.
type SDB struct {
	client *aws.QueryClient
}

// New returns a new SDB client.
func New(creds aws.CredentialsProvider, region string, client *http.Client) *SDB {
	if client == nil {
		client = http.DefaultClient
	}

	endpoint, service, region := endpoints.Lookup("sdb", region)

	return &SDB{
		client: &aws.QueryClient{
			Context: aws.Context{
				Credentials: creds,
				Service:     service,
				Region:      region,
			},
			Client:     client,
			Endpoint:   endpoint,
			APIVersion: "2009-04-15",
		},
	}
}

// BatchDeleteAttributes performs multiple DeleteAttributes operations in a
// single call, which reduces round trips and latencies. This enables
// Amazon SimpleDB to optimize requests, which generally yields better
// throughput. The following limitations are enforced for this operation: 1
// MB request size
func (c *SDB) BatchDeleteAttributes(req *BatchDeleteAttributesRequest) (err error) {
	// NRE
	err = c.client.Do("BatchDeleteAttributes", "POST", "/", req, nil)
	return
}

// BatchPutAttributes the BatchPutAttributes operation creates or replaces
// attributes within one or more items. By using this operation, the client
// can perform multiple PutAttribute operation with a single call. This
// helps yield savings in round trips and latencies, enabling Amazon
// SimpleDB to optimize requests and generally produce better throughput.
// The client may specify the item name with the Item.X.ItemName parameter.
// The client may specify new attributes using a combination of the
// Item.X.Attribute.Y.Name and Item.X.Attribute.Y.Value parameters. The
// client may specify the first attribute for the first item using the
// parameters Item.0.Attribute.0.Name and Item.0.Attribute.0.Value , and
// for the second attribute for the first item by the parameters
// Item.0.Attribute.1.Name and Item.0.Attribute.1.Value , and so on.
// Attributes are uniquely identified within an item by their name/value
// combination. For example, a single item can have the attributes {
// "first_name", "first_value" and { "first_name", "second_value" .
// However, it cannot have two attribute instances where both the
// Item.X.Attribute.Y.Name and Item.X.Attribute.Y.Value are the same.
// Optionally, the requester can supply the Replace parameter for each
// individual value. Setting this value to true will cause the new
// attribute values to replace the existing attribute values. For example,
// if an item has the attributes { 'a', '1' }, { 'b', '2'} and { 'b', '3'
// and the requester does a BatchPutAttributes of 'b', '4' with the Replace
// parameter set to true, the final attributes of the item will be { 'a',
// '1' and { 'b', '4' , replacing the previous values of the 'b' attribute
// with the new value. This operation is vulnerable to exceeding the
// maximum URL size when making a request using the GET method. This
// operation does not support conditions using Expected.X.Name ,
// Expected.X.Value , or Expected.X.Exists . You can execute multiple
// BatchPutAttributes operations and other operations in parallel. However,
// large numbers of concurrent BatchPutAttributes calls can result in
// Service Unavailable (503) responses. The following limitations are
// enforced for this operation: 256 attribute name-value pairs per item
func (c *SDB) BatchPutAttributes(req *BatchPutAttributesRequest) (err error) {
	// NRE
	err = c.client.Do("BatchPutAttributes", "POST", "/", req, nil)
	return
}

// CreateDomain the CreateDomain operation creates a new domain. The domain
// name should be unique among the domains associated with the Access Key
// ID provided in the request. The CreateDomain operation may take 10 or
// more seconds to complete. The client can create up to 100 domains per
// account. If the client requires additional domains, go to
// http://aws.amazon.com/contact-us/simpledb-limit-request/ .
func (c *SDB) CreateDomain(req *CreateDomainRequest) (err error) {
	// NRE
	err = c.client.Do("CreateDomain", "POST", "/", req, nil)
	return
}

// DeleteAttributes deletes one or more attributes associated with an item.
// If all attributes of the item are deleted, the item is deleted.
// DeleteAttributes is an idempotent operation; running it multiple times
// on the same item or attribute does not result in an error response.
// Because Amazon SimpleDB makes multiple copies of item data and uses an
// eventual consistency update model, performing a GetAttributes or Select
// operation (read) immediately after a DeleteAttributes or PutAttributes
// operation (write) might not return updated item data.
func (c *SDB) DeleteAttributes(req *DeleteAttributesRequest) (err error) {
	// NRE
	err = c.client.Do("DeleteAttributes", "POST", "/", req, nil)
	return
}

// DeleteDomain the DeleteDomain operation deletes a domain. Any items (and
// their attributes) in the domain are deleted as well. The DeleteDomain
// operation might take 10 or more seconds to complete.
func (c *SDB) DeleteDomain(req *DeleteDomainRequest) (err error) {
	// NRE
	err = c.client.Do("DeleteDomain", "POST", "/", req, nil)
	return
}

// DomainMetadata returns information about the domain, including when the
// domain was created, the number of items and attributes in the domain,
// and the size of the attribute names and values.
func (c *SDB) DomainMetadata(req *DomainMetadataRequest) (resp *DomainMetadataResult, err error) {
	resp = &DomainMetadataResult{}
	err = c.client.Do("DomainMetadata", "POST", "/", req, resp)
	return
}

// GetAttributes returns all of the attributes associated with the
// specified item. Optionally, the attributes returned can be limited to
// one or more attributes by specifying an attribute name parameter. If the
// item does not exist on the replica that was accessed for this operation,
// an empty set is returned. The system does not return an error as it
// cannot guarantee the item does not exist on other replicas.
func (c *SDB) GetAttributes(req *GetAttributesRequest) (resp *GetAttributesResult, err error) {
	resp = &GetAttributesResult{}
	err = c.client.Do("GetAttributes", "POST", "/", req, resp)
	return
}

// ListDomains the ListDomains operation lists all domains associated with
// the Access Key ID. It returns domain names up to the limit set by
// MaxNumberOfDomains . A NextToken is returned if there are more than
// MaxNumberOfDomains domains. Calling ListDomains successive times with
// the NextToken provided by the operation returns up to MaxNumberOfDomains
// more domain names with each successive operation call.
func (c *SDB) ListDomains(req *ListDomainsRequest) (resp *ListDomainsResult, err error) {
	resp = &ListDomainsResult{}
	err = c.client.Do("ListDomains", "POST", "/", req, resp)
	return
}

// PutAttributes the PutAttributes operation creates or replaces attributes
// in an item. The client may specify new attributes using a combination of
// the Attribute.X.Name and Attribute.X.Value parameters. The client
// specifies the first attribute by the parameters Attribute.0.Name and
// Attribute.0.Value , the second attribute by the parameters
// Attribute.1.Name and Attribute.1.Value , and so on. Attributes are
// uniquely identified in an item by their name/value combination. For
// example, a single item can have the attributes { "first_name",
// "first_value" and { "first_name", second_value" . However, it cannot
// have two attribute instances where both the Attribute.X.Name and
// Attribute.X.Value are the same. Optionally, the requestor can supply the
// Replace parameter for each individual attribute. Setting this value to
// true causes the new attribute value to replace the existing attribute
// value(s). For example, if an item has the attributes { 'a', '1' , { 'b',
// '2'} and { 'b', '3' and the requestor calls PutAttributes using the
// attributes { 'b', '4' with the Replace parameter set to true, the final
// attributes of the item are changed to { 'a', '1' and { 'b', '4' , which
// replaces the previous values of the 'b' attribute with the new value.
// You cannot specify an empty string as an attribute name. Because Amazon
// SimpleDB makes multiple copies of client data and uses an eventual
// consistency update model, an immediate GetAttributes or Select operation
// (read) immediately after a PutAttributes or DeleteAttributes operation
// (write) might not return the updated data. The following limitations are
// enforced for this operation: 256 total attribute name-value pairs per
// item
func (c *SDB) PutAttributes(req *PutAttributesRequest) (err error) {
	// NRE
	err = c.client.Do("PutAttributes", "POST", "/", req, nil)
	return
}

// Select the Select operation returns a set of attributes for ItemNames
// that match the select expression. Select is similar to the standard SQL
// statement. The total size of the response cannot exceed 1 MB in total
// size. Amazon SimpleDB automatically adjusts the number of items returned
// per page to enforce this limit. For example, if the client asks to
// retrieve 2500 items, but each individual item is 10 kB in size, the
// system returns 100 items and an appropriate NextToken so the client can
// access the next page of results. For information on how to construct
// select expressions, see Using Select to Create Amazon SimpleDB Queries
// in the Developer Guide.
func (c *SDB) Select(req *SelectRequest) (resp *SelectResult, err error) {
	resp = &SelectResult{}
	err = c.client.Do("Select", "POST", "/", req, resp)
	return
}

// Attribute is undocumented.
type Attribute struct {
	AlternateNameEncoding  aws.StringValue `query:"AlternateNameEncoding" xml:"AlternateNameEncoding"`
	AlternateValueEncoding aws.StringValue `query:"AlternateValueEncoding" xml:"AlternateValueEncoding"`
	Name                   aws.StringValue `query:"Name" xml:"Name"`
	Value                  aws.StringValue `query:"Value" xml:"Value"`
}

// BatchDeleteAttributesRequest is undocumented.
type BatchDeleteAttributesRequest struct {
	DomainName aws.StringValue `query:"DomainName" xml:"DomainName"`
	Items      []DeletableItem `query:"Item" xml:"Item"`
}

// BatchPutAttributesRequest is undocumented.
type BatchPutAttributesRequest struct {
	DomainName aws.StringValue   `query:"DomainName" xml:"DomainName"`
	Items      []ReplaceableItem `query:"Item" xml:"Item"`
}

// CreateDomainRequest is undocumented.
type CreateDomainRequest struct {
	DomainName aws.StringValue `query:"DomainName" xml:"DomainName"`
}

// DeletableItem is undocumented.
type DeletableItem struct {
	Attributes []Attribute     `query:"Attribute" xml:"Attribute"`
	Name       aws.StringValue `query:"ItemName" xml:"ItemName"`
}

// DeleteAttributesRequest is undocumented.
type DeleteAttributesRequest struct {
	Attributes []Attribute      `query:"Attribute" xml:"Attribute"`
	DomainName aws.StringValue  `query:"DomainName" xml:"DomainName"`
	Expected   *UpdateCondition `query:"Expected" xml:"Expected"`
	ItemName   aws.StringValue  `query:"ItemName" xml:"ItemName"`
}

// DeleteDomainRequest is undocumented.
type DeleteDomainRequest struct {
	DomainName aws.StringValue `query:"DomainName" xml:"DomainName"`
}

// DomainMetadataRequest is undocumented.
type DomainMetadataRequest struct {
	DomainName aws.StringValue `query:"DomainName" xml:"DomainName"`
}

// DomainMetadataResult is undocumented.
type DomainMetadataResult struct {
	AttributeNameCount       aws.IntegerValue `query:"AttributeNameCount" xml:"DomainMetadataResult>AttributeNameCount"`
	AttributeNamesSizeBytes  aws.LongValue    `query:"AttributeNamesSizeBytes" xml:"DomainMetadataResult>AttributeNamesSizeBytes"`
	AttributeValueCount      aws.IntegerValue `query:"AttributeValueCount" xml:"DomainMetadataResult>AttributeValueCount"`
	AttributeValuesSizeBytes aws.LongValue    `query:"AttributeValuesSizeBytes" xml:"DomainMetadataResult>AttributeValuesSizeBytes"`
	ItemCount                aws.IntegerValue `query:"ItemCount" xml:"DomainMetadataResult>ItemCount"`
	ItemNamesSizeBytes       aws.LongValue    `query:"ItemNamesSizeBytes" xml:"DomainMetadataResult>ItemNamesSizeBytes"`
	Timestamp                aws.IntegerValue `query:"Timestamp" xml:"DomainMetadataResult>Timestamp"`
}

// GetAttributesRequest is undocumented.
type GetAttributesRequest struct {
	AttributeNames []string         `query:"AttributeName" xml:"AttributeName"`
	ConsistentRead aws.BooleanValue `query:"ConsistentRead" xml:"ConsistentRead"`
	DomainName     aws.StringValue  `query:"DomainName" xml:"DomainName"`
	ItemName       aws.StringValue  `query:"ItemName" xml:"ItemName"`
}

// GetAttributesResult is undocumented.
type GetAttributesResult struct {
	Attributes []Attribute `query:"Attribute" xml:"GetAttributesResult>Attribute"`
}

// Item is undocumented.
type Item struct {
	AlternateNameEncoding aws.StringValue `query:"AlternateNameEncoding" xml:"AlternateNameEncoding"`
	Attributes            []Attribute     `query:"Attribute" xml:"Attribute"`
	Name                  aws.StringValue `query:"Name" xml:"Name"`
}

// ListDomainsRequest is undocumented.
type ListDomainsRequest struct {
	MaxNumberOfDomains aws.IntegerValue `query:"MaxNumberOfDomains" xml:"MaxNumberOfDomains"`
	NextToken          aws.StringValue  `query:"NextToken" xml:"NextToken"`
}

// ListDomainsResult is undocumented.
type ListDomainsResult struct {
	DomainNames []string        `query:"DomainName" xml:"ListDomainsResult>DomainName"`
	NextToken   aws.StringValue `query:"NextToken" xml:"ListDomainsResult>NextToken"`
}

// PutAttributesRequest is undocumented.
type PutAttributesRequest struct {
	Attributes []ReplaceableAttribute `query:"Attribute" xml:"Attribute"`
	DomainName aws.StringValue        `query:"DomainName" xml:"DomainName"`
	Expected   *UpdateCondition       `query:"Expected" xml:"Expected"`
	ItemName   aws.StringValue        `query:"ItemName" xml:"ItemName"`
}

// ReplaceableAttribute is undocumented.
type ReplaceableAttribute struct {
	Name    aws.StringValue  `query:"Name" xml:"Name"`
	Replace aws.BooleanValue `query:"Replace" xml:"Replace"`
	Value   aws.StringValue  `query:"Value" xml:"Value"`
}

// ReplaceableItem is undocumented.
type ReplaceableItem struct {
	Attributes []ReplaceableAttribute `query:"Attribute" xml:"Attribute"`
	Name       aws.StringValue        `query:"ItemName" xml:"ItemName"`
}

// SelectRequest is undocumented.
type SelectRequest struct {
	ConsistentRead   aws.BooleanValue `query:"ConsistentRead" xml:"ConsistentRead"`
	NextToken        aws.StringValue  `query:"NextToken" xml:"NextToken"`
	SelectExpression aws.StringValue  `query:"SelectExpression" xml:"SelectExpression"`
}

// SelectResult is undocumented.
type SelectResult struct {
	Items     []Item          `query:"Item" xml:"SelectResult>Item"`
	NextToken aws.StringValue `query:"NextToken" xml:"SelectResult>NextToken"`
}

// UpdateCondition is undocumented.
type UpdateCondition struct {
	Exists aws.BooleanValue `query:"Exists" xml:"Exists"`
	Name   aws.StringValue  `query:"Name" xml:"Name"`
	Value  aws.StringValue  `query:"Value" xml:"Value"`
}

// avoid errors if the packages aren't referenced
var _ time.Time

var _ xml.Decoder
var _ = io.EOF
