package MapAPI

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

//Disclaimer

//THIS DaumMapAPI IS FOR KOREANS.
//I'LL ADD GoogleMapAPI soon(or later).
//NO ENGLISH COMMENT PASSING BY

// 네이버API도 곧 추가하겠습니다!!
//
//
// https://developers.kakao.com
/* 우선 가입부터 해서 appkey를 발급받으세요!
 */
// json 파싱용
type MapRequest struct {
	Query string `json:"query"`
	Key   string `json:"appkey"`
}

type Documents struct {
	Meta      map[string]interface{} `json:"meta"`
	Documents []struct {
		PlaceName   string `json:"place_name"`
		PlaceURL    string `json:"place_url"`
		Category    string `json:"category_name"`
		Address     string `json:"address_name"`
		RoadAddress string `json:"road_address_name"`
		Phone       string `json:"phone"`
	} `json:"documents"`
}

// 지도API 호출함수
func CallMapAPI() []byte {

	// 1. 앱키 등록
	jsonreq, err := json.Marshal(MapRequest{
		// 주의!!!!!!!!!!!!!!! 제발 appKey를 Github 등에 올리지 마세요!!!!!!!!!!
		// 제발!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
		Key: "appkey",
	})

	if err != nil {
		log.Fatalln(err)
		recover()
	}

	return jsonreq

}

// 키워드로 주소 검색
func FindMap(where string) {

	var docu Documents

	// API호출
	jsonreq := CallMapAPI()

	//연결준비 (query가 한글이므로 꼭 QueryEscape()필요)
	//다음지도api 키워드 검색 주소입니다.
	mapurl := "https://dapi.kakao.com/v2/local/search/keyword.json?query="
	myurl := mapurl + url.QueryEscape(where)

	/*
		만약 어디 지역 반경 몇km이내 검색이 하고 싶으면 url에
		"https://dapi.kakao.com/v2/local/search/keyword.json?y=37.514322572335935&x=127.06283102249932&radius=20000" \
		이런식으로 값을 주시면 됩니다(https://developers.kakao.com/docs/restapi/local#카테고리-검색)
	*/

	// 1. 한번 url로 값을 전송해 본다.
	req, err := http.NewRequest("GET", myurl, bytes.NewBuffer(jsonreq))
	req.Header.Set("Content-Type", "application/json")
	// * 전에 쓰던 다음 api는 json으로 보내도 됐던 것 같은데... 이제 헤더에
	// 담는 식으로 변경되었습니다. Authorization : KakaoAK  어쩌구 식으로 주시면 됨.
	req.Header.Add("Authorization", "KakaoAK 자신의 키")

	client := &http.Client{}

	// 2. 값을 받아온다
	resp, err := client.Do(req)
	if err != nil {
		log.Panicln(err)
	}
	// 항상 Close함수는 defer로 까먹지 말고 만듭니다(자주 까먹음)
	defer resp.Body.Close()

	// 3. 값을 읽고 리턴해 줍니다.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Panicln(err)
	}

	if err := json.Unmarshal(body, &docu); err != nil {
		log.Fatal(err)
	}

	// 4. 검색된 주소값에 접근할 수 있습니다.
	for _, v := range docu.Documents {

		fmt.Println(v.PlaceName)
		fmt.Println(v.Category)
		fmt.Println(v.Address)
		fmt.Println(v.RoadAddress)
		fmt.Println(v.Phone)

	}

}
