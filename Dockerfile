FROM golang:1.22.2-bookworm
#ARG http_proxy=http:192.168.1.20:8889
#ARG https_proxy=http:192.168.1.20:8889
VOLUME /data
RUN mkdir /data/app
WORKDIR /data/app
COPY . .
RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN go build -o /usr/local/bin/srt main.go
RUN chmod +x install-retry.sh
RUN sed -i 's/deb.debian.org/mirrors4.tuna.tsinghua.edu.cn/g' /etc/apt/sources.list.d/debian.sources
RUN apt update
RUN ./install-retry.sh ffmpeg python3 python3-pip nano
RUN pip config set global.index-url https://pypi.tuna.tsinghua.edu.cn/simple
WORKDIR /data/app/python/debian
RUN pip install -r requirements.txt --break-system-packages
RUN pip install openai-whisper --break-system-packages
WORKDIR /data/app
CMD ["srt"]
# docker build -t whisper:latest .
# docker run -itd --name=whisper -v /mnt/f/ubuntu/jp/en:/srt whisper:latest srt
# docker run -idt --name=trans -v /d/srt:/srt -e APPID={your baidu appid} -e KEY={your baidu key} trans:v1 ash
# docker run -itd --name=trans1 -v /d/srt:/srt use-whisper:v0.0.3
# ENTRYPOINT ["srt"]
# docker build -t trans:latest .
# docker build -t zhangyiming748/whisper:apple0.0.3 .
