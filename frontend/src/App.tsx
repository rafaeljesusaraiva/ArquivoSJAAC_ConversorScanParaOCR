import Head from "@/components/Head.tsx";
import Body from "@/components/Body.tsx";
import {ThemeProvider} from "@/components/Theme-Provider.tsx";

import {Checkbox} from "@/components/ui/checkbox"
import {Label} from "@/components/ui/label"
import {Input} from "@/components/ui/input"
import {Button} from "@/components/ui/button"

import React from "react"
import {Progress} from "@/components/ui/progress"

import logoSJAAC from "@/assets/SJAAC-logo.jpg";

import { ChimeEndTask } from "../wailsjs/go/main/App";

interface StepSectionProps {
    title: string;
    children: React.ReactNode;
    minimal?: boolean;
    end?: boolean;
}

const StepSection: React.FC<StepSectionProps> = ({title, children, minimal, end}) => {
    var mainClass = minimal ? "col-span-4" : "col-span-full"

    return (
        <div className={mainClass}>
            <div className={"grid grid-cols-5"}>
                <div className={"col-span-full"}>
                    <h2 className={"scroll-m-20 border-b pb-2 text-xl tracking-tight first:mt-0"}>
                        {title}
                    </h2>
                </div>
                <div className={"col-span-full m-4"}>
                    {end ? children : (
                        minimal ? children : (
                            <div className={"grid grid-cols-12 items-center gap-2"}>
                                {children}
                            </div>
                        )
                    )}
                </div>
            </div>
        </div>
    )
}

function App() {
    const [progress, _] = React.useState(13)

    return (
        <ThemeProvider defaultTheme="dark" storageKey="vite-ui-theme">
            <Head/>
            {/* Main Grid for UI Elements */}
            <Body className="grid grid-cols-5 grid-rows-7 py-4 mx-8 gap-2">
                {/* Left grid side for user interaction, right side for logo */}
                <StepSection title={"1º Passo - Escolher pasta com scans do documento"} minimal={true}>
                    {/* First UI Line with elements*/}
                    <div className={"grid grid-cols-12 items-center gap-2"}>
                        <Checkbox checked={false} disabled={true} name={"inputFolderCheck"}
                                  className={"col-span-1 scale-150 mx-auto"}/>
                        <Input name={"showInputFolderLocation"} disabled={true} className={"col-span-8"}/>
                        <div className={"col-span-3 border rounded-lg py-2 px-4 text-center"}>
                            <Label htmlFor="input01">Escolher Pasta</Label>
                            <Input name={"inputFolderLocation"} type={"file"} id={"input01"}
                                   className={"hidden"}/>
                        </div>
                    </div>
                    {/* Second UI Line */}
                    <div className={"grid grid-cols-12 items-center gap-2 mt-4"}>
                        <Checkbox name={"folderWithMultipleDocuments"} id={"check01"}
                                  className={"col-span-1 scale-150 mx-auto"}/>
                        <Label htmlFor="check01" className={"col-span-8"}>
                            Pasta com vários documentos
                        </Label>
                        <Button className={"col-span-3"}>Pré-visualizar Imagens</Button>
                    </div>
                </StepSection>
                {/* Image Logo */}
                <div className={"col-span-1"}>
                    <img src={logoSJAAC} alt="Image" className="rounded-md object-cover mx-auto"/>
                </div>
                {/* Second Step */}
                <StepSection title="2º Passo - Escolher pasta de destino do(s) PDF(s)">
                    <Checkbox checked={false} disabled={true} name={"inputFolderCheck"}
                              className={"col-span-1 scale-150 mx-auto"}/>
                    <Input name={"showInputFolderLocation"} disabled={true} className={"col-span-8"}/>
                    <div className={"col-span-3 border rounded-lg py-2 px-4 text-center"}>
                        <Label htmlFor="input01">Escolher Pasta</Label>
                        <Input name={"inputFolderLocation"} type={"file"} id={"input01"}
                               className={"hidden"}/>
                    </div>
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
