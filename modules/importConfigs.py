import os
import configparser

# Define the path to the config file
config_file_path = "config.ini"

# Define the default values
base_config = {
    "screenWidth": "800",
    "screenHeight": "600",
    "startingPosX": "200",
    "startingPosY": "200",
    # Add more variables and their default values here
}

app_config = None

def initialize():
    if os.path.exists(config_file_path):
        app_config = configparser.ConfigParser()
        app_config.read(config_file_path)

        # Check if all required values are present
        for key in base_config:
            if key not in app_config["DEFAULT"]:
                print(f"Error: Required value '{key}' is missing in the config file.")
                exit(1)
    else:
        # Create a new config file with default values
        app_config = configparser.ConfigParser(base_config)
        with open(config_file_path, "w") as file:
            app_config.write(file)

def get(config_name):
    app_config = configparser.ConfigParser()
    app_config.read(config_file_path)
    return app_config.get("DEFAULT", config_name)


def change(config_name, new_value):
    app_config = configparser.ConfigParser()
    app_config.read(config_file_path)
    app_config.set("DEFAULT", config_name, new_value)
    with open(config_file_path, "w") as file:
        app_config.write(file)

# Export the get_config and change_config functions
__all__ = ["initialize", "get", "change"]


