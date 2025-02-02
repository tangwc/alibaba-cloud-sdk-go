package eflo

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

// GetVcc invokes the eflo.GetVcc API synchronously
func (client *Client) GetVcc(request *GetVccRequest) (response *GetVccResponse, err error) {
	response = CreateGetVccResponse()
	err = client.DoAction(request, response)
	return
}

// GetVccWithChan invokes the eflo.GetVcc API asynchronously
func (client *Client) GetVccWithChan(request *GetVccRequest) (<-chan *GetVccResponse, <-chan error) {
	responseChan := make(chan *GetVccResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetVcc(request)
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

// GetVccWithCallback invokes the eflo.GetVcc API asynchronously
func (client *Client) GetVccWithCallback(request *GetVccRequest, callback func(response *GetVccResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetVccResponse
		var err error
		defer close(result)
		response, err = client.GetVcc(request)
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

// GetVccRequest is the request struct for api GetVcc
type GetVccRequest struct {
	*requests.RpcRequest
	VccId string `position:"Body" name:"VccId"`
}

// GetVccResponse is the response struct for api GetVcc
type GetVccResponse struct {
	*responses.BaseResponse
	Code      int     `json:"Code" xml:"Code"`
	Message   string  `json:"Message" xml:"Message"`
	RequestId string  `json:"RequestId" xml:"RequestId"`
	Content   Content `json:"Content" xml:"Content"`
}

// CreateGetVccRequest creates a request to invoke GetVcc API
func CreateGetVccRequest() (request *GetVccRequest) {
	request = &GetVccRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("eflo", "2022-05-30", "GetVcc", "eflo", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetVccResponse creates a response to parse from GetVcc response
func CreateGetVccResponse() (response *GetVccResponse) {
	response = &GetVccResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
