package dm

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

func (client *Client) CreateSign(request *CreateSignRequest) (response *CreateSignResponse, err error) {
	response = CreateCreateSignResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) CreateSignWithChan(request *CreateSignRequest) (<-chan *CreateSignResponse, <-chan error) {
	responseChan := make(chan *CreateSignResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateSign(request)
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

func (client *Client) CreateSignWithCallback(request *CreateSignRequest, callback func(response *CreateSignResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateSignResponse
		var err error
		defer close(result)
		response, err = client.CreateSign(request)
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

type CreateSignRequest struct {
	*requests.RpcRequest
	SignName             string           `position:"Query" name:"SignName"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	Remark               string           `position:"Query" name:"Remark"`
	FileNames            string           `position:"Query" name:"FileNames"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	FromType             requests.Integer `position:"Query" name:"FromType"`
	SignType             requests.Integer `position:"Query" name:"SignType"`
}

type CreateSignResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

func CreateCreateSignRequest() (request *CreateSignRequest) {
	request = &CreateSignRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dm", "2015-11-23", "CreateSign", "", "")
	request.Method = requests.POST
	return
}

func CreateCreateSignResponse() (response *CreateSignResponse) {
	response = &CreateSignResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
