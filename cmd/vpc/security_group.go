package vpc

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"gitlab.ecloud.com/ecloud/ecloudsdkvpc/model"
)

type SecurityGroupFlags struct {
	test              string
	SecurityGroupName string
	SecurityGroupFile string
}

var securityGroupFlags SecurityGroupFlags

func SecurityGroupCommand() *cobra.Command {
	long := ``
	securityGroupCmd := &cobra.Command{
		Use:   "sg",
		Short: "安全组",
		Long:  long,
	}

	securityGroupCmd.AddCommand(
		listCommand(),
		getCommand(),
		ruleGetCommand(),
		ruleCreateCommand(),
	)

	securityGroupCmd.PersistentFlags().StringVarP(&securityGroupFlags.test, "test", "t", "", "测试标志")
	securityGroupCmd.PersistentFlags().StringVar(&securityGroupFlags.SecurityGroupName, "sg-name", "", "安全组名称")
	securityGroupCmd.PersistentFlags().StringVar(&securityGroupFlags.SecurityGroupFile, "file", "", "安全组excel文件")

	return securityGroupCmd
}

func findSecurityGroupID(sgName string) (string, error) {
	if securityGroupFlags.SecurityGroupName == "" {
		logrus.Fatalf("请指定安全组名称")
	}

	var sgID string

	response, err := vpcClient.ListSecurityGroupResp(&model.ListSecurityGroupRespRequest{})
	if err != nil {
		return "", fmt.Errorf("无法列出安全组，原因: %v", err)
	}

	for _, content := range *response.Body.Content {
		if *content.Name == securityGroupFlags.SecurityGroupName {
			sgID = *content.Id
		}
	}

	if sgID == "" {
		logrus.Fatalf("指定的安全组名称不存在")
	}

	return sgID, nil
}
