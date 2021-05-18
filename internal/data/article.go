package data

import (
	"blog/internal/biz"
	"context"
	"time"
)

type articleRepo struct {
	data *Data
}

func NewArticleRepo(data *Data) biz.ArticleRepo {
	return &articleRepo{
		data: data,
	}
}

func (a *articleRepo) CreateArticle(ctx context.Context, article *biz.Article) error {
	_, err := a.data.db.Article.
		Create().
		SetTitle(article.Title).
		SetContent(article.Content).
		SetCreatedBy(article.CreatedBy).
		SetUpdatedBy(article.UpdatedBy).
		SetCityID(article.CityId).
		SetPublishedAt(time.Now()).
		SetType(article.Type).
		Save(ctx)

	return err
}

func (a *articleRepo) UpdateArticle(ctx context.Context, article *biz.Article) error {
	panic("implement me")
}
