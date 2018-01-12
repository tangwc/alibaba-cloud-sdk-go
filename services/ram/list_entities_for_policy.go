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

func (client *Client) ListEntitiesForPolicy(request *ListEntitiesForPolicyRequest) (response *ListEntitiesForPolicyResponse, err error) {
	response = CreateListEntitiesForPolicyResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) ListEntitiesForPolicyWithChan(request *ListEntitiesForPolicyRequest) (<-chan *ListEntitiesForPolicyResponse, <-chan error) {
	responseChan := make(chan *ListEntitiesForPolicyResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListEntitiesForPolicy(request)
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

func (client *Client) ListEntitiesForPolicyWithCallback(request *ListEntitiesForPolicyRequest, callback func(response *ListEntitiesForPolicyResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListEntitiesForPolicyResponse
		var err error
		defer close(result)
		response, err = client.ListEntitiesForPolicy(request)
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

type ListEntitiesForPolicyRequest struct {
	*requests.RpcRequest
	PolicyType string `position:"Query" name:"PolicyType"`
	PolicyName string `position:"Query" name:"PolicyName"`
}

type ListEntitiesForPolicyResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Groups    struct {
		Group []struct {
			GroupName  string `json:"GroupName" xml:"GroupName"`
			Comments   string `json:"Comments" xml:"Comments"`
			AttachDate string `json:"AttachDate" xml:"AttachDate"`
		} `json:"Group" xml:"Group"`
	} `json:"Groups" xml:"Groups"`
	Users struct {
		User []struct {
			UserId      string `json:"UserId" xml:"UserId"`
			UserName    string `json:"UserName" xml:"UserName"`
			DisplayName string `json:"DisplayName" xml:"DisplayName"`
			AttachDate  string `json:"AttachDate" xml:"AttachDate"`
		} `json:"User" xml:"User"`
	} `json:"Users" xml:"Users"`
	Roles struct {
		Role []struct {
			RoleId      string `json:"RoleId" xml:"RoleId"`
			RoleName    string `json:"RoleName" xml:"RoleName"`
			Arn         string `json:"Arn" xml:"Arn"`
			Description string `json:"Description" xml:"Description"`
			AttachDate  string `json:"AttachDate" xml:"AttachDate"`
		} `json:"Role" xml:"Role"`
	} `json:"Roles" xml:"Roles"`
}

func CreateListEntitiesForPolicyRequest() (request *ListEntitiesForPolicyRequest) {
	request = &ListEntitiesForPolicyRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ram", "2015-05-01", "ListEntitiesForPolicy", "", "")
	request.Method = requests.POST
	return
}

func CreateListEntitiesForPolicyResponse() (response *ListEntitiesForPolicyResponse) {
	response = &ListEntitiesForPolicyResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
