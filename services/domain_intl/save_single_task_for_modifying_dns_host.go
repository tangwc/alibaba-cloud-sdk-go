package domain_intl

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

func (client *Client) SaveSingleTaskForModifyingDnsHost(request *SaveSingleTaskForModifyingDnsHostRequest) (response *SaveSingleTaskForModifyingDnsHostResponse, err error) {
	response = CreateSaveSingleTaskForModifyingDnsHostResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) SaveSingleTaskForModifyingDnsHostWithChan(request *SaveSingleTaskForModifyingDnsHostRequest) (<-chan *SaveSingleTaskForModifyingDnsHostResponse, <-chan error) {
	responseChan := make(chan *SaveSingleTaskForModifyingDnsHostResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SaveSingleTaskForModifyingDnsHost(request)
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

func (client *Client) SaveSingleTaskForModifyingDnsHostWithCallback(request *SaveSingleTaskForModifyingDnsHostRequest, callback func(response *SaveSingleTaskForModifyingDnsHostResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SaveSingleTaskForModifyingDnsHostResponse
		var err error
		defer close(result)
		response, err = client.SaveSingleTaskForModifyingDnsHost(request)
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

type SaveSingleTaskForModifyingDnsHostRequest struct {
	*requests.RpcRequest
	Lang       string    `position:"Query" name:"Lang"`
	InstanceId string    `position:"Query" name:"InstanceId"`
	Ip         *[]string `position:"Query" name:"Ip"  type:"Repeated"`
	DnsName    string    `position:"Query" name:"DnsName"`
}

type SaveSingleTaskForModifyingDnsHostResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	TaskNo    string `json:"TaskNo" xml:"TaskNo"`
}

func CreateSaveSingleTaskForModifyingDnsHostRequest() (request *SaveSingleTaskForModifyingDnsHostRequest) {
	request = &SaveSingleTaskForModifyingDnsHostRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Domain-intl", "2017-12-18", "SaveSingleTaskForModifyingDnsHost", "", "")
	request.Method = requests.POST
	return
}

func CreateSaveSingleTaskForModifyingDnsHostResponse() (response *SaveSingleTaskForModifyingDnsHostResponse) {
	response = &SaveSingleTaskForModifyingDnsHostResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
