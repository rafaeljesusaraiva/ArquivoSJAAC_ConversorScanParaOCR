package goScripts

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type DocumentType struct {
	documentPath string
	documentName string
	pagesPath    []string
}

func (d *DocumentType) GetDocumentName() string {
	return d.documentName
}

func getFolders(dir string) ([]string, error) {
	var folders []string
	files, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	for _, file := range files {
		if file.IsDir() {
			folders = append(folders, filepath.Join(dir, file.Name()))
		}
	}
	if len(folders) == 0 {
		return nil, errors.New("Nenhuma pasta encontrada")
	}
	return folders, nil
}

func getFilesOfType(dir string, fileType string) ([]string, error) {
	var files []string
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	for _, entry := range entries {
		if !entry.IsDir() && strings.HasSuffix(entry.Name(), "."+fileType) {
			files = append(files, filepath.Join(dir, entry.Name()))
		}
	}
	if len(files) == 0 {
		return nil, errors.New("Nenhum ficheiro '.'" + fileType + " encontrado")
	}
	return files, nil
}

func GatherDocuments(inputFolderPath string, hasMultipleDocs bool) ([]DocumentType, error) {
	returningDocument := []DocumentType{}

	if hasMultipleDocs {
		folders, err := getFolders(inputFolderPath)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}

		for _, folder := range folders {
			/* Get pages from folder */
			pages, err := getFilesOfType(folder, "tif")
			if err != nil {
				fmt.Println(err)
				return nil, err
			}

			newDocument := DocumentType{
				documentPath: folder,
				documentName: filepath.Base(folder),
				pagesPath:    pages,
			}

			returningDocument = append(returningDocument, newDocument)
		}
	} else {
		pages, err := getFilesOfType(inputFolderPath, "tif")
		if err != nil {
			fmt.Println(err)
			return nil, err
		}

		newDocument := DocumentType{
			documentPath: inputFolderPath,
			documentName: filepath.Base(inputFolderPath),
			pagesPath:    pages,
		}

		returningDocument = append(returningDocument, newDocument)
	}

	return returningDocument, nil
}
