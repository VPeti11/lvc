# LVC - Literally Version Control

**LVC** is a simple CLI tool for basic versioning of your project files. It uses a straightforward folder and file-based system to create and manage versions, with optional integration to Git.

This tool is designed for lightweight version control without the complexity of full Git setups — perfect if you want quick snapshots and an easy way to convert versions into Git repositories later.

---

## Features

* Initialize a simple versioning database file (`lvc.db`)
* Create version snapshots as folders named `Version <number>`
* Copy project files into version folders automatically
* Convert any version folder into a Git repository with one command
* Run arbitrary Git commands inside version folders via CLI
* Minimal dependencies, just Go standard libs and Git CLI

## Installation

Just clone the repo and compile it



### Installer
Download or compile the `installer.go` file and run the command:

    chmod +x installer
    
Then run the installer by typing:

    sudo ./installer
    

### Linux

Make sure you have Go installed:

```
sudo pacman -S go  # or apt, or dnf, your pick
```

Clone the repo and build:

```
go build -o lvc main.go
```

Move it somewhere in your `$PATH` if you're fancy:

```
sudo mv lvc /usr/bin/
```

### Windows

Make sure you have Go and PowerShell. Then compile like this:

```
go build -o lvc.exe main.go
```

Or use WSL like a real developer

---

## Usage

```
lvc <command> [options]
```

### Commands

| Command                    | Description                                  |
| -------------------------- | -------------------------------------------- |
| `init`                     | Create the initial `lvc.db` database file    |
| `create`                   | Create a new version snapshot folder         |
| `convert <version_name>`   | Convert a version folder to a Git repository |
| `git <cmd> [version_name]` | Run a Git command inside a version folder    |

### Examples

Initialize version control:

```
lvc init
```

Create a new version snapshot:

```
lvc create
```

Convert a version named "5" into a Git repo:

```
lvc convert 5
```

Run `git status` on the latest version folder:

```
lvc git status
```

Run `git log` on version "3":

```
lvc git log 3
```

---

## How it works

* The database file `lvc.db` stores the next version number.
* Versions are stored in folders named `Version 1`, `Version 2`, etc.
* When creating a version, LVC copies all project files (except version folders and `lvc.db`) into the new version folder.
* You can convert any version folder into a Git repo with `lvc convert`.
* Run Git commands on any version folder with `lvc git`.

---

## Compatibility

LVC works on Linux, Windows, and macOS — anywhere Go and Git are installed. Note macOS binaries are not provided due to me not being a mac user. 

---

## Dependencies

* Go compiler (for building)
* Git CLI (for Git integration commands)

---

## License

Licensed under the GNU General Public License v3 (GPLv3). See the [LICENSE](LICENSE) file for details.

---
