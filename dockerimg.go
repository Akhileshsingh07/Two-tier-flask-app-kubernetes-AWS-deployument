FROM python:3.9-slim

WORKDIR /app

RUN apt-get update\
    && apt-get upgrade\
	&& apt-get install -y gcc default-libmysqlclient-dev pkg-config\
	&& rm -rf/var/lib/apt/lists/*

COPY requirments.txt .

RUN pip install mqsqlclient 

RUN pip install --no-cache-dir -r requirments.txt 

COPY . . 

CMD ["python", "app.py"]