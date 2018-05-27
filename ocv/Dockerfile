FROM debian:latest

RUN apt-get update && apt-get install -y \
    python3-pip

WORKDIR /app

COPY requirements.txt .
RUN pip3 install -r requirements.txt

COPY . .

EXPOSE 8123

CMD [ "python3", "main.py" ]