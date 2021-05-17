package data

import (
	"blog/internal/biz"
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type articleRepo struct {
	data *Data
	logger *log.Helper
}

func NewArticleRepo(data *Data, logger *log.Helper) biz.ArticleRepo {
	return &articleRepo{
		data: data,
		logger: log.NewHelper("data/article", logger),
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
		Save(ctx)

	return err;
}

func (a *articleRepo) UpdateArticle(ctx context.Context, article *biz.Article) error {
	panic("implement me")
}
