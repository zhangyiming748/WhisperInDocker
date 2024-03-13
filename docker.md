```shell
docker commit -a zen -m "whisper可用" -p whisper useWhisper:v0.0.1
# 需要变量language
# 需要映射/srt

docker run -dit --name=test -v /e/video/srt:/srt -e language=en usewhisper:v0.0.1 srt
docker load -i whisper.tar
docker run -dit --name=test --rm -v /e/video/srt:/srt -e language=ja whisper:v0.0.1 srt
docker run -dit --name=test --rm -v /mnt/d/git/RemoveIntroOutro:/srt -e language=en whisper:latest bash
docker run -dit --name=test --rm -v /d/git/RemoveIntroOutro:/srt -e language=en whisper:latest srt
docker run -dit --name=whisper --rm -v '/f/large/GirlFriend4ever/2020/Pure As Sin/Pure As Sin - Airborne Blowjob Animation:/srt' -e language=en whisper:latest srt
```