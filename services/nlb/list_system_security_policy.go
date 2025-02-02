package nlb

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

// ListSystemSecurityPolicy invokes the nlb.ListSystemSecurityPolicy API synchronously
func (client *Client) ListSystemSecurityPolicy(request *ListSystemSecurityPolicyRequest) (response *ListSystemSecurityPolicyResponse, err error) {
	response = CreateListSystemSecurityPolicyResponse()
	err = client.DoAction(request, response)
	return
}

// ListSystemSecurityPolicyWithChan invokes the nlb.ListSystemSecurityPolicy API asynchronously
func (client *Client) ListSystemSecurityPolicyWithChan(request *ListSystemSecurityPolicyRequest) (<-chan *ListSystemSecurityPolicyResponse, <-chan error) {
	responseChan := make(chan *ListSystemSecurityPolicyResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListSystemSecurityPolicy(request)
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

// ListSystemSecurityPolicyWithCallback invokes the nlb.ListSystemSecurityPolicy API asynchronously
func (client *Client) ListSystemSecurityPolicyWithCallback(request *ListSystemSecurityPolicyRequest, callback func(response *ListSystemSecurityPolicyResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListSystemSecurityPolicyResponse
		var err error
		defer close(result)
		response, err = client.ListSystemSecurityPolicy(request)
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

// ListSystemSecurityPolicyRequest is the request struct for api ListSystemSecurityPolicy
type ListSystemSecurityPolicyRequest struct {
	*requests.RpcRequest
	Channel              string `position:"Body" name:"Channel"`
	OwnerIdLoginEmail    string `position:"Body" name:"OwnerIdLoginEmail"`
	CallerBidLoginEmail  string `position:"Body" name:"CallerBidLoginEmail"`
	CallerUidLoginEmail  string `position:"Body" name:"CallerUidLoginEmail"`
	RequestContent       string `position:"Body" name:"RequestContent"`
	ResourceOwnerAccount string `position:"Body" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Body" name:"OwnerAccount"`
}

// ListSystemSecurityPolicyResponse is the response struct for api ListSystemSecurityPolicy
type ListSystemSecurityPolicyResponse struct {
	*responses.BaseResponse
	RequestId        string            `json:"RequestId" xml:"RequestId"`
	Success          bool              `json:"Success" xml:"Success"`
	Code             string            `json:"Code" xml:"Code"`
	Message          string            `json:"Message" xml:"Message"`
	HttpStatusCode   int               `json:"HttpStatusCode" xml:"HttpStatusCode"`
	DynamicCode      string            `json:"DynamicCode" xml:"DynamicCode"`
	DynamicMessage   string            `json:"DynamicMessage" xml:"DynamicMessage"`
	SecurityPolicies []SecurityPolicie `json:"SecurityPolicies" xml:"SecurityPolicies"`
}

// CreateListSystemSecurityPolicyRequest creates a request to invoke ListSystemSecurityPolicy API
func CreateListSystemSecurityPolicyRequest() (request *ListSystemSecurityPolicyRequest) {
	request = &ListSystemSecurityPolicyRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Nlb", "2022-04-30", "ListSystemSecurityPolicy", "nlb", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListSystemSecurityPolicyResponse creates a response to parse from ListSystemSecurityPolicy response
func CreateListSystemSecurityPolicyResponse() (response *ListSystemSecurityPolicyResponse) {
	response = &ListSystemSecurityPolicyResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
