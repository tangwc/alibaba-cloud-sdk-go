package ens

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

// DescribeDisks invokes the ens.DescribeDisks API synchronously
func (client *Client) DescribeDisks(request *DescribeDisksRequest) (response *DescribeDisksResponse, err error) {
	response = CreateDescribeDisksResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDisksWithChan invokes the ens.DescribeDisks API asynchronously
func (client *Client) DescribeDisksWithChan(request *DescribeDisksRequest) (<-chan *DescribeDisksResponse, <-chan error) {
	responseChan := make(chan *DescribeDisksResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDisks(request)
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

// DescribeDisksWithCallback invokes the ens.DescribeDisks API asynchronously
func (client *Client) DescribeDisksWithCallback(request *DescribeDisksRequest, callback func(response *DescribeDisksResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDisksResponse
		var err error
		defer close(result)
		response, err = client.DescribeDisks(request)
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

// DescribeDisksRequest is the request struct for api DescribeDisks
type DescribeDisksRequest struct {
	*requests.RpcRequest
	Type           string `position:"Query" name:"Type"`
	OrderByParams  string `position:"Query" name:"OrderByParams"`
	DiskName       string `position:"Query" name:"DiskName"`
	DiskChargeType string `position:"Query" name:"DiskChargeType"`
	EnsRegionId    string `position:"Query" name:"EnsRegionId"`
	InstanceId     string `position:"Query" name:"InstanceId"`
	Status         string `position:"Query" name:"Status"`
	SnapshotId     string `position:"Query" name:"SnapshotId"`
	PageNumber     string `position:"Query" name:"PageNumber"`
	PageSize       string `position:"Query" name:"PageSize"`
	DiskIds        string `position:"Query" name:"DiskIds"`
	DiskId         string `position:"Query" name:"DiskId"`
	EnsRegionIds   string `position:"Query" name:"EnsRegionIds"`
	DiskType       string `position:"Query" name:"DiskType"`
	Category       string `position:"Query" name:"Category"`
}

// DescribeDisksResponse is the response struct for api DescribeDisks
type DescribeDisksResponse struct {
	*responses.BaseResponse
	Code       int    `json:"Code" xml:"Code"`
	PageNumber int    `json:"PageNumber" xml:"PageNumber"`
	PageSize   int    `json:"PageSize" xml:"PageSize"`
	RequestId  string `json:"RequestId" xml:"RequestId"`
	TotalCount int    `json:"TotalCount" xml:"TotalCount"`
	Disks      Disks  `json:"Disks" xml:"Disks"`
}

// CreateDescribeDisksRequest creates a request to invoke DescribeDisks API
func CreateDescribeDisksRequest() (request *DescribeDisksRequest) {
	request = &DescribeDisksRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ens", "2017-11-10", "DescribeDisks", "ens", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeDisksResponse creates a response to parse from DescribeDisks response
func CreateDescribeDisksResponse() (response *DescribeDisksResponse) {
	response = &DescribeDisksResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
