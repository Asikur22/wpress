package main

import (
	"fmt"
	"github.com/yani-/wpress"
	"os"
	"path/filepath"
	"strings"
)

func fileNameWithoutExtension(fileName string) string {
	return strings.TrimSuffix(fileName, filepath.Ext(fileName))
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	fmt.Printf("Wpress Extracter.\n")

	if len(os.Args) == 2 {
		pathTofile := os.Args[1]
		fileFolderName := fileNameWithoutExtension(pathTofile)
		fileName := filepath.Base(pathTofile)

		isFolderCreated := os.Mkdir(fileFolderName, 0755)
		check(isFolderCreated)
		fmt.Println(fileFolderName + ", New Folder Created\n")

		newFilePathRaw := filepath.Join(fileFolderName, fileName)
		newFilePath, newFilePathErr := filepath.Abs(newFilePathRaw)
		check(newFilePathErr)
		
		isRenamed := os.Rename(pathTofile, newFilePath)
		check(isRenamed)
		
		isChdir := os.Chdir(fileFolderName)
		check(isChdir)

		archiver, _ := wpress.NewReader(newFilePath)
		_, err := archiver.Extract()
		if err != nil {
			fmt.Println("Error = ")
			fmt.Println(err)
		} else {
			fmt.Println("All done!")
		}

		// fmt.Println("total files = ", i, " files read = ", x);
	} else {
		fmt.Println("Inorder to run the extractor please provide the path to the .wpress file as the first argument.")
	}

	// wpress.Init(archiver);

}
