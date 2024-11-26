package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"time"
)

func (tasks *Tasks) writeCsvFile(fileName string) error {
	file, err := os.Create(fileName)
	
	if err != nil {
		return err
	}

	writer := csv.NewWriter(file)

	for _, record := range *tasks {
		fmt.Printf("%v\n", record.Slice())
		writer.Write(record.Slice())
	}

	writer.Flush()

	return nil	
}

func readCsvFile(fileName string) (Tasks, error) {
	
	file, err := os.Open(fileName)

	if err != nil {
		fmt.Println("Can't Open File : ", fileName)
		return nil, err
	}

	defer file.Close()

	reader := csv.NewReader(file)
	
	t := Tasks{}	

	for {
		record, err := reader.Read()

		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Error Reading Record in file.")
			break
		}
		
		description := record[0]
		status, _ := strconv.ParseBool(record[1])
		createdAt, _ := time.Parse(time.RFC1123, record[2])
		
		var completedAt *time.Time
		if c := record[3] ; c == "nil" {
			completedAt = nil
		} else {
			time, _ := time.Parse(time.RFC1123, c) 
			completedAt = &time
		}
		t = append(t, Task{
			Description: description,
			Status: status,
			CreatedAt: createdAt,
			CompletedAt: completedAt,
		})
	}

	return t, nil
}