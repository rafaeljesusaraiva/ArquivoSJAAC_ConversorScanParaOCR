import Head from "@/components/Head.tsx";
import Body from "@/components/Body.tsx";
import {ThemeProvider} from "@/components/Theme-Provider.tsx";

import {StepSection} from "@/components/StepSection"

import {Checkbox} from "@/components/ui/checkbox"
import {Label} from "@/components/ui/label"
import {Input} from "@/components/ui/input"
import {Button} from "@/components/ui/button"

import React from "react"
import {Progress} from "@/components/ui/progress"

import JSConfetti from 'js-confetti'
import logoSJAAC from "@/assets/SJAAC-logo.jpg";

import {ChimeEndTask, OpenDialog} from "../wailsjs/go/main/App";


function App() {
    const [progress, _] = React.useState(13)

    const [checkOne, setCheckOne] = React.useState(false)
    const [inputFolderDirectory, setInputFolderDirectory] = React.useState("")

    const [checkFolderMultipleDocuments, setCheckFolderMultipleDocuments] = React.useState(0)
    const flipCheckFolderMultipleDocuments = () =>
        setCheckFolderMultipleDocuments(checkFolderMultipleDocuments === 0 ? 1 : 0)

    const [checkTwo, setCheckTwo] = React.useState(false)
    const [inputFolderDirectoryTwo, setInputFolderDirectoryTwo] = React.useState("")

    const [check]

    const [spooked, setSpooked] = React.useState(false);
    const partyArquivo = () => {
        setSpooked(true);
        setTimeout(() => setSpooked(false), 1000); // reset spooked state after 1s
        const confetti = new JSConfetti()
        confetti.addConfetti({
            emojis: ['🖨️', '🖥️', '🎞️', '💾', '📼'],
            emojiSize: 90,
            confettiNumber: 20,
        })
        confetti.addConfetti()
    }

    const internalOpenDialog = async () => {
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

    return (
        <ThemeProvider defaultTheme="dark" storageKey="vite-ui-theme">
            <Head/>
            {/* Main Grid for UI Elements */}
            <Body className="grid grid-cols-5 grid-rows-7 py-4 mx-8 gap-2">
                {/* Left grid side for user interaction, right side for logo */}
                <StepSection title={"1º Passo - Escolher pasta com scans do documento"} minimal={true}>
                    {/* First UI Line with elements*/}
                    <div className={"grid grid-cols-12 items-center gap-2"}>
                        <Checkbox checked={checkOne} disabled={true} name={"inputFolderCheck"}
                                  className={"col-span-1 scale-150 mx-auto"}/>
                        <Input name={"showInputFolderLocation"} value={inputFolderDirectory} disabled={true}
                               className={"col-span-8"}/>
                        <Button className={"col-span-3 text-center"} onClick={internalOpenDialog}>Escolher
                            Pasta</Button>
                    </div>
                    {/* Second UI Line */}
                    <div className={"grid grid-cols-12 items-center gap-2 mt-4"}>
                        <Checkbox name={"folderWithMultipleDocuments"} id={"check01"}
                                  className={"col-span-1 scale-150 mx-auto"}
                                  value={checkFolderMultipleDocuments}
                                  onClick={flipCheckFolderMultipleDocuments}/>
                        <Label htmlFor="check01" className={"col-span-8"}>
                            Pasta com vários documentos
                        </Label>
                        <Button className={"col-span-3"}>Pré-visualizar Imagens</Button>
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
                    <Checkbox checked={false} disabled={true} name={"inputFolderCheck"}
                              className={"col-span-1 scale-150 mx-auto"}/>
                    <Input name={"showInputFolderLocation"} disabled={true} className={"col-span-8"}/>
                    <Button className={"col-span-3 text-center"}>Escolher Pasta</Button>
                </StepSection>
                {/* Third Step */}
                <StepSection title={"3º Passo (opcional) - Escolher extras"}>
                    <Checkbox name={"dontCreateNoOcrPdf"} id={"check02"}
                              className={"col-span-1 scale-150 mx-auto"}/>
                    <Label htmlFor="check02" className={"col-span-3"}>
                        Não criar PDF sem OCR
                    </Label>
                    <Checkbox name={"chimeOnFinish"} id={"check03"}
                              className={"col-span-1 scale-150 mx-auto"}/>
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
                                Progresso da tarefa atual:
                            </h2>
                        </div>
                        <div className={"col-span-full relative"}>
                            <Progress value={progress} className={"h-6"}/>
                            <div className={"absolute"}
                                 style={{
                                     top: "50%",
                                     left: "50%",
                                     translate: "-50% -50%",
                                     filter: "invert(1)"
                                 }}>{progress}%
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
                            <Progress value={progress} className={"h-6"}/>
                            <div className={"absolute"}
                                 style={{
                                     top: "50%",
                                     left: "50%",
                                     translate: "-50% -50%",
                                     filter: "invert(1)"
                                 }}>{progress}%
                            </div>
                        </div>
                    </div>
                    <div className={"grid grid-cols-5 gap-4 mt-4"}>
                        <Button className={"col-span-3"} onClick={ChimeEndTask}>Iniciar Processo</Button>
                        <Button className={"col-span-1"}>Pausar</Button>
                        <Button className={"col-span-1"}>Cancelar</Button>
                    </div>
                </StepSection>
            </Body>
        </ThemeProvider>
    )
}

export default App
