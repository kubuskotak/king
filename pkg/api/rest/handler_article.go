// Package rest is port handler.
package rest

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jinzhu/copier"
	pkgRest "github.com/kubuskotak/asgard/rest"
	pkgTracer "github.com/kubuskotak/asgard/tracer"

	"github.com/kubuskotak/king/pkg/entity"
	"github.com/kubuskotak/king/pkg/persist/crud"
	"github.com/kubuskotak/king/pkg/persist/crud/ent"
	"github.com/kubuskotak/king/pkg/persist/crud/ent/article"
)

// ArticleOption is a struct holding the handler options.
type ArticleOption func(article *Article)

// Article handler instance data.
type Article struct {
	*crud.Database
}

// WithArticleDatabase option function to assign on article.
func WithArticleDatabase(db *crud.Database) ArticleOption {
	return func(a *Article) {
		a.Database = db
	}
}

// NewArticle creates a new article handler instance.
//
//	var articleHandler = rest.NewArticle()
//
//	You can pass optional configuration options by passing a Config struct:
//
//	var adaptor = &adapters.Adapter{}
//	var articleHandler = rest.NewArticle(rest.WithArticleAdapter(adaptor))
func NewArticle(opts ...ArticleOption) *Article {
	// Create a new handler.
	var handler = &Article{}

	// Assign handler options.
	for o := range opts {
		var opt = opts[o]
		opt(handler)
	}

	// Return handler.
	return handler
}

// Register is endpoint group for handler.
func (a *Article) Register(router chi.Router) {
	router.Route("/articles", func(r chi.Router) {
		r.Get("/", pkgRest.HandlerAdapter[ListArticlesRequest](a.ListArticles).JSON)
		r.Post("/", pkgRest.HandlerAdapter[SaveArticleRequest](a.SaveArticle).JSON)
		r.Route("/{id:[0-9-]+}", func(id chi.Router) {
			id.Get("/", pkgRest.HandlerAdapter[GetArticleRequest](a.GetArticle).JSON)
			id.Put("/", pkgRest.HandlerAdapter[SaveArticleRequest](a.SaveArticle).JSON)
			id.Delete("/", pkgRest.HandlerAdapter[DeleteArticleRequest](a.DeleteArticle).JSON)
		})
	})
}

// ListArticles [GET /] articles endpoint func.
func (a *Article) ListArticles(w http.ResponseWriter, r *http.Request) (resp ListArticlesResponse, err error) {
	var (
		ctxSpan, span, l = pkgTracer.StartSpanLogTrace(r.Context(), "ListArticles")
		request          ListArticlesRequest
	)
	defer span.End()
	request, err = pkgRest.GetBind[ListArticlesRequest](r)
	if err != nil {
		l.Error().Err(err).Msg("Bind ListArticles")
		return resp, pkgRest.ErrBadRequest(w, r, err)
	}
	var (
		total    int
		query    = a.Database.Article.Query()
		articles []*ent.Article
		offset   = (request.Page - 1) * request.Limit
		rows     = make([]*entity.Article, len(articles))
	)
	// pagination
	total, err = query.Count(ctxSpan)
	if err != nil {
		return resp, pkgRest.ErrBadRequest(w, r, err)
	}
	pkgRest.Paging(r, pkgRest.Pagination{
		Page:  request.Page,
		Limit: request.Limit,
		Total: total,
	})
	articles, err = query.
		Limit(request.Limit).
		Offset(offset).
		Order(ent.Desc(article.FieldTitle)).
		Where(article.Or(
			article.TitleContains(request.Query),
			article.DescriptionContains(request.Query),
		)).
		All(ctxSpan)
	if err != nil {
		return resp, pkgRest.ErrStatusConflict(w, r, a.Database.ConvertDBError("got an error", err))
	}
	if err = copier.Copy(&rows, &articles); err != nil {
		return resp, pkgRest.ErrBadRequest(w, r, err)
	}

	l.Info().Msg("ListArticles")
	return ListArticlesResponse{
		Articles: rows,
	}, nil
}

// SaveArticle [POST :id] upsert article endpoint func.
func (a *Article) SaveArticle(w http.ResponseWriter, r *http.Request) (resp SaveArticleResponse, err error) {
	var (
		ctxSpan, span, l = pkgTracer.StartSpanLogTrace(r.Context(), "SaveArticle")
		request          SaveArticleRequest
		row              *ent.Article
		artcl            entity.Article
	)
	defer span.End()
	request, err = pkgRest.GetBind[SaveArticleRequest](r)
	if err != nil {
		l.Error().Err(err).Msg("Bind SaveArticle")
		return resp, pkgRest.ErrBadRequest(w, r, err)
	}
	//  upsert
	var client = a.Database.Article
	if request.ID > 0 {
		row, err = client.
			UpdateOneID(request.ID).
			SetTitle(request.Title).
			SetBody(request.Body).
			SetDescription(request.Description).
			SetSlug(request.Slug).
			Save(ctxSpan)
	} else {
		row, err = client.
			Create().
			SetTitle(request.Title).
			SetBody(request.Body).
			SetDescription(request.Description).
			SetSlug(request.Slug).
			Save(ctxSpan)
	}
	if err != nil {
		return resp, pkgRest.ErrStatusConflict(w, r, a.Database.ConvertDBError("got an error", err))
	}
	l.Info().Interface("Article", artcl).Msg("SaveArticle")
	if err = copier.Copy(&artcl, &row); err != nil {
		return resp, pkgRest.ErrBadRequest(w, r, err)
	}
	return SaveArticleResponse{
		Article: artcl,
	}, nil
}

// GetArticle [GET :id] article endpoint func.
func (a *Article) GetArticle(w http.ResponseWriter, r *http.Request) (resp GetArticleResponse, err error) {
	var (
		ctxSpan, span, l = pkgTracer.StartSpanLogTrace(r.Context(), "GetArticle")
		request          GetArticleRequest
		row              *ent.Article
		artcl            entity.Article
	)
	defer span.End()
	request, err = pkgRest.GetBind[GetArticleRequest](r)
	if err != nil {
		l.Error().Err(err).Msg("Bind GetArticle")
		return resp, pkgRest.ErrBadRequest(w, r, err)
	}
	row, err = a.Database.Article.
		Query().
		Where(article.ID(request.Keys.ID)).
		First(ctxSpan)
	if err != nil {
		return resp, pkgRest.ErrStatusConflict(w, r, a.Database.ConvertDBError("got an error", err))
	}
	l.Info().Msg("GetArticleRequest")
	if err = copier.Copy(&artcl, &row); err != nil {
		return resp, pkgRest.ErrBadRequest(w, r, err)
	}
	return GetArticleResponse{
		Article: artcl,
	}, nil
}

// DeleteArticle [DELETE :id] article endpoint func.
func (a *Article) DeleteArticle(w http.ResponseWriter, r *http.Request) (resp DeleteArticleResponse, err error) {
	var (
		ctxSpan, span, l = pkgTracer.StartSpanLogTrace(r.Context(), "DeleteArticle")
		request          DeleteArticleRequest
	)
	defer span.End()
	request, err = pkgRest.GetBind[DeleteArticleRequest](r)
	if err != nil {
		l.Error().Err(err).Msg("Bind DeleteArticle")
		return resp, pkgRest.ErrBadRequest(w, r, err)
	}
	var client = a.Database.Article
	if request.ID < 1 {
		return resp, pkgRest.ErrStatusConflict(w, r, errors.New("record id is"))
	}
	err = client.
		DeleteOneID(request.ID).
		Exec(ctxSpan)
	if err != nil {
		return resp, pkgRest.ErrStatusConflict(w, r, a.Database.ConvertDBError("record", err))
	}
	return DeleteArticleResponse{
		Message: fmt.Sprintf("record deleted successfully: %d", request.ID),
	}, nil
}
