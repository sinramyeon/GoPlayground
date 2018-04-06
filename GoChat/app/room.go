package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/mholt/binding"
	"gopkg.in/mgo.v2/bson"
)

type Room struct {
	ID   bson.ObjectId `bson:"_id" json:"id"`
	Name string        `bson:"name" json:"name"`
}

func (r *Room) FieldMap(req *http.Request) binding.FieldMap {
	return binding.FieldMap{&r.Name: "name"}
}

func createRoom(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	r := new(Room)
	errs := binding.Bind(req, r)
	if errs.Error() != "" {
		return
	}

	session := mongoSession.Copy()
	defer session.Close()

	r.ID = bson.NewObjectId()
	c := session.DB("test").C("rooms")

	// rooms 컬렉션에 room 정보 저장
	if err := c.Insert(r); err != nil {
		// 오류 발생 시 500 에러 반환
		renderer.JSON(w, http.StatusInternalServerError, err)
		return
	}
	// 처리 결과 반환
	renderer.JSON(w, http.StatusCreated, r)

}

func retrieveRooms(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	// 몽고DB 세션 생성
	session := mongoSession.Copy()
	// 몽고DB 세션을 닫는 코드를 defer로 등록
	defer session.Close()

	var rooms []Room
	// 모든 room 정보 조회
	err := session.DB("test").C("rooms").Find(nil).All(&rooms)
	if err != nil {
		// 오류 발생 시 500 에러 반환
		renderer.JSON(w, http.StatusInternalServerError, err)
		return
	}
	// room 조회 결과 반환
	renderer.JSON(w, http.StatusOK, rooms)
}
