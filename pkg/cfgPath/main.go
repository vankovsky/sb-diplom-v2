package cfgPath

import (
	"encoding/json"
	"fmt"
	"os"
)

func Init(filePath string) (*Root, error) {

	f, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("read file: %v", err)
	}

	var root Root
	if err := json.Unmarshal(f, &root); err != nil {
		return nil, fmt.Errorf("error config  parse: %w", err)
	}
	fmt.Println(root.HTTPServer.Port, root.HTTPService.Port)
	if err := root.initiate(); err != nil {
		return nil, err
	}

	return &root, nil
}
