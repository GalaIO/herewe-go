package main

import (
	"crypto/tls"
	"fmt"
	"html"
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
	"41 Nginx核心知识100讲/": "https://d.shikey.com/jike/%E5%B7%B2%E5%AE%8C%E7%BB%93%E7%9A%84%E8%AF%BE%E7%A8%8B/41%20Nginx%E6%A0%B8%E5%BF%83%E7%9F%A5%E8%AF%86100%E8%AE%B2/",
	"42 MySQL实战45讲/":    "https://d.shikey.com/jike/%E5%B7%B2%E5%AE%8C%E7%BB%93%E7%9A%84%E8%AF%BE%E7%A8%8B/42%20MySQL%E5%AE%9E%E6%88%9845%E8%AE%B2/",
	"33 算法面试通关40讲/":     "https://d.shikey.com/jike/%E5%B7%B2%E5%AE%8C%E7%BB%93%E7%9A%84%E8%AF%BE%E7%A8%8B/33%20%E7%AE%97%E6%B3%95%E9%9D%A2%E8%AF%95%E9%80%9A%E5%85%B340%E8%AE%B2/",
	"14 深入浅出gRPC/":      "https://d.shikey.com/jike/%E5%B7%B2%E5%AE%8C%E7%BB%93%E7%9A%84%E8%AF%BE%E7%A8%8B/14%20%E6%B7%B1%E5%85%A5%E6%B5%85%E5%87%BAgRPC/",
	"57 Vue开发实战/":       "https://d.shikey.com/jike/%E5%B7%B2%E5%AE%8C%E7%BB%93%E7%9A%84%E8%AF%BE%E7%A8%8B/57%20Vue%E5%BC%80%E5%8F%91%E5%AE%9E%E6%88%98/",
	"58 10x程序员工作法/":     "https://d.shikey.com/jike/%E5%B7%B2%E5%AE%8C%E7%BB%93%E7%9A%84%E8%AF%BE%E7%A8%8B/58%2010x%E7%A8%8B%E5%BA%8F%E5%91%98%E5%B7%A5%E4%BD%9C%E6%B3%95/",
	//"30软件测试52讲/":     "https://d.shikey.com/jike/%E6%9E%81%E5%AE%A2%E6%97%B6%E9%97%B4%E5%B7%B2%E5%AE%8C%E7%BB%93/30%E8%BD%AF%E4%BB%B6%E6%B5%8B%E8%AF%9552%E8%AE%B2/",
	//"29趣谈网络协议音频修复版/": "https://d.shikey.com/jike/%E6%9E%81%E5%AE%A2%E6%97%B6%E9%97%B4%E5%B7%B2%E5%AE%8C%E7%BB%93/29%E8%B6%A3%E8%B0%88%E7%BD%91%E7%BB%9C%E5%8D%8F%E8%AE%AE%E9%9F%B3%E9%A2%91%E4%BF%AE%E5%A4%8D%E7%89%88/",
	//"06 9小时搞定微信小程序开发（完结）/": "https://d.shikey.com/jike/%E6%9E%81%E5%AE%A2%E6%97%B6%E9%97%B4%E5%B7%B2%E5%AE%8C%E7%BB%93/06%209%E5%B0%8F%E6%97%B6%E6%90%9E%E5%AE%9A%E5%BE%AE%E4%BF%A1%E5%B0%8F%E7%A8%8B%E5%BA%8F%E5%BC%80%E5%8F%91%EF%BC%88%E5%AE%8C%E7%BB%93%EF%BC%89/",
	//"11 大咖说/":              "https://d.shikey.com/jike/%E6%9E%81%E5%AE%A2%E6%97%B6%E9%97%B4%E5%B7%B2%E5%AE%8C%E7%BB%93/11%20%E5%A4%A7%E5%92%96%E8%AF%B4/",

	// downloaded
	//"56 程序员工作法":            "https://d.shikey.com/jike/%E6%9E%81%E5%AE%A2%E6%97%B6%E9%97%B4%E5%B7%B2%E5%AE%8C%E7%BB%93/56%20%E7%A8%8B%E5%BA%8F%E5%91%98%E5%B7%A5%E4%BD%9C%E6%B3%95/",
	//"54 技术领导力300讲-TGO鲲鹏会":  "https://d.shikey.com/jike/%E6%9E%81%E5%AE%A2%E6%97%B6%E9%97%B4%E5%B7%B2%E5%AE%8C%E7%BB%93/54%20%E6%8A%80%E6%9C%AF%E9%A2%86%E5%AF%BC%E5%8A%9B300%E8%AE%B2-TGO%E9%B2%B2%E9%B9%8F%E4%BC%9A/html/",
	//"18 朱赟技术管理课":           "https://d.shikey.com/jike/%E6%9E%81%E5%AE%A2%E6%97%B6%E9%97%B4%E5%B7%B2%E5%AE%8C%E7%BB%93/18%20%E6%9C%B1%E8%B5%9F%E6%8A%80%E6%9C%AF%E7%AE%A1%E7%90%86%E8%AF%BE/",
	//"41秒杀学习资料":             "https://d.shikey.com/jike/%E6%9E%81%E5%AE%A2%E6%97%B6%E9%97%B4%E5%B7%B2%E5%AE%8C%E7%BB%93/41%E7%A7%92%E6%9D%80%E5%AD%A6%E4%B9%A0%E8%B5%84%E6%96%99/",
	//"57 面试现场":              "https://d.shikey.com/jike/%E6%9E%81%E5%AE%A2%E6%97%B6%E9%97%B4%E5%B7%B2%E5%AE%8C%E7%BB%93/57%20%E9%9D%A2%E8%AF%95%E7%8E%B0%E5%9C%BA/",
	//"05 从零开始学架构/":          "https://d.shikey.com/jike/%E6%9E%81%E5%AE%A2%E6%97%B6%E9%97%B4%E5%B7%B2%E5%AE%8C%E7%BB%93/05%20%E4%BB%8E%E9%9B%B6%E5%BC%80%E5%A7%8B%E5%AD%A6%E6%9E%B6%E6%9E%84/",
	//"07 微服务架构核心20讲/":       "https://d.shikey.com/jike/%E6%9E%81%E5%AE%A2%E6%97%B6%E9%97%B4%E5%B7%B2%E5%AE%8C%E7%BB%93/07%20%E5%BE%AE%E6%9C%8D%E5%8A%A1%E6%9E%B6%E6%9E%84%E6%A0%B8%E5%BF%8320%E8%AE%B2/",
	//"14 深入浅出gRPC-李林峰/":     "https://d.shikey.com/jike/%E6%9E%81%E5%AE%A2%E6%97%B6%E9%97%B4%E5%B7%B2%E5%AE%8C%E7%BB%93/14%20%E6%B7%B1%E5%85%A5%E6%B5%85%E5%87%BAgRPC-%E6%9D%8E%E6%9E%97%E5%B3%B0/",
	//"17 持续交付36讲/":          "https://d.shikey.com/jike/%E6%9E%81%E5%AE%A2%E6%97%B6%E9%97%B4%E5%B7%B2%E5%AE%8C%E7%BB%93/17%20%E6%8C%81%E7%BB%AD%E4%BA%A4%E4%BB%9836%E8%AE%B2/",
	//"20 如何设计一个秒杀系统-许令波/":   "https://d.shikey.com/jike/%E6%9E%81%E5%AE%A2%E6%97%B6%E9%97%B4%E5%B7%B2%E5%AE%8C%E7%BB%93/20%20%E5%A6%82%E4%BD%95%E8%AE%BE%E8%AE%A1%E4%B8%80%E4%B8%AA%E7%A7%92%E6%9D%80%E7%B3%BB%E7%BB%9F-%E8%AE%B8%E4%BB%A4%E6%B3%A2/",
	//"23 技术与商业案例解读-徐飞/":     "https://d.shikey.com/jike/%E6%9E%81%E5%AE%A2%E6%97%B6%E9%97%B4%E5%B7%B2%E5%AE%8C%E7%BB%93/23%20%E6%8A%80%E6%9C%AF%E4%B8%8E%E5%95%86%E4%B8%9A%E6%A1%88%E4%BE%8B%E8%A7%A3%E8%AF%BB-%E5%BE%90%E9%A3%9E/",
	//"12 机器学习40讲-王天一":       "https://d.shikey.com/jike/%E6%9E%81%E5%AE%A2%E6%97%B6%E9%97%B4%E5%B7%B2%E5%AE%8C%E7%BB%93/12%20%E6%9C%BA%E5%99%A8%E5%AD%A6%E4%B9%A040%E8%AE%B2-%E7%8E%8B%E5%A4%A9%E4%B8%80/",
	//"50 Linux性能优化实战":       "https://d.shikey.com/jike/%E6%9E%81%E5%AE%A2%E6%97%B6%E9%97%B4%E5%B7%B2%E5%AE%8C%E7%BB%93/50%20Linux%E6%80%A7%E8%83%BD%E4%BC%98%E5%8C%96%E5%AE%9E%E6%88%98/",
	//"58 重学前端/":             "https://d.shikey.com/jike/%E6%9E%81%E5%AE%A2%E6%97%B6%E9%97%B4%E5%B7%B2%E5%AE%8C%E7%BB%93/58%20%E9%87%8D%E5%AD%A6%E5%89%8D%E7%AB%AF/",
	//"60 软件工程之美/":           "https://d.shikey.com/jike/%E6%9E%81%E5%AE%A2%E6%97%B6%E9%97%B4%E5%B7%B2%E5%AE%8C%E7%BB%93/60%20%E8%BD%AF%E4%BB%B6%E5%B7%A5%E7%A8%8B%E4%B9%8B%E7%BE%8E/",
	//"59 程序员进阶攻略（61讲）/":     "https://d.shikey.com/jike/%E6%9E%81%E5%AE%A2%E6%97%B6%E9%97%B4%E5%B7%B2%E5%AE%8C%E7%BB%93/59%20%E7%A8%8B%E5%BA%8F%E5%91%98%E8%BF%9B%E9%98%B6%E6%94%BB%E7%95%A5%EF%BC%8861%E8%AE%B2%EF%BC%89/",
	//"53Go语言从入门到实战/":        "https://d.shikey.com/jike/%E6%9E%81%E5%AE%A2%E6%97%B6%E9%97%B4%E5%B7%B2%E5%AE%8C%E7%BB%93/53Go%E8%AF%AD%E8%A8%80%E4%BB%8E%E5%85%A5%E9%97%A8%E5%88%B0%E5%AE%9E%E6%88%98/",
	//"02 杨晓峰-.java核心技术36讲/": "https://d.shikey.com/jike/%E6%9E%81%E5%AE%A2%E6%97%B6%E9%97%B4%E5%B7%B2%E5%AE%8C%E7%BB%93/02%20%E6%9D%A8%E6%99%93%E5%B3%B0-.java%E6%A0%B8%E5%BF%83%E6%8A%80%E6%9C%AF36%E8%AE%B2/",
	//"03 深入浅出区块链/":          "https://d.shikey.com/jike/%E6%9E%81%E5%AE%A2%E6%97%B6%E9%97%B4%E5%B7%B2%E5%AE%8C%E7%BB%93/03%20%E6%B7%B1%E5%85%A5%E6%B5%85%E5%87%BA%E5%8C%BA%E5%9D%97%E9%93%BE/",
	//"04-白话法律42讲（完结）/":      "https://d.shikey.com/jike/%E6%9E%81%E5%AE%A2%E6%97%B6%E9%97%B4%E5%B7%B2%E5%AE%8C%E7%BB%93/04-%E7%99%BD%E8%AF%9D%E6%B3%95%E5%BE%8B42%E8%AE%B2%EF%BC%88%E5%AE%8C%E7%BB%93%EF%BC%89/",
	//"38微服务架构实战160讲/":       "https://d.shikey.com/jike/%E6%9E%81%E5%AE%A2%E6%97%B6%E9%97%B4%E5%B7%B2%E5%AE%8C%E7%BB%93/38%E5%BE%AE%E6%9C%8D%E5%8A%A1%E6%9E%B6%E6%9E%84%E5%AE%9E%E6%88%98160%E8%AE%B2/",
	//"35Go语言核心36讲/": "https://d.shikey.com/jike/%E6%9E%81%E5%AE%A2%E6%97%B6%E9%97%B4%E5%B7%B2%E5%AE%8C%E7%BB%93/35Go%E8%AF%AD%E8%A8%80%E6%A0%B8%E5%BF%8336%E8%AE%B2/",
	//"46 数据结构与算法之美/": "https://d.shikey.com/jike/%E6%9E%81%E5%AE%A2%E6%97%B6%E9%97%B4%E5%B7%B2%E5%AE%8C%E7%BB%93/46%20%E6%95%B0%E6%8D%AE%E7%BB%93%E6%9E%84%E4%B8%8E%E7%AE%97%E6%B3%95%E4%B9%8B%E7%BE%8E/",
}

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
	//err := fetchFromUrl("./chjina_tmp", "https://pan.chjina.com/Book/?rootId=root")
	//if err != nil {
	//	panic(err)
	//}
}

func fetchPathInit() {
	//pushDirParse("/Users/galaio/Documentss/ziliao", "https://d.shikey.com/jike/%E6%9E%81%E5%AE%A2%E6%97%B6%E9%97%B4%E5%B7%B2%E5%AE%8C%E7%BB%93/")
	for key, val := range whiteNameList {
		pushDirParse("/data1/baidupan/"+key, val)
		fetchDirStat.IncrTotal()
	}

	//args := os.Args
	//if len(args) < 2 {
	//	// read config file
	//	file, err := os.Open("fetchList.dat")
	//	if err != nil {
	//		panic(err)
	//	}
	//	bytes, err := ioutil.ReadAll(file)
	//	if err != nil {
	//		panic(err)
	//	}
	//	type FetchUrlInfo struct {
	//		Name string `json:"name"`
	//		Url string `json:"url"`
	//	}
	//	var list []FetchUrlInfo
	//	err = json.Unmarshal(bytes, &list)
	//	if err != nil {
	//		panic(err)
	//	}
	//	for _, l := range list {
	//		pushDirParse("./chjina_"+l.Name, l.Url+"?rootId=root")
	//	}
	//	return
	//}
	//pushDirParse("./chjina_jiketime", args[1])
}

type RequestInfo struct {
	SaveDir string
	URL     string
	Name    string
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
				fmt.Println("parse dir err, url: ", ri.URL, err)
				pushDirParse(ri.SaveDir, ri.URL)
			} else {
				fetchDirStat.IncrFetched()
			}
		}()
		time.Sleep(time.Millisecond * 2000)
	}
}
func pushDirParse(saveDir, requestPath string) {
	dirChannel <- RequestInfo{
		SaveDir: saveDir,
		URL:     requestPath,
	}
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

func fetchFromUrl(saveDir, requestPath string) error {
	c := newClient()
	req, err := http.NewRequest("GET", requestPath, nil)
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
	fmt.Println("fetch dir:", requestPath, "resp:", resp.Status, "length:", resp.ContentLength)
	// Check server response
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status: %s", resp.Status)
	}

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	re := regexp.MustCompile("<tr><td><A HREF=\"(.*?)\">(.*?)</A></td>")
	submatch := re.FindAllStringSubmatch(string(bytes), -1)
	for _, match := range submatch {
		name := match[2]
		rel := match[1]
		rel = html.UnescapeString(rel)
		newUrl := "https://d.shikey.com" + rel
		if strings.HasSuffix(name, "/") {
			pushDirParse(path.Join(saveDir, name), newUrl)
			fetchDirStat.IncrTotal()
			continue
		}
		pushDownloadFile(newUrl, saveDir, name)
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
