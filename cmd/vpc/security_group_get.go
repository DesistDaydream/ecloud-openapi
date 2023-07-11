package vpc

import (
	"fmt"

	"github.com/sirupsen/logrus"
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
	sgID, err := findSecurityGroupID(securityGroupFlags.SecurityGroupName)
	if err != nil {
		logrus.Fatalf("获取安全组 ID 失败，原因: %v", err)
	}

	request := &model.GetSecurityGroupDetailRespRequest{
		GetSecurityGroupDetailRespPath: &model.GetSecurityGroupDetailRespPath{
			Path:            position.Path{},
			SecurityGroupId: &sgID,
		},
	}

	response, err := vpcClient.GetSecurityGroupDetailResp(request)
	if err == nil {
		fmt.Printf("%+v\n", response)
	} else {
		fmt.Println(err)
	}
}
