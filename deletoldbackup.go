package main

import (
	"log"
	"os"
	"sort"
	"time"
)

func deleteOldBackups(backuptype string, keep int) {
	rsyncLock.Lock()
	log.Println("start deleteting...", backuptype, keep)

	var baseurl string = targetRootDir + "/" + FULL_BACKUP + "/"
	var dateString string = "2006-02-01"
	if backuptype == INCREMENTAL {
		baseurl = targetRootDir + "/" + INCREMENTAL + "/"

	}
	defer rsyncLock.Unlock()
	entries, err := os.ReadDir(baseurl)
	if err != nil {
		log.Println("os.ReadDir(baseurl)", err)
		return
	}

	listOfFolders := make([]time.Time, 0)

	for _, e := range entries {

		t, err := time.Parse(dateString, e.Name())
		if err == nil {
			listOfFolders = append(listOfFolders, t)
		} else {
			log.Println("time.Parse", err)
		}
	}
	sort.Slice(listOfFolders, func(i, j int) bool {

		return listOfFolders[i].Before(listOfFolders[j])
	})
	log.Println("list of foldres", listOfFolders)

	deleteFolderList := make([]time.Time, 0)
	if len(listOfFolders) > keep {

		deleteFolderList = listOfFolders[:len(listOfFolders)-keep]

	}
	for _, folder := range deleteFolderList {
		delFolder := baseurl + folder.Format(dateString)
		log.Println("del", delFolder)
		err := os.RemoveAll(delFolder)
		if err != nil {
			log.Println("could not delete,delFolder", err)
		} else {
			log.Println(delFolder, "is deleted")
		}
	}

}
