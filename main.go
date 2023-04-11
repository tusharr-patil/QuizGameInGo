package main

import (
  "fmt"
  "encoding/csv"
  "os"
)

func readCsvFile(filePath string) ([][]string, error){
  file, err := os.Open(filePath)
  if err != nil {
    return nil, err
  }
  defer file.Close()

  reader := csv.NewReader(file)
  records, err := reader.ReadAll()
  if err != nil {
    return nil, err
  }

  ques := make([][] string, len(records))
  for i, record := range records {
    ques[i] = make([]string, len(record))
    for j, value := range record {
      ques[i][j] = value
    }
  }

  return ques, nil
}

func main() {
  ques, err := readCsvFile("./QuizData.csv")
  if err != nil {
    fmt.Println(err)
  }
  
  n := len(ques)
  totalMarks := 0

  for i, que := range ques {
    var ans string
    fmt.Printf("Que %v: %v", i, que[0])
    fmt.Scan(&ans)
    if que[1] == ans {
      totalMarks++
    }
  }

  fmt.Printf("your score is %v out %v", totalMarks, n)
}
