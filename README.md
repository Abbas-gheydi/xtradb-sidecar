how to restore backups:
run these commands on remote server on testing mysql server
first copy "full backup" and "incremental backup" to new location.
Do not run this command on original folders.
then:
# decompres full backup
xtrabackup --decompress --target-dir=2023-21-02/

#  decompres incremental
xtrabackup --decompress --target-dir=11-23-00/

xtrabackup --prepare --apply-log-only --target-dir=2023-21-02/
xtrabackup --prepare --target-dir=2023-21-02/ --incremental-dir=11-23-00/
sudo rsync -avrP 2023-21-02/ /var/lib/mysql/

