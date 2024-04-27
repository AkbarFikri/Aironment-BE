package model

import "mime/multipart"

type CommunityRequest struct {
	Name           string                `form:"name"`
	Description    string                `form:"description"`
	ProfilePicture *multipart.FileHeader `form:"profile_picture" binding:"required"`
	CoverPicture   *multipart.FileHeader `form:"cover_picture" binding:"required"`
	Price          uint64                `form:"price"`
}

type CommunityDataResponse struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	Description    string `json:"description"`
	ProfilePicture string `json:"profile_picture"`
	CoverPicture   string `json:"cover_picture"`
}

type CommunityDetails struct {
	Detail Details               `json:"detail"`
	Member []UserProfileResponse `json:"members"`
}

type Details struct {
	Description string   `json:"description"`
	Galery      []string `json:"galery"`
}

type Member struct {
	ID          string `json:"id"`
	UserID      string `json:"user_id"`
	CommunityID string `json:"community_id"`
}
