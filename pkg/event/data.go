package event

// FollowEvent 关注
type FollowEvent struct {
	UserId  uint64 `json:"userId"`
	OtherId uint64 `json:"otherId"`
}

// UnFollowEvent 取消关注
type UnFollowEvent struct {
	UserId  uint64 `json:"userId"`
	OtherId uint64 `json:"otherId"`
}

type PostAddEvent struct {
	FollowingId uint64 `json:"followingId"`
	PostId      uint64 `json:"postId"`
}

type PostDeleteEvent struct {
	FollowingId uint64 `json:"followingId"`
	PostId      uint64 `json:"postId"`
}

type ArticleAddEvent struct {
	FollowingId uint64 `json:"followingId"`
	ArticleId   uint64 `json:"articleId"`
}

type ArticleDeleteEvent struct {
	FollowingId uint64 `json:"followingId"`
	ArticleId   uint64 `json:"articleId"`
}

type UserLikeEvent struct {
	UserId     uint64 `json:"userId"`
	EntityId   uint64 `json:"entityId"`
	EntityType string `json:"entityType"`
}

type UserUnLikeEvent struct {
	UserId     uint64 `json:"userId"`
	EntityId   uint64 `json:"entityId"`
	EntityType string `json:"entityType"`
}

type UserFavoriteEvent struct {
	UserId     uint64 `json:"userId"`
	EntityId   uint64 `json:"entityId"`
	EntityType string `json:"entityType"`
}

type CommentCreateEvent struct {
	UserId    uint64 `json:"userId"`
	CommentId uint64 `json:"commentId"`
}

type TopicRecommendEvent struct {
	TopicId   uint64 `json:"topicId"`
	Recommend bool   `json:"recommend"`
}
