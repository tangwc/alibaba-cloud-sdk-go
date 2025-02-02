package oceanbasepro

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

// DescribeCharset invokes the oceanbasepro.DescribeCharset API synchronously
func (client *Client) DescribeCharset(request *DescribeCharsetRequest) (response *DescribeCharsetResponse, err error) {
	response = CreateDescribeCharsetResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeCharsetWithChan invokes the oceanbasepro.DescribeCharset API asynchronously
func (client *Client) DescribeCharsetWithChan(request *DescribeCharsetRequest) (<-chan *DescribeCharsetResponse, <-chan error) {
	responseChan := make(chan *DescribeCharsetResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeCharset(request)
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

// DescribeCharsetWithCallback invokes the oceanbasepro.DescribeCharset API asynchronously
func (client *Client) DescribeCharsetWithCallback(request *DescribeCharsetRequest, callback func(response *DescribeCharsetResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeCharsetResponse
		var err error
		defer close(result)
		response, err = client.DescribeCharset(request)
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

// DescribeCharsetRequest is the request struct for api DescribeCharset
type DescribeCharsetRequest struct {
	*requests.RpcRequest
	TenantMode string `position:"Body" name:"TenantMode"`
}

// DescribeCharsetResponse is the response struct for api DescribeCharset
type DescribeCharsetResponse struct {
	*responses.BaseResponse
	RequestId string        `json:"RequestId" xml:"RequestId"`
	Charset   []CharsetItem `json:"Charset" xml:"Charset"`
}

// CreateDescribeCharsetRequest creates a request to invoke DescribeCharset API
func CreateDescribeCharsetRequest() (request *DescribeCharsetRequest) {
	request = &DescribeCharsetRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("OceanBasePro", "2019-09-01", "DescribeCharset", "oceanbase", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeCharsetResponse creates a response to parse from DescribeCharset response
func CreateDescribeCharsetResponse() (response *DescribeCharsetResponse) {
	response = &DescribeCharsetResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
