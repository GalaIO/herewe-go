package main

import (
	"crypto/tls"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"regexp"
	"strings"
	"sync/atomic"
	"time"
)

var whiteNameList = map[string]string{
	"": "",
}

var host = ""
var saveDir = ""
var defaultRefer = ""
var urlRegx = "<a href=\"(.*?)\">(.*?)</a>"
var useProxy = false
var transPort = &http.Transport{
	TLSClientConfig: &tls.Config{
		InsecureSkipVerify: true,
	},
	IdleConnTimeout: 0, // no limit避免自动关闭链接 导致下载中断
}

// init config
func init() {
	if useProxy {
		os.Setenv("HTTP_PROXY", "http://127.0.0.1:1080")
		os.Setenv("HTTPS_PROXY", "http://127.0.0.1:1080")
		// 使用环境变量的代理
		transPort.Proxy = http.ProxyFromEnvironment
	}
}

func newClient() *http.Client {
	return &http.Client{
		Transport: transPort,
	}
}

func main() {
	fetchPathInit()
	// run a goroutinue to parse dir
	go dirParseLoop()

	// run 3 goroutinue to download files
	for i := 0; i < 3; i++ {
		go downloadFileLoop()
	}

	// run download stat
	go downloadStat()

	select {}
}

func fetchPathInit() {
	for key, val := range whiteNameList {
		pushDirParse(saveDir+key, val)
		fetchDirStat.IncrTotal()
	}
}

type RequestInfo struct {
	SaveDir string
	URL     string
	Name    string
	Count   int64
}

// handler dir parse queue
var dirChannel = make(chan RequestInfo, 10000)

func dirParseLoop() {
	for {
		ri := <-dirChannel
		fmt.Println("find new dir path:", ri.URL, "file:", ri.SaveDir)
		go func() {
			err := fetchFromUrl(ri.SaveDir, ri.URL)
			if err != nil {
				if ri.Count >= 3 {
					fmt.Println("parse dir err, url: ", ri.URL, err, "retry count: ", ri.Count, "stop...")
					return
				}
				fmt.Println("parse dir err, url: ", ri.URL, err, "retry...")
				ri.Count++
				retryParse(ri)
			} else {
				fetchDirStat.IncrFetched()
			}
		}()
		time.Sleep(time.Millisecond * 2000)
	}
}
func pushDirParse(saveDir, dirPath string) {
	dirChannel <- RequestInfo{
		SaveDir: saveDir,
		URL:     dirPath,
	}
}
func retryParse(req RequestInfo) {
	dirChannel <- req
}

// handler dir parse queue
var downloadFileChannel = make(chan RequestInfo, 100000)

func downloadFileLoop() {
	for {
		ri := <-downloadFileChannel
		fmt.Println("find new file path:", ri.URL, "file:", ri.SaveDir)
		err := downloadFile(ri.SaveDir, ri.URL, ri.Name)
		if err != nil {
			fmt.Println("download file err, url: ", ri.URL, err)
			pushDownloadFile(ri.SaveDir, ri.URL, ri.Name)
		} else {
			// stat download info
			fetchFilesStat.IncrFetched()
		}
		time.Sleep(time.Millisecond * 1000)
	}
}
func pushDownloadFile(saveDir, requestPath, name string) {
	downloadFileChannel <- RequestInfo{
		SaveDir: saveDir,
		URL:     requestPath,
		Name:    name,
	}
}

func fetchFromUrl(saveDir, dirPath string) error {
	c := newClient()
	req, err := http.NewRequest("GET", dirPath, nil)
	if err != nil {
		return err
	}
	req.Header.Set("accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	req.Header.Set("referer", defaultRefer)
	req.Header.Set("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 11_0_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36")
	req.AddCookie(&http.Cookie{
		Name:  "timezone",
		Value: "8",
	})
	req.AddCookie(&http.Cookie{
		Name:  "theme",
		Value: "nginx.html",
	})
	resp, err := c.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	fmt.Println("fetch dir:", dirPath, "resp:", resp.Status, "length:", resp.ContentLength)
	// Check server response
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status: %s", resp.Status)
	}

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	re := regexp.MustCompile(urlRegx)
	submatch := re.FindAllStringSubmatch(string(bytes), -1)
	for _, match := range submatch {
		name := match[2]
		url := match[1]
		if strings.EqualFold(url, "./") ||
			strings.EqualFold(url, "../") {
			fmt.Println("err path", url, "ignore")
			continue
		}
		// abs url
		if strings.HasPrefix(url, "/") {
			url = host + url
		} else {
			//related url
			url = dirPath + url
		}
		if strings.HasSuffix(url, "/") {
			pushDirParse(path.Join(saveDir, name), url)
			fetchDirStat.IncrTotal()
			continue
		}
		pushDownloadFile(url, saveDir, name)
		// stat download info
		fetchFilesStat.IncrTotal()
	}

	return nil
}

func downloadFile(url, dir, name string) error {
	// 避免链接重用，导致读取流出现混乱，下载文件尤其是大文件乱码
	c := newClient()
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		os.MkdirAll(dir, os.ModePerm)
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	req.Header.Set("accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	req.Header.Set("referer", "https://d.shikey.com/jike/")
	req.Header.Set("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 11_0_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36")
	req.AddCookie(&http.Cookie{
		Name:  "timezone",
		Value: "8",
	})
	req.AddCookie(&http.Cookie{
		Name:  "theme",
		Value: "nginx.html",
	})
	resp, err := c.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// 创建一个文件用于保存
	newPath := path.Join(dir, name)
	fmt.Println("fetch:", newPath, "resp:", resp.Status, "length:", resp.ContentLength)
	// Check server response
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status: %s", resp.Status)
	}
	if _, err := os.Stat(newPath); os.IsExist(err) {
		fmt.Println("fetch:", newPath, "file existed, skip")
		return nil
	}
	out, err := os.Create(newPath)
	if err != nil {
		return err
	}
	defer out.Close()

	// 然后将响应流和文件流对接起来
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}

// stat modules
type FetchStatInfo struct {
	Total   int32
	Fetched int32
}

func (f *FetchStatInfo) Percent() int {
	if f.Total <= 0 {
		return 0
	}
	return int(float64(f.Fetched) / float64(f.Total) * 100)
}

func (f *FetchStatInfo) IncrTotal() {
	atomic.AddInt32(&f.Total, 1)
}

func (f *FetchStatInfo) IncrFetched() {
	atomic.AddInt32(&f.Fetched, 1)
}

// fetch file stat info
var fetchFilesStat = FetchStatInfo{
	Total:   0,
	Fetched: 0,
}

// fetch dir stat info
var fetchDirStat = FetchStatInfo{
	Total:   0,
	Fetched: 0,
}

func downloadStat() {
	for {
		fmt.Println("fetch dir stat, total:", fetchDirStat.Total, "fetched:", fetchDirStat.Fetched, "percent:", fetchDirStat.Percent(), "%")
		fmt.Println("fetch file stat, total:", fetchFilesStat.Total, "fetched:", fetchFilesStat.Fetched, "percent:", fetchFilesStat.Percent(), "%")
		time.Sleep(10 * time.Second)
	}
}
