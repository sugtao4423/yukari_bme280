# Yukari BME280
BME280で取得した温度と湿度を読み上げてくれる

ゆかりさんに教えてもらいたかったから作った  
ファイル名さえ同じにすれば他のボイスでもいける

## Require
* 温度の取得  
[SWITCHSCIENCE/BME280](https://github.com/SWITCHSCIENCE/BME280)をforkした  
[sugtao4423/BME280](https://github.com/sugtao4423/BME280)をphpから使用

* 読み上げ  
 `aplay` を使用

## Voices
ファイル名 | 音声
--- | ---
voices/num/0〜100.wav | 0〜100
voices/dot.wav | てん（小数点）
voices/now_temp.wav | ただいまの温度は
voices/degrees_hum.wav | 度で、湿度は
voices/percent_end.wav | パーセントです

## Usage
* `php /home/hoge/yukari_bme280/yukari_bme280.php`  
PHP内の `$YUKARI_VOICE_DIR` と `$BME280_PYTHON` を適宜変更すること

* Amazon Dash Buttonと連携
    - `bash /home/hoge/yukari_bme280/dash.sh`  
    スクリプト内の `DASH_IP` と `9行目のPHPのパス` を適宜変更すること

Raspberry Piで動作させることを想定して作成。
