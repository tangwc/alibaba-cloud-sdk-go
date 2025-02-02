package dms_enterprise

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

// GetOpLog invokes the dms_enterprise.GetOpLog API synchronously
func (client *Client) GetOpLog(request *GetOpLogRequest) (response *GetOpLogResponse, err error) {
	response = CreateGetOpLogResponse()
	err = client.DoAction(request, response)
	return
}

// GetOpLogWithChan invokes the dms_enterprise.GetOpLog API asynchronously
func (client *Client) GetOpLogWithChan(request *GetOpLogRequest) (<-chan *GetOpLogResponse, <-chan error) {
	responseChan := make(chan *GetOpLogResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetOpLog(request)
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

// GetOpLogWithCallback invokes the dms_enterprise.GetOpLog API asynchronously
func (client *Client) GetOpLogWithCallback(request *GetOpLogRequest, callback func(response *GetOpLogResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetOpLogResponse
		var err error
		defer close(result)
		response, err = client.GetOpLog(request)
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

// GetOpLogRequest is the request struct for api GetOpLog
type GetOpLogRequest struct {
	*requests.RpcRequest
	StartTime  string           `position:"Query" name:"StartTime"`
	Tid        requests.Integer `position:"Query" name:"Tid"`
	PageNumber requests.Integer `position:"Query" name:"PageNumber"`
	PageSize   requests.Integer `position:"Query" name:"PageSize"`
	Module     string           `position:"Query" name:"Module"`
	EndTime    string           `position:"Query" name:"EndTime"`
}

// GetOpLogResponse is the response struct for api GetOpLog
type GetOpLogResponse struct {
	*responses.BaseResponse
	TotalCount   int64        `json:"TotalCount" xml:"TotalCount"`
	RequestId    string       `json:"RequestId" xml:"RequestId"`
	ErrorCode    string       `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMessage string       `json:"ErrorMessage" xml:"ErrorMessage"`
	Success      bool         `json:"Success" xml:"Success"`
	OpLogDetails OpLogDetails `json:"OpLogDetails" xml:"OpLogDetails"`
}

// CreateGetOpLogRequest creates a request to invoke GetOpLog API
func CreateGetOpLogRequest() (request *GetOpLogRequest) {
	request = &GetOpLogRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dms-enterprise", "2018-11-01", "GetOpLog", "dms-enterprise", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetOpLogResponse creates a response to parse from GetOpLog response
func CreateGetOpLogResponse() (response *GetOpLogResponse) {
	response = &GetOpLogResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
