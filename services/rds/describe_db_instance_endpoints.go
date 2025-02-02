package rds

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

// DescribeDBInstanceEndpoints invokes the rds.DescribeDBInstanceEndpoints API synchronously
func (client *Client) DescribeDBInstanceEndpoints(request *DescribeDBInstanceEndpointsRequest) (response *DescribeDBInstanceEndpointsResponse, err error) {
	response = CreateDescribeDBInstanceEndpointsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDBInstanceEndpointsWithChan invokes the rds.DescribeDBInstanceEndpoints API asynchronously
func (client *Client) DescribeDBInstanceEndpointsWithChan(request *DescribeDBInstanceEndpointsRequest) (<-chan *DescribeDBInstanceEndpointsResponse, <-chan error) {
	responseChan := make(chan *DescribeDBInstanceEndpointsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDBInstanceEndpoints(request)
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

// DescribeDBInstanceEndpointsWithCallback invokes the rds.DescribeDBInstanceEndpoints API asynchronously
func (client *Client) DescribeDBInstanceEndpointsWithCallback(request *DescribeDBInstanceEndpointsRequest, callback func(response *DescribeDBInstanceEndpointsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDBInstanceEndpointsResponse
		var err error
		defer close(result)
		response, err = client.DescribeDBInstanceEndpoints(request)
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

// DescribeDBInstanceEndpointsRequest is the request struct for api DescribeDBInstanceEndpoints
type DescribeDBInstanceEndpointsRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ClientToken          string           `position:"Query" name:"ClientToken"`
	DBInstanceEndpointId string           `position:"Query" name:"DBInstanceEndpointId"`
	DBInstanceId         string           `position:"Query" name:"DBInstanceId"`
}

// DescribeDBInstanceEndpointsResponse is the response struct for api DescribeDBInstanceEndpoints
type DescribeDBInstanceEndpointsResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateDescribeDBInstanceEndpointsRequest creates a request to invoke DescribeDBInstanceEndpoints API
func CreateDescribeDBInstanceEndpointsRequest() (request *DescribeDBInstanceEndpointsRequest) {
	request = &DescribeDBInstanceEndpointsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "DescribeDBInstanceEndpoints", "rds", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeDBInstanceEndpointsResponse creates a response to parse from DescribeDBInstanceEndpoints response
func CreateDescribeDBInstanceEndpointsResponse() (response *DescribeDBInstanceEndpointsResponse) {
	response = &DescribeDBInstanceEndpointsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
