package dytnsapi

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

// TwoElementsVerification invokes the dytnsapi.TwoElementsVerification API synchronously
func (client *Client) TwoElementsVerification(request *TwoElementsVerificationRequest) (response *TwoElementsVerificationResponse, err error) {
	response = CreateTwoElementsVerificationResponse()
	err = client.DoAction(request, response)
	return
}

// TwoElementsVerificationWithChan invokes the dytnsapi.TwoElementsVerification API asynchronously
func (client *Client) TwoElementsVerificationWithChan(request *TwoElementsVerificationRequest) (<-chan *TwoElementsVerificationResponse, <-chan error) {
	responseChan := make(chan *TwoElementsVerificationResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.TwoElementsVerification(request)
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

// TwoElementsVerificationWithCallback invokes the dytnsapi.TwoElementsVerification API asynchronously
func (client *Client) TwoElementsVerificationWithCallback(request *TwoElementsVerificationRequest, callback func(response *TwoElementsVerificationResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *TwoElementsVerificationResponse
		var err error
		defer close(result)
		response, err = client.TwoElementsVerification(request)
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

// TwoElementsVerificationRequest is the request struct for api TwoElementsVerification
type TwoElementsVerificationRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	RouteName            string           `position:"Query" name:"RouteName"`
	Mask                 string           `position:"Query" name:"Mask"`
	ResultCount          string           `position:"Query" name:"ResultCount"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	NumberType           string           `position:"Query" name:"NumberType"`
	AuthCode             string           `position:"Query" name:"AuthCode"`
	InputNumber          string           `position:"Query" name:"InputNumber"`
	Name                 string           `position:"Query" name:"Name"`
	FlowName             string           `position:"Query" name:"FlowName"`
}

// TwoElementsVerificationResponse is the response struct for api TwoElementsVerification
type TwoElementsVerificationResponse struct {
	*responses.BaseResponse
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateTwoElementsVerificationRequest creates a request to invoke TwoElementsVerification API
func CreateTwoElementsVerificationRequest() (request *TwoElementsVerificationRequest) {
	request = &TwoElementsVerificationRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dytnsapi", "2020-02-17", "TwoElementsVerification", "", "")
	request.Method = requests.POST
	return
}

// CreateTwoElementsVerificationResponse creates a response to parse from TwoElementsVerification response
func CreateTwoElementsVerificationResponse() (response *TwoElementsVerificationResponse) {
	response = &TwoElementsVerificationResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
