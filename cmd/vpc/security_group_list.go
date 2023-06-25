package vpc

import (
	"fmt"

	"github.com/spf13/cobra"
	"gitlab.ecloud.com/ecloud/ecloudsdkvpc/model"
)

func listCommand() *cobra.Command {
	long := ``
	listCmd := &cobra.Command{
		Use:   "list",
		Short: "安全组",
		Long:  long,
		Run:   securityGroupList,
	}

	listCmd.AddCommand()

	return listCmd
}

func securityGroupList(cmd *cobra.Command, args []string) {
	request := &model.ListSecurityGroupRespRequest{}
	response, err := vpcClient.ListSecurityGroupResp(request)
	if err == nil {
		fmt.Printf("%+v\n", response)
	} else {
		fmt.Println(err)
	}
}
