import tkinter as tk
import modules.importConfigs
import modules.windows.mainWindow as mainWindow
import modules.windows.showAbout as showAbout
import modules.windows.showConfig as showConfig

# Import configurations from importConfigs.py
appConfig = modules.importConfigs.ImportConfigs()

# Create a tkinter window
window = tk.Tk()
window.title("Conversor Scan Para OCR - Arquivo SJ/AAC")

# Use the configurations to customize the GUI
# Set window size based on configuration
window.geometry(f"{appConfig.get('screenWidth')}x{appConfig.get('screenHeight')}+{appConfig.get('startingPosX')}+{appConfig.get('startingPosY')}")
# Lock the window size
window.resizable(False, False)

# Create a menu bar
menu_bar = tk.Menu(window)
window.config(menu=menu_bar)

# Create a "Ficheiro" menu
ficheiro_menu = tk.Menu(menu_bar, tearoff=0)
menu_bar.add_cascade(label="Ficheiro", menu=ficheiro_menu)

# Create menu items for "Ficheiro"
ficheiro_menu.add_command(label="Opção teste")
ficheiro_menu.add_separator()
ficheiro_menu.add_command(label="Sair", command=window.quit)

# Create an "Extras" menu
extras_menu = tk.Menu(menu_bar, tearoff=0)
menu_bar.add_cascade(label="Opções", menu=extras_menu)
    
extras_menu.add_command(label="Configuração", command=showConfig.ShowConfig(window).show)
extras_menu.add_command(label="Sobre a aplicação", command=showAbout.AboutWindow(window).show)

mainWindow.MainWindow(window)

# Start the tkinter event loop
window.mainloop()
