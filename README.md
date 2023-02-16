# 使用golang调用GPT-3的api

## config/Config.toml配置文件

``` toml
ApiKey = "YOUR APIKEY" #api密匙 这里更改为你的密匙即可

Url = "https://api.openai.com/v1/completions" #请求地址

[GptRuq]
    Model = "text-davinci-003"        #gpt模型
 Prompt = ""                       #输入给chat主体内容
 Temperature = 0.9                 #温度参数，用于控制生成文本的随机程度
 Max_tokens = 150                  #最大tokens数，003最大4000
 Top_p = 1                         #控制生成文本的多样性
 Frequency_penalty = 0             #用于降低生成文本中出现重复的单词或短语的概率
 Presence_penalty =  0.6           #用于降低生成文本中出现不相关单词或短语的概率
 Stop = ["HackerXiao:","Gpt-3-AI:"]  #返回信息的开头结尾                       
 N = 1                             #返回的信息数
```

## 界面

![test](download.png)
