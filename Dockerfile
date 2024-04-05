# 基础镜像
# docker pull golang:1.22.1-bookworm
# docker run -dit --name=test usewhisper:latest -v /mnt/e/video/srt:/srt -e language=en bash
FROM golang:1.22.1-bookworm
#ARG http_proxy=http:192.168.1.5:8889
#ARG https_proxy=http:192.168.1.5:8889
# 用于存储程序和视频字幕文件的文件夹
VOLUME /srt
COPY srt /usr/local/bin/srt
COPY install-retry.sh /usr/local/bin/install-retry.sh
RUN chmod +x /usr/local/bin/install-retry.sh

RUN sed -i 's/deb.debian.org/mirrors4.tuna.tsinghua.edu.cn/g' /etc/apt/sources.list.d/debian.sources
RUN sed -i 's/security.debian.org/mirrors4.tuna.tsinghua.edu.cn/g' /etc/apt/sources.list.d/debian.sources
RUN apt update
RUN install-retry.sh ffmpeg
RUN install-retry.sh python3
RUN install-retry.sh python3-pip
RUN install-retry.sh vim
RUN install-retry.sh nano
RUN install-retry.sh mediainfo
RUN install-retry.sh wget
RUN install-retry.sh git
# RUN apt install -y ffmpeg python3 python3-pip vim nano mediainfo wget git
RUN pip config set global.index-url https://pypi.tuna.tsinghua.edu.cn/simple
RUN pip install openai-whisper --break-system-packages
RUN mkdir /module
RUN wget --tries=10 https://openaipublic.azureedge.net/main/whisper/models/ed3a0b6b1c0edf879ad9b11b1af5a0e6ab5db9205f891f668f8b0e6c6326e34e/base.pt -O /module/base.pt --no-check-certificate
RUN go env -w GO111MODULE=on
# RUN go env -w GOPROXY=https://goproxy.cn,direct
# RUN git clone https://github.com/zhangyiming748/WhisperInDocker.git /root/WhisperInDocker
# WORKDIR /root/WhisperInDocker
# RUN go mod tidy

CMD ["srt"]
# docker build -t trans:v1 .
# docker run -itd --name=trans1 -v /d/srt:/srt zhangyiming748/use-whisper:v0.0.3 bash
# docker run -idt --name=trans -v /d/srt:/srt -e APPID={your baidu appid} -e KEY={your baidu key} trans:v1 ash
# docker run -itd --name=trans1 -v /d/srt:/srt use-whisper:v0.0.3
# ENTRYPOINT ["srt"]
# docker build -t trans:v1 .
# docker run -itd --name=trans1 -v /d/srt:/srt zhangyiming748/use-whisper:v0.0.3 bash
# docker run -idt --name=trans -v /d/srt:/srt -e APPID={your baidu appid} -e KEY={your baidu key} trans:v1 ash
# docker run -itd --name=trans1 -v /d/srt:/srt use-whisper:v0.0.3
# docker build -t zhangyiming748/whisper:apple0.0.3 .
