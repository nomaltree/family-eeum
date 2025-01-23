package handler

import (
	"github.com/gorilla/mux"
	"github.com/nomaltree/family-eeum/feed/service"
)

func FeedHandler(fs service.Service) *mux.Router {
	router := mux.NewRouter()

	//전체 피드 조회 핸들러(공개로 설정된 것만 가져옴)
	router.HandleFunc("/feed/get/feeds", getAllFeedHandler(fs)).Methods("GET") //정상작동

	//특정 가족의 피드 데이터 조회 핸들러
	router.HandleFunc("/feed/get/familyFeeds", getFamilyFeedHandler(fs)).Methods("GET") //정상작동

	//특정 피드 데이터 조회 핸들러
	router.HandleFunc("/feed/get/feed", getFeedHandler(fs)).Methods("GET") //정상작동

	//특정 유저의 피드 데이터 조회 핸들러
	router.HandleFunc("/feed/get/userFeeds", getFeedsHandler(fs)).Methods("GET") //정상작동

	//피드 등록 핸들러
	router.HandleFunc("/feed/insert/feed", insertFeedHandler(fs)).Methods("POST") //정상작동

	//피드 삭제 핸들러
	router.HandleFunc("/feed/delete/feed", deleteFeedHandler(fs)).Methods("PUT") //정상작동

	//피드 검색 핸들러(공개로 설정된 것만 가져옴)
	router.HandleFunc("/feed/search/feed", searchFeedHandler(fs)).Methods("GET") //정상작동

	//피드 수정 핸들러
	router.HandleFunc("/feed/update/feed", updateFeedHandler(fs)).Methods("PUT") //정상작동

	//피드 태그 등록 핸들러
	router.HandleFunc("/feed/insert/feedTags", insertFeedTagHandler(fs)).Methods("POST") //정상작동

	//특정 피드의 댓글둘을 조회하는 핸들러
	router.HandleFunc("/feed/get/feedComments", getFeedCommentsHandler(fs)).Methods("GET") //정상작동

	//특정 사용자가 작성한 댓글들을 조회하는 핸들러
	router.HandleFunc("/feed/get/userComments", getUserFeedCommentsHandler(fs)).Methods("GET") //정상작동

	//댓글 등록 핸들러
	router.HandleFunc("/feed/insert/feedcomment", insertFeedCommentHandler(fs)).Methods("POST") //정상작동

	//댓글 삭제 핸들러
	router.HandleFunc("/feed/delete/feedcomment", deleteFeedCommentHandler(fs)).Methods("PUT") //정상작동

	//특정 사용자의 즐겨찾기 피드 조회 핸들러(좋아요 기능)
	router.HandleFunc("/feed/get/feedbookmark", getUserFeedBookmarkHandler(fs)).Methods("GET") //정상작동

	//즐겨찾기 등록 핸들러(좋아요 등록)
	router.HandleFunc("/feed/insert/feedbookmark", insertFeedBookmarkHandler(fs)).Methods("POST") //정상작동

	//즐겨찾기 삭제 핸들러(좋아요 삭제)
	router.HandleFunc("/feed/delete/feedbookmark", deleteFeedBookmarkHandler(fs)).Methods("PUT") //정상작동

	return router
}
