<?php

$YUKARI_VOICE_DIR = '/home/tao/yukari_bme280/voices/';
$BME280_PYTHON = '/home/tao/bme280_tao.py';

exec("sudo python ${BME280_PYTHON}", $out);

$temp = explode(',', $out[0])[0];
$hum = explode(',', $out[0])[1];

$voicePath = array("${YUKARI_VOICE_DIR}/now_temp.wav");
$voicePath = getAudioPath($voicePath, $temp);
array_push($voicePath, "${YUKARI_VOICE_DIR}/degrees_hum.wav");
$voicePath = getAudioPath($voicePath, $hum);
array_push($voicePath, "${YUKARI_VOICE_DIR}/percent_end.wav");

$result = implode(' ', $voicePath);

exec("sudo aplay ${result}");

function getAudioPath($array, $num){
    global $YUKARI_VOICE_DIR;
    $integer = explode('.', $num)[0];
    $decimal = explode('.', $num)[1];

    array_push($array,
        "${YUKARI_VOICE_DIR}/num/${integer}.wav",
        "${YUKARI_VOICE_DIR}/dot.wav"
    );

    $decimalArr = str_split($decimal);
    for($i = 0; $i < count($decimalArr); $i++){
        array_push($array, "${YUKARI_VOICE_DIR}/num/{$decimalArr[$i]}.wav");
    }
    return $array;
}
