from os import path, listdir
import tkinter as tk
from PIL import Image, ImageTk

class FilesWindow:
    def __init__(self, root, file_path, multiple_folders=False):
        self.root = root
        self.file_path = file_path
        self.multiple_folders = multiple_folders
        self.files_window = None
        self.current_image = None

    def show(self):
        # if no filepath is provided, show an error message
        if not self.file_path:
            tk.messagebox.showerror("Error", "No file path provided.")
            return
        
        try:
            if self.files_window:
                self.files_window.destroy()

        finally:
            self.files_window = tk.Toplevel(self.root)
            self.files_window.title("Pré-visualizar Imagens")
            self.files_window.geometry("800x600")
            # self.files_window.resizable(False, False)

            # Create frames for the image and the list of files
            self.image_frame = tk.Frame(self.files_window)
            self.image_frame.pack(side="left", fill="both", expand=True)

            self.list_frame = tk.Frame(self.files_window)
            self.list_frame.pack(side="right", fill="both", expand=True)

            # Create the image label
            self.image_label = tk.Label(self.image_frame)
            self.image_label.pack(fill="both", expand=True)

            # Create the listbox
            self.file_listbox = tk.Listbox(self.list_frame, selectmode=tk.SINGLE, width=30, yscrollcommand=True, xscrollcommand=False)
            self.file_listbox.pack(fill="both", expand=True)

            # Fill the listbox with files
            self.populate_listbox()

            self.file_listbox.bind("<<ListboxSelect>>", self.display_image)

            # Center the window
            self.files_window.update_idletasks()
            screen_width = self.root.winfo_screenwidth()
            screen_height = self.root.winfo_screenheight()
            width = self.files_window.winfo_reqwidth()
            height = self.files_window.winfo_reqheight()
            x = (screen_width/2) - (width/2)
            y = (screen_height/2) - (height/2)
            self.files_window.geometry(f"+{int(x)}+{int(y)}")

    # # display a frame with folder name and number of files in the folder
    # def list_folderSection(self, foldername, no_files):
    #     frame = tk.Frame(self.list_frame)
    #     frame.pack(fill="x")

    #     folder_label = tk.Label(frame, text=foldername)
    #     folder_label.pack(side="left")

    #     no_files_label = tk.Label(frame, text=f"({no_files} files)")
    #     no_files_label.pack(side="right")

    # # display a frame with file name and file size, also receiving file path to on-click show the image in the image_label
    # def list_fileSection(self, filename, filesize, filepath):
    #     frame = tk.Frame(self.list_frame)
    #     frame.pack(fill="x", padx=20)

    #     file_label = tk.Label(frame, text=filename)
    #     file_label.pack(side="left")

    #     file_size_label = tk.Label(frame, text=f"({filesize} bytes)")
    #     file_size_label.pack(side="right")

    #     frame.bind("<Button-1>", lambda event: self.display_image(event))
            
    # display a frame with folder name and number of files in the folder
    def list_folderSection(self, foldername, no_files):
        # frame = tk.Frame(self.list_frame)
        frame = tk.Frame()
        frame.pack(fill="x")

        folder_label = tk.Label(frame, text=foldername)
        folder_label.pack(side="left")

        no_files_label = tk.Label(frame, text=f"({no_files} files)")
        no_files_label.pack(side="right")

    # display a frame with file name and file size, also receiving file path to on-click show the image in the image_label
    def list_fileSection(self, filename, filesize, filepath):
        self.file_listbox.insert(tk.END, filename)

    def populate_listbox(self):
        if self.multiple_folders:
            # Get all folder names in directory
            folders = [folder for folder in listdir(self.file_path) if path.isdir(path.join(self.file_path, folder))]
            # Loop through the folders and show the folder name in the list with the number of files in the folder and the file list after
            for folder in folders:
                folder_path = path.join(self.file_path, folder)
                files = [file for file in listdir(folder_path) if path.isfile(path.join(folder_path, file))]
                self.list_folderSection(folder, len(files))
                for file in files:
                    file_path = path.join(folder_path, file)
                    self.list_fileSection(file, path.getsize(file_path), file_path)
        else:
            # Show folder name and number of files in the folder
            self.list_folderSection(self.file_path, len([file for file in listdir(self.file_path) if path.isfile(path.join(self.file_path, file))]))
            for file in [file for file in listdir(self.file_path) if path.isfile(path.join(self.file_path, file))]:
                file_path = path.join(self.file_path, file)
                self.list_fileSection(file, path.getsize(file_path), file_path)

    def display_image(self, event):
        # Get the selected file
        selection = self.file_listbox.curselection()
        if selection:
            index = selection[0]
            file_name = self.file_listbox.get(index)

            # Construct the full path to the file
            if self.multiple_folders:
                file_path = path.join(self.file_path, file_name, "image.tiff")
            else:
                file_path = path.join(self.file_path, file_name)

            # Check if the file exists
            if path.isfile(file_path):
                try:
                    # Load the image
                    image = Image.open(file_path)
                    photo_image = ImageTk.PhotoImage(image)

                    # Update the image label
                    self.image_label.config(image=photo_image)
                    self.image_label.image = photo_image

                    # Keep track of the current image
                    self.current_image = photo_image
                except:
                    # If the file is not an image, show an error message
                    tk.messagebox.showerror("Error", "Invalid image file.")
            else:
                # If the file doesn't exist, show an error message
                tk.messagebox.showerror("Error", "File not found.")