package cmd

import (
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/DesistDaydream/ecloud-openapi/cmd/vpc"
	"github.com/DesistDaydream/ecloud-openapi/config"
	ecloudclient "github.com/DesistDaydream/ecloud-openapi/pkg/ecloud_client"

	logging "github.com/DesistDaydream/logging/pkg/logrus_init"
)

type Flags struct {
	AuthFile string
	Username string
	Region   string
}

func AddFlags(f *Flags) {

}

var (
	flags    Flags
	logFlags logging.LogrusFlags
)

func Execute() {
	app := newApp()
	err := app.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func newApp() *cobra.Command {
	long := ``

	var RootCmd = &cobra.Command{
		Use:   "ecloud-openapi",
		Short: "通过移动云 OpenAPI 管理资源的工具",
		Long:  long,
		// PersistentPreRun: rootPersistentPreRun,
	}

	cobra.OnInitialize(initConfig)

	RootCmd.PersistentFlags().StringVarP(&flags.AuthFile, "auth-file", "F", "config/my_config.yaml", "认证信息文件")
	RootCmd.PersistentFlags().StringVarP(&flags.Username, "username", "u", "", "用户名")
	RootCmd.PersistentFlags().StringVarP(&flags.Region, "region", "r", "CIDC-RP-29", "区域")

	logging.AddFlags(&logFlags)

	// 添加子命令
	RootCmd.AddCommand(
		vpc.CreateCommand(),
	)

	return RootCmd
}

// 执行每个 root 下的子命令时，都需要执行的函数
func initConfig() {
	// 初始化日志
	if err := logging.LogrusInit(&logFlags); err != nil {
		logrus.Fatal("初始化日志失败", err)
	}

	// 检查 clientFlags.AuthFile 文件是否存在
	if _, err := os.Stat(flags.AuthFile); os.IsNotExist(err) {
		logrus.Fatalf("打开【%v】文件失败: %v", flags.AuthFile, err)
	}
	// 获取认证信息
	auth := config.NewAuthInfo(flags.AuthFile)

	// 判断传入的用户是否存在在认证信息中
	if !auth.IsUserExist(flags.Username) {
		logrus.Fatalf("认证信息中不存在 %v 用户, 请检查认证信息文件或命令行参数的值", flags.Username)
	}
	ecloudclient.Info = &ecloudclient.EcloudClientInfo{
		AK:     auth.AuthList[flags.Username].AccessKey,
		SK:     auth.AuthList[flags.Username].SecretKey,
		Region: flags.Region,
	}
}
