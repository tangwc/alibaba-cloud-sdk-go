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

func (client *Client) GetMainDomainName(request *GetMainDomainNameRequest) (response *GetMainDomainNameResponse, err error) {
	response = CreateGetMainDomainNameResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) GetMainDomainNameWithChan(request *GetMainDomainNameRequest) (<-chan *GetMainDomainNameResponse, <-chan error) {
	responseChan := make(chan *GetMainDomainNameResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetMainDomainName(request)
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

func (client *Client) GetMainDomainNameWithCallback(request *GetMainDomainNameRequest, callback func(response *GetMainDomainNameResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetMainDomainNameResponse
		var err error
		defer close(result)
		response, err = client.GetMainDomainName(request)
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

type GetMainDomainNameRequest struct {
	*requests.RpcRequest
	UserClientIp string `position:"Query" name:"UserClientIp"`
	InputString  string `position:"Query" name:"InputString"`
	Lang         string `position:"Query" name:"Lang"`
}

type GetMainDomainNameResponse struct {
	*responses.BaseResponse
	RequestId   string `json:"RequestId" xml:"RequestId"`
	DomainName  string `json:"DomainName" xml:"DomainName"`
	RR          string `json:"RR" xml:"RR"`
	DomainLevel int    `json:"DomainLevel" xml:"DomainLevel"`
}

func CreateGetMainDomainNameRequest() (request *GetMainDomainNameRequest) {
	request = &GetMainDomainNameRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Alidns", "2015-01-09", "GetMainDomainName", "", "")
	request.Method = requests.POST
	return
}

func CreateGetMainDomainNameResponse() (response *GetMainDomainNameResponse) {
	response = &GetMainDomainNameResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
