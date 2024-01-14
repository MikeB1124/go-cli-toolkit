package csvUtil

import (
	"encoding/csv"
	"fmt"
	"os"
)

func ReadCSVFile(filePath string) ([]map[string]string, error) {
	// Open the CSV file
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}
	defer file.Close()

	// Create a CSV reader
	reader := csv.NewReader(file)

	// Read the header
	header, err := reader.Read()
	if err != nil {
		fmt.Println("Error reading header:", err)
		return nil, err
	}

	// Initialize an array of dictionaries
	var data []map[string]string

	// Read and process each record
	for {
		record, err := reader.Read()
		if err != nil {
			break // Stop reading at EOF
		}

		// Create a dictionary for the current record
		recordDict := make(map[string]string)
		for i, value := range record {
			recordDict[header[i]] = value
		}

		// Append the dictionary to the array
		data = append(data, recordDict)
	}

	// // Print the array of dictionaries
	// for _, recordDict := range data {
	// 	fmt.Println(recordDict)
	// }

	return data, nil
}
