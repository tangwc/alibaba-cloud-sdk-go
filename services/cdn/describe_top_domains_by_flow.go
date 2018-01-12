package cdn

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

func (client *Client) DescribeTopDomainsByFlow(request *DescribeTopDomainsByFlowRequest) (response *DescribeTopDomainsByFlowResponse, err error) {
	response = CreateDescribeTopDomainsByFlowResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) DescribeTopDomainsByFlowWithChan(request *DescribeTopDomainsByFlowRequest) (<-chan *DescribeTopDomainsByFlowResponse, <-chan error) {
	responseChan := make(chan *DescribeTopDomainsByFlowResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeTopDomainsByFlow(request)
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

func (client *Client) DescribeTopDomainsByFlowWithCallback(request *DescribeTopDomainsByFlowRequest, callback func(response *DescribeTopDomainsByFlowResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeTopDomainsByFlowResponse
		var err error
		defer close(result)
		response, err = client.DescribeTopDomainsByFlow(request)
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

type DescribeTopDomainsByFlowRequest struct {
	*requests.RpcRequest
	Limit         requests.Integer `position:"Query" name:"Limit"`
	EndTime       string           `position:"Query" name:"EndTime"`
	StartTime     string           `position:"Query" name:"StartTime"`
	OwnerId       requests.Integer `position:"Query" name:"OwnerId"`
	SecurityToken string           `position:"Query" name:"SecurityToken"`
}

type DescribeTopDomainsByFlowResponse struct {
	*responses.BaseResponse
	RequestId         string `json:"RequestId" xml:"RequestId"`
	StartTime         string `json:"StartTime" xml:"StartTime"`
	EndTime           string `json:"EndTime" xml:"EndTime"`
	DomainCount       int    `json:"DomainCount" xml:"DomainCount"`
	DomainOnlineCount int    `json:"DomainOnlineCount" xml:"DomainOnlineCount"`
	TopDomains        struct {
		TopDomain []struct {
			DomainName     string `json:"DomainName" xml:"DomainName"`
			Rank           int    `json:"Rank" xml:"Rank"`
			TotalTraffic   string `json:"TotalTraffic" xml:"TotalTraffic"`
			TrafficPercent string `json:"TrafficPercent" xml:"TrafficPercent"`
			MaxBps         int    `json:"MaxBps" xml:"MaxBps"`
			MaxBpsTime     string `json:"MaxBpsTime" xml:"MaxBpsTime"`
			TotalAccess    int    `json:"TotalAccess" xml:"TotalAccess"`
		} `json:"TopDomain" xml:"TopDomain"`
	} `json:"TopDomains" xml:"TopDomains"`
}

func CreateDescribeTopDomainsByFlowRequest() (request *DescribeTopDomainsByFlowRequest) {
	request = &DescribeTopDomainsByFlowRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cdn", "2014-11-11", "DescribeTopDomainsByFlow", "", "")
	request.Method = requests.POST
	return
}

func CreateDescribeTopDomainsByFlowResponse() (response *DescribeTopDomainsByFlowResponse) {
	response = &DescribeTopDomainsByFlowResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
