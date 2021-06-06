// +build windows

package main

import (
	"log"
	"os/exec"

	"github.com/itchyny/volume-go"
)

func main() {
	volume := 30
	url := "https://mynoise.net/NoiseMachines/whiteNoiseGenerator.php?l=14222840464228211305&m=&d=0"

	log.Println("start whitenoise")
	setVolumeGoal(volume)
	openBrowser(url)

	// infinite read input loop
	// reader := bufio.NewReader(os.Stdin)
	// for {
	// 	_, err := reader.ReadString('\n')
	// 	if err != nil {
	// 		log.Fatalf("get input failed: %+v", err)
	// 	}
	// }
}

func setVolumeGoal(vol int) {
	err := volume.SetVolume(vol)
	if err != nil {
		log.Fatalf("set volume failed: %+v", err)
	}
	log.Printf("set volume to: %d\n", vol)

	readVol, err := volume.GetVolume()
	if err != nil {
		log.Fatalf("get volume failed: %+v", err)
	}
	log.Printf("current volume: %d\n", readVol)
}

func openBrowser(url string) {
	cmd := exec.Command("rundll32", "url.dll,FileProtocolHandler", url)
	if err := cmd.Start(); err != nil {
		log.Fatalf("open browser failed: %+v", err)
	}
}
