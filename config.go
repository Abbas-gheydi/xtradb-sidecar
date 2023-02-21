package main

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigName("config")              // name of config file (without extension)
	viper.SetConfigType("yaml")                // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("/etc/xtradb-sidecar") // path to look for the config file in
	viper.AddConfigPath("$HOME/.appname")      // call multiple times to add many search paths
	viper.AddConfigPath(".")                   // optionally look for config in the working directory
	err := viper.ReadInConfig()                // Find and read the config file
	if err != nil {                            // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	//fmt.Println(viper.Get("image"))
	userName = viper.GetString("userName")
	//password = viper.GetString("password")
	password = os.Getenv("XT_PASSWORD")
	targetRootDir = viper.GetString("targetRootDir")
	fullBackupCronJobTimes = viper.GetString("fullBackupCronJobTimes")
	incrementalCronJobTimes = viper.GetString("incrementalCronJobTimes")
	enableIncremental = viper.GetBool("enableIncremental")
	enableCompresstion = viper.GetBool("enableCompresstion")
	mysqlport = viper.GetString("mysqlport")
	enableDeleteOldBackup = viper.GetBool("enableDeleteOldBackup")
	keepBackupCount = viper.GetInt("keepBackupCount")
	sshPort = viper.GetString("sshPort")
	sshidentityFile = viper.GetString("sshidentityFile")
	sshUsername = viper.GetString("sshUsername")
	sshServerAddr = viper.GetString("sshServerAddr")
	RemoteDir = viper.GetString("RemoteDir")
	enablerSyncToRemote = viper.GetBool("enablerSyncToRemote")

	xtrabackupCmd = []string{
		//"xtrabackup",
		"--host",
		"127.0.0.1",
		"--backup",
		"--user",
		userName,
		"--password=" + password,
		"--port",
		mysqlport,
	}

	rsyncCmd = []string{
		"-ah",
		"--delete",
		"-e",
		"ssh -oStrictHostKeyChecking=no -p " + sshPort + " -i " + sshidentityFile,
		targetRootDir,
		sshUsername + "@" + sshServerAddr + ":" + RemoteDir,
	}
	fmt.Println(mysqlport)

	fmt.Println(userName)
}
