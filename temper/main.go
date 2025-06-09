package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func checkIfDirExists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}

func getUserInput(dirpaths []os.DirEntry) int {
	var chooseDirId int

	for {
		fmt.Println(">>> temper")
		for i, dir := range dirpaths {
			fmt.Printf(">>> %d %s\n", i+1, dir.Name())
		}
		fmt.Printf(">>> choose %d - %d\n", 1, len(dirpaths))

		fmt.Scanln(&chooseDirId)
		if chooseDirId >= 1 && chooseDirId <= len(dirpaths) {
			break
		}
		fmt.Print(">>> choose invalid input\n\n")
	}

	return chooseDirId
}

func main() {
	fmt.Println("\nChecking if ~/.config/temper exists...")
	configPath := filepath.Join(os.Getenv("HOME"), ".config", "temper")

	if !checkIfDirExists(configPath) {
		panic("~/.config/temper does not exist")
	}

	fmt.Println("\nDirectories in ~/.config/temper:")
	dirs, err := os.ReadDir(configPath)
	if err != nil {
		panic(err)
	}

	var dirEntries []os.DirEntry
	for _, dir := range dirs {
		if dir.Name() != ".git" {
			dirEntries = append(dirEntries, dir)
		}
	}

	dirId := getUserInput(dirEntries)

	selectedDirPath := filepath.Join(configPath, dirEntries[dirId-1].Name())
	currentPath, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	err = CopyDir(selectedDirPath, currentPath)
	if err != nil {
		panic(err)
	}
	fmt.Println(">>> copied :D")
}
