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

// ListDevopsProjectTasks invokes the devops_rdc.ListDevopsProjectTasks API synchronously
func (client *Client) ListDevopsProjectTasks(request *ListDevopsProjectTasksRequest) (response *ListDevopsProjectTasksResponse, err error) {
	response = CreateListDevopsProjectTasksResponse()
	err = client.DoAction(request, response)
	return
}

// ListDevopsProjectTasksWithChan invokes the devops_rdc.ListDevopsProjectTasks API asynchronously
func (client *Client) ListDevopsProjectTasksWithChan(request *ListDevopsProjectTasksRequest) (<-chan *ListDevopsProjectTasksResponse, <-chan error) {
	responseChan := make(chan *ListDevopsProjectTasksResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListDevopsProjectTasks(request)
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

// ListDevopsProjectTasksWithCallback invokes the devops_rdc.ListDevopsProjectTasks API asynchronously
func (client *Client) ListDevopsProjectTasksWithCallback(request *ListDevopsProjectTasksRequest, callback func(response *ListDevopsProjectTasksResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListDevopsProjectTasksResponse
		var err error
		defer close(result)
		response, err = client.ListDevopsProjectTasks(request)
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

// ListDevopsProjectTasksRequest is the request struct for api ListDevopsProjectTasks
type ListDevopsProjectTasksRequest struct {
	*requests.RpcRequest
	ProjectIds string `position:"Body" name:"ProjectIds"`
	OrgId      string `position:"Body" name:"OrgId"`
}

// ListDevopsProjectTasksResponse is the response struct for api ListDevopsProjectTasks
type ListDevopsProjectTasksResponse struct {
	*responses.BaseResponse
	Successful bool   `json:"Successful" xml:"Successful"`
	ErrorCode  string `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMsg   string `json:"ErrorMsg" xml:"ErrorMsg"`
	RequestId  string `json:"RequestId" xml:"RequestId"`
	Object     []Task `json:"Object" xml:"Object"`
}

// CreateListDevopsProjectTasksRequest creates a request to invoke ListDevopsProjectTasks API
func CreateListDevopsProjectTasksRequest() (request *ListDevopsProjectTasksRequest) {
	request = &ListDevopsProjectTasksRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("devops-rdc", "2020-03-03", "ListDevopsProjectTasks", "1.9.6", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListDevopsProjectTasksResponse creates a response to parse from ListDevopsProjectTasks response
func CreateListDevopsProjectTasksResponse() (response *ListDevopsProjectTasksResponse) {
	response = &ListDevopsProjectTasksResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
