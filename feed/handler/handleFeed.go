package handler

import (
	"encoding/json"
	"github.com/nomaltree/family-eeum/feed/service"
	"github.com/nomaltree/family-eeum/feed/types"
	"net/http"
	"strconv"
)

// 전체 피드 조회 메소드(공개로 설정된 것만 가져옴)
func getAllFeedHandler(fs service.Service) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		feeds, err := fs.GetAllFeeds()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		json.NewEncoder(w).Encode(feeds)
	}
}

// 특정 가족의 피드 데이터 조회 메소드
func getFamilyFeedHandler(fs service.Service) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var query = r.URL.Query()
		var familyId int
		familyId, err := strconv.Atoi(query["familyId"][0])
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		feeds, err := fs.GetFamilyFeeds(familyId)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		json.NewEncoder(w).Encode(feeds)
	}
}

// 특정 피드 데이터 조회 메소드
func getFeedHandler(fs service.Service) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var query = r.URL.Query()
		var feedId int
		feedId, err := strconv.Atoi(query["feedId"][0])
		Feed, err := fs.GetFeed(feedId)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
		json.NewEncoder(w).Encode(Feed)
	}
}

// 특정 유저의 피드 데이터 조회 메소드
func getFeedsHandler(fs service.Service) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var query = r.URL.Query()
		var userId int
		userId, err := strconv.Atoi(query["userId"][0])
		Feeds, err := fs.GetUserFeeds(userId)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
		json.NewEncoder(w).Encode(Feeds)
	}
}

// 피드 등록 메소드
func insertFeedHandler(fs service.Service) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var Feed types.Feed
		err := json.NewDecoder(r.Body).Decode(&Feed)
		if err != nil {
			http.Error(w, "Bad Request", http.StatusBadRequest)
		}
		id, err := fs.InsertFeed(Feed)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
		Feed.FeedId = id
		json.NewEncoder(w).Encode(Feed)
	}
}

// 피드 삭제 메소드
func deleteFeedHandler(fs service.Service) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var feed types.Feed
		err := json.NewDecoder(r.Body).Decode(&feed)
		if err != nil {
			http.Error(w, "Bad Request", http.StatusBadRequest)
		}
		err = fs.DeleteFeedTags(feed.FeedId)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
		err = fs.DeleteFeed(feed.FeedId)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
		json.NewEncoder(w).Encode(feed)
	}
}

// 피드 검색 메소드
func searchFeedHandler(fs service.Service) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var query = r.URL.Query()
		var search string
		search = query["search"][0]
		feeds, err := fs.SearchFeed(search)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
		json.NewEncoder(w).Encode(feeds)
	}
}

// 피드 수정 메소드
func updateFeedHandler(fs service.Service) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var Feed types.Feed
		err := json.NewDecoder(r.Body).Decode(&Feed)
		if err != nil {
			http.Error(w, "Bad Request", http.StatusBadRequest)
		}
		err = fs.UpdateFeed(Feed)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
		json.NewEncoder(w).Encode(Feed)
	}
}

// 피드 태그 등록 메소드
func insertFeedTagHandler(fs service.Service) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var tags types.FeedTag
		err := json.NewDecoder(r.Body).Decode(&tags)
		if err != nil {
			http.Error(w, "Bad Request", http.StatusBadRequest)
		}
		_, err = fs.InsertFeedTag(tags.FeedId, tags.Content)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
		json.NewEncoder(w).Encode(tags)
	}
}

// 특정 피드의 댓글들을 조회하는 메소드
func getFeedCommentsHandler(fs service.Service) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var query = r.URL.Query()
		var feedId int
		feedId, err := strconv.Atoi(query["feedId"][0])
		if err != nil {
			http.Error(w, "Bad Request", http.StatusBadRequest)
		}
		comments, err := fs.GetOneFeedComments(feedId)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
		json.NewEncoder(w).Encode(comments)
	}
}

// 특정 사용자가 작성한 댓글들을 조회하는 메소드
func getUserFeedCommentsHandler(fs service.Service) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var query = r.URL.Query()
		var userId int
		userId, err := strconv.Atoi(query["userId"][0])
		if err != nil {
			http.Error(w, "Bad Request", http.StatusBadRequest)
		}
		comments, err := fs.GetOneUserFeedComments(userId)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
		json.NewEncoder(w).Encode(comments)
	}
}

// 댓글 등록 메소드
func insertFeedCommentHandler(fs service.Service) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var Comment types.FeedComment
		err := json.NewDecoder(r.Body).Decode(&Comment)
		if err != nil {
			http.Error(w, "Bad Request", http.StatusBadRequest)
		}
		id, err := fs.InsertFeedComment(Comment)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
		Comment.Feedcommentid = id
		json.NewEncoder(w).Encode(Comment)
	}
}

// 댓글 삭제 메소드
func deleteFeedCommentHandler(fs service.Service) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var comment types.FeedComment
		err := json.NewDecoder(r.Body).Decode(&comment)
		if err != nil {
			http.Error(w, "Bad Request", http.StatusBadRequest)
		}
		err = fs.DeleteFeedComment(comment)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
		json.NewEncoder(w).Encode(comment)
	}
}

// 특정 사용자의 즐겨찾기 피드 조회 핸들러(좋아요 기능)
func getUserFeedBookmarkHandler(fs service.Service) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var query = r.URL.Query()
		var userId int
		userId, err := strconv.Atoi(query["userId"][0])
		feeds, err := fs.GetUserBookmarkFeed(userId)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
		json.NewEncoder(w).Encode(feeds)
	}
}

// 즐겨찾기 등록 핸들러(좋아요 등록)
func insertFeedBookmarkHandler(fs service.Service) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var bookmark types.FeedBookmark
		err := json.NewDecoder(r.Body).Decode(&bookmark)
		if err != nil {
			http.Error(w, "Bad Request", http.StatusBadRequest)
		}
		id, err := fs.InsertFeedBookmark(bookmark.UserId, bookmark.FeedId)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
		bookmark.FeedBookmarkId = id
		json.NewEncoder(w).Encode(bookmark)
	}
}

// 즐겨찾기 삭제 메소드
func deleteFeedBookmarkHandler(fs service.Service) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var bookmark types.FeedBookmark
		err := json.NewDecoder(r.Body).Decode(&bookmark)
		if err != nil {
			http.Error(w, "Bad Request", http.StatusBadRequest)
		}
		err = fs.DeleteFeedBookmark(bookmark.UserId, bookmark.FeedId)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
		json.NewEncoder(w).Encode(bookmark)
	}
}
