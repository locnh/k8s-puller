FROM docker:stable

RUN apk add --update --no-cache python3
RUN python3 -m ensurepip
RUN pip3 install --no-cache --upgrade pip schedule

ADD puller.py /usr/local/bin/puller.py

ENTRYPOINT ["python3", "/usr/local/bin/puller.py"]