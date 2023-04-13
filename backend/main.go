package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
	"os"

	"github.com/gorilla/mux"
)

type hMaterialRequest struct{
	Q1Ans string `json:"q1Ans"`
	Q2Ans string `json:"q2Ans"`
	Q3Ans string `json:"q3Ans"`
	Q4Ans string `json:"q4Ans"`
	Q5Ans string `json:"q5Ans"`
	Q6Ans string `json:"q6Ans"`
	Q7Ans string `json:"q7Ans"`
	Q8Ans string `json:"q8Ans"`
	Q9Ans string `json:"q9Ans"`
	Q10Ans string `json:"q10Ans"`
	Q11Ans string `json:"q11Ans"`
	Q12Ans string `json:"q12Ans"`
	Q13Ans string `json:"q13Ans"`
	Q14Ans string `json:"q14Ans"`
	Q15Ans string `json:"q15Ans"`
	Q16Ans string `json:"q16Ans"`
	Q17Ans string `json:"q17Ans"`
	Q18Ans string `json:"q18Ans"`
	Q19Ans string `json:"q19Ans"`
	Q20Ans string `json:"q20Ans"`
	Name string `json:"name"`
}

type hMaterialResponse struct{
	HMaterialScore float64 `json:"hMaterialScore"`
	BMaterialScore int `json:"bMaterialScore"`
	PlayBoyScore int `json:"playBoyScore"`
}

func initializeRouter(){
	r := mux.NewRouter()
	log.Println(r)
	r.HandleFunc("/hMaterial", CheckHMaterial).Methods("POST", "OPTIONS")	
	r.HandleFunc("/okay", PrintOkay).Methods("GET", "OPTIONS")
	port := os.Getenv("PORT")

	if port == ""{
		port = "8080"
	}
	log.Println("Server running on port : ", port)
	log.Fatal(http.ListenAndServe(":" + port, r))

}

func PrintOkay(w http.ResponseWriter,r *http.Request){
	log.Println("Okay, request received")
	test := "Connected"
	json.NewEncoder(w).Encode(&test)
}

func CheckHMaterial(w http.ResponseWriter,r *http.Request){
	log.Println("API started")
	w.Header().Set("Content-Type", "text/html; charset=ascii")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers","Content-Type,access-control-allow-origin, access-control-allow-headers")
	now := time.Now()
	var hMaterialInstance hMaterialRequest
	json.NewDecoder(r.Body).Decode(&hMaterialInstance)
	log.Println(hMaterialInstance, "<- params")
	var hMaterialResponseInstance hMaterialResponse
	_, err := hMaterialResponseInstance.calculateHMaterial(&hMaterialInstance)
	if err != nil {
		log.Println(err.Error())
	}
	log.Println(hMaterialResponseInstance.HMaterialScore, " Final score for hmaterial")
	// hMaterialResponseInstance.calculatePercentage(score, 200)

	// hMaterialResponseInstance.HMaterialScore = scorePer
	hMaterialResponseInstance.BMaterialScore = 00
	hMaterialResponseInstance.PlayBoyScore = 00
	log.Println(time.Since(now), " <- time took for API")
	json.NewEncoder(w).Encode(&hMaterialResponseInstance)
}

func (h *hMaterialResponse)calculateHMaterial(hMaterialInstance *hMaterialRequest) (int, error) {
var hMaterialScore int
// var err error
// total score 200

log.Println("Here")
h.AddDatainStruct(h.checkScore(hMaterialInstance.Q1Ans, &h.HMaterialScore, false))
h.AddDatainStruct(h.checkScore(hMaterialInstance.Q2Ans, &h.HMaterialScore, false))
h.AddDatainStruct(h.checkScore(hMaterialInstance.Q3Ans, &h.HMaterialScore, false))
h.AddDatainStruct(h.checkScore(hMaterialInstance.Q4Ans, &h.HMaterialScore, false))
h.AddDatainStruct(h.checkScore(hMaterialInstance.Q5Ans, &h.HMaterialScore, false))
h.AddDatainStruct(h.checkScore(hMaterialInstance.Q6Ans, &h.HMaterialScore, false))
h.AddDatainStruct(h.checkScore(hMaterialInstance.Q7Ans, &h.HMaterialScore, false))
h.AddDatainStruct(h.checkScore(hMaterialInstance.Q8Ans, &h.HMaterialScore, false))
h.AddDatainStruct(h.checkScore(hMaterialInstance.Q9Ans, &h.HMaterialScore, false))
h.AddDatainStruct(h.checkScore(hMaterialInstance.Q10Ans, &h.HMaterialScore,false))
h.AddDatainStruct(h.checkScore(hMaterialInstance.Q11Ans, &h.HMaterialScore,false))
h.AddDatainStruct(h.checkScore(hMaterialInstance.Q12Ans, &h.HMaterialScore,false))
h.AddDatainStruct(h.checkScore(hMaterialInstance.Q13Ans, &h.HMaterialScore,false))
h.AddDatainStruct(h.checkScore(hMaterialInstance.Q14Ans, &h.HMaterialScore,false))
h.AddDatainStruct(h.checkScore(hMaterialInstance.Q15Ans, &h.HMaterialScore,false))
h.AddDatainStruct(h.checkScore(hMaterialInstance.Q16Ans, &h.HMaterialScore,false))
h.AddDatainStruct(h.checkScore(hMaterialInstance.Q17Ans, &h.HMaterialScore,false))
h.AddDatainStruct(h.checkScore(hMaterialInstance.Q18Ans, &h.HMaterialScore,false))
h.AddDatainStruct(h.checkScore(hMaterialInstance.Q19Ans, &h.HMaterialScore,false))
h.AddDatainStruct(h.checkScore(hMaterialInstance.Q20Ans, &h.HMaterialScore,true))

return hMaterialScore, nil
}

func (h *hMaterialResponse)checkScore(value string ,score *float64, done bool) <- chan float64{
// SomeTimes - False - True
	response := make(chan float64)
	finalScore := *score 

	go func () {
	if value == "True"{
	finalScore += 10
	}else if value == "False"{
	finalScore += 0
	}else if value == "SomeTimes"{
	finalScore += 5
	}else{

	}
	log.Println(done, "here 2")
	response <- finalScore
	if done {
		log.Println("here 3")
		close(response)
	}
	}()


	log.Println("here 5")
	return response
}

func (h *hMaterialResponse)AddDatainStruct(in <- chan float64){
	log.Println("here 6")
	go func (){
		for t := range in {
			log.Println(h.HMaterialScore, t,"<- value in chan")
			h.HMaterialScore = h.HMaterialScore + float64(t)
			log.Println(h.HMaterialScore, "<- value in chan")
		}
	}()
return
}


func (h hMaterialResponse)calculatePercentage(score, totalScore int)float64{
var percentage float64

h.HMaterialScore = float64(score) / float64(totalScore) * 100

return percentage
}

func main() {
	log.Println("Server running at : 8080")
	initializeRouter()
}
