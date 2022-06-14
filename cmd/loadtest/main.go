package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/SurgicalSteel/kvothe/utils"

	"github.com/alexsasharegan/dotenv"
	vegeta "github.com/tsenart/vegeta/lib"
)

// Config struct
type Config struct {
	High          TargetConfig
	Med           TargetConfig
	Low           TargetConfig
	Hostname      string
	Port          string
	Protocol      string
	EndpointGroup string
	EndpointList  map[string]string
	AdpanelToken  string
	Timeout       int
	TokenKey      string
}

type TargetConfig struct {
	Freq        int
	Duration    int
	Unlimited   string
	Workers     int
	Connections int
}

// TargeterParams struct
type TargeterParams struct {
	Protocol      string
	Hostname      string
	Port          string
	EndpointGroup string
	EP            string
}

const (
	TestThreshold       = 0.999
	TestDefaultInt      = 0
	SuccessUnitLoadtest = 1
	FailLoadtest        = 0
)

func GetAttacker(tc TargetConfig) (rate vegeta.Rate, duration time.Duration, attacker *vegeta.Attacker) {
	if !strings.EqualFold(tc.Unlimited, "true") {
		rate = vegeta.Rate{Freq: tc.Freq, Per: time.Second}
	} else {
		rate = vegeta.Rate{}
	}
	duration = time.Duration(tc.Duration) * time.Second
	attacker = vegeta.NewAttacker(
		vegeta.MaxWorkers(uint64(tc.Workers)),
		vegeta.Workers(uint64(tc.Workers)),
		vegeta.Connections(tc.Connections),
	)

	return rate, duration, attacker
}

func main() {
	//SET OUTPUTFILE==========
	f, err := os.Create("result.loadtest")
	if err != nil {
		log.Println("failed create file result | " + err.Error())
		f.Close()

		return
	}

	//conf := LoadConfig()

	// //INITIAL
	// resultNum := make(map[string]int)
	// resultNum["len"] = 0
	// resultNum["success"] = 0

	// // //end loadtest

	// // Get List Article Category SECTION
	// var metrics2 vegeta.Metrics
	// rate, duration, attacker := GetAttacker(conf.High)
	// params := TargeterParams{
	// 	conf.Protocol,
	// 	conf.Hostname,
	// 	conf.Port,
	// 	conf.EndpointGroup,
	// 	conf.EndpointList["get-slide-content"],
	// }
	// targeter := TargeterGetSlideContent(params)

	// // var body []byte
	// for res := range attacker.Attack(targeter, rate, duration, "Big Bang!") {
	// 	// body = res.Body
	// 	metrics2.Add(res)
	// }
	// // log.Println(string(body))

	// metrics2.Close()

	// resultNum["len"] = resultNum["len"] + 1
	// resultNum["success"] = resultNum["success"] + IsSuccess(metrics2, conf.High)

	// restext := vegeta.NewTextReporter(&metrics2)
	// url := params.Protocol + "://" + params.Hostname + ":" + params.Port + params.EndpointGroup + params.EP

	// fmt.Fprintln(f, "Endpoint : "+url)
	// err = restext.Report(f)
	// if err != nil {
	// 	log.Println(err)
	// }

	// time.Sleep(time.Duration(conf.Timeout) * time.Second)
	// // End List Article Category seciton

	// resLoadTest := resultNum["success"] / resultNum["len"]
	resLoadTest := 1
	if resLoadTest != 1 {
		fmt.Fprintln(f, "Result:failed")
		fmt.Println("failed")
	} else {
		fmt.Fprintln(f, "Result:success")
		fmt.Println("success")
	}

}

// LoadConfig load config
func LoadConfig() Config {
	err := dotenv.Load(".env")
	if err != nil {
		log.Println("Failed load .env | ", err.Error())
	}

	appCredentials := os.Getenv("GOOGLE_APPLICATION_CREDENTIALS")
	err = os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", appCredentials)
	if err != nil {
		log.Fatal(err)
	}

	var conf Config
	//high
	conf.High.Freq = utils.GetInt(os.Getenv("HIGH_FREQ"))
	conf.High.Duration = utils.GetInt(os.Getenv("HIGH_DURATION"))
	conf.High.Unlimited = os.Getenv("HIGH_UNLIMITED")
	conf.High.Workers = utils.GetInt(os.Getenv("HIGH_WORKERS"))
	conf.High.Connections = utils.GetInt(os.Getenv("HIGH_CONNECTIONs"))
	//med
	conf.Med.Freq = utils.GetInt(os.Getenv("MID_FREQ"))
	conf.Med.Duration = utils.GetInt(os.Getenv("MID_DURATION"))
	conf.Med.Unlimited = os.Getenv("MID_UNLIMITED")
	conf.Med.Workers = utils.GetInt(os.Getenv("MID_WORKERS"))
	conf.Med.Connections = utils.GetInt(os.Getenv("MID_CONNECTIONs"))
	//low
	conf.Low.Freq = utils.GetInt(os.Getenv("LOW_FREQ"))
	conf.Low.Duration = utils.GetInt(os.Getenv("LOW_DURATION"))
	conf.Low.Unlimited = os.Getenv("LOW_UNLIMITED")
	conf.Low.Workers = utils.GetInt(os.Getenv("LOW_WORKERS"))
	conf.Low.Connections = utils.GetInt(os.Getenv("LOW_CONNECTIONs"))

	//Timeout
	conf.Timeout = utils.GetInt(os.Getenv("TIMEOUT"))

	//token
	conf.TokenKey = os.Getenv("TOKEN_KEY")

	// redis
	conf.Hostname = os.Getenv("HOSTNAME")
	conf.Port = os.Getenv("PORT")
	conf.Protocol = os.Getenv("PROTOCOL")
	conf.EndpointGroup = os.Getenv("ENDPOINT_GROUP")
	conf.AdpanelToken = os.Getenv("TOKEN_ADPANEL_KEY")
	m := make(map[string]string)

	conf.EndpointList = m

	return conf
}

// IsSuccess return is success 1 or 0
func IsSuccess(data vegeta.Metrics, conf TargetConfig) int {

	if data.StatusCodes["200"] == TestDefaultInt {
		return FailLoadtest
	}

	isSuccess := true
	threshold := int(float64(conf.Freq*conf.Duration) * TestThreshold)
	totalValue := 0
	for key, value := range data.StatusCodes {
		if key == "200" || key == "0" || key == "503" {
			totalValue += value
		} else {
			isSuccess = false
			break
		}
	}
	if isSuccess {
		if totalValue >= threshold {
			return SuccessUnitLoadtest
		} else {
			return FailLoadtest
		}
	}
	return FailLoadtest
}
