package main

import (
	"fmt"
	"log"
	"sync"
)

func main() {

	fmt.Printf(`
	targetRootDir: %v
	fullBackupCronJobTimes: %v
	incrementalCronJobTimes: %v
	enableIncremental %v
	enableCompresstion: %v
	enableDeleteOldBackup: %v
	keepBackupCount: %v
	`,
		targetRootDir,
		fullBackupCronJobTimes,
		incrementalCronJobTimes,
		enableIncremental,
		enableCompresstion,
		enableDeleteOldBackup,
		keepBackupCount,
	)
	log.Println("xtradb-sidecar started.")

	var wg sync.WaitGroup

	wg.Add(1)
	go startBackup(FULL_BACKUP)

	if enableIncremental {
		go startBackup(INCREMENTAL)
	}

	wg.Wait()

}
