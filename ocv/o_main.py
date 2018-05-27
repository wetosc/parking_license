import cv2
from cv2 import (VideoCapture, imwrite)
# initialize the camera
cam = VideoCapture(0)   # 0 -> index of camera
img = None
s, img = cam.read()

print(cam)
print(s)
print(img)
imwrite("filename.jpg", img)

import socket
sock = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
sock.connect(('localhost', 9000))

sock.send('hello')


