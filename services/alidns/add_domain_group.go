package alidns

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

func (client *Client) AddDomainGroup(request *AddDomainGroupRequest) (response *AddDomainGroupResponse, err error) {
	response = CreateAddDomainGroupResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) AddDomainGroupWithChan(request *AddDomainGroupRequest) (<-chan *AddDomainGroupResponse, <-chan error) {
	responseChan := make(chan *AddDomainGroupResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AddDomainGroup(request)
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

func (client *Client) AddDomainGroupWithCallback(request *AddDomainGroupRequest, callback func(response *AddDomainGroupResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AddDomainGroupResponse
		var err error
		defer close(result)
		response, err = client.AddDomainGroup(request)
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

type AddDomainGroupRequest struct {
	*requests.RpcRequest
	GroupName    string `position:"Query" name:"GroupName"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
	Lang         string `position:"Query" name:"Lang"`
}

type AddDomainGroupResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	GroupId   string `json:"GroupId" xml:"GroupId"`
	GroupName string `json:"GroupName" xml:"GroupName"`
}

func CreateAddDomainGroupRequest() (request *AddDomainGroupRequest) {
	request = &AddDomainGroupRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Alidns", "2015-01-09", "AddDomainGroup", "", "")
	request.Method = requests.POST
	return
}

func CreateAddDomainGroupResponse() (response *AddDomainGroupResponse) {
	response = &AddDomainGroupResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
