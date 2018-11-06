package ipsum

import (
  "net/http"
  "io/ioutil"
  "strings"
  "math/rand"
  "time"
)

var URL string = "https://gist.githubusercontent.com/nbw/fd9a78b0afed83d77c876da3166d6c52/raw/82f142562cf50b0f6fb8010f890b2f934093553e/animals.txt"

func RandomIpsum(numWords int) (string) {
    var wordString string
    var n, wordListLen int

    rand.Seed(time.Now().Unix())

    wordList := *getWordList()
    wordListLen = len(wordList)

    for i := 0; i < numWords; i++ {
      n = rand.Int() % wordListLen
      wordString += wordList[n]
      wordString += " "
    }

    return wordString
}

func getWordList()(*[]string){
    resp, err := http.Get(URL)
    if err != nil {
        return nil
    }
    defer resp.Body.Close()

    bodyBytes, err := ioutil.ReadAll(resp.Body)
    if err != nil {
      return nil
    }

    words := strings.Fields(string(bodyBytes))

    return &words
}

