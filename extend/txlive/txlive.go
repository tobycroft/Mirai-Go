package txlive

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

const ()

var APPID string
var BIZID string
var TencentKey string
var Key string

/**
 * 禁止推流
 * @liuhan
 */
func CloseStream(streamId string) bool {
	fmt.Println("streamId", streamId)
	baseTime := time.Now()
	sign, endDay := GetSign(baseTime)
	param := map[string]string{
		"appid":               APPID,
		"interface":           "Live_Channel_SetStatus",
		"Param.s.channel_id":  BIZID + "_" + streamId,
		"t":                   endDay,
		"sign":                sign,
		"Param.n.abstime_end": endDay,
		"Param.s.action":      "forbid",
	}
	baseUrl := "http://fcgi.video.qcloud.com/common_access"
	s := HTTPGet(baseUrl, param)
	fmt.Println("s", string(s))
	return true
}

/**
 * 恢复推流
 * @liuhan
 */
func RecoveryStream(streamId string) bool {
	fmt.Println("streamId", streamId)
	baseTime := time.Now()
	sign, endDay := GetSign(baseTime)
	param := map[string]string{
		"appid":              APPID,
		"interface":          "Live_Channel_SetStatus",
		"Param.s.channel_id": BIZID + "_" + streamId,
		"t":                  endDay,
		"sign":               sign,
		"Param.n.status":     strconv.Itoa(1),
	}
	baseUrl := "http://fcgi.video.qcloud.com/common_access"
	s := HTTPGet(baseUrl, param)
	fmt.Println("s", string(s))
	return true
}

/**
 * 生成直播流
 * @liuhan
 */
func BuildStream(rtmp, rtmp_key, hls, hls_key, streamname string) (string, string) {
	txTime := BuildTxTime()
	//txTime = "5EB2DEFF"
	//txSecret := BuildTxSecret(streamname)
	stream1 := "rtmp://" + rtmp + "/live/" + streamname + "?txSecret=" + BuildTxSecret(streamname, rtmp_key) + "&txTime=" + txTime
	stream2 := "http://" + hls + "/live/" + streamname + ".m3u8?txSecret=" + BuildTxSecret(streamname, hls_key) + "&txTime=" + txTime
	return stream1, stream2
}

/**
 * 生成txTime
 * @liuhan
 */
func BuildTxTime() string {
	timestamp := time.Now().Unix()
	timestamp = timestamp + 86400
	basetime := strconv.FormatInt(timestamp, 16)
	return basetime
}

/**
 * 生成txSecret
 * @liuhan
 */
func BuildTxSecret(streamname, KEY string) string {
	stream_id := streamname
	//stream_id = ar_id
	txTime := BuildTxTime()
	//txTime="5EB2DEFF"
	kst := KEY + stream_id + txTime
	return Md5(kst)
}

/**
 * 房间人数
 * @liuhan
 */
func AudienceCount() {
	baseTime := time.Now()
	sign, endDay := GetSign(baseTime)
	param := map[string]string{
		"cmd":       APPID,
		"interface": "Get_LiveStat",
		"t":         endDay,
		"sign":      sign,
	}
	baseUrl := "http://statcgi.video.qcloud.com/common_access"
	s := HTTPGet(baseUrl, param)
	fmt.Println("s", string(s))
	getLiveStat := &GetLiveStat{}
	_ = json.Unmarshal(s, &getLiveStat)
	fmt.Println("getLiveStat", getLiveStat)
	fmt.Println("getLiveStat.Output.StreamInfo", getLiveStat.Output.StreamInfo)

}

type GetLiveStat struct {
	Errmsg  string  `json:"errmsg"`
	Message string  `json:"message"`
	Output  Output  `json:"output"`
	Ret     float64 `json:"ret"`
	Retcode float64 `json:"retcode"`
}
type Output struct {
	DataTime       string       `json:"data_time"`
	StreamCount    int          `json:"stream_count"`
	StreamInfo     []StreamInfo `json:"stream_info"`
	TotalBandwidth float64      `json:"total_bandwidth"`
	TotalOnline    int          `json:"total_online"`
}

type StreamInfo struct {
	Bandwidth  float64 `json:"bandwidth"`
	ClientIp   string  `json:"client_ip"`
	Flr        int     `json:"flr"`
	Fps        int     `json:"fps"`
	Online     int     `json:"online"`
	ServerIp   string  `json:"server_ip"`
	Speed      int     `json:"speed"`
	StreamName string  `json:"stream_name"`
	Time       int     `json:"time"`
}

/**
 * 截图回调
 * @liuhan
 */
//func ScreenshotCallback(c *f.Context) {
//	body, _ := ioutil.ReadAll(c.Request.Body)
//	fmt.Println("stringbody", string(body))
//	screenshot := Screenshot{}
//	_ = json.Unmarshal(body, &screenshot)
//	fmt.Println("screenshot", screenshot)
//	fmt.Println("ChannelId", screenshot.ChannelId)
//	fmt.Println("PicFullUrl", screenshot.PicFullUrl)
//	c.String(0, "code")
//}

type Screenshot struct {
	ChannelId    string `json:"channel_id"`
	CreateTime   int    `json:"create_time"`
	EventType    int    `json:"event_type"`
	FileSize     int    `json:"file_size"`
	Height       int    `json:"height"`
	PicExtraInfo int    `json:"pic_extra_info"`
	PicFullUrl   string `json:"pic_full_url"`
	PicUrl       string `json:"pic_url"`
	Sign         string `json:"sign"`
	StreamId     string `json:"stream_id"`
	T            int    `json:"t"`
	Width        int    `json:"width"`
}

/**
 * 推流回调
 * @liuhan
 */
//func StreamCallback(c *f.Context) {
//	body, _ := ioutil.ReadAll(c.Request.Body)
//	fmt.Println("BBBBBBBBBBBBB", string(body))
//	streamMsg := StreamMsg{}
//	_ = json.Unmarshal(body, &streamMsg)
//	fmt.Println("streamMsg", streamMsg)
//	//event_type = 0 代表断流，event_type = 1 代表推流
//	fmt.Println("ChannelId", streamMsg.ChannelId)
//	fmt.Println("EventType", streamMsg.EventType)
//	arrStr := strings.Split(streamMsg.ChannelId, "_") //主播Id加房间Id
//	arr := strings.Split(arrStr[1], "-")
//	if streamMsg.EventType == 0 || streamMsg.EventType == 1 {
//		id, _ := strconv.ParseInt(arr[1], 10, 64)
//		go GoStreamCallback(id, uint64(streamMsg.EventType))
//	}
//	c.String(0, "code")
//}

type StreamMsg struct {
	App         string `json:"app"`
	Appid       int    `json:"appid"`
	Appname     string `json:"appname"`
	ChannelId   string `json:"channel_id"`
	Errcode     int    `json:"errcode"`
	Errmsg      string `json:"errmsg"`
	EventTime   int    `json:"event_time"`
	EventType   int    `json:"event_type"`
	IdcId       int    `json:"idc_id"`
	Node        string `json:"node"`
	Sequence    string `json:"sequence"`
	SetId       int    `json:"set_id"`
	Sign        string `json:"sign"`
	StreamId    string `json:"stream_id"`
	StreamParam string `json:"stream_param"`
	T           int    `json:"t"`
	UserIp      string `json:"user_ip"`
}

//腾讯云 签名
func GetSign(baseTime time.Time) (sign, endFay string) {
	date := baseTime.Add(24 * time.Hour).Unix()
	fmt.Println("date", date)
	endOfDayStr := strconv.FormatInt(date, 10)
	h := md5.New()
	h.Write([]byte(TencentKey + endOfDayStr))
	cipherStr := h.Sum(nil)
	hex.EncodeToString(cipherStr)
	return hex.EncodeToString(cipherStr), endOfDayStr
}
func Md5(str string) string {
	data := []byte(str)
	hash := md5.New()
	hash.Write(data)
	md := hash.Sum(nil)
	mdStr := hex.EncodeToString(md)
	return mdStr
}
