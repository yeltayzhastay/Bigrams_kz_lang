package main

import (
    "encoding/csv"
    "fmt"
    "log"
    "os"
    "strings"
)

func readCsvFile(filePath string) [][]string {
    f, err := os.Open(filePath)
    if err != nil {
        log.Fatal("Unable to read input file " + filePath, err)
    }
    defer f.Close()

    csvReader := csv.NewReader(f)
    records, err := csvReader.ReadAll()
    if err != nil {
        log.Fatal("Unable to parse file as CSV for " + filePath, err)
    }

    return records
}

func writeCsvFile(filePath string, file [][]string){
  csvOut, err := os.Create(filePath)
  if err != nil {
    log.Fatal("Unable to open output")
  }
  w := csv.NewWriter(csvOut)
  defer csvOut.Close()
  w.WriteAll(file)
}

func comparison(st1 string, st2 string) string{
  a1:= []rune(st1)
  a2:= []rune(st2)
  length:= 0
  if(len(a1) > len(a2)){
    length = len(a2)
  }else{
    length = len(a1)
  }
  result := ""
  for i:=0;i < length; i++{
    if(a1[i] == a2[i]){
      result = result + string(a1[i])
    }else{
      break
    }
  }
  return result
}

func stemmingkazakh(docc [][]string, doc string) string{
    alldocin:= docc
    docin:= strings.Fields(doc)
    results:= ""
    for i:=0; i < len(alldocin); i++{
      for j:=0; j < len(docin); j++{
        s:= comparison(alldocin[i][0], docin[j])
        if(len(s) > 6){
          docin[j] = s
        }
      }
    }

    for j:=0; j < len(docin); j++{
      results += docin[j] + " "
    }
    return results
}

func preproccessing(uniq [][]string, dfs [][]string) [][]string{
  for i:=1; i<len(dfs); i++{
    dfs[i][0] = stemmingkazakh(uniq, dfs[i][0])
    fmt.Println(i)
  }
  return dfs
}

func main() {
    maindata := readCsvFile("C:/Users/Zhastay/_Projects in Univer(Python jupyter)/Bigrams_kz_lang/Corpus/corpus_cleaned.csv")
    uniquedata := readCsvFile("C:/Users/Zhastay/_Projects in Univer(Python jupyter)/Bigrams_kz_lang/Corpus/unique_cleaned_words.csv")
    fmt.Println("go")
    //fmt.Println(stemmingkazakh(uniquedata, maindata[1][0]))
    ending := preproccessing(uniquedata, maindata)
    fmt.Println("end")
    writeCsvFile("C:/Users/Zhastay/_Projects in Univer(Python jupyter)/Bigrams_kz_lang/stemming_fast_process/corpus_stemmed.csv", ending)
}
