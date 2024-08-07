import os
import configparser

class ImportConfigs:
    _instance = None
    _config_data = None

    def __new__(cls):
        if cls._instance is None:
            cls._instance = super(ImportConfigs, cls).__new__(cls)
            cls._instance.init_config()
            cls._instance.load_config_data()
        return cls._instance
    
    def load_config_data(self):
        self._config_data = self.app_config["DEFAULT"]
    
    def init_config(self):
        self.app_config = None
        # Define the path to the config file
        self.config_file_path = "config.ini"
        # Define the default values
        self.base_config = {
            "screenWidth": "800",
            "screenHeight": "600",
            "startingPosX": "200",
            "startingPosY": "200",
            # Add more variables and their default values here
        }
        if os.path.exists(self.config_file_path):
            self.app_config = configparser.ConfigParser()
            self.app_config.read(self.config_file_path)

            # Check if all required values are present
            for key in self.base_config:
                if key not in self.app_config["DEFAULT"]:
                    print(f"Error: Required value '{key}' is missing in the config file.")
                    exit(1)
        else:
            # Create a new config file with default values
            self.app_config = configparser.ConfigParser(self.base_config)
            with open(self.config_file_path, "w") as file:
                self.app_config.write(file)
        
        self._config_data = self.app_config["DEFAULT"]

    def getAll(self):
        # app_config = configparser.ConfigParser()
        # app_config.read(self.config_file_path)
        # return self.app_config.items("DEFAULT")
        return self._config_data.items()

    def get(self, config_name):
        # app_config = configparser.ConfigParser()
        # app_config.read(self.config_file_path)
        # return self.app_config.get("DEFAULT", config_name)
        if not self._config_data:
            self.init_config()
            self.load_config_data()
        return self._config_data.get(config_name)

    def change(self, config_name, new_value):
        # app_config = configparser.ConfigParser()
        # app_config.read(self.config_file_path)
        # self.app_config.set("DEFAULT", config_name, new_value)
        # with open(self.config_file_path, "w") as file:
        #     self.app_config.write(file)
        self._config_data[config_name] = new_value
        self.app_config.set("DEFAULT", config_name, new_value)
        with open(self.config_file_path, "w") as file:
            self.app_config.write(file)

# Export the get_config and change_config functions
# __all__ = ["initialize", "get", "change"]


