package goScripts

import (
	"github.com/disintegration/imaging"
	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/types"
	"log"
	"os"
	"path/filepath"
)

func changeFilePathAndType(oldFilePath string, newDir string, newFileType string) (string, error) {
	// Get the file name without the extension
	fileName := filepath.Base(oldFilePath)
	fileNameWithoutExt := fileName[:len(fileName)-len(filepath.Ext(fileName))]

	// Create the new file path
	newFilePath := filepath.Join(newDir, fileNameWithoutExt+"."+newFileType)

	return newFilePath, nil
}

func ConvertDocumentPagesToJpg(documentToProcess DocumentType, temporaryFolder string, taskId int, taskProgress *OverallProgress) (DocumentType, error) {
	documentInJPG := DocumentType{
		documentName: documentToProcess.documentName,
		pagesPath:    []string{},
		documentPath: "",
	}

	/* count how many files to convert */
	totalFiles := len(documentToProcess.pagesPath)
	/* create in float64 the value to increment after each image has been converted */
	incrementValue := 100.0 / float64(totalFiles)

	/* check if temporary folder exists, if not create one */
	tempDocumentFolderPath := filepath.Join(temporaryFolder, documentToProcess.documentName)
	// Check if the folder exists
	if _, err := os.Stat(tempDocumentFolderPath); os.IsNotExist(err) {
		// Create the folder if it doesn't exist
		err := os.Mkdir(tempDocumentFolderPath, 0755)
		if err != nil {
			log.Fatalf("failed to create directory: %v", err)
			return DocumentType{}, err
		}
	}

	pagesInJPG := []string{}
	for _, page := range documentToProcess.pagesPath {
		img, err := imaging.Open(page)
		if err != nil {
			log.Fatalf("failed to open image: %v", err)
			return DocumentType{}, err
		}

		newPagePath, err := changeFilePathAndType(page, tempDocumentFolderPath, "jpg")
		if err != nil {
			log.Fatalf("failed to change file path and type: %v", err)
			return DocumentType{}, err
		}

		err = imaging.Save(img, newPagePath, imaging.JPEGQuality(90))
		if err != nil {
			log.Fatalf("failed to save image: %v", err)
			return DocumentType{}, err
		}

		pagesInJPG = append(pagesInJPG, newPagePath)
		taskProgress.UpdateTaskProgress(taskId, taskProgress.Tasks[taskId].Progress+incrementValue)
	}

	documentInJPG.pagesPath = pagesInJPG
	documentInJPG.documentPath = tempDocumentFolderPath

	return documentInJPG, nil
}

func ConvertJpgsToPdf(documentToProcess DocumentType, temporaryFolder string, taskId int, taskProgress *OverallProgress, simplePdf bool, outputFolder string) (pdfPath string, error error) {
	tempDocumentFolderPath := filepath.Join(temporaryFolder, "pdfs_simple")
	err := os.Mkdir(tempDocumentFolderPath, 0755)
	if err != nil {
		log.Fatalf("failed to create directory: %v", err)
		return "", err
	}

	documentFilePath := filepath.Join(tempDocumentFolderPath, documentToProcess.documentName+".pdf")

	imp, _ := api.Import("form:A4, pos:c, s:1.0", types.POINTS)
	api.ImportImagesFile(documentToProcess.pagesPath, documentFilePath, imp, nil)

	if simplePdf {
		simpleFilePath := filepath.Join(outputFolder, documentToProcess.documentName+"_simple.pdf")
		api.ImportImagesFile(documentToProcess.pagesPath, simpleFilePath, imp, nil)
	}

	taskProgress.UpdateTaskProgress(taskId, 100)

	return documentFilePath, nil
}
