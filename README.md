# LVC (Literally Version Control)

This is a simple version control system. It allows you to initialize a database, create versions, convert the latest version into a Git repository, and execute Git commands within a version's folder.

## Usage

### Initialize Database

lcv init

### Create Version

lcv create

Creates a new version. It copies all files and folders except `lvc.db` and folders starting with "Version " into a new version folder.

# Version Control System (VCS) in Go

This is a simple version control system implemented in Go. It allows you to initialize a database, create versions, convert versions into Git repositories, and execute Git commands within specific version folders.

## Usage

### Initialize Database

lcv init

Initializes the database.

### Create Version

lcv create

Creates a new version. It copies all files and folders except `lvc.db` and folders starting with "Version " into a new version folder.

### Convert Version to Git Repository

lcv convert <version_name>

Converts the specified version into a Git repository. Replace `<version_name>` with the name of the version you want to convert.

### Execute Git Command in Version's Folder

lcv git <git_command> [<version_name>]

Executes a Git command within the specified version's folder. Replace `<git_command>` with the desired Git command. If `<version_name>` is provided, the command is executed within that version's folder; otherwise, it defaults to the latest version.

## Note
- Run `init` before using any other command.
- The `git` command can be used with or without specifying a version name.
- The `convert` command converts a specific version into a Git repository.

