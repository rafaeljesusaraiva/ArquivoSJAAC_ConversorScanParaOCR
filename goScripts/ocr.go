package goScripts

import (
	"bufio"
	"fmt"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

func ConvertPdfToOcr(stopChannel chan bool, pdfPath string, outputFilePath string, taskId int, taskProgress *OverallProgress) error {
	cmd := exec.Command("ocrmypdf", "--output-type", "pdf", "--rotate-pages", "--deskew", "--force-ocr", "-l", "por", pdfPath, outputFilePath)

	stderr, _ := cmd.StderrPipe()
	_ = cmd.Start()

	scanner := bufio.NewScanner(stderr)
	percentRe := regexp.MustCompile(`(\d+)%`) // Regular expression to match percentages
	scanningProgress := 0
	ocrProgress := 0
	deflatingProgress := 0
	for scanner.Scan() {
		m := scanner.Text()
		fmt.Println(m)

		// Check for stop signal
		select {
		case <-stopChannel:
			// Stop the command
			_ = cmd.Process.Kill()
			return nil
		default:
			// Continue executing the command
		}

		// Check if the line starts with "Scanning contents"
		if strings.HasPrefix(m, "Start processing") {
			// Extract percentage from the line
			match := percentRe.FindStringSubmatch(m)
			if len(match) > 1 {
				scanningProgress, _ = strconv.Atoi(match[1])
				// Update task progress (25% of main task progress)
				taskProgress.UpdateTaskProgress(taskId, float64(scanningProgress/4))
			}
		}

		// Check if the line starts with "OCR"
		if strings.HasPrefix(m, "OCR") {
			// Extract percentage from the line
			match := percentRe.FindStringSubmatch(m)
			if len(match) > 1 {
				ocrProgress, _ = strconv.Atoi(match[1])
				// Update task progress (50% of main task progress, offset by 25%)
				taskProgress.UpdateTaskProgress(taskId, float64(25+ocrProgress/2))
			}
		}

		// Check if "Recompressing JPEGs" has shown up
		if strings.HasPrefix(m, "Recompressing JPEGs") {
			// Update task progress (75% of main task progress)
			taskProgress.UpdateTaskProgress(taskId, 75)
		}

		// Check if the line starts with "Deflating JPEGs"
		if strings.HasPrefix(m, "Deflating JPEGs") {
			// Extract percentage from the line
			match := percentRe.FindStringSubmatch(m)
			if len(match) > 1 {
				deflatingProgress, _ = strconv.Atoi(match[1])
				// Update task progress (12.5% of main task progress, offset by 75%)
				taskProgress.UpdateTaskProgress(taskId, float64(75+deflatingProgress/8))
			}
		}

		// Check if "JBIG2" has shown up
		if strings.HasPrefix(m, "JBIG2") {
			// Update task progress (87.5% of main task progress)
			taskProgress.UpdateTaskProgress(taskId, 87.5)
		}

		// Check if "Total file size ratio" has shown up
		if strings.HasPrefix(m, "Total file size ratio") {
			// Update task progress (100% of main task progress)
			taskProgress.UpdateTaskProgress(taskId, 100)
		}
	}
	_ = cmd.Wait()

	return nil
}
