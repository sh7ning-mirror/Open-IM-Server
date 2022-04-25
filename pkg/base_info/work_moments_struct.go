package base_info

import "Open_IM/pkg/proto/office"

type CreateOneWorkMomentReq struct {
	office.CreateOneWorkMomentReq
}

type CreateOneWorkMomentResp struct {
	CommResp
}

type DeleteOneWorkMomentReq struct {
	office.DeleteOneWorkMomentReq
}

type DeleteOneWorkMomentResp struct {
	CommResp
}

type LikeOneWorkMomentReq struct {
	office.LikeOneWorkMomentReq
}

type LikeOneWorkMomentResp struct {
	CommResp
}

type CommentOneWorkMomentReq struct {
	office.CommentOneWorkMomentReq
}

type CommentOneWorkMomentResp struct {
	CommResp
}

type WorkMomentsUserCommonReq struct {
	PageNumber  int32  `json:"pageNumber" binding:"required"`
	ShowNumber  int32  `json:"showNumber" binding:"required"`
	OperationID string `json:"operationID" binding:"required"`
}

type GetWorkMomentByIDReq struct {
	office.GetWorkMomentByIDReq
}

type WorkMoment struct {
	WorkMomentID string            `json:"workMomentID"`
	UserID       string            `json:"userID"`
	Content      string            `json:"content"`
	LikeUserList []*WorkMomentUser `json:"likeUsers"`
	Comments     []*Comment        `json:"comments"`
	FaceURL      string            `json:"faceUrl"`
	UserName     string            `json:"userName"`
	//Permission            int32       `json:"permission"`
	//PermissionUserIDList  []string    `json:"permissionUserIDList"`
	//PermissionGroupIDList []string    `json:"permissionGroupIDList"`
	AtUserList []*WorkMomentUser `json:"atUsers"`
	CreateTime int32             `json:"createTime"`
}

type WorkMomentUser struct {
	UserID   string `json:"userID"`
	UserName string `json:"userName"`
}

type Comment struct {
	UserID        string `json:"userID"`
	UserName      string `json:"userName"`
	ReplyUserID   string `json:"replyUserID"`
	ReplyUserName string `json:"replyUserName"`
	ContentID     string `json:"contentID"`
	Content       string `json:"content"`
	CreateTime    int32  `json:"createTime"`
}

type GetWorkMomentByIDResp struct {
	CommResp
	Data struct {
		WorkMoment *WorkMoment `json:"workMoment"`
	} `json:"data"`
}

type GetUserWorkMomentsReq struct {
	WorkMomentsUserCommonReq
}

type GetUserWorkMomentsResp struct {
	CommResp
	Data struct {
		WorkMoments []*WorkMoment `json:"workMoments"`
		CurrentPage int32         `json:"currentPage"`
		ShowNumber  int32         `json:"showNumber"`
	} `json:"data"`
}

type GetUserFriendWorkMomentsReq struct {
	WorkMomentsUserCommonReq
}

type GetUserFriendWorkMomentsResp struct {
	CommResp
	Data struct {
		WorkMoments []*WorkMoment `json:"workMoments"`
		CurrentPage int32         `json:"currentPage"`
		ShowNumber  int32         `json:"showNumber"`
	} `json:"data"`
}

type SetUserWorkMomentsLevelReq struct {
	office.SetUserWorkMomentsLevelReq
}

type SetUserWorkMomentsLevelResp struct {
	CommResp
}
