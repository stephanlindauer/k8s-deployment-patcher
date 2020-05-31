package main

import (
	"github.com/rebuy-de/rebuy-go-sdk/cmdutil"
	"github.com/sirupsen/logrus"
	"github.com/stephanlindauer/k8s-deployment-patcher/cmd"
)

func main() {
	defer cmdutil.HandleExit()
	if err := cmd.NewRootCommand().Execute(); err != nil {
		logrus.Fatal(err)
	}
}
