package vpc

import "github.com/spf13/cobra"

type SecurityGroupFlags struct {
	test string
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

	return securityGroupCmd
}
