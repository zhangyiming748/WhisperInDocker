# 基础镜像
# docker pull golang:1.22.1-bookworm
# docker run -dit --name=test usewhisper:latest -v /mnt/e/video/srt:/srt -e language=en bash
FROM golang:1.22.1-bookworm
# 用于存储程序和视频字幕文件的文件夹
VOLUME /srt
COPY srt /usr/local/bin/srt
# RUN sed -i 's/deb.debian.org/mirrors4.tuna.tsinghua.edu.cn/g' /etc/apt/sources.list.d/debian.sources
RUN apt update && apt full-upgrade
RUN apt install -y ffmpeg python3 python3-pip vim nano mediainfo wget git
# RUN pip config set global.index-url https://pypi.tuna.tsinghua.edu.cn/simple
RUN pip install openai-whisper --break-system-packages
RUN mkdir /module
RUN wget https://openaipublic.azureedge.net/main/whisper/models/ed3a0b6b1c0edf879ad9b11b1af5a0e6ab5db9205f891f668f8b0e6c6326e34e/base.pt -O /module/base.pt --no-check-certificate
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