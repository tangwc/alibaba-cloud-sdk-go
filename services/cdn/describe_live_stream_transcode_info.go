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

func (client *Client) DescribeLiveStreamTranscodeInfo(request *DescribeLiveStreamTranscodeInfoRequest) (response *DescribeLiveStreamTranscodeInfoResponse, err error) {
	response = CreateDescribeLiveStreamTranscodeInfoResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) DescribeLiveStreamTranscodeInfoWithChan(request *DescribeLiveStreamTranscodeInfoRequest) (<-chan *DescribeLiveStreamTranscodeInfoResponse, <-chan error) {
	responseChan := make(chan *DescribeLiveStreamTranscodeInfoResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeLiveStreamTranscodeInfo(request)
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

func (client *Client) DescribeLiveStreamTranscodeInfoWithCallback(request *DescribeLiveStreamTranscodeInfoRequest, callback func(response *DescribeLiveStreamTranscodeInfoResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeLiveStreamTranscodeInfoResponse
		var err error
		defer close(result)
		response, err = client.DescribeLiveStreamTranscodeInfo(request)
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

type DescribeLiveStreamTranscodeInfoRequest struct {
	*requests.RpcRequest
	DomainTranscodeName string           `position:"Query" name:"DomainTranscodeName"`
	OwnerId             requests.Integer `position:"Query" name:"OwnerId"`
	SecurityToken       string           `position:"Query" name:"SecurityToken"`
}

type DescribeLiveStreamTranscodeInfoResponse struct {
	*responses.BaseResponse
	RequestId           string `json:"RequestId" xml:"RequestId"`
	DomainTranscodeList struct {
		DomainTranscodeInfo []struct {
			TranscodeApp      string `json:"TranscodeApp" xml:"TranscodeApp"`
			TranscodeId       string `json:"TranscodeId" xml:"TranscodeId"`
			TranscodeName     string `json:"TranscodeName" xml:"TranscodeName"`
			TranscodeRecord   string `json:"TranscodeRecord" xml:"TranscodeRecord"`
			TranscodeSnapshot string `json:"TranscodeSnapshot" xml:"TranscodeSnapshot"`
			TranscodeTemplate string `json:"TranscodeTemplate" xml:"TranscodeTemplate"`
		} `json:"DomainTranscodeInfo" xml:"DomainTranscodeInfo"`
	} `json:"DomainTranscodeList" xml:"DomainTranscodeList"`
}

func CreateDescribeLiveStreamTranscodeInfoRequest() (request *DescribeLiveStreamTranscodeInfoRequest) {
	request = &DescribeLiveStreamTranscodeInfoRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cdn", "2014-11-11", "DescribeLiveStreamTranscodeInfo", "", "")
	request.Method = requests.POST
	return
}

func CreateDescribeLiveStreamTranscodeInfoResponse() (response *DescribeLiveStreamTranscodeInfoResponse) {
	response = &DescribeLiveStreamTranscodeInfoResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
