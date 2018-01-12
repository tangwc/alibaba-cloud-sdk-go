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

func (client *Client) DescribeExtensiveDomainData(request *DescribeExtensiveDomainDataRequest) (response *DescribeExtensiveDomainDataResponse, err error) {
	response = CreateDescribeExtensiveDomainDataResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) DescribeExtensiveDomainDataWithChan(request *DescribeExtensiveDomainDataRequest) (<-chan *DescribeExtensiveDomainDataResponse, <-chan error) {
	responseChan := make(chan *DescribeExtensiveDomainDataResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeExtensiveDomainData(request)
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

func (client *Client) DescribeExtensiveDomainDataWithCallback(request *DescribeExtensiveDomainDataRequest, callback func(response *DescribeExtensiveDomainDataResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeExtensiveDomainDataResponse
		var err error
		defer close(result)
		response, err = client.DescribeExtensiveDomainData(request)
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

type DescribeExtensiveDomainDataRequest struct {
	*requests.RpcRequest
	EndTime         string           `position:"Query" name:"EndTime"`
	PageSize        requests.Integer `position:"Query" name:"PageSize"`
	StartTime       string           `position:"Query" name:"StartTime"`
	PageNumber      requests.Integer `position:"Query" name:"PageNumber"`
	ExtensiveDomain string           `position:"Query" name:"ExtensiveDomain"`
	OwnerId         requests.Integer `position:"Query" name:"OwnerId"`
	SecurityToken   string           `position:"Query" name:"SecurityToken"`
}

type DescribeExtensiveDomainDataResponse struct {
	*responses.BaseResponse
	RequestId       string `json:"RequestId" xml:"RequestId"`
	ExtensiveDomain string `json:"ExtensiveDomain" xml:"ExtensiveDomain"`
	DataInterval    string `json:"DataInterval" xml:"DataInterval"`
	StartTime       string `json:"StartTime" xml:"StartTime"`
	EndTime         string `json:"EndTime" xml:"EndTime"`
	PageNumber      string `json:"PageNumber" xml:"PageNumber"`
	TotalCount      string `json:"TotalCount" xml:"TotalCount"`
	PageSize        string `json:"PageSize" xml:"PageSize"`
	DataPerInterval struct {
		UsageData []struct {
			ExactDomain string `json:"ExactDomain" xml:"ExactDomain"`
			TimeStamp   string `json:"TimeStamp" xml:"TimeStamp"`
			Acc         string `json:"Acc" xml:"Acc"`
			Flow        string `json:"Flow" xml:"Flow"`
		} `json:"UsageData" xml:"UsageData"`
	} `json:"DataPerInterval" xml:"DataPerInterval"`
}

func CreateDescribeExtensiveDomainDataRequest() (request *DescribeExtensiveDomainDataRequest) {
	request = &DescribeExtensiveDomainDataRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cdn", "2014-11-11", "DescribeExtensiveDomainData", "", "")
	request.Method = requests.POST
	return
}

func CreateDescribeExtensiveDomainDataResponse() (response *DescribeExtensiveDomainDataResponse) {
	response = &DescribeExtensiveDomainDataResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
