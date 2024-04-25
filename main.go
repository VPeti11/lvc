package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
)

const (
	databaseFileName = "lvc.db"
	versionFolder    = "Version"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go [init|create|convert|git <git_command> [<version_name>]]")
		return
	}

	command := os.Args[1]

	switch command {
	case "init":
		err := initializeDatabase()
		if err != nil {
			fmt.Println("Error initializing database:", err)
		} else {
			fmt.Println("Database initialized successfully.")
		}
	case "create":
		err := createVersion()
		if err != nil {
			fmt.Println("Error creating version:", err)
		} else {
			fmt.Println("Version created successfully.")
		}
	case "convert":
		if len(os.Args) < 3 {
			fmt.Println("Usage: go run main.go convert <version_name>")
			return
		}
		versionName := os.Args[2]
		err := convertVersionToGit(versionName)
		if err != nil {
			fmt.Println("Error converting version to Git repository:", err)
		} else {
			fmt.Println("Version converted to Git repository successfully.")
		}
	case "git":
		if len(os.Args) < 3 {
			fmt.Println("Usage: go run main.go git <git_command> [<version_name>]")
			return
		}
		gitCommand := os.Args[2]
		var versionName string
		if len(os.Args) >= 4 {
			versionName = os.Args[3]
		}
		err := runGitCommand(gitCommand, versionName)
		if err != nil {
			fmt.Println("Error running git command:", err)
		} else {
			fmt.Println("Git command executed successfully.")
		}
	default:
		fmt.Println("Unknown command.")
	}
}

func initializeDatabase() error {
	// Check if the database file already exists
	_, err := os.Stat(databaseFileName)
	if os.IsNotExist(err) {
		// Create the database file
		file, err := os.Create(databaseFileName)
		if err != nil {
			return err
		}
		defer file.Close()

		// Write initial value (1) to the file
		err = writeDatabaseValue(1)
		if err != nil {
			return err
		}
	} else if err != nil {
		return err
	}
	return nil
}

func createVersion() error {
	// Read the current value from the database
	currentValue, err := readDatabaseValue()
	if err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("please run 'init' first")
		}
		return err
	}

	// Create version folder if it doesn't exist
	versionFolderName := fmt.Sprintf("%s %d", versionFolder, currentValue)
	err = os.Mkdir(versionFolderName, 0755)
	if err != nil && !os.IsExist(err) {
		return err
	}

	// Copy contents of lvc.db to the version folder
	err = copyAllFiles(".", versionFolderName)
	if err != nil {
		return err
	}

	// Increment the database value
	err = writeDatabaseValue(currentValue + 1)
	if err != nil {
		return err
	}

	return nil
}

func convertVersionToGit(versionName string) error {
	versionFolderName := fmt.Sprintf("%s %s", versionFolder, versionName)
	err := runGitInitAndAdd(versionFolderName)
	if err != nil {
		return err
	}
	return nil
}

func runGitInitAndAdd(folder string) error {
	cmd := exec.Command("git", "init")
	cmd.Dir = folder
	err := runCmdAndWait(cmd)
	if err != nil {
		return err
	}

	cmd = exec.Command("git", "add", "*")
	cmd.Dir = folder
	err = runCmdAndWait(cmd)
	if err != nil {
		return err
	}

	return nil
}

func runGitCommand(gitCommand, versionName string) error {
	var folder string
	if versionName != "" {
		folder = fmt.Sprintf("%s %s", versionFolder, versionName)
	} else {
		// Read the current value from the database
		currentValue, err := readDatabaseValue()
		if err != nil {
			if os.IsNotExist(err) {
				return fmt.Errorf("please run 'init' first")
			}
			return err
		}
		folder = fmt.Sprintf("%s %d", versionFolder, currentValue-1)
	}

	cmd := exec.Command("git", gitCommand)
	cmd.Dir = folder
	err := runCmdAndWait(cmd)
	if err != nil {
		return err
	}

	return nil
}

func readDatabaseValue() (int, error) {
	file, err := os.Open(databaseFileName)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	if scanner.Scan() {
		valueStr := strings.TrimSpace(scanner.Text())
		value, err := strconv.Atoi(valueStr)
		if err != nil {
			return 0, err
		}
		return value, nil
	}

	return 0, fmt.Errorf("empty database")
}

func writeDatabaseValue(value int) error {
	file, err := os.OpenFile(databaseFileName, os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = fmt.Fprintf(file, "%d\n", value)
	if err != nil {
		return err
	}

	return nil
}

func copyAllFiles(srcDir, destDir string) error {
	return filepath.Walk(srcDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if path == databaseFileName || strings.HasPrefix(info.Name(), versionFolder) {
			return nil
		}

		dest := filepath.Join(destDir, strings.TrimPrefix(path, srcDir))
		if info.IsDir() {
			return os.MkdirAll(dest, 0755)
		}

		source, err := os.Open(path)
		if err != nil {
			return err
		}
		defer source.Close()

		destination, err := os.Create(dest)
		if err != nil {
			return err
		}
		defer destination.Close()

		_, err = io.Copy(destination, source)
		return err
	})
}

func runCmdAndWait(cmd *exec.Cmd) error {
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
