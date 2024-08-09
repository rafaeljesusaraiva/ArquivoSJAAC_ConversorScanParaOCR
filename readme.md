# Conversor Scan para PDF+OCR

# README

## About

This template comes with Vite, React, TypeScript, TailwindCSS and shadcn/ui.

Built with `Wails v2.5.1` and [shadcn's CLI](https://ui.shadcn.com/docs/cli)

### Using the Template
```console
wails init -n project-name -t https://github.com/Mahcks/wails-vite-react-tailwind-shadcnui-ts
```

```console
cd frontend
```

```console
npm install
```

### Installing Components
To install components, use shadcn's CLI tool to install

More info here: https://ui.shadcn.com/docs/cli#add

Example:
```console
npx shadcn-ui@latest add [component]
```

## Live Development

To run in live development mode, run `wails dev` in the project directory. In another terminal, go into the `frontend`
directory and run `npm run dev`. The frontend dev server will run on http://localhost:34115. Connect to this in your
browser and connect to your application.

## Building

To build a redistributable, production mode package, use `wails build`.


### A FAZER
- Listagem de ficheiros da pasta selecionada
- Pré-visualização de ficheiros .TIFF
- Pré-visualização de PDFs tratados (sem OCR)
- Opção de janela com consola de logs
- Implementar sinais e eventos para transmissão de dados entre widgets
- Multithreading em pyqt6 ???

### Pesquisa

---

## Configuração antiga (python)

### Criar ambiente python:
``` bash
python3 -m venv arquivoCabraPythonEnv 
source arquivoCabraPythonEnv/bin/activate
pip3 install configparser tk pillow
```

### Inicializar ambiente python:
```bash
source arquivoCabraPythonEnv/bin/activate
```

### Multithreads & Multiprocessing

https://medium.com/codex/an-introduction-to-multiprocessing-using-python-165f51f83c0d





 