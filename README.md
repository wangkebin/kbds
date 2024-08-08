# KBDS is a tool to gather your local file properties into a DB table.
This is a tool I use to detect duplicate files. 
  After 10+ years of computer usage and backup from different laptops/desktops at different stages, I have tons of files.
Files being backed up in different folders at different years would be redundant and create confusion when I just want the latest version. 
It would be a pity if you deleted the latest version by mistake. 

## How to use. 
Currently, it uses the MySQL database to store information. Here is a way to find duplicated files in the DB

The following query finds all duplications by name only
```
SELECT COUNT(*), d.name, d.loc FROM dirs d GROUP BY d.name having COUNT(*) > 2 ORDER BY size DESC;
```
## Performance
It walked a folder with 173k files and wrote those records into a MySQL DB on a local area network in less than 12 seconds. 

## Future Features
### Auto update
Hook up with file system update operation to update entries. 
### UI for file operations
Once duplicate files or directories are detected, let the user delete/merge them into one location

