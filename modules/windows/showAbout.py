from os import path
import tkinter as tk
from PIL import Image, ImageTk
import webbrowser

logoPath = path.join(path.dirname(path.abspath(__file__)), '..', '..', 'gui', 'logoIcons', '256x256.png')
textAboutApp = """
Aplicação desenvolvida no âmbito do projeto do arquivo digital do Arquivo da SJ/AAC, para otimizar o processo de conversão de documentos digitalizados em PDFs com texto pesquisável (OCR).
"""

def callbackRepo(e):
    webbrowser.open_new("https://github.com/rafaeljesusaraiva/ArquivoSJAAC_ConversorScanParaOCR")

class AboutWindow:
    def __init__(self, root):
        self.root = root
        self.about_window = None

    def center_window(self):
        screen_width = self.root.winfo_screenwidth()
        screen_height = self.root.winfo_screenheight()
        width = self.about_window.winfo_reqwidth()
        height = self.about_window.winfo_reqheight()
        x = (screen_width/2) - (width/2)
        y = (screen_height/2) - (height/2)
        self.about_window.geometry(f"+{int(x)}+{int(y)}")

    def show(self):
        try:
            if self.about_window:
                self.about_window.destroy()
        finally:
            self.about_window = tk.Toplevel(self.root)
            self.about_window.title("Sobre a aplicação")
            self.about_window.resizable(False, False)  # enable resizing
            self.about_window.withdraw()  # make the window invisible

            # create a frame to hold the image and text
            frame = tk.Frame(self.about_window)
            frame.pack(fill="both", expand=True)

            # load the image
            image = Image.open(logoPath)  # replace with your image file

            # resize the image to fit the window
            width, height = image.size
            new_width = 200  # adjust this value to set the desired width
            new_height = int(height * (new_width / width))
            image = image.resize((new_width, new_height))

            # create a label with the bitmap image
            photo_image = ImageTk.PhotoImage(image)
            image_label = tk.Label(frame, image=photo_image)
            image_label.image = photo_image
            image_label.pack(pady=10, fill="x")  # fill horizontally

            # create a text label with lorem ipsum text
            introText_label = tk.Label(frame, text=textAboutApp, wraplength=300)
            introText_label.pack(pady=10, fill="x")  # fill horizontally

            repositoryTitle_label = tk.Label(frame, text="Repositório GitHub", fg="blue", cursor="hand2")
            repositoryTitle_label.pack(pady=10, fill="x")  # fill horizontally
            repositoryTitle_label.bind("<Button-1>", callbackRepo)

            # create an exit button
            exit_button = tk.Button(frame, text="Fechar", command=self.about_window.destroy)
            exit_button.pack(pady=10, fill="x")  # fill horizontally

            # update the window size
            self.about_window.update_idletasks()

            # center the window
            self.center_window()

            # make the window visible
            self.about_window.deiconify()


