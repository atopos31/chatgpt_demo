package main

import (
	"bufio"
	"bytes"
	"chatgpt/Global"
	"chatgpt/Viper"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	Viper.Config() //加载配置文件
	ReqData := Global.GptConfig.GptRuq
	Cache := "" //缓存区，方标连续对话
	inpu := bufio.NewReader(os.Stdin)
	re := false

	fmt.Println("********************************************")
	fmt.Println("作者：hackerxiao,欢迎访问：hackerxiao.online")
	fmt.Printf("\n调用Gpt-3 API，开始和ChatGPT聊天吧!\n输入exit退出程序\n")
	fmt.Printf("*********************************************\n\n")

	for !re {
		fmt.Print(Global.GptConfig.GptRuq.Stop[0])

		input, _ := inpu.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "exit" {
			break
		}
		if input == "" {
			continue
		}

		ReqData.Prompt = Cache + "\n" + Global.GptConfig.GptRuq.Stop[0] + input + "\n" + Global.GptConfig.GptRuq.Stop[1]
		Cache = ReqData.Prompt
		aiRspText := ""
		jsonData, err := json.Marshal(ReqData)
		if err != nil {
			fmt.Printf("json数据格式化失败: %s\n", err)
			return
		}
		retry := 0
		for {
			// 创建一个http.Client对象
			client := &http.Client{}
			req, err := http.NewRequest("POST", Global.GptConfig.Url, bytes.NewBuffer(jsonData))
			if err != nil {
				fmt.Printf("创建HTTP请求出错: %s\n", err)
				return
			}
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("Authorization", "Bearer "+Global.GptConfig.ApiKey)

			//发送请求
			response, err := client.Do(req)
			if err != nil {
				fmt.Printf("HTTP请求失败：%s\n", err)
				re = true
				break
			} else {
				defer response.Body.Close()
				var rspData Global.GptRsp_S

				body, err := io.ReadAll(response.Body)
				if err != nil {
					panic(err)
				}

				err = json.Unmarshal(body, &rspData)
				if err != nil {
					fmt.Printf("返回数据解析为json失败: %s\n", err)
					re = true
					break
				} else {
					if rspData.Error.Message != "" && retry < 10 {
						fmt.Printf("ChatGPT连接出错:%v\n", rspData.Error.Message)
						retry++
						time.Sleep(time.Second)
						continue
					}

					Cache += rspData.Choices[0].Text
					aiRspText += rspData.Choices[0].Text

					// 判断回答是否结束
					if rspData.Choices[0].Finish_reason != "length" {
						break
					}

					// 继续循环
					ReqData.Prompt = Cache
					Global.GptConfig.GptRuq.N++
					jsonData, err = json.Marshal(ReqData)
					if err != nil {
						fmt.Printf("正文格式化为json失败: %s\n", err)
						re = true
						break
					}

				}

			}
		}
		if re {
			break
		}
		fmt.Printf("   %v%v\n", strings.TrimSpace(Global.GptConfig.GptRuq.Stop[1]), aiRspText) //打印返回的文本

	}
}
