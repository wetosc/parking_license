from openalpr import Alpr
from source import DataSource
from notify import Notifier

class App:

    def __init__(self, id):
        alpr = Alpr("eu", "/etc/openalpr/openalpr.conf","../runtime_data")
        if not alpr.is_loaded():
            print("Error loading OpenALPR")
            exit()
        alpr.set_top_n(20)

        self.alpr = alpr
        self.notifier = Notifier()
        self.id = id
        self.source = DataSource(self)

    def on_image_input(self, image_path):
        result = self.alpr.recognize_file(image_path)
        self.send_data(result)
        # self.print_file_result(image_path, result)
        
    def print_file_result(self, file_name, data):
        print("=====\nFile {}:\n---".format(file_name))
        results = data["results"]
        if len(results) == 0 :
            print("No results found")
        else:
            print(results[0]["plate"])

    def send_data(self, data):
        results = data["results"]
        if len(results) == 0:
            return    
        json = {
            "id": self.id,
            "plate": results[0]["plate"],
            "isEntering": True 
        }
        self.notifier.post(json)

    def exit(self):
        self.alpr.unload()


if __name__ == "__main__":
    app = App("f83447f6-69d4-41f5-9ed1-b07484d8c968")
    try:
        while True:
            pass
    except KeyboardInterrupt:
        print('Exit')
        app.exit()