package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/itchyny/volume-go"
)

func main() {
	volume := 30
	log.Println("start whitenoise")
	setVolumeGoal(volume)

	reader := bufio.NewReader(os.Stdin)

	for {
		_, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalf("get input failed: %+v", err)
		}
	}
}

func setVolumeGoal(vol int) {
	err := volume.SetVolume(vol)
	if err != nil {
		log.Fatalf("set volume failed: %+v", err)
	}
	fmt.Printf("set volume to: %d\n", vol)

	readVol, err := volume.GetVolume()
	if err != nil {
		log.Fatalf("get volume failed: %+v", err)
	}
	fmt.Printf("current volume: %d\n", readVol)
}
