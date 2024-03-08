```shell
docker commit -a zen -m "whisper可用" -p whisper useWhisper:v0.0.1
# 需要变量language
# 需要映射/srt

docker run -dit --name=test -v /e/video/srt:/srt -e language=en usewhisper:v0.0.1 srt
docker load -i whisper.tar
docker run -dit --name=test --rm -v /e/video/srt:/srt -e language=en whisper:v0.0.1 srt
docker run -dit --name=test --rm -v /path/to/video:/srt -e language=en whisper:latest srt
```