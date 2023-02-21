package main

import (
	"fmt"
	"time"
)

var (
	userName                string
	password                string
	targetRootDir           string
	fullBackupCronJobTimes  string
	incrementalCronJobTimes string
	enableIncremental       bool
	enableCompresstion      bool
	mysqlport               string
	enableDeleteOldBackup   bool
	keepBackupCount         int
)

const (
	FULL_BACKUP string = "FULL_BACKUP"
	INCREMENTAL string = "INCREMENTAL"
)

func getfullbackupDir() string {
	return fmt.Sprintf("%v/%v/%v",
		targetRootDir,
		FULL_BACKUP,
		time.Now().Format("2006-02-01"),
	)
}
func getIncrementalDir() string {
	return fmt.Sprintf("%v/%v/%v/%v",
		targetRootDir,
		INCREMENTAL,
		time.Now().Format("2006-02-01"),
		time.Now().Format("15-04-05"))

}
