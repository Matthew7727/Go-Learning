package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
)

type FileManger struct {
	InputFilePath string
	OutputFilePath string
}

func (fm FileManger) ReadLines() ([]string, error){
	file, err := os.Open(fm.InputFilePath)

	if err != nil {
		return nil, err
	}

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	err = scanner.Err()
	if err != nil {
		file.Close()
		return nil, err
	}

	file.Close()
	return lines, nil
}

func (fm FileManger) WriteJson( data interface{}) error{
	file , err := os.Create(fm.OutputFilePath)

	if err != nil {
		file.Close()
		return errors.New("failed to create file")
	}

	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)

	if err != nil {
		file.Close()
		return errors.New("failed to convert data to JSON")
		
	}

	file.Close()
	return nil

}

func New(inputPath, outputPath string) FileManger {
	return FileManger{
		InputFilePath: inputPath,
		OutputFilePath: outputPath,
	}

}