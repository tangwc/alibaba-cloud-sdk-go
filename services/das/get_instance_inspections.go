package das

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

// GetInstanceInspections invokes the das.GetInstanceInspections API synchronously
func (client *Client) GetInstanceInspections(request *GetInstanceInspectionsRequest) (response *GetInstanceInspectionsResponse, err error) {
	response = CreateGetInstanceInspectionsResponse()
	err = client.DoAction(request, response)
	return
}

// GetInstanceInspectionsWithChan invokes the das.GetInstanceInspections API asynchronously
func (client *Client) GetInstanceInspectionsWithChan(request *GetInstanceInspectionsRequest) (<-chan *GetInstanceInspectionsResponse, <-chan error) {
	responseChan := make(chan *GetInstanceInspectionsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetInstanceInspections(request)
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

// GetInstanceInspectionsWithCallback invokes the das.GetInstanceInspections API asynchronously
func (client *Client) GetInstanceInspectionsWithCallback(request *GetInstanceInspectionsRequest, callback func(response *GetInstanceInspectionsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetInstanceInspectionsResponse
		var err error
		defer close(result)
		response, err = client.GetInstanceInspections(request)
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

// GetInstanceInspectionsRequest is the request struct for api GetInstanceInspections
type GetInstanceInspectionsRequest struct {
	*requests.RpcRequest
	EndTime         string `position:"Query" name:"EndTime"`
	SearchMap       string `position:"Query" name:"SearchMap"`
	StartTime       string `position:"Query" name:"StartTime"`
	InstanceArea    string `position:"Query" name:"InstanceArea"`
	ResourceGroupId string `position:"Query" name:"ResourceGroupId"`
	Engine          string `position:"Query" name:"Engine"`
	PageNo          string `position:"Query" name:"PageNo"`
	PageSize        string `position:"Query" name:"PageSize"`
}

// GetInstanceInspectionsResponse is the response struct for api GetInstanceInspections
type GetInstanceInspectionsResponse struct {
	*responses.BaseResponse
	Message   string                       `json:"Message" xml:"Message"`
	RequestId string                       `json:"RequestId" xml:"RequestId"`
	Code      string                       `json:"Code" xml:"Code"`
	Success   string                       `json:"Success" xml:"Success"`
	Data      DataInGetInstanceInspections `json:"Data" xml:"Data"`
}

// CreateGetInstanceInspectionsRequest creates a request to invoke GetInstanceInspections API
func CreateGetInstanceInspectionsRequest() (request *GetInstanceInspectionsRequest) {
	request = &GetInstanceInspectionsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("DAS", "2020-01-16", "GetInstanceInspections", "", "")
	request.Method = requests.POST
	return
}

// CreateGetInstanceInspectionsResponse creates a response to parse from GetInstanceInspections response
func CreateGetInstanceInspectionsResponse() (response *GetInstanceInspectionsResponse) {
	response = &GetInstanceInspectionsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
