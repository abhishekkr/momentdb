package main

import (
	"fmt"

	"github.com/abhishekkr/goshare"

	momentdb_splitter "./splitter"
)

func banner() {
	fmt.Println("**********************************************************************************")
	fmt.Println("**.   *****    *******************************************************.'*   ******")
	fmt.Println("**. *  ***  *  ***.'   ***.'  *****    **.'     **.'  ****  **'..   **. ***  *****")
	fmt.Println("**. **  *  **  **. ***  **. *  ***  *  **. *******. *  ***  ****. ****. ****  ****")
	fmt.Println("**. ***   ***  **. ***  **. **  *  **  **.    ****. **  **  ****. ****. *****  ***")
	fmt.Println("**. *********  **. ***  **. ***   ***  **. *******. ***  *  ****. ****. ****  ****")
	fmt.Println("**. *********  ***..   ***. *********  **.      **. ****    ****. ****. ***  *****")
	fmt.Println("**. *********  *******************************************************. *   ******")
	fmt.Println("**********************************************************************************")
}

func main() {
	banner()
	config := goshare.ConfigFromFlags()

	fmt.Println(config, "\n\n")
	momentdb_splitter.StartEngines("./node-config.json.sample")
	/*
		goshare.GoShareEngine(config)
	*/
	goshare.DoYouWannaContinue()
}
