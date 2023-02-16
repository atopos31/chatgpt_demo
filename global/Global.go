package Global

type GptRuq struct {
	Model             string    `mapstructure:"model" json:"model"`                       //gpt模型
	Prompt            string    `mapstructure:"prompt" json:"prompt"`                     //输入给chat主体内容
	Temperature       float64   `mapstructure:"temperature" json:"temperature"`           //温度参数，用于控制生成文本的随机程度
	Max_tokens        int       `mapstructure:"max_tokens" json:"max_tokens"`             //最大tokens数，003最大4000
	Top_p             int       `mapstructure:"top_p" json:"top_p"`                       //控制生成文本的多样性
	Frequency_penalty int       `mapstructure:"frequency" json:"frequency_penalty"`       //用于降低生成文本中出现重复的单词或短语的概率
	Presence_penalty  float64   `mapstructure:"presence_penalty" json:"presence_penalty"` //用于降低生成文本中出现不相关单词或短语的概率
	Stop              [2]string `mapstructure:"stop" json:"stop"`                         //返回信息的开头结尾
	N                 int       `mapstructure:"n" json:"n"`                               //生成的文本数量
}

type GptChoices struct {
	Text          string `json:"text"`          //生成的文本内容
	Finish_reason string `json:"finish_reason"` //生成文本的停止原因
}

type GptError struct {
	Message string `json:"message"` //报错信息
}

type GptRsp_S struct {
	Choices []GptChoices `json:"choices"` //包含生成的文本和相关信息的数组
	Error   GptError     `json:"error"`
}

type GptConfig_S struct {
	ApiKey string `mapstructure:"ApiKey" json:"apikey"`
	Url string `mapstructure:"Url" json:"url"`
	GptRuq GptRuq `mapstructure:"GptRuq"`
}

var GptConfig GptConfig_S
var GptRsp GptRsp_S
