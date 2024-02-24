# KBDS is a tool to gather your local file properties into a DB table.
This is a tool I use to detect duplicate files. 
  After 10+ years of computer usage and backup from different laptops/desktops at different stage, I have tons of files.
Files being backed up at different folders at different years would be redundant and create confussion when I just want the latest version. 
It would be a pity if you delete the latest version by mistake. 

## How to usage. 
Currently it uses MySQL database to store these info. Here are a way to find dup files in the DB

The following query finds all dups by name only
```
SELECT COUNT(*), d.name, d.loc FROM dirs d GROUP BY d.name having COUNT(*) > 2 ORDER BY size DESC;
```

## Future Features
### Auto update
Hook up with file system update to update entries. 
### UI for file operations
Once dup files/dirs detected, let the user delete/merge them into one location

