package vpc

import (
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
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
	request := &model.ListSecurityGroupRuleRequest{
		ListSecurityGroupRuleQuery: &model.ListSecurityGroupRuleQuery{
			SecurityGroupId: &securityGroupFlags.SecurityGroupID,
		},
	}

	response, err := vpcClient.ListSecurityGroupRule(request)
	if err != nil {
		fmt.Println(err)
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"方向", "协议", "描述", "CIDR", "port"})

	for _, r := range *response.Body.Content {
		var (
			des  string
			cidr string
			port string
		)

		if r.Description != nil {
			des = *r.Description
		} else {
			des = ""
		}

		if r.RemoteIpPrefix != nil {
			cidr = *r.RemoteIpPrefix
		} else {
			cidr = ""
		}

		if r.MinPortRange != nil || r.MaxPortRange != nil {
			port = fmt.Sprintf("%v-%v", *r.MinPortRange, *r.MaxPortRange)
		} else {
			port = ""
		}

		table.Append([]string{string(*r.Direction), string(*r.Protocol), des, cidr, port})
	}

	table.Render()
}
