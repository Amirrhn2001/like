package domain

type LikeDislike struct {
	Id        string `json:"id" bson:"_id,omitempty"`
	RefId     string `json:"ref_uuid" bson:"ref_uuid,omitempty"`
	UserId    string `json:"user_uuid" bson:"user_uuid,omitempty"`
	CreatedAt int64  `json:"created_at" bson:"created_at,omitempty"`
	UpdatedAt int64  `json:"updated_at" bson:"updated_at,omitempty"`
	Status    int    `json:"status" bson:"status,omitempty"` // -1.removed 1.like 2.dislike
}

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Details string `json:"details,omitempty"`
}

type Filter struct {
	Field    string `json:"field"`
	Value    any    `json:"value"`
	Operator string `json:"operator"`
	// logic
	// id
}

type Pagination struct {
	Skip  int64 `json:"skip"`
	Limit int64 `json:"limit"`
}

type Response struct {
	Error *Error         `json:"errors,omitempty"`
	Data  any            `json:"data,omitempty"`
	Meta  map[string]any `json:"meta,omitempty"`
}
