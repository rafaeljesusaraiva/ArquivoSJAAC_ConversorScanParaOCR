package main

import (
	"context"
	"fmt"
	"github.com/pkg/browser"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"os/exec"
	sysruntime "runtime"
)

func checkRuntime() string {
	switch sysruntime.GOOS {
	case "darwin", "linux", "windows":
		return sysruntime.GOOS
	default:
		return "Unsupported platform"
	}
}

func isPythonInstalled() bool {
	var cmd string
	switch checkRuntime() {
	case "darwin", "linux":
		cmd = "which python3"
	case "windows":
		cmd = "where python3"
	default:
		fmt.Printf("Unsupported platform: %s\n", sysruntime.GOOS)
		return false
	}

	_, err := exec.Command("cmd", "/C", cmd).Output()
	if sysruntime.GOOS != "windows" {
		_, err = exec.Command("sh", "-c", cmd).Output()
	}

	return err == nil
}

func isOcrmypdfInstalled() bool {
	var cmd string
	switch checkRuntime() {
	case "darwin", "linux":
		cmd = "which ocrmypdf"
	case "windows":
		cmd = "where ocrmypdf"
	default:
		fmt.Printf("Unsupported platform: %s\n", sysruntime.GOOS)
		return false
	}

	_, err := exec.Command("cmd", "/C", cmd).Output()
	if sysruntime.GOOS != "windows" {
		_, err = exec.Command("sh", "-c", cmd).Output()
	}

	return err == nil
}

func CheckDependencies(ctx context.Context) {
	// Check which OS is running, can be windows, linux or darwin
	if !isPythonInstalled() {
		selection, _ := runtime.MessageDialog(ctx, runtime.MessageDialogOptions{
			Title:         "Dependências não encontradas",
			Message:       "Para executar esta aplicação, é necessário instalar o Python (python3) e o OCRmyPDF.",
			Buttons:       []string{"Fechar Aplicação", "Instalar Python", "Instalar OCRmyPDF"},
			DefaultButton: "Instalar Python",
		})
		if selection == "Instalar Python" {
			browser.OpenURL("https://www.python.org/downloads/")
			runtime.Quit(ctx)
		} else if selection == "Instalar OCRmyPDF" {
			switch checkRuntime() {
			case "macOS":
				browser.OpenURL("https://ocrmypdf.readthedocs.io/en/latest/installation.html#installing-on-macos")
			case "linux":
				browser.OpenURL("https://ocrmypdf.readthedocs.io/en/latest/installation.html#installing-on-linux")
			case "windows":
				browser.OpenURL("https://ocrmypdf.readthedocs.io/en/latest/installation.html#installing-on-windows")
			}
		} else {
			runtime.Quit(ctx)
		}
	}

	if !isOcrmypdfInstalled() {
		selection, _ := runtime.MessageDialog(ctx, runtime.MessageDialogOptions{
			Title:         "Dependências não encontradas",
			Message:       "Para executar esta aplicação, é necessário instalar o OCRmyPDF.",
			Buttons:       []string{"Fechar Aplicação", "Instalar OCRmyPDF"},
			DefaultButton: "Instalar OCRmyPDF",
		})
		if selection == "Instalar OCRmyPDF" {
			switch checkRuntime() {
			case "macOS":
				browser.OpenURL("https://ocrmypdf.readthedocs.io/en/latest/installation.html#installing-on-macos")
			case "linux":
				browser.OpenURL("https://ocrmypdf.readthedocs.io/en/latest/installation.html#installing-on-linux")
			case "windows":
				browser.OpenURL("https://ocrmypdf.readthedocs.io/en/latest/installation.html#installing-on-windows")
			}
		} else {
			runtime.Quit(ctx)
		}
	}
}
