package main

import (
	"encoding/json"
	"fmt"
	"github.com/valyala/fasthttp"
	"io"
	"io/ioutil"
	"math"
	"net/http"
	"net/url"
	"os"
	"strconv"
)

const (
	BaseParams string = "iid=7173886861033883398&device_id=7109589686008153606&ac=wifi&channel=googleplay&aid=1233&app_name=musical_ly&version_code=270204&version_name=27.2.4&device_platform=android&ab_version=27.2.4&ssmix=a&device_type=SM-N975F&device_brand=samsung&language=en&os_api=25&os_version=7.1.2&openudid=e424163b4fff5720&manifest_version_code=2022702040&resolution=800*1200&dpi=266&update_version_code=2022702040&_rticket=1670305079987&app_type=normal&sys_region=US&mcc_mnc=214214&timezone_name=America%2FChicago&ts=1670305079&timezone_offset=-21600&build_number=27.2.4&region=US&uoo=0&app_language=en&carrier_region=ES&locale=en&op_region=ES&ac2=wifi&host_abi=armeabi-v7a&cdid=2133c73a-770c-4aca-afe8-2daefb9e6946"
)

func main() {
	fmt.Println("[✨] ttDownload - v1.0")

	var keyword string
	fmt.Print("[❓] keyword: ")
	_, err := fmt.Scanln(&keyword)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	var totalCount int
	fmt.Print("[❓] count: ")
	_, err = fmt.Scanln(&totalCount)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// create our working directory, if it exists, count the files in the directory and set the offset so we don't download the same videos again
	offset := 0
	if exists(keyword) {
		files, _ := ioutil.ReadDir(keyword)
		offset = len(files)
	} else {
		err = os.Mkdir(keyword, os.ModePerm)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	}

	baseOffset := offset

	fmt.Println("[✨] grabbing videos " + strconv.Itoa(offset+1) + "-" + strconv.Itoa(offset+totalCount) + "...")

	videoIndex := 1
	for i := 0; i < int(math.Ceil(float64(totalCount)/30.0)); i++ { // looks scary, but the upper bound of this is simply the ceiling of the count divided by 30.0, the maximum amount of results the API can get us in a request
		// update our offset
		offset += i * 30

		// the count parameter tells the server how many videos to send back. there's a hard cap of 30 per request.
		count := 30

		// if this is the last run, we might not need to get a whole 30 videos
		if i == int(math.Ceil(float64(totalCount)/30.0))-1 && totalCount%30 != 0 {
			count = totalCount % 30
		}

		// generate the necessary signature headers
		var xGorgon, xKhronos = generateSignature(BaseParams + "&count=" + strconv.Itoa(count) + "&offset=" + strconv.Itoa(offset) + "&keyword=" + url.QueryEscape(keyword))

		// set up and send our request with our search term
		var request = fasthttp.AcquireRequest()
		var response = fasthttp.AcquireResponse()
		request.Header.SetMethod("POST")
		request.SetRequestURI("https://search19-normal-c-useast1a.tiktokv.com/aweme/v1/search/item/?" + BaseParams + "&count=" + strconv.Itoa(count) + "&offset=" + strconv.Itoa(offset) + "&keyword=" + url.QueryEscape(keyword))
		request.Header.Set("User-Agent", "com.zhiliaoapp.musically/2022702040 (Linux; U; Android 7.1.2; en; SM-N975F; Build/N2G48H;tt-ok/3.12.13.1)")
		request.Header.Set("X-Khronos", strconv.FormatInt(xKhronos, 10))
		request.Header.Set("X-Gorgon", xGorgon)
		err := fasthttp.Do(request, response)
		if err != nil {
			fmt.Println(err)
			return
		}

		// unmarshal the big JSON response to search, an instance of a struct used to store this JSON response
		var search SearchData
		err2 := json.Unmarshal(response.Body(), &search)
		if err != nil {
			fmt.Println(err2)
			return
		}

		// make sure we release the request and response!
		fasthttp.ReleaseRequest(request)
		fasthttp.ReleaseResponse(response)

		// loop through all the returned videos, print some metadata and download each one
		for _, aweme := range search.AwemeList {
			fmt.Printf("#%d / author: @%s / plays: %d / likes: %d / desc: %s\n", videoIndex, aweme.Author.UniqueID, aweme.Statistics.PlayCount, aweme.Statistics.DiggCount, aweme.Desc)

			downloadURL := aweme.Video.PlayAddr.URLList[0]

			video, _ := os.Create(keyword + "/" + strconv.Itoa(baseOffset+videoIndex) + ".mp4")

			resp, _ := http.Get(downloadURL)
			_, err := io.Copy(video, resp.Body)
			if err != nil {
				fmt.Println(err.Error())
				return
			}

			err = resp.Body.Close()
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			err = video.Close()
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			videoIndex++
		}
	}

}
