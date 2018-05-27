import websocket
import time

class MessageListener:

    def __init__(self, address, greeting):
        self.address = address
        self.greeting = greeting
    
    def connect(self):
        ws = websocket.create_connection(self.address)
        ws.connect(self.address)
        ws.send(self.greeting)
        self.ws = ws
    
    def listen(self):
        while True:
            t0 = time.time()
            # self.ws.send("lalala")
            result = self.ws.recv()
            print("Received {} in {} s".format(result, time.time()-t0))
    
    def exit(self):
        self.ws.close()

if __name__ == "__main__":
    try:
        address = "wss://echo.websocket.org"
        first_message = "Nieooow"
        
        app = MessageListener(address, first_message)
        app.connect()
        app.listen()
    except KeyboardInterrupt:
        exit()
