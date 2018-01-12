package ram

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

func (client *Client) GetGroup(request *GetGroupRequest) (response *GetGroupResponse, err error) {
	response = CreateGetGroupResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) GetGroupWithChan(request *GetGroupRequest) (<-chan *GetGroupResponse, <-chan error) {
	responseChan := make(chan *GetGroupResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetGroup(request)
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

func (client *Client) GetGroupWithCallback(request *GetGroupRequest, callback func(response *GetGroupResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetGroupResponse
		var err error
		defer close(result)
		response, err = client.GetGroup(request)
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

type GetGroupRequest struct {
	*requests.RpcRequest
	GroupName string `position:"Query" name:"GroupName"`
}

type GetGroupResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Group     struct {
		GroupName  string `json:"GroupName" xml:"GroupName"`
		Comments   string `json:"Comments" xml:"Comments"`
		CreateDate string `json:"CreateDate" xml:"CreateDate"`
		UpdateDate string `json:"UpdateDate" xml:"UpdateDate"`
	} `json:"Group" xml:"Group"`
}

func CreateGetGroupRequest() (request *GetGroupRequest) {
	request = &GetGroupRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ram", "2015-05-01", "GetGroup", "", "")
	request.Method = requests.POST
	return
}

func CreateGetGroupResponse() (response *GetGroupResponse) {
	response = &GetGroupResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
