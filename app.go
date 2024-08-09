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
	totalTasks      int
	tasksProgress   goScripts.OverallProgress
	documents       []goScripts.DocumentType
	temporaryFolder string
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

	a.frontendData = AppDataType{
		inputFolderPath:      "",
		outputFolderPath:     "",
		hasMultipleDocs:      false,
		doNotCreateSimplePdf: false,
	}
	a.totalTasks = 0
	a.tasksProgress = goScripts.OverallProgress{
		TotalProgress: 0,
		CurrentTask:   0,
		Tasks:         []goScripts.TaskProgress{},
	}
	a.documents = []goScripts.DocumentType{}

	temporaryFolder, errTemporaryFolder := os.MkdirTemp("", "tempFolder")
	if errTemporaryFolder != nil {
		log.Fatal(errTemporaryFolder)
	}
	a.temporaryFolder = temporaryFolder
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
	/* check which latest task is not completed */
	for i := len(a.tasksProgress.Tasks) - 1; i >= 0; i-- {
		if a.tasksProgress.Tasks[i].Progress < 100 {
			return a.tasksProgress.Tasks[i].Progress
		}
	}
	return 0
}

/* return value of total progress of every task */
func (a *App) GetTotalProgress() float64 {
	return 0
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

	var err error
	a.documents, err = goScripts.GatherDocuments(a.frontendData.inputFolderPath, a.frontendData.hasMultipleDocs)
	if err != nil {
		fmt.Println(err)
		return
	}

	/* Loop through documents */
	for i, document := range a.documents {
		/* Convert pages to JPG into temporary folder */
		taskName := "A converter páginas para JPG - " + document.GetDocumentName()
		taskidConvertToJpg := a.tasksProgress.AddTask(taskName)
		var tempDoc goScripts.DocumentType
		tempDoc, err = goScripts.ConvertDocumentPagesToJpg(document, a.temporaryFolder, taskidConvertToJpg, &a.tasksProgress)
		a.documents[i] = tempDoc
	}

	/* Convert images to PDF */
	var pdfPaths []string
	for _, document := range a.documents {
		/* Convert pages to PDF */
		taskName := "A converter páginas para PDF - " + document.GetDocumentName()
		taskidConvertToPdf := a.tasksProgress.AddTask(taskName)

		var pdfPath string
		pdfPath, err = goScripts.ConvertJpgsToPdf(document, a.temporaryFolder, taskidConvertToPdf, &a.tasksProgress)
		pdfPaths = append(pdfPaths, pdfPath)
	}

	/* print pdfPaths */
	fmt.Println(pdfPaths)

	/* Convert PDF to OCR */

	/* Chime in Success with Alert Window */

	/* Reset Values */
}
