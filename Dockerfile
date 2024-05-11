FROM golang:1.22.2-bookworm
#ARG http_proxy=http:192.168.1.20:8889
#ARG https_proxy=http:192.168.1.20:8889
VOLUME /data
RUN mkdir /root/app
WORKDIR /root/app

COPY . .
RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN go build -o /usr/local/bin/CreateSubtitle main.go
RUN chmod +x install-retry.sh
RUN sed -i 's/deb.debian.org/mirrors4.tuna.tsinghua.edu.cn/g' /etc/apt/sources.list.d/debian.sources
RUN apt update
RUN apt install -y dos2unix
RUN dos2unix ./install-retry.sh
RUN mv install-retry.sh /usr/local/bin
RUN install-retry.sh ffmpeg python3 python3-pip nano
RUN pip config set global.index-url https://pypi.tuna.tsinghua.edu.cn/simple

RUN tar xvf pip.tar
RUN pip install  --find-links=./pip openai-whisper --break-system-packages

RUN dos2unix entrypoint.sh
RUN chmod +x entrypoint.sh
# CMD ["whisper"]
ENTRYPOINT ["/root/app/entrypoint.sh"]
# docker build --progress=plain -t whisper:latest .
# docker run -dit --rm --name=whisper_en  -v '/c/Users/zen/Videos/test:/data' -e language=English whisper:latest