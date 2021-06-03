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

// GetPipelineInstanceInfo invokes the devops_rdc.GetPipelineInstanceInfo API synchronously
func (client *Client) GetPipelineInstanceInfo(request *GetPipelineInstanceInfoRequest) (response *GetPipelineInstanceInfoResponse, err error) {
	response = CreateGetPipelineInstanceInfoResponse()
	err = client.DoAction(request, response)
	return
}

// GetPipelineInstanceInfoWithChan invokes the devops_rdc.GetPipelineInstanceInfo API asynchronously
func (client *Client) GetPipelineInstanceInfoWithChan(request *GetPipelineInstanceInfoRequest) (<-chan *GetPipelineInstanceInfoResponse, <-chan error) {
	responseChan := make(chan *GetPipelineInstanceInfoResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetPipelineInstanceInfo(request)
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

// GetPipelineInstanceInfoWithCallback invokes the devops_rdc.GetPipelineInstanceInfo API asynchronously
func (client *Client) GetPipelineInstanceInfoWithCallback(request *GetPipelineInstanceInfoRequest, callback func(response *GetPipelineInstanceInfoResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetPipelineInstanceInfoResponse
		var err error
		defer close(result)
		response, err = client.GetPipelineInstanceInfo(request)
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

// GetPipelineInstanceInfoRequest is the request struct for api GetPipelineInstanceInfo
type GetPipelineInstanceInfoRequest struct {
	*requests.RpcRequest
	FlowInstanceId string           `position:"Body" name:"FlowInstanceId"`
	UserPk         string           `position:"Body" name:"UserPk"`
	OrgId          string           `position:"Body" name:"OrgId"`
	PipelineId     requests.Integer `position:"Body" name:"PipelineId"`
}

// GetPipelineInstanceInfoResponse is the response struct for api GetPipelineInstanceInfo
type GetPipelineInstanceInfoResponse struct {
	*responses.BaseResponse
	ErrorCode    string `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
	RequestId    string `json:"RequestId" xml:"RequestId"`
	Success      bool   `json:"Success" xml:"Success"`
	Object       Object `json:"Object" xml:"Object"`
}

// CreateGetPipelineInstanceInfoRequest creates a request to invoke GetPipelineInstanceInfo API
func CreateGetPipelineInstanceInfoRequest() (request *GetPipelineInstanceInfoRequest) {
	request = &GetPipelineInstanceInfoRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("devops-rdc", "2020-03-03", "GetPipelineInstanceInfo", "1.9.6", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetPipelineInstanceInfoResponse creates a response to parse from GetPipelineInstanceInfo response
func CreateGetPipelineInstanceInfoResponse() (response *GetPipelineInstanceInfoResponse) {
	response = &GetPipelineInstanceInfoResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
