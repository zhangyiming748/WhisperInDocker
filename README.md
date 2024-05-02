# WhisperInDocker

使用openaiwhisper批量生成字幕

# model

https://huggingface.co/ggerganov/whisper.cpp/tree/main

|Size|Parameters|Englishonlymodel|Multilingualmodel|RequiredVRAM|Relativespeed|
|:---:|:---:|:---:|:---:|:---:|:---:|
|Size|Parameters|Englishonlymodel|Multilingualmodel|RequiredVRAM|Relativespeed|
|tiny|39M|[tiny.en](https://openaipublic.azureedge.net/main/whisper/models/d3dd57d32accea0b295c96e26691aa14d8822fac7d9d27d5dc00b4ca2826dd03/tiny.en.pt)|[tiny](https://openaipublic.azureedge.net/main/whisper/models/65147644a518d12f04e32d6f3b26facc3f8dd46e5390956a9424a650c0ce22b9/tiny.pt)|~1GB|~32x|
|base|74M|[base.en](https://openaipublic.azureedge.net/main/whisper/models/25a8566e1d0c1e2231d1c762132cd20e0f96a85d16145c3a00adf5d1ac670ead/base.en.pt)|[base](https://openaipublic.azureedge.net/main/whisper/models/ed3a0b6b1c0edf879ad9b11b1af5a0e6ab5db9205f891f668f8b0e6c6326e34e/base.pt)|~1GB|~16x|
|small|244M|[small.en](https://openaipublic.azureedge.net/main/whisper/models/f953ad0fd29cacd07d5a9eda5624af0f6bcf2258be67c92b79389873d91e0872/small.en.pt)|[small](https://openaipublic.azureedge.net/main/whisper/models/9ecf779972d90ba49c06d968637d720dd632c55bbf19d441fb42bf17a411e794/small.pt)|~2GB|~6x|
|medium|769M|[medium.en](https://openaipublic.azureedge.net/main/whisper/models/d7440d1dc186f76616474e0ff0b3b6b879abc9d1a4926b7adfa41db2d497ab4f/medium.en.pt)|[medium](https://openaipublic.azureedge.net/main/whisper/models/345ae4da62f9b3d59415adc60127b97c714f32e89e936602e85993674d08dcb1/medium.pt)|~5GB|~2x|
|large|1550M|N/A|[large](https://openaipublic.azureedge.net/main/whisper/models/e4b87e7e0bf463eb8e6956e646f1e277e901512310def2c24bf0e11bd3c28e9a/large.pt)|~10GB|1x|



# usage 
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
docker run -dit --name=whisper_en  -v '/mnt/f/ubuntu/jp/en:/srt' -e language=English whisper:latest srt
docker run -dit --name=whisper_en -v '/Users/zen/Movies:/srt' -e language=English -e pattern=m4a whisper:latest bash
docker run -dit --name=whisper_sp --rm -v '/c/Users/zen/Videos/sp:/srt' -e pattern=webm -e language=Spanish zhangyiming748/whisper:v0.0.1 srt
docker run -dit --name=whisper_ja --rm -v '/f/srt:/srt' -e language=Japanese whisper:latest srt
docker run -dit --name=whisper_ja --rm -v '/f/ubuntu/jp/ja:/data' -e language=Japanese whisper:latest srt
docker run -dit --name=whisper_de --rm -v '/f/Telegram/srt/cut/en:/srt' -e language=German whisper:latest srt
docker run -dit --name=whisper_ru --rm -v '/f/alist/bilibili/ru:/srt' -e language=Russian whisper:latest srt
docker run -dit --name=whisper_en --cpus=1 --memory=2048M --rm -v '/c/Users/zen/Videos/Nicole Aniston:/srt' -v /d/git/WhisperInDocker:/data -e language=English whisper:latest srt

```