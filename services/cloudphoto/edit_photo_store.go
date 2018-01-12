package cloudphoto

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

func (client *Client) EditPhotoStore(request *EditPhotoStoreRequest) (response *EditPhotoStoreResponse, err error) {
	response = CreateEditPhotoStoreResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) EditPhotoStoreWithChan(request *EditPhotoStoreRequest) (<-chan *EditPhotoStoreResponse, <-chan error) {
	responseChan := make(chan *EditPhotoStoreResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.EditPhotoStore(request)
		responseChan <- response
		errChan <- err
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

func (client *Client) EditPhotoStoreWithCallback(request *EditPhotoStoreRequest, callback func(response *EditPhotoStoreResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *EditPhotoStoreResponse
		var err error
		defer close(result)
		response, err = client.EditPhotoStore(request)
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

type EditPhotoStoreRequest struct {
	*requests.RpcRequest
	DefaultQuota      requests.Integer `position:"Query" name:"DefaultQuota"`
	Remark            string           `position:"Query" name:"Remark"`
	AutoCleanDays     requests.Integer `position:"Query" name:"AutoCleanDays"`
	StoreName         string           `position:"Query" name:"StoreName"`
	DefaultTrashQuota requests.Integer `position:"Query" name:"DefaultTrashQuota"`
	AutoCleanEnabled  string           `position:"Query" name:"AutoCleanEnabled"`
}

type EditPhotoStoreResponse struct {
	*responses.BaseResponse
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Action    string `json:"Action" xml:"Action"`
}

func CreateEditPhotoStoreRequest() (request *EditPhotoStoreRequest) {
	request = &EditPhotoStoreRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudPhoto", "2017-07-11", "EditPhotoStore", "cloudphoto", "openAPI")
	request.Method = requests.POST
	return
}

func CreateEditPhotoStoreResponse() (response *EditPhotoStoreResponse) {
	response = &EditPhotoStoreResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
