package vcs

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

// CreateVideoComposeTask invokes the vcs.CreateVideoComposeTask API synchronously
// api document: https://help.aliyun.com/api/vcs/createvideocomposetask.html
func (client *Client) CreateVideoComposeTask(request *CreateVideoComposeTaskRequest) (response *CreateVideoComposeTaskResponse, err error) {
	response = CreateCreateVideoComposeTaskResponse()
	err = client.DoAction(request, response)
	return
}

// CreateVideoComposeTaskWithChan invokes the vcs.CreateVideoComposeTask API asynchronously
// api document: https://help.aliyun.com/api/vcs/createvideocomposetask.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateVideoComposeTaskWithChan(request *CreateVideoComposeTaskRequest) (<-chan *CreateVideoComposeTaskResponse, <-chan error) {
	responseChan := make(chan *CreateVideoComposeTaskResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateVideoComposeTask(request)
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

// CreateVideoComposeTaskWithCallback invokes the vcs.CreateVideoComposeTask API asynchronously
// api document: https://help.aliyun.com/api/vcs/createvideocomposetask.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateVideoComposeTaskWithCallback(request *CreateVideoComposeTaskRequest, callback func(response *CreateVideoComposeTaskResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateVideoComposeTaskResponse
		var err error
		defer close(result)
		response, err = client.CreateVideoComposeTask(request)
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

// CreateVideoComposeTaskRequest is the request struct for api CreateVideoComposeTask
type CreateVideoComposeTaskRequest struct {
	*requests.RpcRequest
	CorpId          string `position:"Body" name:"CorpId"`
	DomainName      string `position:"Body" name:"DomainName"`
	PicUrlList      string `position:"Body" name:"PicUrlList"`
	AudioUrl        string `position:"Body" name:"AudioUrl"`
	BucketName      string `position:"Body" name:"BucketName"`
	ImageParameters string `position:"Body" name:"ImageParameters"`
	VideoRate       string `position:"Body" name:"VideoRate"`
	VideoFormat     string `position:"Body" name:"VideoFormat"`
}

// CreateVideoComposeTaskResponse is the response struct for api CreateVideoComposeTask
type CreateVideoComposeTaskResponse struct {
	*responses.BaseResponse
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateCreateVideoComposeTaskRequest creates a request to invoke CreateVideoComposeTask API
func CreateCreateVideoComposeTaskRequest() (request *CreateVideoComposeTaskRequest) {
	request = &CreateVideoComposeTaskRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vcs", "2020-05-15", "CreateVideoComposeTask", "vcs", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateVideoComposeTaskResponse creates a response to parse from CreateVideoComposeTask response
func CreateCreateVideoComposeTaskResponse() (response *CreateVideoComposeTaskResponse) {
	response = &CreateVideoComposeTaskResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
