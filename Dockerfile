FROM golang:1.22.3-bookworm
#ARG http_proxy=http:192.168.1.20:8889
#ARG https_proxy=http:192.168.1.20:8889
# 声明视频文件夹
VOLUME /data
# 声明程序文件夹
# VOLUME /app
# 声明模型文件夹
VOLUME /model

RUN mkdir /root/app
WORKDIR /root/app
COPY . /root/app
RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN go build -o /usr/local/bin/CreateSubtitle main.go

COPY debian.sources /etc/apt/sources.list.d/
RUN chmod +x install-retry.sh
RUN apt update
RUN apt install -y dos2unix
RUN dos2unix ./install-retry.sh
RUN mv install-retry.sh /usr/local/bin
RUN install-retry.sh ffmpeg python3 python3-pip nano
RUN apt clean

RUN pip config set global.index-url https://pypi.tuna.tsinghua.edu.cn/simple
RUN pip install openai-whisper --break-system-packages --no-cache-dir

RUN dos2unix entrypoint.sh
RUN chmod +x entrypoint.sh
# CMD ["whisper"]
ENTRYPOINT ["/root/app/entrypoint.sh"]
# docker build --progress=plain -t whisper:latest .
# docker run -dit --rm --name=whisper_en  -v '/c/Users/zen/Videos/test:/data' -e language=English whisper:latest
# sudo docker run -dit --cpus=4 --memory=4096M --name test -v /home/zen/github/WhisperInDocker/model:/model -v /home/zen/docker/yt-dlp:/data -e root=/data -e language=English -e pattern=webm -e model=large whisper:latest