import tkinter as tk
from tkinter import filedialog
from tkinter import ttk
import modules.windows.showFiles as showFiles

CHECK_ICON = "✔"
CROSS_ICON = "✘"

class MainWindow:
    def __init__(self, root):
        self.root = root
        self.stepOne_directory = None  # variable to store the selected folder
        self.stepOne_multipleDocs = tk.BooleanVar()  # variable to store if the user wants to convert multiple documents in the same folder
        self.stepTwo_directory = None
        self.stepThree_notCreateSimplePDF = tk.BooleanVar()
        self.stepThree_soundOnFinish = tk.BooleanVar()
        self.show()

    def show(self):
         # Setup framing for each group of widgets
        ## Frame for first step
        self.StepOne_FrameMain = tk.Frame(self.root)
        self.StepOne_FrameMain.pack(fill=tk.X, padx=10, pady=10)

        self.StepOne_Label = tk.Label(self.StepOne_FrameMain, text="1º Passo - Escolher pasta com scans do documento")
        self.StepOne_Label.grid(row = 0, column = 0, sticky = "W", pady = 10, columnspan=3)

        ### show circle with checkmark (with green background) or cross (with red background), it updates when the user selects a folder with text centered to middle (horizontal and vertical)
        self.StepOne_LabelCheckmark = tk.Label(self.StepOne_FrameMain, text=CROSS_ICON, fg="white", bg="red", width=2, height=1)
        self.StepOne_LabelCheckmark.grid(row = 1, column = 0, sticky = "W", pady = 2, padx = 5)

        ### show the path of the selected folder, updates path after user selects a folder
        self.StepOne_LabelPath = tk.Label(self.StepOne_FrameMain, text="Nenhuma pasta selecionada", border=1, relief="solid", width=50, height=1)
        self.StepOne_LabelPath.grid(row = 1, column = 1, sticky = "W", pady = 2, padx=5, columnspan=2)

        self.StepOne_Button = tk.Button(self.StepOne_FrameMain, text="Escolher pasta", command=self.stepOne_selectFolder)
        self.StepOne_Button.grid(row = 1, column = 4, sticky = "W", pady = 2, padx=5)

        ### show logo of app as image (located at projectRoot/gui/logoIcons/128x128.png) beside the StepOne_FrameMain
        # self.StepOne_Logo = tk.PhotoImage(file="gui/logoIcons/128x128.png")
        # self.StepOne_LabelLogo = tk.Label(self.StepOne_FrameMain, image=self.StepOne_Logo)
        # self.StepOne_LabelLogo.grid(row = 0, column = 5, sticky = "W", rowspan=2)

        self.stepOne_FolderMultipleDocs = tk.Checkbutton(self.StepOne_FrameMain, text='Pasta com múltiplos documentos',variable=self.stepOne_multipleDocs, onvalue=True, offvalue=False)
        self.stepOne_FolderMultipleDocs.grid(row = 2, column = 0, sticky = "W", pady = 2, padx=5, columnspan=3)

        self.StepOne_ShowFolder = tk.Button(self.StepOne_FrameMain, text="Pré-visualizar Imagens", command=self.stepOne_showFolder)
        self.StepOne_ShowFolder.grid(row = 2, column = 4, sticky = "W", pady = 2, padx=5)

        ## Frame for second step
        self.StepTwo_FrameMain = tk.Frame(self.root)
        self.StepTwo_FrameMain.pack(fill=tk.X, padx=10, pady=10)

        self.StepTwo_Label = tk.Label(self.StepTwo_FrameMain, text="2º Passo - Escolher pasta de destino do(s) PDF(s)")
        self.StepTwo_Label.grid(row = 0, column = 0, sticky = "W", pady = 10, columnspan=3)

        ### show circle with checkmark (with green background) or cross (with red background), it updates when the user selects a folder with text centered to middle (horizontal and vertical)
        self.StepTwo_LabelCheckmark = tk.Label(self.StepTwo_FrameMain, text=CROSS_ICON, fg="white", bg="red", width=2, height=1)
        self.StepTwo_LabelCheckmark.grid(row = 1, column = 0, sticky = "W", pady = 2, padx = 5)

        ### show the path of the selected folder, updates path after user selects a folder
        self.StepTwo_LabelPath = tk.Label(self.StepTwo_FrameMain, text="Nenhuma pasta selecionada", border=1, relief="solid", width=50, height=1)
        self.StepTwo_LabelPath.grid(row = 1, column = 1, sticky = "W", pady = 2, padx=5, columnspan=2)

        self.StepTwo_Button = tk.Button(self.StepTwo_FrameMain, text="Escolher pasta", command=self.stepTwo_selectFolder)
        self.StepTwo_Button.grid(row = 1, column = 4, sticky = "W", pady = 2, padx=5)

        ## Frame for third step
        self.StepThree_FrameMain = tk.Frame(self.root)
        self.StepThree_FrameMain.pack(fill=tk.X, padx=10, pady=10)

        self.StepThree_Label = tk.Label(self.StepThree_FrameMain, text="3º Passo (opcional) - Escolher extras")
        self.StepThree_Label.grid(row = 0, column = 0, sticky = "W", pady = 10, columnspan=3)

        ### Shows two checkboxes (side by side)
        self.StepThree_Checkbox1 = tk.Checkbutton(self.StepThree_FrameMain, text='Não criar PDF sem OCR', variable=self.stepThree_notCreateSimplePDF, onvalue=True, offvalue=False)
        self.StepThree_Checkbox1.grid(row = 1, column = 0, sticky = "W", pady = 2, padx=5)

        self.StepThree_Checkbox2 = tk.Checkbutton(self.StepThree_FrameMain, text='Dar aviso sonoro ao terminar', variable=self.stepThree_soundOnFinish, onvalue=True, offvalue=False)
        self.StepThree_Checkbox2.grid(row = 1, column = 1, sticky = "W", pady = 2, padx=5)

        ## Frame for Progress Area
        self.ProgressArea_FrameMain = tk.Frame(self.root)
        self.ProgressArea_FrameMain.pack(fill=tk.X, padx=10, pady=10, expand=True)
        self.ProgressArea_FrameMain.columnconfigure(0, weight=1)

        self.ProgressArea_Label = tk.Label(self.ProgressArea_FrameMain, text="Progresso da tarefa atual:")
        self.ProgressArea_Label.grid(row = 0, column = 0, sticky = "W", pady = 10, columnspan=3)

        ## Show progress bar that stretches width of window for current task
        self.ProgressArea_ProgressBar = ttk.Progressbar(self.ProgressArea_FrameMain, orient=tk.HORIZONTAL, mode='determinate')
        self.ProgressArea_ProgressBar.grid(row = 1, column = 0, sticky = "WE", pady = 2, padx=5, columnspan=7)
        #self.ProgressArea_FrameMain.columnconfigure(1, weight=1)

        ## Show progress bar for whole request
        self.ProgressArea_LabelTotal = tk.Label(self.ProgressArea_FrameMain, text="Progresso geral:")
        self.ProgressArea_LabelTotal.grid(row = 2, column = 0, sticky = "W", pady = 10, columnspan=3)

        self.ProgressArea_ProgressBarTotal = ttk.Progressbar(self.ProgressArea_FrameMain, orient=tk.HORIZONTAL, mode='determinate')
        self.ProgressArea_ProgressBarTotal.grid(row = 3, column = 0, sticky = "WE", pady = 2, padx=5, columnspan=7)
        #self.ProgressArea_FrameMain.columnconfigure(3, weight=1)

        ## show button to start the conversion process
        self.StartButton = tk.Button(self.ProgressArea_FrameMain, text="Iniciar Conversão", command=self.startConversion)
        self.StartButton.grid(row = 4, column = 0, sticky = "NESW", pady = 2, padx=5, columnspan=3)

        ## show button to pause the conversion process
        self.PauseButton = tk.Button(self.ProgressArea_FrameMain, text="Pausar Conversão")
        self.PauseButton.grid(row = 4, column = 4, sticky = "W", pady = 2, padx=5, columnspan=1)

        ## show button to cancel the conversion process
        self.CancelButton = tk.Button(self.ProgressArea_FrameMain, text="Cancelar Conversão")
        self.CancelButton.grid(row = 4, column = 6, sticky = "W", pady = 2, padx=5, columnspan=1)


    def stepOne_selectFolder(self):
        self.stepOne_directory = filedialog.askdirectory()
        if self.stepOne_directory:
            self.StepOne_LabelPath.config(text=self.stepOne_directory)
            self.StepOne_LabelCheckmark.config(text=CHECK_ICON, fg="white", bg="green")
        else:
            self.StepOne_LabelPath.config(text="Nenhuma pasta selecionada")
            self.StepOne_LabelCheckmark.config(text=CROSS_ICON, fg="white", bg="red")

    def stepOne_showFolder(self):
        showFiles.FilesWindow(self.root, self.stepOne_directory, self.stepOne_multipleDocs.get()).show()

    def stepTwo_selectFolder(self):
        self.stepTwo_directory = filedialog.askdirectory()
        if self.stepTwo_directory:
            self.StepTwo_LabelPath.config(text=self.stepTwo_directory)
            self.StepTwo_LabelCheckmark.config(text=CHECK_ICON, fg="white", bg="green")
        else:
            self.StepTwo_LabelPath.config(text="Nenhuma pasta selecionada")
            self.StepTwo_LabelCheckmark.config(text=CROSS_ICON, fg="white", bg="red")

    def startConversion(self):
        pass




