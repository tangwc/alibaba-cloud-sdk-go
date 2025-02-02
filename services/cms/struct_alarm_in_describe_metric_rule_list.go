package cms

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

// AlarmInDescribeMetricRuleList is a nested struct in cms response
type AlarmInDescribeMetricRuleList struct {
	SilenceTime                  int                            `json:"SilenceTime" xml:"SilenceTime"`
	MetricName                   string                         `json:"MetricName" xml:"MetricName"`
	Webhook                      string                         `json:"Webhook" xml:"Webhook"`
	ContactGroups                string                         `json:"ContactGroups" xml:"ContactGroups"`
	SourceType                   string                         `json:"SourceType" xml:"SourceType"`
	Namespace                    string                         `json:"Namespace" xml:"Namespace"`
	MailSubject                  string                         `json:"MailSubject" xml:"MailSubject"`
	NoEffectiveInterval          string                         `json:"NoEffectiveInterval" xml:"NoEffectiveInterval"`
	EffectiveInterval            string                         `json:"EffectiveInterval" xml:"EffectiveInterval"`
	RuleName                     string                         `json:"RuleName" xml:"RuleName"`
	AlertState                   string                         `json:"AlertState" xml:"AlertState"`
	Period                       string                         `json:"Period" xml:"Period"`
	RuleId                       string                         `json:"RuleId" xml:"RuleId"`
	GroupName                    string                         `json:"GroupName" xml:"GroupName"`
	GroupId                      string                         `json:"GroupId" xml:"GroupId"`
	Dimensions                   string                         `json:"Dimensions" xml:"Dimensions"`
	EnableState                  bool                           `json:"EnableState" xml:"EnableState"`
	GroupBy                      string                         `json:"GroupBy" xml:"GroupBy"`
	Resources                    string                         `json:"Resources" xml:"Resources"`
	NoDataPolicy                 string                         `json:"NoDataPolicy" xml:"NoDataPolicy"`
	Options                      string                         `json:"Options" xml:"Options"`
	DynamicAlertSensitivity      string                         `json:"DynamicAlertSensitivity" xml:"DynamicAlertSensitivity"`
	DynamicAlertHistoryDataRange string                         `json:"DynamicAlertHistoryDataRange" xml:"DynamicAlertHistoryDataRange"`
	RuleType                     string                         `json:"RuleType" xml:"RuleType"`
	Escalations                  Escalations                    `json:"Escalations" xml:"Escalations"`
	CompositeExpression          CompositeExpression            `json:"CompositeExpression" xml:"CompositeExpression"`
	Prometheus                   Prometheus                     `json:"Prometheus" xml:"Prometheus"`
	Labels                       LabelsInDescribeMetricRuleList `json:"Labels" xml:"Labels"`
}
