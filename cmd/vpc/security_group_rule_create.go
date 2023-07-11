package vpc

import (
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

	return ruleCreateCmd
}

func securityGroupRuleCreate(cmd *cobra.Command, args []string) {
	if securityGroupFlags.SecurityGroupFile == "" {
		logrus.Fatalf("请指定要读取安全组规则的 excel 文件。")
	}

	sgID, err := findSecurityGroupID(securityGroupFlags.SecurityGroupName)
	if err != nil {
		logrus.Fatalf("获取安全组 ID 失败，原因: %v", err)
	}

	// logrus.Debugf("安全组ID为 %v", sgID)

	// body, err := fileparse.GetSecurityGroupRules(securityGroupFlags.SecurityGroupFile, securityGroupFlags.SecurityGroupName)
	// if err != nil {
	// 	logrus.Fatalf("获取待创建的安全组规则失败，原因: %v", err)
	// }

	// for _, rule := range body {
	// 	response, err := vpcClient.CreateSecurityRule(&model.CreateSecurityRuleRequest{
	// 		CreateSecurityRuleBody: &model.CreateSecurityRuleBody{
	// 			SecurityGroupId: &sgID,
	// 			RemoteType:      rule.RemoteType,
	// 			Protocol:        rule.Protocol,
	// 			EtherType:       rule.EtherType,
	// 			Description:     &rule.Description,
	// 			RemoteIpPrefix:  &rule.RemoteIpPrefix,
	// 			Direction:       rule.Direction,
	// 		},
	// 	})
	// 	if err != nil {
	// 		logrus.Errorf("创建规则失败，原因: %v", err)
	// 	}

	// 	logrus.Infoln(response)
	// }

	remoteType := "cidr"
	protocol := "ANY"
	etherType := "IPv4"
	description := "测试描述"
	direction := "ingress"
	remoteIpPrefix := "1.1.1.2/32"

	response, err := vpcClient.CreateSecurityRule(&model.CreateSecurityRuleRequest{
		CreateSecurityRuleBody: &model.CreateSecurityRuleBody{
			SecurityGroupId: &sgID,
			RemoteType:      (*model.CreateSecurityRuleBodyRemoteTypeEnum)(&remoteType),
			Protocol:        (*model.CreateSecurityRuleBodyProtocolEnum)(&protocol),
			EtherType:       (*model.CreateSecurityRuleBodyEtherTypeEnum)(&etherType),
			Description:     &description,
			RemoteIpPrefix:  &remoteIpPrefix,
			Direction:       (*model.CreateSecurityRuleBodyDirectionEnum)(&direction),
		},
	})
	if err != nil {
		logrus.Errorf("创建规则失败，原因: %v", err)
	}

	logrus.Infoln(response)
}
