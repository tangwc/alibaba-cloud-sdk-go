package arms

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

func (client *Client) WhereInDimQuery(request *WhereInDimQueryRequest) (response *WhereInDimQueryResponse, err error) {
	response = CreateWhereInDimQueryResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) WhereInDimQueryWithChan(request *WhereInDimQueryRequest) (<-chan *WhereInDimQueryResponse, <-chan error) {
	responseChan := make(chan *WhereInDimQueryResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.WhereInDimQuery(request)
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

func (client *Client) WhereInDimQueryWithCallback(request *WhereInDimQueryRequest, callback func(response *WhereInDimQueryResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *WhereInDimQueryResponse
		var err error
		defer close(result)
		response, err = client.WhereInDimQuery(request)
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

type WhereInDimQueryRequest struct {
	*requests.RpcRequest
	Limit         requests.Integer             `position:"Query" name:"Limit"`
	IntervalInSec requests.Integer             `position:"Query" name:"IntervalInSec"`
	DatasetId     requests.Integer             `position:"Query" name:"DatasetId"`
	ReduceTail    requests.Boolean             `position:"Query" name:"ReduceTail"`
	MinTime       requests.Integer             `position:"Query" name:"MinTime"`
	MaxTime       requests.Integer             `position:"Query" name:"MaxTime"`
	IsDrillDown   requests.Boolean             `position:"Query" name:"IsDrillDown"`
	DateStr       string                       `position:"Query" name:"DateStr"`
	WhereInValues *[]string                    `position:"Query" name:"WhereInValues"  type:"Repeated"`
	Dimensions    *[]WhereInDimQueryDimensions `position:"Query" name:"Dimensions"  type:"Repeated"`
	WhereInKey    string                       `position:"Query" name:"WhereInKey"`
	OrderByKey    string                       `position:"Query" name:"OrderByKey"`
	Measures      *[]string                    `position:"Query" name:"Measures"  type:"Repeated"`
}

type WhereInDimQueryDimensions struct {
	Key   string `name:"Key"`
	Value string `name:"Value"`
	Type  string `name:"Type"`
}

type WhereInDimQueryResponse struct {
	*responses.BaseResponse
	Data string `json:"Data" xml:"Data"`
}

func CreateWhereInDimQueryRequest() (request *WhereInDimQueryRequest) {
	request = &WhereInDimQueryRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ARMS", "2016-11-25", "WhereInDimQuery", "", "")
	request.Method = requests.POST
	return
}

func CreateWhereInDimQueryResponse() (response *WhereInDimQueryResponse) {
	response = &WhereInDimQueryResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
