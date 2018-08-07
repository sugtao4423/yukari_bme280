package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

const YUKARI_VOICE_DIR = "/home/tao/yukari_bme280/voices/"
const BME280_PYTHON = "/home/tao/bme280_tao.py"

func main() {
	out, err := exec.Command("sudo", "python", BME280_PYTHON).Output()
	if err != nil {
		fmt.Errorf("Error: failed get nums from bme280")
		os.Exit(1)
	}
	temp := strings.Split((string)(out), ",")[0]
	hum := strings.Split((string)(out), ",")[1]

	voicePath := []string{"sudo", "aplay", YUKARI_VOICE_DIR + "/now_temp.wav"}
	appendVoicePath(&voicePath, temp)
	voicePath = append(voicePath, YUKARI_VOICE_DIR+"/degrees_hum.wav")
	appendVoicePath(&voicePath, hum)
	voicePath = append(voicePath, YUKARI_VOICE_DIR+"/percent_end.wav")

	exec.Command(voicePath[0], voicePath[1:]...).Run()
}

func appendVoicePath(array *[]string, num string) {
	integer := strings.Split(num, ".")[0]
	decimal := strings.Split(num, ".")[1]

	*array = append(*array,
		YUKARI_VOICE_DIR+"/num/"+integer+".wav",
		YUKARI_VOICE_DIR+"/dot.wav")

	decimalArr := strings.Split(decimal, "")
	for _, v := range decimalArr {
		*array = append(*array, YUKARI_VOICE_DIR+"/num/"+v+".wav")
	}
}
