import requests as req

class Notifier:
    
    def post(self, data):
        try:
            req.post("http://mockgate:8080", data=data)
            print(str(data))
        except req.ConnectionError as e:
            print("Error sendind the car data message: " + str(e)) 
        