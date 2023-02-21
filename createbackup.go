package main

import (
	"log"
	"os"
	"os/exec"

	cron "github.com/robfig/cron/v3"
)

func createBackupDir(backupType string) {

	if backupType == INCREMENTAL {

		if err := os.MkdirAll(getIncrementalDir(), os.ModePerm); err != nil {
			log.Println(err)
		}

	} else {
		if err := os.MkdirAll(getfullbackupDir(), os.ModePerm); err != nil {
			log.Println(err)
		}

	}

}

func startBackup(backupType string) {
	createBackups := func() {
		log.Println("start: ", backupType)
		createBackupDir(backupType)
		cmd := exec.Command("xtrabackup", createXtraBackupCommand(backupType)...)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			log.Println("exec.Command( createXtraBackupCommand(backupType)...)", err)
		} else {

			if enableDeleteOldBackup {

				go deleteOldBackups(backupType, keepBackupCount)
			}

			if enablerSyncToRemote {
				go rsync()
			}

		}

	}
	cronJobTime := fullBackupCronJobTimes
	if backupType == INCREMENTAL {
		cronJobTime = incrementalCronJobTimes

	}
	c := cron.New()
	c.AddFunc(cronJobTime, createBackups)
	c.Start()

}
