package main

import (
	"encoding/json"
	"fmt"
	"github.com/valyala/fasthttp"
	"io"
	"math"
	"net/http"
	"net/url"
	"os"
	"strconv"
)

const (
	BaseParams string = "version_code=1.0.1&app_name=tiktok_snail&channel=App%20Store&device_id=4564563&aid=364225&os_version=16.2&device_platform=iphone&iid=7386407102867523334&device_brand=iphone&device_type=iPhone10,6"
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
		files, _ := os.ReadDir(keyword)
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
		fmt.Println("[🔍] searching for videos " + strconv.Itoa(offset+i*30+1) + "-" + strconv.Itoa(offset+(i+1)*30) + "...")
		// update our offset
		offset += i * 30

		// the count parameter tells the server how many videos to send back. there's a hard cap of 30 per request.
		count := 30

		// if this is the last run, we might not need to get a whole 30 videos
		if i == int(math.Ceil(float64(totalCount)/30.0))-1 && totalCount%30 != 0 {
			count = totalCount % 30
		}

		// set up and send our request with our search term
		var request = fasthttp.AcquireRequest()
		var response = fasthttp.AcquireResponse()

		request.Header.SetMethod("GET")
		request.SetRequestURI("https://api16-normal-c-alisg.tiktokv.com/aweme/v1/search/item/?" + BaseParams + "&count=" + strconv.Itoa(count) + "&offset=" + strconv.Itoa(offset) + "&keyword=" + url.QueryEscape(keyword))
		request.Header.Set("Cookie", "multi_sids=7386406673554703370%3A; passport_csrf_token=3c7388e9d6bffc7bdd8afa250a4a7bfa; passport_csrf_token_default=3c7388e9d6bffc7bdd8afa250a4a7bfa; lol")
		request.Header.Set("User-Agent", "Whee 1.0.1 rv:10102 (iPhone; iOS 16.2; en_US) Cronet")
		request.Header.Set("X-Argus", "igh7V38ke2stD2IEtD/yxbdK5VyH3Fk7paEj63GmsaiFbjsNZu2F13jMtG5BhvUyQva96w8B3FP0nScetWB4p0plAj0Ucl9QNCdvOGUzQS3rv3p6pIQ6/K4V/vYiEVVNZ+/q/Un1Yc2uqqPS/MfJwXRdrZOYf8wqPeEZPcVMIur48TX2kzXUHs6aoW38uzeA/0/zUBLXmlawAPxcWfHCgSARymhPYy40BT6IO1LL3wSxB0Fo5rDmSt3wlrMhOahjrrVwd/iXAVoKXb4R25HRT+jVM6GPQKNMR/TrEPvl51ckTxmFwIMtVIz6pUwlk+WCoD5tb6mArCyEBQ60ZCPrQ/ElrtLxPa7teBPWi2lPcWrX/SQnNFliJa/JtEKCibjnlpCbgoDEYQxCe0iG/d6aaf+eX7L5pk/isLAITXOvVvqSGsafBKY5Qh6lXL4dCET5mlN31p9ACq7U05gtl9Gy17nqG+fcpQuFdkv0Il7t5G+vOClcu7563jvp3zTlevmlX0qI1Y+OZXo4c4oyJ6AgyOw1")
		request.Header.Set("X-Gorgon", "840460220800b76857197a4198d6e74216d14534b9b5464b8cd9")
		request.Header.Set("X-Khronos", "1719783511")
		request.Header.Set("X-Ladon", "8b25KJ6jgQlLCnQatueMjCayo+M9J9+1JPKp7/rWz+7nHBu+")
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
			fmt.Printf("#%d / author: @%s / plays: %d / likes: %d / desc: %s\n", videoIndex, aweme.Author.UniqueId, aweme.Statistics.PlayCount, aweme.Statistics.DiggCount, aweme.Desc)

			downloadURL := aweme.Video.PlayAddr.UrlList[0]

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
