package cloudformation

import (
	"encoding/json"
)

// AWSOpsWorksLayer_ShutdownEventConfiguration AWS CloudFormation Resource (AWS::OpsWorks::Layer.ShutdownEventConfiguration)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-layer-lifecycleeventconfiguration-shutdowneventconfiguration.html
type AWSOpsWorksLayer_ShutdownEventConfiguration struct {

	// DelayUntilElbConnectionsDrained AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-layer-lifecycleeventconfiguration-shutdowneventconfiguration.html#cfn-opsworks-layer-lifecycleconfiguration-shutdowneventconfiguration-delayuntilelbconnectionsdrained
	DelayUntilElbConnectionsDrained *Value `json:"DelayUntilElbConnectionsDrained,omitempty"`

	// ExecutionTimeout AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-layer-lifecycleeventconfiguration-shutdowneventconfiguration.html#cfn-opsworks-layer-lifecycleconfiguration-shutdowneventconfiguration-executiontimeout
	ExecutionTimeout *Value `json:"ExecutionTimeout,omitempty"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSOpsWorksLayer_ShutdownEventConfiguration) AWSCloudFormationType() string {
	return "AWS::OpsWorks::Layer.ShutdownEventConfiguration"
}

func (r *AWSOpsWorksLayer_ShutdownEventConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(*r)
}
