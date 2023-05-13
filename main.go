package main

import (
  "fmt"
  "encoding/csv"
  "os"
  "time"
  "flag"
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
  timeLimit := flag.Int("limit", 30, "the time limit for the quiz in seconds")
  flag.Parse()
  timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

  for i, que := range ques {
    fmt.Printf("Que %v: %v \n", i, que[0])
    ansCh := make(chan string)

    go func() {
      var ans string
      fmt.Scanf("%s \n", &ans)
      ansCh <- ans
    }()

    select {
    case <-timer.C: 
      fmt.Printf("\nYou Scored %d out %d \n", totalMarks, n)
      return 
    case ans := <-ansCh:
      if ans == que[1] {
        totalMarks++
      }
    }
  }
  fmt.Printf("\nYou Scored %d out %d \n", totalMarks, n)
}
