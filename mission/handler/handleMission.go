package handler

import (
	"encoding/json"
	"fmt"
	"github.com/nomaltree/family-eeum/mission/service"
	"github.com/nomaltree/family-eeum/mission/types"
	"log"
	"net/http"
	"strconv"
)

// 미션 모델 데이터 담아가기 메소드
func getModelHandler(ms service.Service) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var query = r.URL.Query()
		modelId, err := strconv.Atoi(query["modelId"][0])
		if err != nil {
			log.Fatal(err)
		}
		models, err := ms.GetModelMission(modelId)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
		json.NewEncoder(w).Encode(models)
	}
}

// 미션 모델 추천 메소드(미션 알고리즘)
func getModelsHandler(ms service.Service) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var DISC []string
		if err := json.NewDecoder(r.Body).Decode(&DISC); err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
		models, err := ms.GetMissionModels(DISC)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
		json.NewEncoder(w).Encode(models)
	}
}

// 특정 유저의 미션 데이터 조회 핸들러
func getUserMissionHandler(ms service.Service) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var userId int
		var missionType string
		var query = r.URL.Query()
		userId, err := strconv.Atoi(query["userId"][0])
		if err != nil {
			log.Fatal(err)
		}
		missionType = query["missionType"][0]
		missions, err := ms.GetAllUserMissions(userId, missionType)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
		json.NewEncoder(w).Encode(missions)
	}
}

// 특정 유저의 미션 데이터 조회 메소드
// 이 경우 공개 설정된 미션만 가져옴
func getOtherUserMissionHandler(ms service.Service) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var userId int
		var query = r.URL.Query()
		userId, err := strconv.Atoi(query["userId"][0])
		if err != nil {
			log.Fatal(err)
		}
		missions, err := ms.GetAllOtherMissions(userId)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
		json.NewEncoder(w).Encode(missions)
	}
}

// 나를 제외한 가족 구성원이 참여하는 미션 조회 메소드
func getFamilyMissionHandler(ms service.Service) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var userId int
		var familyId int
		var query = r.URL.Query()
		userId, err := strconv.Atoi(query["userId"][0])
		if err != nil {
			log.Fatal(err)
		}
		familyId, err = strconv.Atoi(query["familyId"][0])
		if err != nil {
			log.Fatal(err)
		}
		missions, err := ms.GetFamilyMission(familyId, userId)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
		json.NewEncoder(w).Encode(missions)
	}
}

// id로 특정 미션 데이터 조회 메소드
func getOneMissionHandler(ms service.Service) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var missionId int
		var query = r.URL.Query()
		missionId, err := strconv.Atoi(query["missionId"][0])
		if err != nil {
			log.Fatal(err)
		}
		Mission, err := ms.GetOneMission(missionId)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
		json.NewEncoder(w).Encode(Mission)
	}
}

// 특정 가족들의 미션 데이터 조회 메소드
func getFamilyMissionsHandler(ms service.Service) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var familyId int
		var query = r.URL.Query()
		familyId, err := strconv.Atoi(query["familyId"][0])
		if err != nil {
			log.Fatal(err)
		}
		missions, err := ms.GetAllFamilyMission(familyId)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
		json.NewEncoder(w).Encode(missions)
	}
}

// 미션 등록 메소드
func insertMissionHandler(ms service.Service) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var Mission types.Mission
		if err := json.NewDecoder(r.Body).Decode(&Mission); err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
		id, err := ms.InsertMission(Mission)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
		Mission.MissionId = id
		json.NewEncoder(w).Encode(Mission)
	}
}

// 미션 삭제 메소드
func deleteMissionHandler(ms service.Service) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var mission types.Mission
		if err := json.NewDecoder(r.Body).Decode(&mission); err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
		err := ms.DeleteMission(mission.MissionId)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
		json.NewEncoder(w).Encode(mission)
	}
}

// 미션 수정 메소드
func updateMissionHandler(ms service.Service) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var Mission types.Mission
		if err := json.NewDecoder(r.Body).Decode(&Mission); err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
		err := ms.UpdateMission(Mission)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
		json.NewEncoder(w).Encode(Mission)
	}
}

// 미션 검색 메소드
func searchMissionHandler(ms service.Service) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var query = r.URL.Query()
		search := query["search"][0]
		missionType := query["missionType"][0]
		missions, err := ms.SearchMissions(search, missionType)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
		json.NewEncoder(w).Encode(missions)
	}
}

// 미션을 완료상태로 바꾸는 메소드(미션 완료 인증)
func updateCompleteHandler(ms service.Service) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var missionId int
		if err := json.NewDecoder(r.Body).Decode(&missionId); err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
		err := ms.UpdateMissionComplete(missionId)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
		json.NewEncoder(w).Encode(missionId)
	}
}

// 미션을 공개로 설정하는 메소드
func isPubHandler(ms service.Service) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var missionId int
		if err := json.NewDecoder(r.Body).Decode(&missionId); err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
		err := ms.IsPub(missionId)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
		json.NewEncoder(w).Encode(missionId)
	}
}

// 미션을 비공개로 설정하는 메소드
func isNotPubHandler(ms service.Service) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var missionId int
		if err := json.NewDecoder(r.Body).Decode(&missionId); err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
		err := ms.IsNotPub(missionId)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
		json.NewEncoder(w).Encode(missionId)
	}
}

// 특정 미션에 참여하는 가족 구성원 정보 조회 메소드
func getFamMissionHandler(ms service.Service) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var familyId int
		var userId int
		var missionId int
		var query = r.URL.Query()
		familyId, err := strconv.Atoi(query["familyId"][0])
		if err != nil {
			log.Fatal(err)
		}
		userId, err = strconv.Atoi(query["userId"][0])
		if err != nil {
			log.Fatal(err)
		}
		missionId, err = strconv.Atoi(query["missionId"][0])
		if err != nil {
			log.Fatal(err)
		}
		familys, err := ms.GetFamMission(familyId, userId, missionId)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
		json.NewEncoder(w).Encode(familys)
	}
}

// 특정 미션 등록시 참여하는 가족 구성원 등록 메소드
func insertFamMissionHandler(ms service.Service) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var fam types.FamMission
		json.NewDecoder(r.Body).Decode(&fam)
		id, err := ms.InsertFamMission(fam.FamilyId, fam.UserId, fam.MissionId)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
		json.NewEncoder(w).Encode(id)
	}
}

// 특정 미션 삭제시 참여하는 가족 구성원 삭제 메소드
func deleteFamMissionHandler(ms service.Service) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var fam types.FamMission
		json.NewDecoder(r.Body).Decode(&fam)
		err := ms.DeleteOneFamMission(fam.UserId, fam.MissionId)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	}
}

// 특정 미션의 댓글들을 조회하는 메소드
func getMissionCommentsHandler(ms service.Service) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var missionId int
		var query = r.URL.Query()
		missionId, err := strconv.Atoi(query["missionId"][0])
		if err != nil {
			log.Fatal(err)
		}
		comments, err := ms.GetOneMissionComments(missionId)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
		json.NewEncoder(w).Encode(comments)
	}
}

// 특정 사용자가 작성한 댓글들을 조회하는 메소드
func getUserCommentsHandler(ms service.Service) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var userId int
		var query = r.URL.Query()
		fmt.Println(query["userId"])
		userId, err := strconv.Atoi(query["userId"][0])
		if err != nil {
			log.Fatal(err)
		}
		comments, err := ms.GetOneUserMissionComments(userId)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
		json.NewEncoder(w).Encode(comments)
	}
}

// 댓글 등록 메소드
func insertMissionCommentHandler(ms service.Service) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var Comment types.MissionComment
		if err := json.NewDecoder(r.Body).Decode(&Comment); err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
		id, err := ms.InsertMissionComment(Comment)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
		Comment.MissionCommentId = id
		json.NewEncoder(w).Encode(Comment)
	}
}

// 댓글 삭제 메소드
func deleteMissionCommentHandler(ms service.Service) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var comment types.MissionComment
		if err := json.NewDecoder(r.Body).Decode(&comment); err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
		err := ms.DeleteMissionComment(comment)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
		json.NewEncoder(w).Encode(comment)
	}
}
