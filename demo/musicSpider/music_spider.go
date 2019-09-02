package musicSpider

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

type Music struct {
	Type   string `json:"type"`
	Link   string `json:"link"`
	Songid string `json:"songid"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Url    string `json:"url"`
	Pic    string `json:"pic"`
}

func SpiderJayMusic() {
	url := "https://lvyueyang.github.io/web-music/system/musicList.json?timestamp=1567146846833"
	ret, err := http.Get(url)
	if err != nil {
		fmt.Println("Get content failed:", err)
		return
	}
	content, err := ioutil.ReadAll(ret.Body)

	var music []Music
	err = json.Unmarshal(content, &music)
	if err != nil {
		fmt.Println(err)
		return
	}

	ch := make(chan int, 3)
	lenCh := make(chan int, len(music))
	for i, m := range music {
		go downloadMusic(i, m, ch, lenCh)
	}

	var i int
	for {
		select {
		case <-lenCh:
			i++
			fmt.Printf("Finish [%d].\n", i)
			if i >= len(music) {
				fmt.Println("Download finished.")
				return
			}
		}
	}
}

func downloadMusic(i int, m Music, ch chan int, lenCh chan int) {
	ch <- 1
	lenCh <- 1
	url := fmt.Sprintf("http://jay-music.lvyueyang.top/%s.mp3", m.Songid)
	fmt.Println("Starting to download music:", m.Title)
	ret, err := http.Get(url)
	defer ret.Body.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	filename := fmt.Sprintf("./music/%d-%s(%s).mp3", i+1, m.Title, m.Author)
	fd, err := os.Create(filename)
	defer fd.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	//实现句柄拷贝
	_, err = io.Copy(fd, ret.Body)
	if err != nil {
		fmt.Println("copy error:", err)
		return
	}
	fmt.Printf("%s has been downloaded.\n", filename)
	<-ch
}
