package vpc

import (
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"gitlab.ecloud.com/ecloud/ecloudsdkvpc/model"
)

func listCommand() *cobra.Command {
	long := ``
	listCmd := &cobra.Command{
		Use:   "list",
		Short: "查询安全组规则列表",
		Long:  long,
		Run:   securityGroupGet,
	}

	listCmd.AddCommand()

	return listCmd
}

func securityGroupGet(cmd *cobra.Command, args []string) {
	request := &model.ListSecurityGroupRespRequest{}
	response, err := vpcClient.ListSecurityGroupResp(request)
	if err != nil {
		fmt.Println(err)
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"SG-ID", "名称", "描述"})

	for _, content := range *response.Body.Content {
		table.Append([]string{*content.Id, *content.Name, *content.Description})
	}

	table.Render()
}
