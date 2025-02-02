package resourcecenter

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// ListResourceTypes invokes the resourcecenter.ListResourceTypes API synchronously
func (client *Client) ListResourceTypes(request *ListResourceTypesRequest) (response *ListResourceTypesResponse, err error) {
	response = CreateListResourceTypesResponse()
	err = client.DoAction(request, response)
	return
}

// ListResourceTypesWithChan invokes the resourcecenter.ListResourceTypes API asynchronously
func (client *Client) ListResourceTypesWithChan(request *ListResourceTypesRequest) (<-chan *ListResourceTypesResponse, <-chan error) {
	responseChan := make(chan *ListResourceTypesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListResourceTypes(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// ListResourceTypesWithCallback invokes the resourcecenter.ListResourceTypes API asynchronously
func (client *Client) ListResourceTypesWithCallback(request *ListResourceTypesRequest, callback func(response *ListResourceTypesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListResourceTypesResponse
		var err error
		defer close(result)
		response, err = client.ListResourceTypes(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// ListResourceTypesRequest is the request struct for api ListResourceTypes
type ListResourceTypesRequest struct {
	*requests.RpcRequest
	Query          *[]string `position:"Query" name:"Query"  type:"Repeated"`
	ResourceType   string    `position:"Query" name:"ResourceType"`
	AcceptLanguage string    `position:"Query" name:"AcceptLanguage"`
}

// ListResourceTypesResponse is the response struct for api ListResourceTypes
type ListResourceTypesResponse struct {
	*responses.BaseResponse
	RequestId      string         `json:"RequestId" xml:"RequestId"`
	Success        bool           `json:"Success" xml:"Success"`
	ErrorCode      string         `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMessage   string         `json:"ErrorMessage" xml:"ErrorMessage"`
	DynamicCode    string         `json:"DynamicCode" xml:"DynamicCode"`
	DynamicMessage string         `json:"DynamicMessage" xml:"DynamicMessage"`
	ResourceTypes  []ResourceType `json:"ResourceTypes" xml:"ResourceTypes"`
}

// CreateListResourceTypesRequest creates a request to invoke ListResourceTypes API
func CreateListResourceTypesRequest() (request *ListResourceTypesRequest) {
	request = &ListResourceTypesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ResourceCenter", "2022-12-01", "ListResourceTypes", "", "")
	request.Method = requests.POST
	return
}

// CreateListResourceTypesResponse creates a response to parse from ListResourceTypes response
func CreateListResourceTypesResponse() (response *ListResourceTypesResponse) {
	response = &ListResourceTypesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
