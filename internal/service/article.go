
package service

import(
	"blog/internal/biz"
	"context"
	"github.com/go-kratos/kratos/v2/log"

	pb "blog/api/blog"
)

type ArticleService struct {
	pb.UnimplementedArticleServer

	uc *biz.ArticleUsecase
	log log.Logger
}

func NewArticleService(uc *biz.ArticleUsecase, logger log.Logger) *ArticleService {
	return &ArticleService{uc: uc, log: log.NewHelper("service/greeter", logger)}
}

func (s *ArticleService) CreateArticle(ctx context.Context, req *pb.CreateArticleRequest) (*pb.CreateArticleReply, error) {
	err := s.uc.Create(ctx, &biz.Article{
		Title: req.GetTitle(),
		Content: req.GetContent(),
		Type: req.GetType(),
		CreatedBy: req.GetCreatedBy(),
		UpdatedBy: req.GetUpdatedBy(),
		PublishedAt: req.GetPublishedAt(),
	})
	return &pb.CreateArticleReply{Message: "OK" + req.GetTitle()}, err
}
func (s *ArticleService) UpdateArticle(ctx context.Context, req *pb.UpdateArticleRequest) (*pb.UpdateArticleReply, error) {
	return &pb.UpdateArticleReply{}, nil
}
func (s *ArticleService) DeleteArticle(ctx context.Context, req *pb.DeleteArticleRequest) (*pb.DeleteArticleReply, error) {
	return &pb.DeleteArticleReply{}, nil
}
func (s *ArticleService) GetArticle(ctx context.Context, req *pb.GetArticleRequest) (*pb.GetArticleReply, error) {
	return &pb.GetArticleReply{}, nil
}
func (s *ArticleService) ListArticle(ctx context.Context, req *pb.ListArticleRequest) (*pb.ListArticleReply, error) {
	return &pb.ListArticleReply{}, nil
}
