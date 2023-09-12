package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

func MustParseURL(rawUrl string) *url.URL {
	parsedURL, err := url.Parse(rawUrl)
	if err != nil {
		panic(err)
	}
	return parsedURL
}

func basicAuth(username, password string) string {
	auth := username + ":" + password
	return "Basic " + base64.StdEncoding.EncodeToString([]byte(auth))
}

type ResponseList struct {
	Data struct {
		ProxyList []string `json:"proxy_list"`
	} `json:"data"`
}

func getProxyIPList(api string) ([]string, error) {
	resp, err := http.Get(api)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var response ResponseList
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	proxyList := response.Data.ProxyList
	ipPortList := make([]string, 0)
	usefulTimeList := make([]string, 0)
	// 现有的 tdps 机房 ip
	ipList := []string{
		"182.106.136.9",
		"182.106.136.6",
		"219.150.218.21",
		"123.160.10.195",
		"221.229.212.174",
		"221.229.212.173",
		"221.229.212.172",
		"221.229.212.171",
		"221.229.212.158",
		"221.229.212.170",
		"221.229.212.156",
		"221.229.212.145",
		"182.106.136.217",
		"182.106.136.210",
	}

	for _, jdeProxy := range ipList {
		for _, proxy := range proxyList {
			ip := strings.Split(proxy, ":")[0]
			if jdeProxy == ip {
				parts := strings.Split(proxy, ",")
				if len(parts) == 2 {
					ipPortList = append(ipPortList, parts[0])
					usefulTimeList = append(usefulTimeList, parts[1])
				}
			}
		}
	}

	// 获取到的 tdps 数量
	tdpsNum := strconv.Itoa(len(ipPortList))
	// 打印提取的数据
	fmt.Println("获取到的 TDPS IP 数量:", tdpsNum)
	log.Println("获取到的 TDPS IP 数量:", tdpsNum)
	fmt.Println("IP 地址和端口列表:", ipPortList)
	log.Println("IP 地址和端口列表:", ipPortList)
	fmt.Println("IP 有效时间列表:", usefulTimeList)
	log.Println("IP 有效时间列表:", usefulTimeList)
	return ipPortList, nil
}

func main() {
	// 记录程序开始时间
	elapsedStartTime := time.Now()

	// 创建日志文件
	// 保留之前的日志: file, err := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	currentTime := time.Now().Format("2006-01-02 15:04:05")
	fmt.Printf("当前时间: %s\n", currentTime)
	currentDate := time.Now().Format("2006-01-02")
	logFileName := "./logs/" + currentDate + ".log" // 自定义日志文件名
	file, err := os.Create(logFileName)
	if err != nil {
		log.Fatal("Cannot create log file:", err)
	}
	defer file.Close()

	// 设置日志输出到文件
	log.SetOutput(file)

	// 获取 ip 的个数
	numIp := "5"
	secretId := "ote062zrad5uei90j50l"
	signature := "w8w78s47sx867mu0cmwxwi92bsd98ikz"
	// tdps 权重
	transferIp := "1"
	// crazy 狂暴模式, 一次最多提取 5000 个 ip
	crazyModel := "1"
	api := "https://test1.kuaidaili.com/api/getdps/?secret_id=" + secretId + "&num=" + numIp + "&signature=" + signature + "&f_et=1&format=json&sep=1&transferip=" + transferIp + "&crazy=" + crazyModel
	fmt.Println(api)

	getIpStartTime := time.Now()
	ipPortList, e := getProxyIPList(api)
	// 记录结束时间
	getIpEndTime := time.Now()
	// 计算运行时间
	getIpTotalTime := getIpEndTime.Sub(getIpStartTime)
	fmt.Printf("Get ip time: %s\n", getIpTotalTime)
	if e != nil {
		fmt.Println("Error:", e)
		log.Println("Error:", e)
		return
	}

	// 待请求的 URL
	urls := []string{
		//"https://www.tianyancha.com/company/1609647270",
		//"https://www.xiaohongshu.com/discovery/item/62c2751d0000000004004094",
		//"https://www.tianyancha.com/company/1609647270",
		//"https://beijing.anjuke.com/community/p3/",
		////"http://yangkeduo.com/proxy/api/api/caterham/query/fenlei_gyl_group?pdduid=4&page_sn=10003&opt_type=4&count=20&support_types=0_4&opt_name=&opt_id=14&offset=40",
		//"https://baijiahao.baidu.com/s?id=1759777615087439821&wfr=spider&for=pc",
		"https://www.baidu.com",
	}

	// 自定义请求头
	headers := map[string]string{
		"User-Agent":    "MyCustomUserAgent",
		"Custom-Header": "CustomValue",
	}

	// 整体响应状态码为 200 的次数
	var successOkStatusCodeAll int

	// 设置要发送的请求数量和每秒内的请求数
	numRequests := 100
	requestsPerSecond := 500

	// 创建一个通道来控制每秒内的请求数
	rateLimiter := make(chan struct{}, requestsPerSecond)

	for _, proxyIp := range ipPortList {
		// 单个 ip 响应状态码为 200 的次数
		var successOkStatusCode int

		// 记录请求开始时间
		ipRequestStartTime := time.Now()

		// 代理服务器地址和端口
		proxyAddress := "http://" + proxyIp

		// 打印获取到的 ip
		fmt.Printf("Current use ip: %s\n", proxyIp)

		// 代理认证信息
		proxyUsername := "d4594765938"
		proxyPassword := "n39kmzu8"

		// 创建一个代理认证的 HTTP 客户端
		client := &http.Client{
			// 8s Timeout
			Timeout: 8 * time.Second,
			Transport: &http.Transport{
				Proxy: http.ProxyURL(MustParseURL(proxyAddress)),
				ProxyConnectHeader: http.Header{
					"Proxy-Authorization": []string{basicAuth(proxyUsername, proxyPassword)},
				},
			},
		}

		var wg sync.WaitGroup
		wg.Add(numRequests)

		// 发起并发请求
		for i := 0; i < numRequests; i++ {
			rateLimiter <- struct{}{} // 获取令牌

			go func(requestNum int) {
				defer wg.Done()

				// 生成随机索引
				randomIndex := rand.Intn(len(urls))

				// 根据随机索引选择一个字符串
				reqUrl := urls[randomIndex]

				// 创建自定义请求
				req, er := http.NewRequest("GET", reqUrl, nil)
				if er != nil {
					fmt.Println("Error creating request:", er)
					log.Println("Error creating request:", er)
					return
				}

				// 设置自定义请求头
				for key, value := range headers {
					req.Header.Set(key, value)
				}

				// 发起请求
				resp, r := client.Do(req)

				if r != nil {
					fmt.Println("Error:", r)
					log.Println("Error:", r)
					return
				}

				defer resp.Body.Close()

				if resp.StatusCode == http.StatusOK {
					successOkStatusCode++
				}

				fmt.Printf("Response from %s: Status %s\n", reqUrl, resp.Status)
				log.Printf("Response from %s: Status %s\n", reqUrl, resp.Status)
			}(i + 1)

			// 释放令牌
			<-rateLimiter

			// 控制每秒发出的请求数
			if (i+1)%requestsPerSecond == 0 && i+1 < numRequests {
				time.Sleep(time.Second)
			}
		}

		// 等待所有请求完成
		wg.Wait()

		// 计算单个 ip 请求成功率
		successOkStatusCodeRate := float64(successOkStatusCode) / float64(numRequests) * 100
		fmt.Printf("Current ip success 200 rate: %.2f%%\n", successOkStatusCodeRate)
		log.Printf("Current ip success 200 rate: %.2f%%\n", successOkStatusCodeRate)

		// 记录程序结束时间
		ipRequestEndTime := time.Now()
		// 计算运行时间
		ipRequestTotalTime := ipRequestEndTime.Sub(ipRequestStartTime)
		fmt.Printf("Ip requests time: %s\n", ipRequestTotalTime)
		log.Printf("Ip requests time: %s\n", ipRequestTotalTime)

		// 整体请求成功次数
		successOkStatusCodeAll += successOkStatusCode

		// 记录到目前为止所花费的时间
		currentSpendTime := time.Now()
		// 计算运行时间
		currentElapsedTime := currentSpendTime.Sub(elapsedStartTime)
		fmt.Printf("Current elapsed time: %s\n", currentElapsedTime)
		log.Printf("Current elapsed time: %s\n", currentElapsedTime)
	}

	// 请求总数
	numAllRequests := numRequests * len(ipPortList)

	// 计算整体请求成功率
	successOkStatusCodeRate := float64(successOkStatusCodeAll) / float64(numAllRequests) * 100
	fmt.Printf("All requests num: %d, Success 200 rate all: %.2f%%\n", numAllRequests, successOkStatusCodeRate)
	log.Printf("All requests num: %d, Success 200 rate all: %.2f%%\n", numAllRequests, successOkStatusCodeRate)

	// 记录程序结束时间
	elapsedEndTime := time.Now()
	// 计算运行时间
	elapsedTotalTime := elapsedEndTime.Sub(elapsedStartTime)
	fmt.Printf("Total elapsed time: %s\n", elapsedTotalTime)
	log.Printf("Total elapsed time: %s\n", elapsedTotalTime)
}
