package mse

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

// ListNamingTrack invokes the mse.ListNamingTrack API synchronously
func (client *Client) ListNamingTrack(request *ListNamingTrackRequest) (response *ListNamingTrackResponse, err error) {
	response = CreateListNamingTrackResponse()
	err = client.DoAction(request, response)
	return
}

// ListNamingTrackWithChan invokes the mse.ListNamingTrack API asynchronously
func (client *Client) ListNamingTrackWithChan(request *ListNamingTrackRequest) (<-chan *ListNamingTrackResponse, <-chan error) {
	responseChan := make(chan *ListNamingTrackResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListNamingTrack(request)
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

// ListNamingTrackWithCallback invokes the mse.ListNamingTrack API asynchronously
func (client *Client) ListNamingTrackWithCallback(request *ListNamingTrackRequest, callback func(response *ListNamingTrackResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListNamingTrackResponse
		var err error
		defer close(result)
		response, err = client.ListNamingTrack(request)
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

// ListNamingTrackRequest is the request struct for api ListNamingTrack
type ListNamingTrackRequest struct {
	*requests.RpcRequest
	MseSessionId   string           `position:"Query" name:"MseSessionId"`
	StartTs        requests.Integer `position:"Query" name:"StartTs"`
	PageNum        requests.Integer `position:"Query" name:"PageNum"`
	NamespaceId    string           `position:"Query" name:"NamespaceId"`
	RequestPars    string           `position:"Query" name:"RequestPars"`
	EndTs          requests.Integer `position:"Query" name:"EndTs"`
	PageSize       requests.Integer `position:"Query" name:"PageSize"`
	ServiceName    string           `position:"Query" name:"ServiceName"`
	Group          string           `position:"Query" name:"Group"`
	Ip             string           `position:"Query" name:"Ip"`
	Reverse        requests.Boolean `position:"Query" name:"Reverse"`
	InstanceId     string           `position:"Query" name:"InstanceId"`
	AcceptLanguage string           `position:"Query" name:"AcceptLanguage"`
}

// ListNamingTrackResponse is the response struct for api ListNamingTrack
type ListNamingTrackResponse struct {
	*responses.BaseResponse
	HttpCode   string  `json:"HttpCode" xml:"HttpCode"`
	PageSize   int64   `json:"PageSize" xml:"PageSize"`
	PageNumber int64   `json:"PageNumber" xml:"PageNumber"`
	TotalCount int64   `json:"TotalCount" xml:"TotalCount"`
	RequestId  string  `json:"RequestId" xml:"RequestId"`
	Message    string  `json:"Message" xml:"Message"`
	ErrorCode  string  `json:"ErrorCode" xml:"ErrorCode"`
	Success    bool    `json:"Success" xml:"Success"`
	Traces     []Trace `json:"Traces" xml:"Traces"`
}

// CreateListNamingTrackRequest creates a request to invoke ListNamingTrack API
func CreateListNamingTrackRequest() (request *ListNamingTrackRequest) {
	request = &ListNamingTrackRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("mse", "2019-05-31", "ListNamingTrack", "mse", "openAPI")
	request.Method = requests.GET
	return
}

// CreateListNamingTrackResponse creates a response to parse from ListNamingTrack response
func CreateListNamingTrackResponse() (response *ListNamingTrackResponse) {
	response = &ListNamingTrackResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
