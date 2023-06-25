package vpc

import (
	"fmt"

	"github.com/spf13/cobra"
	"gitlab.ecloud.com/ecloud/ecloudsdkvpc/model"
)

func ruleGetCommand() *cobra.Command {
	long := ``
	ruleGetCmd := &cobra.Command{
		Use:   "rule-get",
		Short: "查询安全组规则列表",
		Long:  long,
		Run:   securityGroupRuleGet,
	}

	ruleGetCmd.AddCommand()

	return ruleGetCmd
}

func securityGroupRuleGet(cmd *cobra.Command, args []string) {
	securityGroupId := "11bb4bcc-cd95-4bed-9def-d2df9d528768"

	request := &model.ListSecurityGroupRuleRequest{
		ListSecurityGroupRuleQuery: &model.ListSecurityGroupRuleQuery{
			SecurityGroupId: &securityGroupId,
		},
	}

	response, err := vpcClient.ListSecurityGroupRule(request)
	if err == nil {
		fmt.Printf("%+v\n", response)
	} else {
		fmt.Println(err)
	}
}
