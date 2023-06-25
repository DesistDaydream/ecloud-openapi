package vpc

import (
	ecloudclient "github.com/DesistDaydream/ecloud-openapi/pkg/ecloud_client"
	"github.com/spf13/cobra"
	"gitlab.ecloud.com/ecloud/ecloudsdkcore/config"
	"gitlab.ecloud.com/ecloud/ecloudsdkvpc"
)

type VPCFlags struct {
	test string
}

var (
	vpcFlags  VPCFlags
	vpcClient *ecloudsdkvpc.Client
)

func CreateCommand() *cobra.Command {
	long := ``
	vpcCmd := &cobra.Command{
		Use:   "vpc",
		Short: "VPC",
		Long:  long,
	}

	cobra.OnInitialize(initConfig)

	vpcCmd.AddCommand(
		SecurityGroupCommand(),
	)
	vpcCmd.PersistentFlags().StringVarP(&vpcFlags.test, "test", "t", "", "测试标志")

	return vpcCmd
}

func initConfig() {

	config := &config.Config{
		AccessKey: &ecloudclient.Info.AK,
		SecretKey: &ecloudclient.Info.SK,
		PoolId:    &ecloudclient.Info.Region,
	}
	vpcClient = ecloudsdkvpc.NewClient(config)
}
