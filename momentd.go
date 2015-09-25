package main

import (
	"fmt"

	"github.com/abhishekkr/goshare"

	momentdConfig "github.com/abhishekkr/momentdb/config"
	momentdbSplitter "github.com/abhishekkr/momentdb/splitter"
)

func banner() {
	asciiBanner := `
	*******************************************************************************************
	**.   *****    *******************************************************.'*   *****     *****
	**. *  ***  *  ***.'   ***.'  *****    **.'     **.'  ****  **'..   **. ***  ****  ***  ***
	**. **  *  **  **. ***  **. *  ***  *  **. *******. *  ***  ****. ****. ****  ***  **  ****
	**. ***   ***  **. ***  **. **  *  **  **.    ****. **  **  ****. ****. *****  **  *  *****
	**. *********  **. ***  **. ***   ***  **. *******. ***  *  ****. ****. ****  ***  **  ****
	**. *********  ***..   ***. *********  **.      **. ****    ****. ****. ***  ****  ***  ***
	**. *********  *******************************************************. *   *****      ****
	*******************************************************************************************
	`
	fmt.Println(asciiBanner)
}

func main() {
	banner()
	config := momentdConfig.ConfigFromFlags()

	switch config["type"] {
	case "splitter":
		momentdbSplitter.StartEngines(config["splitter"])

	case "goshare":
		goshare.GoShareEngine(config)

	default:
		panic(`Ah! Either this type of MomentDB service is WIP or just plain joke.
		Check our docs or just raise a bug if you think it need to work.`)
	}

	goshare.DoYouWannaContinue()
}
