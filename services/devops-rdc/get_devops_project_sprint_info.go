package devops_rdc

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

// GetDevopsProjectSprintInfo invokes the devops_rdc.GetDevopsProjectSprintInfo API synchronously
func (client *Client) GetDevopsProjectSprintInfo(request *GetDevopsProjectSprintInfoRequest) (response *GetDevopsProjectSprintInfoResponse, err error) {
	response = CreateGetDevopsProjectSprintInfoResponse()
	err = client.DoAction(request, response)
	return
}

// GetDevopsProjectSprintInfoWithChan invokes the devops_rdc.GetDevopsProjectSprintInfo API asynchronously
func (client *Client) GetDevopsProjectSprintInfoWithChan(request *GetDevopsProjectSprintInfoRequest) (<-chan *GetDevopsProjectSprintInfoResponse, <-chan error) {
	responseChan := make(chan *GetDevopsProjectSprintInfoResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetDevopsProjectSprintInfo(request)
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

// GetDevopsProjectSprintInfoWithCallback invokes the devops_rdc.GetDevopsProjectSprintInfo API asynchronously
func (client *Client) GetDevopsProjectSprintInfoWithCallback(request *GetDevopsProjectSprintInfoRequest, callback func(response *GetDevopsProjectSprintInfoResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetDevopsProjectSprintInfoResponse
		var err error
		defer close(result)
		response, err = client.GetDevopsProjectSprintInfo(request)
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

// GetDevopsProjectSprintInfoRequest is the request struct for api GetDevopsProjectSprintInfo
type GetDevopsProjectSprintInfoRequest struct {
	*requests.RpcRequest
	SprintId string `position:"Body" name:"SprintId"`
	OrgId    string `position:"Body" name:"OrgId"`
}

// GetDevopsProjectSprintInfoResponse is the response struct for api GetDevopsProjectSprintInfo
type GetDevopsProjectSprintInfoResponse struct {
	*responses.BaseResponse
	Successful bool   `json:"Successful" xml:"Successful"`
	ErrorCode  string `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMsg   string `json:"ErrorMsg" xml:"ErrorMsg"`
	RequestId  string `json:"RequestId" xml:"RequestId"`
	Object     Object `json:"Object" xml:"Object"`
}

// CreateGetDevopsProjectSprintInfoRequest creates a request to invoke GetDevopsProjectSprintInfo API
func CreateGetDevopsProjectSprintInfoRequest() (request *GetDevopsProjectSprintInfoRequest) {
	request = &GetDevopsProjectSprintInfoRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("devops-rdc", "2020-03-03", "GetDevopsProjectSprintInfo", "1.9.6", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetDevopsProjectSprintInfoResponse creates a response to parse from GetDevopsProjectSprintInfo response
func CreateGetDevopsProjectSprintInfoResponse() (response *GetDevopsProjectSprintInfoResponse) {
	response = &GetDevopsProjectSprintInfoResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
