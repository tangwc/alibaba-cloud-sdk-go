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

func (client *Client) DescribeDomainFlowData(request *DescribeDomainFlowDataRequest) (response *DescribeDomainFlowDataResponse, err error) {
	response = CreateDescribeDomainFlowDataResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) DescribeDomainFlowDataWithChan(request *DescribeDomainFlowDataRequest) (<-chan *DescribeDomainFlowDataResponse, <-chan error) {
	responseChan := make(chan *DescribeDomainFlowDataResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDomainFlowData(request)
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

func (client *Client) DescribeDomainFlowDataWithCallback(request *DescribeDomainFlowDataRequest, callback func(response *DescribeDomainFlowDataResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDomainFlowDataResponse
		var err error
		defer close(result)
		response, err = client.DescribeDomainFlowData(request)
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

type DescribeDomainFlowDataRequest struct {
	*requests.RpcRequest
	EndTime        string           `position:"Query" name:"EndTime"`
	StartTime      string           `position:"Query" name:"StartTime"`
	Interval       string           `position:"Query" name:"Interval"`
	FixTimeGap     string           `position:"Query" name:"FixTimeGap"`
	LocationNameEn string           `position:"Query" name:"LocationNameEn"`
	IspNameEn      string           `position:"Query" name:"IspNameEn"`
	OwnerId        requests.Integer `position:"Query" name:"OwnerId"`
	DomainType     string           `position:"Query" name:"DomainType"`
	DomainName     string           `position:"Query" name:"DomainName"`
	SecurityToken  string           `position:"Query" name:"SecurityToken"`
	TimeMerge      string           `position:"Query" name:"TimeMerge"`
}

type DescribeDomainFlowDataResponse struct {
	*responses.BaseResponse
	RequestId           string `json:"RequestId" xml:"RequestId"`
	DomainName          string `json:"DomainName" xml:"DomainName"`
	DataInterval        string `json:"DataInterval" xml:"DataInterval"`
	StartTime           string `json:"StartTime" xml:"StartTime"`
	EndTime             string `json:"EndTime" xml:"EndTime"`
	FlowDataPerInterval struct {
		DataModule []struct {
			TimeStamp            string `json:"TimeStamp" xml:"TimeStamp"`
			Value                string `json:"Value" xml:"Value"`
			DomesticValue        string `json:"DomesticValue" xml:"DomesticValue"`
			OverseasValue        string `json:"OverseasValue" xml:"OverseasValue"`
			DynamicValue         string `json:"DynamicValue" xml:"DynamicValue"`
			DynamicDomesticValue string `json:"DynamicDomesticValue" xml:"DynamicDomesticValue"`
			DynamicOverseasValue string `json:"DynamicOverseasValue" xml:"DynamicOverseasValue"`
			StaticValue          string `json:"StaticValue" xml:"StaticValue"`
			StaticDomesticValue  string `json:"StaticDomesticValue" xml:"StaticDomesticValue"`
			StaticOverseasValue  string `json:"StaticOverseasValue" xml:"StaticOverseasValue"`
		} `json:"DataModule" xml:"DataModule"`
	} `json:"FlowDataPerInterval" xml:"FlowDataPerInterval"`
}

func CreateDescribeDomainFlowDataRequest() (request *DescribeDomainFlowDataRequest) {
	request = &DescribeDomainFlowDataRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cdn", "2014-11-11", "DescribeDomainFlowData", "", "")
	request.Method = requests.POST
	return
}

func CreateDescribeDomainFlowDataResponse() (response *DescribeDomainFlowDataResponse) {
	response = &DescribeDomainFlowDataResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
