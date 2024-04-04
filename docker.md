```shell
docker commit -a zen -m "whisper可用" -p whisper useWhisper:v0.0.1
# 需要变量language
# 需要映射/srt

docker run -dit --name=test -v /e/video/srt:/srt -e language=en usewhisper:v0.0.1 srt
docker load -i whisper.tar
docker run -dit --name=test --rm -v /e/video/srt:/srt -e language=ja whisper:v0.0.1 srt
docker run -dit --name=test --rm -v /mnt/d/git/RemoveIntroOutro:/srt -e language=en whisper:latest bash
docker run -dit --name=test --rm -v /d/git/RemoveIntroOutro:/srt -e language=en whisper:latest srt
docker run -dit --name=whisper --rm -v '/f/large  /未整理/[22sht.me]sdde-584:/srt' -e language=ja whisper:latest srt
docker run -dit --name=whisper_en --rm -v '/f/Telegram/en:/srt' -e language=English whisper:latest srt
docker run -dit --name=whisper_en --rm -v '/d/video/en:/srt' -e language=English zhangyiming748/whisper:v0.0.1 srt
docker run -dit --name=whisper_en --rm -v '/f/Telegram/en:/srt' -e language=English whisper:latest srt
docker run -dit --name=whisper_ja --rm -v '/f/large/未整理/本庄优花:/srt' -e language=Japanese whisper:latest srt
docker run -dit --name=whisper_de --rm -v '/f/Telegram/srt/cut/en:/srt' -e language=German whisper:latest srt
docker run -dit --name=whisper_ru --rm -v '/f/alist/bilibili/ru:/srt' -e language=Russian whisper:latest srt
docker run -dit --name=whisper_en --cpus=1 --memory=2048M --rm -v '/c/Users/zen/Videos/Nicole Aniston:/srt' -v /d/git/WhisperInDocker:/data -e language=English whisper:latest srt

```