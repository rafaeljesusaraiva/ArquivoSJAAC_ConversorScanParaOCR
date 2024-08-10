import Head from "@/components/Head.tsx";
import Body from "@/components/Body.tsx";
import {ThemeProvider} from "@/components/Theme-Provider.tsx";

import {StepSection} from "@/components/StepSection"

import {Checkbox} from "@/components/ui/checkbox"
import {Label} from "@/components/ui/label"
import {Input} from "@/components/ui/input"
import {Button} from "@/components/ui/button"

import React, {useEffect} from "react"
import {Progress} from "@/components/ui/progress"

import JSConfetti from 'js-confetti'
import logoSJAAC from "@/assets/SJAAC-logo.jpg";

import {ChimeEndTask, IsConversionRunning, GetTaskProgress, GetTaskName, GetTotalProgress, OpenDialog, ProcessBegin, StopConversion, ToggleChime, ToggleSimplePdf} from "../wailsjs/go/main/App";

function blowConfetti() {
    const confetti = new JSConfetti()
    confetti.addConfetti({
        emojis: ['🖨️', '🖥️', '🎞️', '💾', '📼'],
        emojiSize: 90,
        confettiNumber: 20,
    })
    confetti.addConfetti()
}

function App() {
    const [progressMain, setProgressMain] = React.useState(0)
    const [progressTask, setProgressTask] = React.useState(0)
    const [currentTaskName, setCurrentTaskName] = React.useState("")
    const [isConversionRunning, setIsConversionRunning] = React.useState(false)

    const [checkOne, setCheckOne] = React.useState(false)
    const [inputFolderDirectory, setInputFolderDirectory] = React.useState("")

    const [checkFolderMultipleDocuments, setCheckFolderMultipleDocuments] = React.useState(false)
    const flipCheckFolderMultipleDocuments = () =>
        setCheckFolderMultipleDocuments(!checkFolderMultipleDocuments)

    const [checkTwo, setCheckTwo] = React.useState(false)
    const [inputFolderDirectoryTwo, setInputFolderDirectoryTwo] = React.useState("")

    const [checkNoPdfWithoutOcr, setCheckNoPdfWithoutOcr] = React.useState(false)
    const flipCheckNoPdfWithoutOcr = async () => {
        let value = await ToggleSimplePdf();
        setCheckNoPdfWithoutOcr(value)
    }

    const [chimeOnFinish, setChimeOnFinish] = React.useState(true)
    const flipChimeOnFinish = async () => {
        let value = await ToggleChime();
        setChimeOnFinish(value)
    }

    const [spooked, setSpooked] = React.useState(false);
    const partyArquivo = () => {
        setSpooked(true);
        setTimeout(() => setSpooked(false), 1000); // reset spooked state after 1s
        blowConfetti()
        ChimeEndTask()
    }
    const firstOpenDialog = async () => {
        try {
            var directory = await OpenDialog("Escolher pasta")

            if (directory !== "") {
                setCheckOne(true)
                setInputFolderDirectory(directory)
            } else {
                setCheckOne(false)
                setInputFolderDirectory("Exemplo: C:/Users/Arquivo/...")
            }
        } catch (error) {
            console.error(error)
        }
    }
    const secondOpenDialog = async () => {
        try {
            var directory = await OpenDialog("Escolher pasta")

            if (directory !== "") {
                setCheckTwo(true)
                setInputFolderDirectoryTwo(directory)
            } else {
                setCheckTwo(false)
                setInputFolderDirectoryTwo("Exemplo: C:/Users/Arquivo/...")
            }
        } catch (error) {
            console.error(error)
        }
    }

    const startProcess = async () => {
        try {
            await ProcessBegin(inputFolderDirectory, inputFolderDirectoryTwo, checkFolderMultipleDocuments, checkNoPdfWithoutOcr)
        } catch (error) {
            console.error(error)
        }
    }
    const stopProcess = async () => {
        try {
            await StopConversion()
        } catch (error) {
            console.error(error)
        }
    }

    useEffect(() => {
        const interval = setInterval(async () => {
            const taskProgress = await GetTaskProgress() | 0;
            const mainProgress = await GetTotalProgress() | 0;
            const taskName = await GetTaskName() || "";
            const isRunning = await IsConversionRunning() || false;
            setProgressTask(taskProgress);
            setProgressMain(mainProgress);
            setCurrentTaskName(taskName);
            setIsConversionRunning(isRunning);
        }, 500);

        // Clearing the interval
        return () => clearInterval(interval);
    }, [progressTask, progressMain]);

    return (
        <ThemeProvider defaultTheme="dark" storageKey="vite-ui-theme">
            <Head/>
            {/* Main Grid for UI Elements */}
            <Body className="grid grid-cols-5 grid-rows-7 py-4 mx-8 gap-2">
                {/* Left grid side for user interaction, right side for logo */}
                <StepSection title={"1º Passo - Escolher pasta com scans do documento"} minimal={true}>
                    {/* First UI Line with elements*/}
                    <div className={"grid grid-cols-12 items-center gap-2"}>
                        <Checkbox checked={checkOne} disabled={true}
                                  className={"col-span-1 scale-150 mx-auto"}/>
                        <Input name={"showInputFolderLocation"} value={inputFolderDirectory} disabled={true}
                               className={"col-span-8"}/>
                        <Button className={"col-span-3 text-center"} onClick={firstOpenDialog}>Escolher
                            Pasta</Button>
                    </div>
                    {/* Second UI Line */}
                    <div className={"grid grid-cols-12 items-center gap-2 mt-4"}>
                        <Checkbox className={"col-span-1 scale-150 mx-auto"} id={"check01"}
                                  checked={checkFolderMultipleDocuments}
                                  onClick={flipCheckFolderMultipleDocuments}/>
                        <Label htmlFor="check01" className={"col-span-8"}>
                            Pasta com vários documentos
                        </Label>
                        <Button className={"col-span-3 line-through text-xs"} disabled={true}>Pré-visualizar Imagens</Button>
                    </div>
                </StepSection>
                {/* Image Logo */}
                <div className={"col-span-1"}>
                    <img src={logoSJAAC} alt="Image"
                         className={spooked ? 'spooked rounded-md object-cover mx-auto' : 'rounded-md object-cover mx-auto'}
                         onClick={partyArquivo}/>
                </div>
                {/* Second Step */}
                <StepSection title="2º Passo - Escolher pasta de destino do(s) PDF(s)">
                    <Checkbox checked={checkTwo} disabled={true} name={"inputFolderCheck"}
                              className={"col-span-1 scale-150 mx-auto"}/>
                    <Input name={"showInputFolderLocation"} value={inputFolderDirectoryTwo} disabled={true} className={"col-span-8"}/>
                    <Button className={"col-span-3 text-center"} onClick={secondOpenDialog}>Escolher Pasta</Button>
                </StepSection>
                {/* Third Step */}
                <StepSection title={"3º Passo (opcional) - Escolher extras"}>
                    <Checkbox id={"check02"}
                              checked={checkNoPdfWithoutOcr}
                              onClick={flipCheckNoPdfWithoutOcr}
                              className={"col-span-1 scale-150 mx-auto"}/>
                    <Label htmlFor="check02" className={"col-span-3"}>
                        Criar PDF sem OCR
                    </Label>
                    <Checkbox id={"check03"}
                              checked={chimeOnFinish}
                              onClick={flipChimeOnFinish}
                              className={"col-span-1 scale-150 mx-auto"}
                    />
                    <Label htmlFor="check03" className={"col-span-3"}>
                        Dar aviso sonoro ao terminar
                    </Label>
                </StepSection>
                <div className={"col-span-full"}></div>
                <div className={"col-span-full"}></div>
                {/* Fourth Step */}
                <StepSection title={"4º Passo - Iniciar processo"} end={true}>
                    <div className={"grid grid-cols-5"}>
                        <div className={"col-span-full"}>
                            <h2 className={"scroll-m-20 border-b pb-2 tracking-tight first:mt-0"}>
                                Progresso da tarefa atual: {currentTaskName}
                            </h2>
                        </div>
                        <div className={"col-span-full relative"}>
                            <Progress value={progressTask} className={"h-6"}/>
                            <div className={"absolute"}
                                 style={{
                                     top: "50%",
                                     left: "50%",
                                     translate: "-50% -50%",
                                     filter: "invert(1)"
                                 }}>{progressTask}%
                            </div>
                        </div>
                    </div>
                    <div className={"grid grid-cols-5 mt-4"}>
                        <div className={"col-span-full"}>
                            <h2 className={"scroll-m-20 border-b pb-2 tracking-tight first:mt-0"}>
                                Progresso geral:
                            </h2>
                        </div>
                        <div className={"col-span-full relative"}>
                            <Progress value={progressMain} className={"h-6"}/>
                            <div className={"absolute"}
                                 style={{
                                     top: "50%",
                                     left: "50%",
                                     translate: "-50% -50%",
                                     filter: "invert(1)"
                                 }}>{progressMain}%
                            </div>
                        </div>
                    </div>
                    <div className={"grid grid-cols-5 gap-4 mt-4"}>
                        <Button className={"col-span-3"} onClick={startProcess} disabled={isConversionRunning}>Iniciar Processo</Button>
                        <Button className={"col-span-1"} disabled={!isConversionRunning}>Pausar</Button>
                        <Button className={"col-span-1"} onClick={stopProcess} disabled={!isConversionRunning}>Cancelar</Button>
                    </div>
                </StepSection>
            </Body>
        </ThemeProvider>
    )
}

export default App
