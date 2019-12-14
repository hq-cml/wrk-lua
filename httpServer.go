package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
    "os"
    "strconv"
    "sync/atomic"
)

var helloCnt int32
var postCnt int32

type Node struct {
    ObjId int `json:"objId"`
    Score int `json:"score"`
    BizIds []int `json:"bizIds"`
    BizLevel int `json:"bizLevel"`
    UniqId string `json:"uniqId"`
}

func cnt(w http.ResponseWriter, req *http.Request) {
    s := strconv.Itoa(int(helloCnt))
    w.Write([]byte(s))
}

func postcnt(w http.ResponseWriter, req *http.Request) {
    s := strconv.Itoa(int(postCnt))
    w.Write([]byte(s))
}

func post(w http.ResponseWriter, req *http.Request) {
    atomic.AddInt32(&postCnt, 1)
    b, err := ioutil.ReadAll(req.Body)
    if err != nil {
        os.Exit(1)
    }
    defer req.Body.Close()

    node := Node{}
    err = json.Unmarshal(b, &node)
    if err != nil {
        fmt.Println(err)
        fmt.Println(string(b))
    } else {
        fmt.Println("ok--", string(b))
    }
    w.Write([]byte("ok"))
}

func hello(w http.ResponseWriter, req *http.Request) {
    atomic.AddInt32(&helloCnt, 1)
    //helloCnt ++
    w.Write([]byte("Hello"))
}

func main() {
    http.HandleFunc("/post", post)
    http.HandleFunc("/hello", hello)
    http.HandleFunc("/cnt", cnt)
    http.HandleFunc("/postcnt", postcnt)
    http.ListenAndServe(":9527", nil)
}
