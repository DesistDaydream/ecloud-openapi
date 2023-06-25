package vpc

import (
	"fmt"

	"github.com/spf13/cobra"
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
	"gitlab.ecloud.com/ecloud/ecloudsdkvpc/model"
)

func getCommand() *cobra.Command {
	long := ``
	getCmd := &cobra.Command{
		Use:   "get",
		Short: "查询安全组详情",
		Long:  long,
		Run:   securityGroupList,
	}

	getCmd.AddCommand()

	return getCmd
}

func securityGroupList(cmd *cobra.Command, args []string) {
	securityGroupId := "11bb4bcc-cd95-4bed-9def-d2df9d528768"

	request := &model.GetSecurityGroupDetailRespRequest{
		GetSecurityGroupDetailRespPath: &model.GetSecurityGroupDetailRespPath{
			Path:            position.Path{},
			SecurityGroupId: &securityGroupId,
		},
	}

	response, err := vpcClient.GetSecurityGroupDetailResp(request)
	if err == nil {
		fmt.Printf("%+v\n", response)
	} else {
		fmt.Println(err)
	}
}
