package dbs

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

// DescribeRestoreTaskList invokes the dbs.DescribeRestoreTaskList API synchronously
func (client *Client) DescribeRestoreTaskList(request *DescribeRestoreTaskListRequest) (response *DescribeRestoreTaskListResponse, err error) {
	response = CreateDescribeRestoreTaskListResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeRestoreTaskListWithChan invokes the dbs.DescribeRestoreTaskList API asynchronously
func (client *Client) DescribeRestoreTaskListWithChan(request *DescribeRestoreTaskListRequest) (<-chan *DescribeRestoreTaskListResponse, <-chan error) {
	responseChan := make(chan *DescribeRestoreTaskListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeRestoreTaskList(request)
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

// DescribeRestoreTaskListWithCallback invokes the dbs.DescribeRestoreTaskList API asynchronously
func (client *Client) DescribeRestoreTaskListWithCallback(request *DescribeRestoreTaskListRequest, callback func(response *DescribeRestoreTaskListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeRestoreTaskListResponse
		var err error
		defer close(result)
		response, err = client.DescribeRestoreTaskList(request)
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

// DescribeRestoreTaskListRequest is the request struct for api DescribeRestoreTaskList
type DescribeRestoreTaskListRequest struct {
	*requests.RpcRequest
	ClientToken    string           `position:"Query" name:"ClientToken"`
	BackupPlanId   string           `position:"Query" name:"BackupPlanId"`
	PageNum        requests.Integer `position:"Query" name:"PageNum"`
	OwnerId        string           `position:"Query" name:"OwnerId"`
	StartTimestamp requests.Integer `position:"Query" name:"StartTimestamp"`
	EndTimestamp   requests.Integer `position:"Query" name:"EndTimestamp"`
	RestoreTaskId  string           `position:"Query" name:"RestoreTaskId"`
	PageSize       requests.Integer `position:"Query" name:"PageSize"`
}

// DescribeRestoreTaskListResponse is the response struct for api DescribeRestoreTaskList
type DescribeRestoreTaskListResponse struct {
	*responses.BaseResponse
	HttpStatusCode int                            `json:"HttpStatusCode" xml:"HttpStatusCode"`
	PageNum        int                            `json:"PageNum" xml:"PageNum"`
	RequestId      string                         `json:"RequestId" xml:"RequestId"`
	ErrCode        string                         `json:"ErrCode" xml:"ErrCode"`
	Success        bool                           `json:"Success" xml:"Success"`
	ErrMessage     string                         `json:"ErrMessage" xml:"ErrMessage"`
	TotalPages     int                            `json:"TotalPages" xml:"TotalPages"`
	TotalElements  int                            `json:"TotalElements" xml:"TotalElements"`
	PageSize       int                            `json:"PageSize" xml:"PageSize"`
	Items          ItemsInDescribeRestoreTaskList `json:"Items" xml:"Items"`
}

// CreateDescribeRestoreTaskListRequest creates a request to invoke DescribeRestoreTaskList API
func CreateDescribeRestoreTaskListRequest() (request *DescribeRestoreTaskListRequest) {
	request = &DescribeRestoreTaskListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dbs", "2019-03-06", "DescribeRestoreTaskList", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeRestoreTaskListResponse creates a response to parse from DescribeRestoreTaskList response
func CreateDescribeRestoreTaskListResponse() (response *DescribeRestoreTaskListResponse) {
	response = &DescribeRestoreTaskListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
