package vcs

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

// RecordsItemInListMotorAlgorithmResults is a nested struct in vcs response
type RecordsItemInListMotorAlgorithmResults struct {
	CorpId           string  `json:"CorpId" xml:"CorpId"`
	DataSourceId     string  `json:"DataSourceId" xml:"DataSourceId"`
	LeftTopX         float64 `json:"LeftTopX" xml:"LeftTopX"`
	LeftTopY         float64 `json:"LeftTopY" xml:"LeftTopY"`
	MotorId          string  `json:"MotorId" xml:"MotorId"`
	PicUrlPath       string  `json:"PicUrlPath" xml:"PicUrlPath"`
	PlateNumber      string  `json:"PlateNumber" xml:"PlateNumber"`
	RightBottomX     float64 `json:"RightBottomX" xml:"RightBottomX"`
	RightBottomY     float64 `json:"RightBottomY" xml:"RightBottomY"`
	ShotTime         string  `json:"ShotTime" xml:"ShotTime"`
	TargetPicUrlPath string  `json:"TargetPicUrlPath" xml:"TargetPicUrlPath"`
	MotorStyle       string  `json:"MotorStyle" xml:"MotorStyle"`
	MotorModel       string  `json:"MotorModel" xml:"MotorModel"`
	MotorColor       string  `json:"MotorColor" xml:"MotorColor"`
	MotorClass       string  `json:"MotorClass" xml:"MotorClass"`
	MotorBrand       string  `json:"MotorBrand" xml:"MotorBrand"`
	PlateColor       string  `json:"PlateColor" xml:"PlateColor"`
	PlateClass       string  `json:"PlateClass" xml:"PlateClass"`
	SafetyBelt       string  `json:"SafetyBelt" xml:"SafetyBelt"`
	Calling          string  `json:"Calling" xml:"Calling"`
}
