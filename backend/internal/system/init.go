package system

import (
	"fmt"
	"os"
)

const BootBaner string = `
  _____                 ____             _     
 |_   _|   _ _ __   ___|  _ \  ___   ___| | __ 
   | || | | | '_ \ / _ \ | | |/ _ \ / __| |/ / 
   | || |_| | | | |  __/ |_| | (_) | (__|   <  
   |_| \__,_|_| |_|\___|____/ \___/ \___|_|\_\ 
 ----------------------------------------------
`

func Init(){
	fmt.Print(BootBaner)
	_, err := os.Stat("./config/config.json")

	if err != nil {
		fmt.Println("Initialization in progress...")
		
	}
}
