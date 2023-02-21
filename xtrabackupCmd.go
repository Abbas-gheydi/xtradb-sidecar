package main

var xtrabackupCmd []string

/* = []string{
	"xtrabackup",
	"--host",
	"127.0.0.1",
	"--backup",
	"--user",
	userName,
	"--password=" + password,
	"--port",
	mysqlport,
}
*/

func createXtraBackupCommand(backupType string) []string {

	var targetdir string
	fullbackupDir := getfullbackupDir()

	incrementalBackupDir := getIncrementalDir()

	cmd := make([]string, 0)
	cmd = append(cmd, xtrabackupCmd...)

	switch backupType {
	case FULL_BACKUP:
		targetdir = "--target-dir=" + fullbackupDir
		cmd = append(cmd, targetdir)

	case INCREMENTAL:
		targetdir = "--target-dir=" + incrementalBackupDir
		cmd = append(cmd, "--incremental-basedir="+fullbackupDir)
		cmd = append(cmd, targetdir)

	}

	if enableCompresstion {
		cmd = append(cmd, "--compress")

	}
	return cmd

}
