package main

import (
	"fmt"

	"github.com/abhishekkr/goshare"
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
	goshare.GoShareEngine(config)
	goshare.DoYouWannaContinue()
}
