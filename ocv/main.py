import cv2
import base64
import socket
import time

class App: 

    def __init__(self, PORT):
        time.sleep(1)
        self.PORT = PORT
        self.cap = cv2.VideoCapture(0)
        self.last_time = time.time()

    def take_photo(self):
        retval, image = self.cap.read()
        data = cv2.imencode('.jpg', image)[1].tostring()
        return data

    def send_image(self, encimg):
        sock = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
        sock.connect(('govision', self.PORT))
        sock.send(encimg)
        sock.close()
        
    def tick(self):
        if time.time() - self.last_time > 1:
            img = self.take_photo()
            self.send_image(img)
            self.last_time = time.time()
        else:
            time.sleep(0.1)
            self.deplete_buffer()

    def deplete_buffer(self):
        self.cap.read()

    def close(self):
        self.cap.release()


if __name__ == "__main__":
    app = App(8123)
    print("START")
    while True:
        app.tick()

    app.close()