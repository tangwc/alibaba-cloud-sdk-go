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

// GetMultiAccountResourceConfiguration invokes the resourcecenter.GetMultiAccountResourceConfiguration API synchronously
func (client *Client) GetMultiAccountResourceConfiguration(request *GetMultiAccountResourceConfigurationRequest) (response *GetMultiAccountResourceConfigurationResponse, err error) {
	response = CreateGetMultiAccountResourceConfigurationResponse()
	err = client.DoAction(request, response)
	return
}

// GetMultiAccountResourceConfigurationWithChan invokes the resourcecenter.GetMultiAccountResourceConfiguration API asynchronously
func (client *Client) GetMultiAccountResourceConfigurationWithChan(request *GetMultiAccountResourceConfigurationRequest) (<-chan *GetMultiAccountResourceConfigurationResponse, <-chan error) {
	responseChan := make(chan *GetMultiAccountResourceConfigurationResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetMultiAccountResourceConfiguration(request)
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

// GetMultiAccountResourceConfigurationWithCallback invokes the resourcecenter.GetMultiAccountResourceConfiguration API asynchronously
func (client *Client) GetMultiAccountResourceConfigurationWithCallback(request *GetMultiAccountResourceConfigurationRequest, callback func(response *GetMultiAccountResourceConfigurationResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetMultiAccountResourceConfigurationResponse
		var err error
		defer close(result)
		response, err = client.GetMultiAccountResourceConfiguration(request)
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

// GetMultiAccountResourceConfigurationRequest is the request struct for api GetMultiAccountResourceConfiguration
type GetMultiAccountResourceConfigurationRequest struct {
	*requests.RpcRequest
	ResourceId       string `position:"Query" name:"ResourceId"`
	ResourceType     string `position:"Query" name:"ResourceType"`
	AccountId        string `position:"Query" name:"AccountId"`
	ResourceRegionId string `position:"Query" name:"ResourceRegionId"`
}

// GetMultiAccountResourceConfigurationResponse is the response struct for api GetMultiAccountResourceConfiguration
type GetMultiAccountResourceConfigurationResponse struct {
	*responses.BaseResponse
}

// CreateGetMultiAccountResourceConfigurationRequest creates a request to invoke GetMultiAccountResourceConfiguration API
func CreateGetMultiAccountResourceConfigurationRequest() (request *GetMultiAccountResourceConfigurationRequest) {
	request = &GetMultiAccountResourceConfigurationRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ResourceCenter", "2022-12-01", "GetMultiAccountResourceConfiguration", "", "")
	request.Method = requests.POST
	return
}

// CreateGetMultiAccountResourceConfigurationResponse creates a response to parse from GetMultiAccountResourceConfiguration response
func CreateGetMultiAccountResourceConfigurationResponse() (response *GetMultiAccountResourceConfigurationResponse) {
	response = &GetMultiAccountResourceConfigurationResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
