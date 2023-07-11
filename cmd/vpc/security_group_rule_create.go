package vpc

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"gitlab.ecloud.com/ecloud/ecloudsdkvpc/model"
)

func ruleCreateCommand() *cobra.Command {
	long := ``
	ruleCreateCmd := &cobra.Command{
		Use:   "rule-create",
		Short: "创建安全组规则",
		Long:  long,
		Run:   securityGroupRuleCreate,
	}

	ruleCreateCmd.AddCommand()

	return ruleCreateCmd
}

func securityGroupRuleCreate(cmd *cobra.Command, args []string) {
	sgID, err := findSecurityGroupID(securityGroupFlags.SecurityGroupName)
	if err != nil {
		logrus.Fatalf("获取安全组 ID 失败，原因: %v", err)
	}

	remoteType := "cidr"
	protocol := "ANY"
	etherType := "IPv4"
	description := "测试描述"
	direction := "ingress"
	remoteIpPrefix := "1.1.1.2/32"

	request := &model.CreateSecurityRuleRequest{
		CreateSecurityRuleBody: &model.CreateSecurityRuleBody{
			SecurityGroupId: &sgID,
			RemoteType:      (*model.CreateSecurityRuleBodyRemoteTypeEnum)(&remoteType),
			Protocol:        (*model.CreateSecurityRuleBodyProtocolEnum)(&protocol),
			EtherType:       (*model.CreateSecurityRuleBodyEtherTypeEnum)(&etherType),
			Description:     &description,
			RemoteIpPrefix:  &remoteIpPrefix,
			Direction:       (*model.CreateSecurityRuleBodyDirectionEnum)(&direction),
		},
	}

	response, err := vpcClient.CreateSecurityRule(request)
	if err == nil {
		fmt.Printf("%+v\n", response)
	} else {
		fmt.Println(err)
	}
}
