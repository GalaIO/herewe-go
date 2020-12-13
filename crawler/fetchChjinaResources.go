package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path"
	"time"
)

const (
	MIME_OF_JSON          = "application/json"
	MIME_OF_GOOGLE_FOLDER = "application/vnd.google-apps.folder"
)

type FileInfo struct {
	Id           string `json:"id"`
	MimeType     string `json:"mimeType"`
	ModifiedTime string `json:"modifiedTime"`
	Name         string `json:"name"`
}

type DirInfo struct {
	Files []FileInfo `json:"files"`
}

var useProxy = false
var c *http.Client

// init config
func init() {
	if useProxy {
		os.Setenv("HTTP_PROXY", "http://127.0.0.1:1080")
		os.Setenv("HTTPS_PROXY", "http://127.0.0.1:1080")
	}
	c = &http.Client{}
	if useProxy {
		c = &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: true,
				},
				// 使用环境变量的代理
				Proxy: http.ProxyFromEnvironment,
			},
		}
	}
}

func main() {
	fetchPathInit()
	// run a goroutinur to parse dir
	go dirParseLoop()

	// run 3 goroutinue to download files
	for i := 0; i < 3; i++ {
		go downloadFileLoop()
	}
	select {}
	//err := fetchFromUrl("./chjina_tmp", "https://pan.chjina.com/Book/?rootId=root")
	//if err != nil {
	//	panic(err)
	//}
}

func fetchPathInit() {
	args := os.Args
	if len(args) < 2 {
		// read config file
		file, err := os.Open("fetchList.dat")
		if err != nil {
			panic(err)
		}
		bytes, err := ioutil.ReadAll(file)
		if err != nil {
			panic(err)
		}
		type FetchUrlInfo struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		}
		var list []FetchUrlInfo
		err = json.Unmarshal(bytes, &list)
		if err != nil {
			panic(err)
		}
		for _, l := range list {
			pushDirParse("./chjina_"+l.Name, l.Url+"?rootId=root")
		}
		return
	}
	//pushDirParse("./chjina_jiketime", "https://pan.chjina.com/%E5%85%B1%E4%BA%AB/%E8%AF%BE%E7%A8%8B/%E6%9E%81%E5%AE%A2%E6%97%B6%E9%97%B4%E5%B7%B2%E5%AE%8C%E7%BB%93/?rootId=root")
	pushDirParse("./chjina_jiketime", args[1])
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
		fmt.Println("find new dir dirPath:", ri.URL, "file:", ri.SaveDir)
		go func() {
			err := fetchFromUrl(ri.SaveDir, ri.URL)
			if err != nil {
				fmt.Println("parse dir err, url: ", ri.URL, err)
				pushDirParse(ri.SaveDir, ri.URL)
			}
		}()
		time.Sleep(time.Millisecond * 1000)
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
		fmt.Println("find new filePath:", ri.SaveDir, "file:", ri.URL)
		err := downloadFile(ri.SaveDir, ri.URL, ri.Name)
		if err != nil {
			fmt.Println("download file err, url: ", ri.URL, err)
			pushDownloadFile(ri.SaveDir, ri.URL, ri.Name)
		}
		time.Sleep(time.Millisecond * 500)
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
	u, err := url.Parse(requestPath)
	if err != nil {
		return err
	}

	preUrl := u.Scheme + "://" + u.Host + u.Path
	req, err := http.NewRequest("POST", requestPath, nil)
	if err != nil {
		return err
	}
	resp, err := c.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	//if MIME_OF_JSON != resp.Header.Get("content-type") {
	//	return errors.New(fmt.Sprintf("get %s requestPath, return wrong format", requestPath))
	//}
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	var dirInfo DirInfo
	err = json.Unmarshal(bytes, &dirInfo)
	if err != nil {
		return err
	}

	for _, f := range dirInfo.Files {
		newPath := path.Join(saveDir, f.Name)
		if MIME_OF_GOOGLE_FOLDER == f.MimeType {
			newUrl := preUrl + url.PathEscape(f.Name) + "/?" + u.Query().Encode()
			pushDirParse(newPath, newUrl)
			//fmt.Println("find new dir dirPath:", newUrl, "file:", newPath)
			//err = fetchFromUrl(newPath, newUrl)
			//if err != nil {
			//	fmt.Println("fetch dir files err, requestPath: ", newPath)
			//}
			continue
		}

		newUrl := preUrl + f.Name + "?" + u.Query().Encode()
		pushDownloadFile(newUrl, saveDir, f.Name)
		//fmt.Println("find new filePath:", newUrl, "file:", newPath)
		//err := downloadFile(newUrl, saveDir, f.Name)
		//if err != nil {
		//	fmt.Println("download err, requestPath: ", newPath)
		//}
	}

	return nil
}

func downloadFile(url, dir, name string) error {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		os.MkdirAll(dir, os.ModePerm)
	}
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// 创建一个文件用于保存
	out, err := os.Create(path.Join(dir, name))
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
