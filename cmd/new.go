package cmd

import (
	"github.com/lanytcc/spot/cloud/qcloud"

	"github.com/spf13/cobra"
)

func cmdNew() *cobra.Command {
	var count int64
	var netaccess, windows bool
	var image string
	c := &cobra.Command{
		Use:     "new",
		Aliases: []string{"up", "create"},
		Short:   "新建腾讯云虚拟机",
		RunE: func(_ *cobra.Command, _ []string) error {
			client := qcloud.NewClient()
			return client.Create(count, netaccess, windows, image)
		},
	}
	c.Flags().Int64VarP(&count, "count", "c", 1, "虚拟机数量")
	c.Flags().BoolVar(&netaccess, "net", true, "是否开启公网访问, 单节点生效")
	c.Flags().BoolVar(&windows, "windows", false, "windows节点")
	c.Flags().StringVarP(&image, "image", "i", "", "指定linux镜像")
	return c
}
