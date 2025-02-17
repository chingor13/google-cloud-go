// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go_gapic. DO NOT EDIT.

package compute

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"net/url"
	"sort"

	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/googleapi"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	httptransport "google.golang.org/api/transport/http"
	computepb "google.golang.org/genproto/googleapis/cloud/compute/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

var newNetworkEndpointGroupsClientHook clientHook

// NetworkEndpointGroupsCallOptions contains the retry settings for each method of NetworkEndpointGroupsClient.
type NetworkEndpointGroupsCallOptions struct {
	AggregatedList         []gax.CallOption
	AttachNetworkEndpoints []gax.CallOption
	Delete                 []gax.CallOption
	DetachNetworkEndpoints []gax.CallOption
	Get                    []gax.CallOption
	Insert                 []gax.CallOption
	List                   []gax.CallOption
	ListNetworkEndpoints   []gax.CallOption
	TestIamPermissions     []gax.CallOption
}

// internalNetworkEndpointGroupsClient is an interface that defines the methods availaible from Google Compute Engine API.
type internalNetworkEndpointGroupsClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	AggregatedList(context.Context, *computepb.AggregatedListNetworkEndpointGroupsRequest, ...gax.CallOption) *NetworkEndpointGroupsScopedListPairIterator
	AttachNetworkEndpoints(context.Context, *computepb.AttachNetworkEndpointsNetworkEndpointGroupRequest, ...gax.CallOption) (*Operation, error)
	Delete(context.Context, *computepb.DeleteNetworkEndpointGroupRequest, ...gax.CallOption) (*Operation, error)
	DetachNetworkEndpoints(context.Context, *computepb.DetachNetworkEndpointsNetworkEndpointGroupRequest, ...gax.CallOption) (*Operation, error)
	Get(context.Context, *computepb.GetNetworkEndpointGroupRequest, ...gax.CallOption) (*computepb.NetworkEndpointGroup, error)
	Insert(context.Context, *computepb.InsertNetworkEndpointGroupRequest, ...gax.CallOption) (*Operation, error)
	List(context.Context, *computepb.ListNetworkEndpointGroupsRequest, ...gax.CallOption) *NetworkEndpointGroupIterator
	ListNetworkEndpoints(context.Context, *computepb.ListNetworkEndpointsNetworkEndpointGroupsRequest, ...gax.CallOption) *NetworkEndpointWithHealthStatusIterator
	TestIamPermissions(context.Context, *computepb.TestIamPermissionsNetworkEndpointGroupRequest, ...gax.CallOption) (*computepb.TestPermissionsResponse, error)
}

// NetworkEndpointGroupsClient is a client for interacting with Google Compute Engine API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// The NetworkEndpointGroups API.
type NetworkEndpointGroupsClient struct {
	// The internal transport-dependent client.
	internalClient internalNetworkEndpointGroupsClient

	// The call options for this service.
	CallOptions *NetworkEndpointGroupsCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *NetworkEndpointGroupsClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *NetworkEndpointGroupsClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *NetworkEndpointGroupsClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// AggregatedList retrieves the list of network endpoint groups and sorts them by zone.
func (c *NetworkEndpointGroupsClient) AggregatedList(ctx context.Context, req *computepb.AggregatedListNetworkEndpointGroupsRequest, opts ...gax.CallOption) *NetworkEndpointGroupsScopedListPairIterator {
	return c.internalClient.AggregatedList(ctx, req, opts...)
}

// AttachNetworkEndpoints attach a list of network endpoints to the specified network endpoint group.
func (c *NetworkEndpointGroupsClient) AttachNetworkEndpoints(ctx context.Context, req *computepb.AttachNetworkEndpointsNetworkEndpointGroupRequest, opts ...gax.CallOption) (*Operation, error) {
	return c.internalClient.AttachNetworkEndpoints(ctx, req, opts...)
}

// Delete deletes the specified network endpoint group. The network endpoints in the NEG and the VM instances they belong to are not terminated when the NEG is deleted. Note that the NEG cannot be deleted if there are backend services referencing it.
func (c *NetworkEndpointGroupsClient) Delete(ctx context.Context, req *computepb.DeleteNetworkEndpointGroupRequest, opts ...gax.CallOption) (*Operation, error) {
	return c.internalClient.Delete(ctx, req, opts...)
}

// DetachNetworkEndpoints detach a list of network endpoints from the specified network endpoint group.
func (c *NetworkEndpointGroupsClient) DetachNetworkEndpoints(ctx context.Context, req *computepb.DetachNetworkEndpointsNetworkEndpointGroupRequest, opts ...gax.CallOption) (*Operation, error) {
	return c.internalClient.DetachNetworkEndpoints(ctx, req, opts...)
}

// Get returns the specified network endpoint group. Gets a list of available network endpoint groups by making a list() request.
func (c *NetworkEndpointGroupsClient) Get(ctx context.Context, req *computepb.GetNetworkEndpointGroupRequest, opts ...gax.CallOption) (*computepb.NetworkEndpointGroup, error) {
	return c.internalClient.Get(ctx, req, opts...)
}

// Insert creates a network endpoint group in the specified project using the parameters that are included in the request.
func (c *NetworkEndpointGroupsClient) Insert(ctx context.Context, req *computepb.InsertNetworkEndpointGroupRequest, opts ...gax.CallOption) (*Operation, error) {
	return c.internalClient.Insert(ctx, req, opts...)
}

// List retrieves the list of network endpoint groups that are located in the specified project and zone.
func (c *NetworkEndpointGroupsClient) List(ctx context.Context, req *computepb.ListNetworkEndpointGroupsRequest, opts ...gax.CallOption) *NetworkEndpointGroupIterator {
	return c.internalClient.List(ctx, req, opts...)
}

// ListNetworkEndpoints lists the network endpoints in the specified network endpoint group.
func (c *NetworkEndpointGroupsClient) ListNetworkEndpoints(ctx context.Context, req *computepb.ListNetworkEndpointsNetworkEndpointGroupsRequest, opts ...gax.CallOption) *NetworkEndpointWithHealthStatusIterator {
	return c.internalClient.ListNetworkEndpoints(ctx, req, opts...)
}

// TestIamPermissions returns permissions that a caller has on the specified resource.
func (c *NetworkEndpointGroupsClient) TestIamPermissions(ctx context.Context, req *computepb.TestIamPermissionsNetworkEndpointGroupRequest, opts ...gax.CallOption) (*computepb.TestPermissionsResponse, error) {
	return c.internalClient.TestIamPermissions(ctx, req, opts...)
}

// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type networkEndpointGroupsRESTClient struct {
	// The http endpoint to connect to.
	endpoint string

	// The http client.
	httpClient *http.Client

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewNetworkEndpointGroupsRESTClient creates a new network endpoint groups rest client.
//
// The NetworkEndpointGroups API.
func NewNetworkEndpointGroupsRESTClient(ctx context.Context, opts ...option.ClientOption) (*NetworkEndpointGroupsClient, error) {
	clientOpts := append(defaultNetworkEndpointGroupsRESTClientOptions(), opts...)
	httpClient, endpoint, err := httptransport.NewClient(ctx, clientOpts...)
	if err != nil {
		return nil, err
	}

	c := &networkEndpointGroupsRESTClient{
		endpoint:   endpoint,
		httpClient: httpClient,
	}
	c.setGoogleClientInfo()

	return &NetworkEndpointGroupsClient{internalClient: c, CallOptions: &NetworkEndpointGroupsCallOptions{}}, nil
}

func defaultNetworkEndpointGroupsRESTClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("https://compute.googleapis.com"),
		internaloption.WithDefaultMTLSEndpoint("https://compute.mtls.googleapis.com"),
		internaloption.WithDefaultAudience("https://compute.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
	}
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *networkEndpointGroupsRESTClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", versionClient, "gax", gax.Version, "rest", "UNKNOWN")
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *networkEndpointGroupsRESTClient) Close() error {
	// Replace httpClient with nil to force cleanup.
	c.httpClient = nil
	return nil
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *networkEndpointGroupsRESTClient) Connection() *grpc.ClientConn {
	return nil
}

// AggregatedList retrieves the list of network endpoint groups and sorts them by zone.
func (c *networkEndpointGroupsRESTClient) AggregatedList(ctx context.Context, req *computepb.AggregatedListNetworkEndpointGroupsRequest, opts ...gax.CallOption) *NetworkEndpointGroupsScopedListPairIterator {
	it := &NetworkEndpointGroupsScopedListPairIterator{}
	req = proto.Clone(req).(*computepb.AggregatedListNetworkEndpointGroupsRequest)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	it.InternalFetch = func(pageSize int, pageToken string) ([]NetworkEndpointGroupsScopedListPair, string, error) {
		resp := &computepb.NetworkEndpointGroupAggregatedList{}
		if pageToken != "" {
			req.PageToken = proto.String(pageToken)
		}
		if pageSize > math.MaxInt32 {
			req.MaxResults = proto.Uint32(math.MaxInt32)
		} else if pageSize != 0 {
			req.MaxResults = proto.Uint32(uint32(pageSize))
		}
		baseUrl, _ := url.Parse(c.endpoint)
		baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/aggregated/networkEndpointGroups", req.GetProject())

		params := url.Values{}
		if req != nil && req.Filter != nil {
			params.Add("filter", fmt.Sprintf("%v", req.GetFilter()))
		}
		if req != nil && req.IncludeAllScopes != nil {
			params.Add("includeAllScopes", fmt.Sprintf("%v", req.GetIncludeAllScopes()))
		}
		if req != nil && req.MaxResults != nil {
			params.Add("maxResults", fmt.Sprintf("%v", req.GetMaxResults()))
		}
		if req != nil && req.OrderBy != nil {
			params.Add("orderBy", fmt.Sprintf("%v", req.GetOrderBy()))
		}
		if req != nil && req.PageToken != nil {
			params.Add("pageToken", fmt.Sprintf("%v", req.GetPageToken()))
		}
		if req != nil && req.ReturnPartialSuccess != nil {
			params.Add("returnPartialSuccess", fmt.Sprintf("%v", req.GetReturnPartialSuccess()))
		}

		baseUrl.RawQuery = params.Encode()

		httpReq, err := http.NewRequest("GET", baseUrl.String(), nil)
		if err != nil {
			return nil, "", err
		}

		// Set the headers
		for k, v := range c.xGoogMetadata {
			httpReq.Header[k] = v
		}

		httpReq.Header["Content-Type"] = []string{"application/json"}
		httpRsp, err := c.httpClient.Do(httpReq)
		if err != nil {
			return nil, "", err
		}
		defer httpRsp.Body.Close()

		if err = googleapi.CheckResponse(httpRsp); err != nil {
			return nil, "", err
		}

		buf, err := ioutil.ReadAll(httpRsp.Body)
		if err != nil {
			return nil, "", err
		}

		unm.Unmarshal(buf, resp)
		it.Response = resp

		elems := make([]NetworkEndpointGroupsScopedListPair, 0, len(resp.GetItems()))
		for k, v := range resp.GetItems() {
			elems = append(elems, NetworkEndpointGroupsScopedListPair{k, v})
		}
		sort.Slice(elems, func(i, j int) bool { return elems[i].Key < elems[j].Key })

		return elems, resp.GetNextPageToken(), nil
	}

	fetch := func(pageSize int, pageToken string) (string, error) {
		items, nextPageToken, err := it.InternalFetch(pageSize, pageToken)
		if err != nil {
			return "", err
		}
		it.items = append(it.items, items...)
		return nextPageToken, nil
	}

	it.pageInfo, it.nextFunc = iterator.NewPageInfo(fetch, it.bufLen, it.takeBuf)
	it.pageInfo.MaxSize = int(req.GetMaxResults())
	it.pageInfo.Token = req.GetPageToken()

	return it
}

// AttachNetworkEndpoints attach a list of network endpoints to the specified network endpoint group.
func (c *networkEndpointGroupsRESTClient) AttachNetworkEndpoints(ctx context.Context, req *computepb.AttachNetworkEndpointsNetworkEndpointGroupRequest, opts ...gax.CallOption) (*Operation, error) {
	m := protojson.MarshalOptions{AllowPartial: true}
	body := req.GetNetworkEndpointGroupsAttachEndpointsRequestResource()
	jsonReq, err := m.Marshal(body)
	if err != nil {
		return nil, err
	}

	baseUrl, _ := url.Parse(c.endpoint)
	baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/zones/%v/networkEndpointGroups/%v/attachNetworkEndpoints", req.GetProject(), req.GetZone(), req.GetNetworkEndpointGroup())

	params := url.Values{}
	if req != nil && req.RequestId != nil {
		params.Add("requestId", fmt.Sprintf("%v", req.GetRequestId()))
	}

	baseUrl.RawQuery = params.Encode()

	httpReq, err := http.NewRequest("POST", baseUrl.String(), bytes.NewReader(jsonReq))
	if err != nil {
		return nil, err
	}
	httpReq = httpReq.WithContext(ctx)
	// Set the headers
	for k, v := range c.xGoogMetadata {
		httpReq.Header[k] = v
	}
	httpReq.Header["Content-Type"] = []string{"application/json"}

	httpRsp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer httpRsp.Body.Close()

	if err = googleapi.CheckResponse(httpRsp); err != nil {
		return nil, err
	}

	buf, err := ioutil.ReadAll(httpRsp.Body)
	if err != nil {
		return nil, err
	}

	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	rsp := &computepb.Operation{}

	if err := unm.Unmarshal(buf, rsp); err != nil {
		return nil, maybeUnknownEnum(err)
	}
	op := &Operation{proto: rsp}
	return op, err
}

// Delete deletes the specified network endpoint group. The network endpoints in the NEG and the VM instances they belong to are not terminated when the NEG is deleted. Note that the NEG cannot be deleted if there are backend services referencing it.
func (c *networkEndpointGroupsRESTClient) Delete(ctx context.Context, req *computepb.DeleteNetworkEndpointGroupRequest, opts ...gax.CallOption) (*Operation, error) {
	baseUrl, _ := url.Parse(c.endpoint)
	baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/zones/%v/networkEndpointGroups/%v", req.GetProject(), req.GetZone(), req.GetNetworkEndpointGroup())

	params := url.Values{}
	if req != nil && req.RequestId != nil {
		params.Add("requestId", fmt.Sprintf("%v", req.GetRequestId()))
	}

	baseUrl.RawQuery = params.Encode()

	httpReq, err := http.NewRequest("DELETE", baseUrl.String(), nil)
	if err != nil {
		return nil, err
	}
	httpReq = httpReq.WithContext(ctx)
	// Set the headers
	for k, v := range c.xGoogMetadata {
		httpReq.Header[k] = v
	}
	httpReq.Header["Content-Type"] = []string{"application/json"}

	httpRsp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer httpRsp.Body.Close()

	if err = googleapi.CheckResponse(httpRsp); err != nil {
		return nil, err
	}

	buf, err := ioutil.ReadAll(httpRsp.Body)
	if err != nil {
		return nil, err
	}

	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	rsp := &computepb.Operation{}

	if err := unm.Unmarshal(buf, rsp); err != nil {
		return nil, maybeUnknownEnum(err)
	}
	op := &Operation{proto: rsp}
	return op, err
}

// DetachNetworkEndpoints detach a list of network endpoints from the specified network endpoint group.
func (c *networkEndpointGroupsRESTClient) DetachNetworkEndpoints(ctx context.Context, req *computepb.DetachNetworkEndpointsNetworkEndpointGroupRequest, opts ...gax.CallOption) (*Operation, error) {
	m := protojson.MarshalOptions{AllowPartial: true}
	body := req.GetNetworkEndpointGroupsDetachEndpointsRequestResource()
	jsonReq, err := m.Marshal(body)
	if err != nil {
		return nil, err
	}

	baseUrl, _ := url.Parse(c.endpoint)
	baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/zones/%v/networkEndpointGroups/%v/detachNetworkEndpoints", req.GetProject(), req.GetZone(), req.GetNetworkEndpointGroup())

	params := url.Values{}
	if req != nil && req.RequestId != nil {
		params.Add("requestId", fmt.Sprintf("%v", req.GetRequestId()))
	}

	baseUrl.RawQuery = params.Encode()

	httpReq, err := http.NewRequest("POST", baseUrl.String(), bytes.NewReader(jsonReq))
	if err != nil {
		return nil, err
	}
	httpReq = httpReq.WithContext(ctx)
	// Set the headers
	for k, v := range c.xGoogMetadata {
		httpReq.Header[k] = v
	}
	httpReq.Header["Content-Type"] = []string{"application/json"}

	httpRsp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer httpRsp.Body.Close()

	if err = googleapi.CheckResponse(httpRsp); err != nil {
		return nil, err
	}

	buf, err := ioutil.ReadAll(httpRsp.Body)
	if err != nil {
		return nil, err
	}

	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	rsp := &computepb.Operation{}

	if err := unm.Unmarshal(buf, rsp); err != nil {
		return nil, maybeUnknownEnum(err)
	}
	op := &Operation{proto: rsp}
	return op, err
}

// Get returns the specified network endpoint group. Gets a list of available network endpoint groups by making a list() request.
func (c *networkEndpointGroupsRESTClient) Get(ctx context.Context, req *computepb.GetNetworkEndpointGroupRequest, opts ...gax.CallOption) (*computepb.NetworkEndpointGroup, error) {
	baseUrl, _ := url.Parse(c.endpoint)
	baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/zones/%v/networkEndpointGroups/%v", req.GetProject(), req.GetZone(), req.GetNetworkEndpointGroup())

	httpReq, err := http.NewRequest("GET", baseUrl.String(), nil)
	if err != nil {
		return nil, err
	}
	httpReq = httpReq.WithContext(ctx)
	// Set the headers
	for k, v := range c.xGoogMetadata {
		httpReq.Header[k] = v
	}
	httpReq.Header["Content-Type"] = []string{"application/json"}

	httpRsp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer httpRsp.Body.Close()

	if err = googleapi.CheckResponse(httpRsp); err != nil {
		return nil, err
	}

	buf, err := ioutil.ReadAll(httpRsp.Body)
	if err != nil {
		return nil, err
	}

	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	rsp := &computepb.NetworkEndpointGroup{}

	if err := unm.Unmarshal(buf, rsp); err != nil {
		return nil, maybeUnknownEnum(err)
	}
	return rsp, nil
}

// Insert creates a network endpoint group in the specified project using the parameters that are included in the request.
func (c *networkEndpointGroupsRESTClient) Insert(ctx context.Context, req *computepb.InsertNetworkEndpointGroupRequest, opts ...gax.CallOption) (*Operation, error) {
	m := protojson.MarshalOptions{AllowPartial: true}
	body := req.GetNetworkEndpointGroupResource()
	jsonReq, err := m.Marshal(body)
	if err != nil {
		return nil, err
	}

	baseUrl, _ := url.Parse(c.endpoint)
	baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/zones/%v/networkEndpointGroups", req.GetProject(), req.GetZone())

	params := url.Values{}
	if req != nil && req.RequestId != nil {
		params.Add("requestId", fmt.Sprintf("%v", req.GetRequestId()))
	}

	baseUrl.RawQuery = params.Encode()

	httpReq, err := http.NewRequest("POST", baseUrl.String(), bytes.NewReader(jsonReq))
	if err != nil {
		return nil, err
	}
	httpReq = httpReq.WithContext(ctx)
	// Set the headers
	for k, v := range c.xGoogMetadata {
		httpReq.Header[k] = v
	}
	httpReq.Header["Content-Type"] = []string{"application/json"}

	httpRsp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer httpRsp.Body.Close()

	if err = googleapi.CheckResponse(httpRsp); err != nil {
		return nil, err
	}

	buf, err := ioutil.ReadAll(httpRsp.Body)
	if err != nil {
		return nil, err
	}

	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	rsp := &computepb.Operation{}

	if err := unm.Unmarshal(buf, rsp); err != nil {
		return nil, maybeUnknownEnum(err)
	}
	op := &Operation{proto: rsp}
	return op, err
}

// List retrieves the list of network endpoint groups that are located in the specified project and zone.
func (c *networkEndpointGroupsRESTClient) List(ctx context.Context, req *computepb.ListNetworkEndpointGroupsRequest, opts ...gax.CallOption) *NetworkEndpointGroupIterator {
	it := &NetworkEndpointGroupIterator{}
	req = proto.Clone(req).(*computepb.ListNetworkEndpointGroupsRequest)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	it.InternalFetch = func(pageSize int, pageToken string) ([]*computepb.NetworkEndpointGroup, string, error) {
		resp := &computepb.NetworkEndpointGroupList{}
		if pageToken != "" {
			req.PageToken = proto.String(pageToken)
		}
		if pageSize > math.MaxInt32 {
			req.MaxResults = proto.Uint32(math.MaxInt32)
		} else if pageSize != 0 {
			req.MaxResults = proto.Uint32(uint32(pageSize))
		}
		baseUrl, _ := url.Parse(c.endpoint)
		baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/zones/%v/networkEndpointGroups", req.GetProject(), req.GetZone())

		params := url.Values{}
		if req != nil && req.Filter != nil {
			params.Add("filter", fmt.Sprintf("%v", req.GetFilter()))
		}
		if req != nil && req.MaxResults != nil {
			params.Add("maxResults", fmt.Sprintf("%v", req.GetMaxResults()))
		}
		if req != nil && req.OrderBy != nil {
			params.Add("orderBy", fmt.Sprintf("%v", req.GetOrderBy()))
		}
		if req != nil && req.PageToken != nil {
			params.Add("pageToken", fmt.Sprintf("%v", req.GetPageToken()))
		}
		if req != nil && req.ReturnPartialSuccess != nil {
			params.Add("returnPartialSuccess", fmt.Sprintf("%v", req.GetReturnPartialSuccess()))
		}

		baseUrl.RawQuery = params.Encode()

		httpReq, err := http.NewRequest("GET", baseUrl.String(), nil)
		if err != nil {
			return nil, "", err
		}

		// Set the headers
		for k, v := range c.xGoogMetadata {
			httpReq.Header[k] = v
		}

		httpReq.Header["Content-Type"] = []string{"application/json"}
		httpRsp, err := c.httpClient.Do(httpReq)
		if err != nil {
			return nil, "", err
		}
		defer httpRsp.Body.Close()

		if err = googleapi.CheckResponse(httpRsp); err != nil {
			return nil, "", err
		}

		buf, err := ioutil.ReadAll(httpRsp.Body)
		if err != nil {
			return nil, "", err
		}

		unm.Unmarshal(buf, resp)
		it.Response = resp
		return resp.GetItems(), resp.GetNextPageToken(), nil
	}

	fetch := func(pageSize int, pageToken string) (string, error) {
		items, nextPageToken, err := it.InternalFetch(pageSize, pageToken)
		if err != nil {
			return "", err
		}
		it.items = append(it.items, items...)
		return nextPageToken, nil
	}

	it.pageInfo, it.nextFunc = iterator.NewPageInfo(fetch, it.bufLen, it.takeBuf)
	it.pageInfo.MaxSize = int(req.GetMaxResults())
	it.pageInfo.Token = req.GetPageToken()

	return it
}

// ListNetworkEndpoints lists the network endpoints in the specified network endpoint group.
func (c *networkEndpointGroupsRESTClient) ListNetworkEndpoints(ctx context.Context, req *computepb.ListNetworkEndpointsNetworkEndpointGroupsRequest, opts ...gax.CallOption) *NetworkEndpointWithHealthStatusIterator {
	it := &NetworkEndpointWithHealthStatusIterator{}
	req = proto.Clone(req).(*computepb.ListNetworkEndpointsNetworkEndpointGroupsRequest)
	m := protojson.MarshalOptions{AllowPartial: true, UseProtoNames: false}
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	it.InternalFetch = func(pageSize int, pageToken string) ([]*computepb.NetworkEndpointWithHealthStatus, string, error) {
		resp := &computepb.NetworkEndpointGroupsListNetworkEndpoints{}
		if pageToken != "" {
			req.PageToken = proto.String(pageToken)
		}
		if pageSize > math.MaxInt32 {
			req.MaxResults = proto.Uint32(math.MaxInt32)
		} else if pageSize != 0 {
			req.MaxResults = proto.Uint32(uint32(pageSize))
		}
		jsonReq, err := m.Marshal(req)
		if err != nil {
			return nil, "", err
		}

		baseUrl, _ := url.Parse(c.endpoint)
		baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/zones/%v/networkEndpointGroups/%v/listNetworkEndpoints", req.GetProject(), req.GetZone(), req.GetNetworkEndpointGroup())

		params := url.Values{}
		if req != nil && req.Filter != nil {
			params.Add("filter", fmt.Sprintf("%v", req.GetFilter()))
		}
		if req != nil && req.MaxResults != nil {
			params.Add("maxResults", fmt.Sprintf("%v", req.GetMaxResults()))
		}
		if req != nil && req.OrderBy != nil {
			params.Add("orderBy", fmt.Sprintf("%v", req.GetOrderBy()))
		}
		if req != nil && req.PageToken != nil {
			params.Add("pageToken", fmt.Sprintf("%v", req.GetPageToken()))
		}
		if req != nil && req.ReturnPartialSuccess != nil {
			params.Add("returnPartialSuccess", fmt.Sprintf("%v", req.GetReturnPartialSuccess()))
		}

		baseUrl.RawQuery = params.Encode()

		httpReq, err := http.NewRequest("POST", baseUrl.String(), bytes.NewReader(jsonReq))
		if err != nil {
			return nil, "", err
		}

		// Set the headers
		for k, v := range c.xGoogMetadata {
			httpReq.Header[k] = v
		}

		httpReq.Header["Content-Type"] = []string{"application/json"}
		httpRsp, err := c.httpClient.Do(httpReq)
		if err != nil {
			return nil, "", err
		}
		defer httpRsp.Body.Close()

		if err = googleapi.CheckResponse(httpRsp); err != nil {
			return nil, "", err
		}

		buf, err := ioutil.ReadAll(httpRsp.Body)
		if err != nil {
			return nil, "", err
		}

		unm.Unmarshal(buf, resp)
		it.Response = resp
		return resp.GetItems(), resp.GetNextPageToken(), nil
	}

	fetch := func(pageSize int, pageToken string) (string, error) {
		items, nextPageToken, err := it.InternalFetch(pageSize, pageToken)
		if err != nil {
			return "", err
		}
		it.items = append(it.items, items...)
		return nextPageToken, nil
	}

	it.pageInfo, it.nextFunc = iterator.NewPageInfo(fetch, it.bufLen, it.takeBuf)
	it.pageInfo.MaxSize = int(req.GetMaxResults())
	it.pageInfo.Token = req.GetPageToken()

	return it
}

// TestIamPermissions returns permissions that a caller has on the specified resource.
func (c *networkEndpointGroupsRESTClient) TestIamPermissions(ctx context.Context, req *computepb.TestIamPermissionsNetworkEndpointGroupRequest, opts ...gax.CallOption) (*computepb.TestPermissionsResponse, error) {
	m := protojson.MarshalOptions{AllowPartial: true}
	body := req.GetTestPermissionsRequestResource()
	jsonReq, err := m.Marshal(body)
	if err != nil {
		return nil, err
	}

	baseUrl, _ := url.Parse(c.endpoint)
	baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/zones/%v/networkEndpointGroups/%v/testIamPermissions", req.GetProject(), req.GetZone(), req.GetResource())

	httpReq, err := http.NewRequest("POST", baseUrl.String(), bytes.NewReader(jsonReq))
	if err != nil {
		return nil, err
	}
	httpReq = httpReq.WithContext(ctx)
	// Set the headers
	for k, v := range c.xGoogMetadata {
		httpReq.Header[k] = v
	}
	httpReq.Header["Content-Type"] = []string{"application/json"}

	httpRsp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer httpRsp.Body.Close()

	if err = googleapi.CheckResponse(httpRsp); err != nil {
		return nil, err
	}

	buf, err := ioutil.ReadAll(httpRsp.Body)
	if err != nil {
		return nil, err
	}

	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	rsp := &computepb.TestPermissionsResponse{}

	if err := unm.Unmarshal(buf, rsp); err != nil {
		return nil, maybeUnknownEnum(err)
	}
	return rsp, nil
}

// NetworkEndpointGroupsScopedListPair is a holder type for string/*computepb.NetworkEndpointGroupsScopedList map entries
type NetworkEndpointGroupsScopedListPair struct {
	Key   string
	Value *computepb.NetworkEndpointGroupsScopedList
}

// NetworkEndpointGroupsScopedListPairIterator manages a stream of NetworkEndpointGroupsScopedListPair.
type NetworkEndpointGroupsScopedListPairIterator struct {
	items    []NetworkEndpointGroupsScopedListPair
	pageInfo *iterator.PageInfo
	nextFunc func() error

	// Response is the raw response for the current page.
	// It must be cast to the RPC response type.
	// Calling Next() or InternalFetch() updates this value.
	Response interface{}

	// InternalFetch is for use by the Google Cloud Libraries only.
	// It is not part of the stable interface of this package.
	//
	// InternalFetch returns results from a single call to the underlying RPC.
	// The number of results is no greater than pageSize.
	// If there are no more results, nextPageToken is empty and err is nil.
	InternalFetch func(pageSize int, pageToken string) (results []NetworkEndpointGroupsScopedListPair, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *NetworkEndpointGroupsScopedListPairIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *NetworkEndpointGroupsScopedListPairIterator) Next() (NetworkEndpointGroupsScopedListPair, error) {
	var item NetworkEndpointGroupsScopedListPair
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *NetworkEndpointGroupsScopedListPairIterator) bufLen() int {
	return len(it.items)
}

func (it *NetworkEndpointGroupsScopedListPairIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}
