package cs

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

func (client *Client) DescribeAgilityTunnelCerts(request *DescribeAgilityTunnelCertsRequest) (response *DescribeAgilityTunnelCertsResponse, err error) {
	response = CreateDescribeAgilityTunnelCertsResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) DescribeAgilityTunnelCertsWithChan(request *DescribeAgilityTunnelCertsRequest) (<-chan *DescribeAgilityTunnelCertsResponse, <-chan error) {
	responseChan := make(chan *DescribeAgilityTunnelCertsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeAgilityTunnelCerts(request)
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

func (client *Client) DescribeAgilityTunnelCertsWithCallback(request *DescribeAgilityTunnelCertsRequest, callback func(response *DescribeAgilityTunnelCertsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeAgilityTunnelCertsResponse
		var err error
		defer close(result)
		response, err = client.DescribeAgilityTunnelCerts(request)
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

type DescribeAgilityTunnelCertsRequest struct {
	*requests.RoaRequest
	Token string `position:"Path" name:"Token"`
}

type DescribeAgilityTunnelCertsResponse struct {
	*responses.BaseResponse
}

func CreateDescribeAgilityTunnelCertsRequest() (request *DescribeAgilityTunnelCertsRequest) {
	request = &DescribeAgilityTunnelCertsRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("CS", "2015-12-15", "DescribeAgilityTunnelCerts", "/agility/[Token]/agent_certs", "", "")
	request.Method = requests.GET
	return
}

func CreateDescribeAgilityTunnelCertsResponse() (response *DescribeAgilityTunnelCertsResponse) {
	response = &DescribeAgilityTunnelCertsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
