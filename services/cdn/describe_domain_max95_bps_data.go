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

// DescribeDomainMax95BpsData invokes the cdn.DescribeDomainMax95BpsData API synchronously
func (client *Client) DescribeDomainMax95BpsData(request *DescribeDomainMax95BpsDataRequest) (response *DescribeDomainMax95BpsDataResponse, err error) {
	response = CreateDescribeDomainMax95BpsDataResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDomainMax95BpsDataWithChan invokes the cdn.DescribeDomainMax95BpsData API asynchronously
func (client *Client) DescribeDomainMax95BpsDataWithChan(request *DescribeDomainMax95BpsDataRequest) (<-chan *DescribeDomainMax95BpsDataResponse, <-chan error) {
	responseChan := make(chan *DescribeDomainMax95BpsDataResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDomainMax95BpsData(request)
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

// DescribeDomainMax95BpsDataWithCallback invokes the cdn.DescribeDomainMax95BpsData API asynchronously
func (client *Client) DescribeDomainMax95BpsDataWithCallback(request *DescribeDomainMax95BpsDataRequest, callback func(response *DescribeDomainMax95BpsDataResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDomainMax95BpsDataResponse
		var err error
		defer close(result)
		response, err = client.DescribeDomainMax95BpsData(request)
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

// DescribeDomainMax95BpsDataRequest is the request struct for api DescribeDomainMax95BpsData
type DescribeDomainMax95BpsDataRequest struct {
	*requests.RpcRequest
	DomainName string `position:"Query" name:"DomainName"`
	EndTime    string `position:"Query" name:"EndTime"`
	StartTime  string `position:"Query" name:"StartTime"`
	Cycle      string `position:"Query" name:"Cycle"`
	TimePoint  string `position:"Query" name:"TimePoint"`
}

// DescribeDomainMax95BpsDataResponse is the response struct for api DescribeDomainMax95BpsData
type DescribeDomainMax95BpsDataResponse struct {
	*responses.BaseResponse
	EndTime          string `json:"EndTime" xml:"EndTime"`
	StartTime        string `json:"StartTime" xml:"StartTime"`
	RequestId        string `json:"RequestId" xml:"RequestId"`
	DomesticMax95Bps string `json:"DomesticMax95Bps" xml:"DomesticMax95Bps"`
	OverseasMax95Bps string `json:"OverseasMax95Bps" xml:"OverseasMax95Bps"`
	Max95Bps         string `json:"Max95Bps" xml:"Max95Bps"`
	DomainName       string `json:"DomainName" xml:"DomainName"`
}

// CreateDescribeDomainMax95BpsDataRequest creates a request to invoke DescribeDomainMax95BpsData API
func CreateDescribeDomainMax95BpsDataRequest() (request *DescribeDomainMax95BpsDataRequest) {
	request = &DescribeDomainMax95BpsDataRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cdn", "2018-05-10", "DescribeDomainMax95BpsData", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeDomainMax95BpsDataResponse creates a response to parse from DescribeDomainMax95BpsData response
func CreateDescribeDomainMax95BpsDataResponse() (response *DescribeDomainMax95BpsDataResponse) {
	response = &DescribeDomainMax95BpsDataResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
