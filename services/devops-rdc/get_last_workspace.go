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

// GetLastWorkspace invokes the devops_rdc.GetLastWorkspace API synchronously
func (client *Client) GetLastWorkspace(request *GetLastWorkspaceRequest) (response *GetLastWorkspaceResponse, err error) {
	response = CreateGetLastWorkspaceResponse()
	err = client.DoAction(request, response)
	return
}

// GetLastWorkspaceWithChan invokes the devops_rdc.GetLastWorkspace API asynchronously
func (client *Client) GetLastWorkspaceWithChan(request *GetLastWorkspaceRequest) (<-chan *GetLastWorkspaceResponse, <-chan error) {
	responseChan := make(chan *GetLastWorkspaceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetLastWorkspace(request)
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

// GetLastWorkspaceWithCallback invokes the devops_rdc.GetLastWorkspace API asynchronously
func (client *Client) GetLastWorkspaceWithCallback(request *GetLastWorkspaceRequest, callback func(response *GetLastWorkspaceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetLastWorkspaceResponse
		var err error
		defer close(result)
		response, err = client.GetLastWorkspace(request)
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

// GetLastWorkspaceRequest is the request struct for api GetLastWorkspace
type GetLastWorkspaceRequest struct {
	*requests.RpcRequest
	RealPk string `position:"Body" name:"RealPk"`
	OrgId  string `position:"Body" name:"OrgId"`
}

// GetLastWorkspaceResponse is the response struct for api GetLastWorkspace
type GetLastWorkspaceResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	ErrorCode    string `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
	Success      bool   `json:"Success" xml:"Success"`
	Object       string `json:"Object" xml:"Object"`
}

// CreateGetLastWorkspaceRequest creates a request to invoke GetLastWorkspace API
func CreateGetLastWorkspaceRequest() (request *GetLastWorkspaceRequest) {
	request = &GetLastWorkspaceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("devops-rdc", "2020-03-03", "GetLastWorkspace", "1.9.6", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetLastWorkspaceResponse creates a response to parse from GetLastWorkspace response
func CreateGetLastWorkspaceResponse() (response *GetLastWorkspaceResponse) {
	response = &GetLastWorkspaceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
