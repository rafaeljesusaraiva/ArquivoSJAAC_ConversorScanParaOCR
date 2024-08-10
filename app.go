package main

import (
	"ConversorScanParaOCR/goScripts"
	"context"
	"fmt"
	"github.com/gopxl/beep/v2"
	"github.com/gopxl/beep/v2/mp3"
	"github.com/gopxl/beep/v2/speaker"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"log"
	"os"
	"time"
)

// App struct
type App struct {
	ctx             context.Context
	frontendData    AppDataType
	tasksProgress   goScripts.OverallProgress
	documents       []goScripts.DocumentType
	temporaryFolder string
	isRunning       bool
	stopChannel     chan bool
	userChime       bool
	userSimplePdf   bool
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called at application startup
func (a *App) startup(ctx context.Context) {
	// Perform your setup here
	// Check if python and ocrmypdf are installed depending on the OS
	CheckDependencies(ctx)

	// App variables
	a.ctx = ctx

	a.stopChannel = make(chan bool)

	a.userChime = true
	a.userSimplePdf = false

	a.frontendData = AppDataType{
		inputFolderPath:      "",
		outputFolderPath:     "",
		hasMultipleDocs:      false,
		doNotCreateSimplePdf: false,
	}
	a.tasksProgress = goScripts.OverallProgress{
		Tasks:      []goScripts.TaskProgress{},
		TotalTasks: 0,
	}
	a.documents = []goScripts.DocumentType{}

	temporaryFolder, errTemporaryFolder := os.MkdirTemp("", "tempFolder")
	if errTemporaryFolder != nil {
		log.Fatal(errTemporaryFolder)
	}
	a.temporaryFolder = temporaryFolder
	a.isRunning = false
}

// domReady is called after front-end resources have been loaded
func (a App) domReady(ctx context.Context) {
	// Add your action here
}

// beforeClose is called when the application is about to quit,
// either by clicking the window close button or calling runtime.Quit.
// Returning true will cause the application to continue, false will continue shutdown as normal.
func (a *App) beforeClose(ctx context.Context) (prevent bool) {
	/* delete contents in temporary folder created on startup */
	goScripts.CleanupTemporaryFolder(a.temporaryFolder)
	return false
}

// shutdown is called at application termination
func (a *App) shutdown(ctx context.Context) {
	// Perform your teardown here
}

func (a *App) ChimeEndTask() {
	f, err := os.Open("audio/completedAudio.mp3")
	if err != nil {
		log.Fatal(err)
	}

	streamer, format, err := mp3.Decode(f)
	if err != nil {
		log.Fatal(err)
	}
	defer streamer.Close()
	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

	done := make(chan bool)
	speaker.Play(beep.Seq(streamer, beep.Callback(func() {
		done <- true
	})))
	<-done
}

func (a *App) ToggleChime() bool {
	a.userChime = !a.userChime
	return a.userChime
}

func (a *App) ToggleSimplePdf() bool {
	a.userSimplePdf = !a.userSimplePdf
	return a.userSimplePdf
}

/* function that opens system dialog and returns folder path selected */
func (a *App) OpenDialog(dialogTitle string) string {
	dialogOptions := runtime.OpenDialogOptions{
		Title: dialogTitle,
	}
	directory, _ := runtime.OpenDirectoryDialog(a.ctx, dialogOptions)

	/* check if error */
	if directory == "" {
		return ""
	}

	return directory
}

/* return value of current task running */
func (a *App) GetTaskProgress() float64 {
	if !a.isRunning {
		return 0
	}
	/* check which latest task is not completed */
	for i := len(a.tasksProgress.Tasks) - 1; i >= 0; i-- {
		if a.tasksProgress.Tasks[i].Progress < 100 {
			return a.tasksProgress.Tasks[i].Progress
		}
	}
	return 0
}

func (a *App) GetTaskName() string {
	if !a.isRunning {
		return ""
	}
	/* check which latest task is not completed */
	for i := len(a.tasksProgress.Tasks) - 1; i >= 0; i-- {
		if a.tasksProgress.Tasks[i].Progress < 100 {
			return a.tasksProgress.Tasks[i].Name
		}
	}
	return ""
}

/* return value of total progress of every task */
func (a *App) GetTotalProgress() float64 {
	if !a.isRunning {
		return 0
	}
	return a.tasksProgress.CalculateMainProgress()
}

func (a *App) IsConversionRunning() bool {
	return a.isRunning
}

func (a *App) StopConversion() {
	a.stopChannel <- true
}

type AppDataType struct {
	inputFolderPath      string
	outputFolderPath     string
	hasMultipleDocs      bool
	doNotCreateSimplePdf bool
}

func (a *App) ProcessBegin(inputFolderDirectory string, outputFolderDirectory string, hasMultipleDocs bool, doNotCreateSimplePdf bool) {
	a.frontendData.inputFolderPath = inputFolderDirectory
	a.frontendData.outputFolderPath = outputFolderDirectory
	a.frontendData.hasMultipleDocs = hasMultipleDocs
	a.frontendData.doNotCreateSimplePdf = doNotCreateSimplePdf

	a.isRunning = true

	// Check if inputFolderPath and outputFolderPath are not empty strings
	if a.frontendData.inputFolderPath == "" || a.frontendData.outputFolderPath == "" {
		// show error messagewindow
		runtime.MessageDialog(a.ctx, runtime.MessageDialogOptions{
			Title:   "Erro",
			Message: "Por favor, selecione uma pasta com scans e uma pasta para guardar o PDF.",
		})
		return
	}

	var err error
	a.documents, err = goScripts.GatherDocuments(a.frontendData.inputFolderPath, a.frontendData.hasMultipleDocs)
	if err != nil {
		fmt.Println(err)
		return
	}

	a.tasksProgress.TotalTasks = len(a.documents) * 3

	// Loop through documents
	for i, document := range a.documents {
		// Convert pages to JPG into temporary folder
		taskName := "A converter páginas para JPG - " + document.GetDocumentName()
		taskidConvertToJpg := a.tasksProgress.AddTask(taskName)
		var tempDoc goScripts.DocumentType
		tempDoc, err = goScripts.ConvertDocumentPagesToJpg(document, a.temporaryFolder, taskidConvertToJpg, &a.tasksProgress)
		a.documents[i] = tempDoc
	}

	// Convert images to PDF
	var pdfPaths []string
	for _, document := range a.documents {
		taskName := "A converter páginas para PDF - " + document.GetDocumentName()
		taskidConvertToPdf := a.tasksProgress.AddTask(taskName)

		var pdfPath string
		pdfPath, err = goScripts.ConvertJpgsToPdf(document, a.temporaryFolder, taskidConvertToPdf, &a.tasksProgress, a.userSimplePdf, a.frontendData.outputFolderPath)
		pdfPaths = append(pdfPaths, pdfPath)
	}

	// print pdfPaths
	// fmt.Println(pdfPaths)

	// Convert PDF to OCR using ocrmypdf by calling external shell
	// loop through pdfPaths
	for i, pdfPath := range pdfPaths {
		taskName := "A converter PDF para OCR - " + a.documents[i].GetDocumentName()
		taskidConvertToOcr := a.tasksProgress.AddTask(taskName)

		outputFilePath := a.frontendData.outputFolderPath + "/" + a.documents[i].GetDocumentName() + ".pdf"

		err = goScripts.ConvertPdfToOcr(a.stopChannel, pdfPath, outputFilePath, taskidConvertToOcr, &a.tasksProgress)
	}

	// Chime in Success with Alert Window
	if a.userChime {
		a.ChimeEndTask()
		runtime.MessageDialog(a.ctx, runtime.MessageDialogOptions{
			Title:   "Sucesso",
			Message: "Processo concluído com sucesso.",
		})
	}

	a.isRunning = false

	// Reset Values
	a.frontendData = AppDataType{
		inputFolderPath:      "",
		outputFolderPath:     "",
		hasMultipleDocs:      false,
		doNotCreateSimplePdf: false,
	}
	a.tasksProgress.TotalTasks = 0
	a.tasksProgress.Reset()

	/* delete contents in temporary folder created on startup */
	goScripts.CleanupTemporaryFolder(a.temporaryFolder)

}
