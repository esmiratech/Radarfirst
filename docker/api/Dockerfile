FROM python:3.10

WORKDIR /opt/api

COPY ./init/requirements.txt .

RUN pip install --no-cache-dir -r requirements.txt

COPY ./init/ .

CMD ["python", "app.py"]