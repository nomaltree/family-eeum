package handler

import (
	"github.com/gorilla/mux"
	"github.com/nomaltree/family-eeum/mission/service"
)

func MissionHandler(ms service.Service) *mux.Router {
	router := mux.NewRouter()

	//미션 모델 데이터 담아가기 핸들러
	router.HandleFunc("/mission/get/model", getModelHandler(ms)).Methods("GET") //정상작동

	//미션 모델 추천 핸들러(추천 알고리즘)
	router.HandleFunc("/mission/get/models", getModelsHandler(ms)).Methods("POST") //정상작동

	//특정 유저의 미션 데이터 조회 핸들러
	router.HandleFunc("/mission/get/userMissions", getUserMissionHandler(ms)).Methods("GET") //정상작동

	//특정 유저의 미션 데이터 조회 핸들러(내가 다른 사용자의 것을 조회하는 경우)
	//이 경우 공개 설정된 미션만 가져옴
	router.HandleFunc("/mission/get/otherMissions", getOtherUserMissionHandler(ms)).Methods("GET") //정상작동

	//나를 제외한 가족 구성원이 참여하는 미션 조회 핸들러
	router.HandleFunc("/mission/get/familyMission", getFamilyMissionHandler(ms)).Methods("GET") //정상작동

	//id로 특정 미션 데이터 조회 핸들러
	router.HandleFunc("/mission/get/mission", getOneMissionHandler(ms)).Methods("GET") //정상작동

	//특정 가족들의 미션 데이터 조회 핸들러
	router.HandleFunc("/mission/get/familyMissions", getFamilyMissionsHandler(ms)).Methods("GET") //정상작동

	//미션 등록 핸들러
	router.HandleFunc("/mission/insert/mission", insertMissionHandler(ms)).Methods("POST") //정상작동

	//미션 삭제 핸들러
	router.HandleFunc("/mission/delete/mission", deleteMissionHandler(ms)).Methods("PUT") //정상작동

	//미션 수정 핸들러
	router.HandleFunc("/mission/update/mission", updateMissionHandler(ms)).Methods("PUT") //정상작동

	//미션 검색 핸들러(공개 설정된 것만 가져옴)
	router.HandleFunc("/mission/search/mission", searchMissionHandler(ms)).Methods("GET") //정상작동

	//미션을 완료상태로 바꾸는 핸들러(미션 완료 인증)
	router.HandleFunc("/mission/update/complete", updateCompleteHandler(ms)).Methods("PUT") //정상작동

	//미션을 공개로 설정하는 핸들러
	router.HandleFunc("/mission/isPub/mission", isPubHandler(ms)).Methods("PUT") //정상작동

	//미션을 비공개로 설정하는 핸들러
	router.HandleFunc("/mission/isNotPub/mission", isNotPubHandler(ms)).Methods("PUT") //정상작동

	//특정 미션에 참여하는 가족 구성원 정보 조회 핸들러
	router.HandleFunc("/mission/get/fammission", getFamMissionHandler(ms)).Methods("GET") //정상작동

	//특정 미션 등록시 참여하는 가족 구성원 등록 핸들러(fam_mission 테이블에 등록)
	router.HandleFunc("/mission/insert/fammission", insertFamMissionHandler(ms)).Methods("POST") //정상작동

	//특정 미션 삭제시 참여하는 가족 구성원 삭제 핸들러(fam_mission 테이블에서 삭제)
	router.HandleFunc("/mission/delete/fammission", deleteFamMissionHandler(ms)).Methods("PUT") //정상작동

	//특정 미션의 댓글들을 조회하는 핸들러
	router.HandleFunc("/mission/get/missionComments", getMissionCommentsHandler(ms)).Methods("GET") //정상작동

	//특정 사용자가 작성한 댓글들을 조회하는 핸들러
	router.HandleFunc("/mission/get/userComments", getUserCommentsHandler(ms)).Methods("GET") //정상작동

	//댓글 등록 핸들러
	router.HandleFunc("/mission/insert/missioncomment", insertMissionCommentHandler(ms)).Methods("POST") //정상작동

	//댓글 삭제 핸들러
	router.HandleFunc("/mission/delete/missioncomment", deleteMissionCommentHandler(ms)).Methods("PUT") //정상작동

	return router
}
