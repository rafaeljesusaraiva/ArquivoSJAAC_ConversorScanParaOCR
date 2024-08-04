import tkinter as tk
import modules.importConfigs as appConfig

class ShowConfig:
    def __init__(self, root):
        self.root = root
        self.config_window = None
        self.configurations = self.loadConfigs()

    def loadConfigs(self):
        # Load configurations from the config file
        tdata = appConfig.ImportConfigs()
        return tdata.getAll()

    def center_window(self):
        screen_width = self.root.winfo_screenwidth()
        screen_height = self.root.winfo_screenheight()
        width = self.config_window.winfo_reqwidth()
        height = self.config_window.winfo_reqheight()
        x = (screen_width/2) - (width/2)
        y = (screen_height/2) - (height/2)
        self.config_window.geometry(f"+{int(x)}+{int(y)}")

    def exit(self):
        # save settings locally and try to signal main app to update accordingly???
        self.config_window.destroy()

    def show(self):
        try:
            if self.config_window:
                self.config_window.destroy()
        finally:
            self.config_window = tk.Toplevel(self.root)
            self.config_window.title("Configurações")
            self.config_window.resizable(False, False)  # enable resizing
            self.config_window.withdraw()  # make the window invisible

             # create a frame to hold the image and text
            frame = tk.Frame(self.config_window)
            frame.pack(fill="both", expand=True)

            w1 = tk.Scale(frame, from_=0, to=42, tickinterval=8)
            w1.set(19)
            # s1 = tk.Scale(frame, variable = v1, from_ = 1, to = 100, orient = "HORIZONTAL") 

            # create an exit button
            exit_button = tk.Button(frame, text="Guardar e Sair", command=self.exit)
            exit_button.pack(pady=10, fill="x")  # fill horizontally

            # update the window size
            self.config_window.update_idletasks()

            # center the window
            self.center_window()

            # make the window visible
            self.config_window.deiconify()
# w2 = Scale(master, from_=0, to=200, orient=HORIZONTAL)
# w2.pack()