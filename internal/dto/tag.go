package dto

import (
	"github.com/scSZn/blog/consts"
	"github.com/scSZn/blog/internal/model"
)

// 展示给前端的Tag信息
type TagVO struct {
	TagID        string `json:"tag_id"`
	TagName      string `json:"tag_name"`
	ArticleCount uint32 `json:"article_count"`
	Status       uint8  `json:"status"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
	DeletedAt    string `json:"deleted_at"`
	IsDel        bool   `json:"is_del"`
}

func GenTagVOFromModelTag(tags []*model.Tag) []*TagVO {
	result := make([]*TagVO, 0, len(tags))
	for _, tag := range tags {
		createdAt := tag.CreatedAt.Format(consts.TimeFormatLayout)
		deletedAt := tag.DeletedAt.Time.Format(consts.TimeFormatLayout)
		updatedAt := tag.UpdatedAt.Format(consts.TimeFormatLayout)
		newTag := &TagVO{
			TagID:        tag.TagID,
			TagName:      tag.TagName,
			ArticleCount: tag.ArticleCount,
			Status:       tag.Status,
			CreatedAt:    createdAt,
			UpdatedAt:    updatedAt,
			DeletedAt:    deletedAt,
			IsDel:        tag.IsDel,
		}
		result = append(result, newTag)
	}
	return result
}
