package article

import (
	"context"
	"net/http/httptest"
	utils "twof/blog-api/initiator"
	zaplog "twof/blog-api/internal/log"

	"twof/blog-api/internal/constant/model"

	"github.com/cucumber/godog"
	"github.com/gin-gonic/gin"
)

type Response struct {
	MetaData interface{} `json:"meta_data,omitempty"`
	Data     interface{} `json:"data"`
}

type ErrData struct {
	Error interface{} `json:"error"`
}

func (a *article) resetResponse() {
	a.resp = httptest.NewRecorder()
}

type article struct {
	article           model.Article
	server            *gin.Engine
	resp              *httptest.ResponseRecorder
	tempResp          *httptest.ResponseRecorder
	success           Response
	err               ErrData
	adminaccesstocken string
	superaccesstocken string
}

func (a *article) theUserGetsTheArticle(arg1 *godog.Table) error {
	return godog.ErrPending
}

func (a *article) theUserSearchForTheArticle() error {
	return godog.ErrPending
}

func (a *article) thereIsArticleWithId(arg1 string) error {
	return godog.ErrPending
}

func InitializeScenario(ctx *godog.TestSuiteContext) {

	var a = &article{}

	ctx.BeforeSuite(func() {
		a.article = model.Article{}

		zaplog.InitLogger()
		defer zaplog.SugerLogger.Sync()
		utils.Init()
	})

	scCtx := ctx.ScenarioContext()
	scCtx.Before(func(ctx context.Context, sc *godog.Scenario) (context.Context, error) {

		a.resetResponse()
		a.article = model.Article{}
		a.err = ErrData{}
		a.success = Response{}
		return ctx, nil
	})
	scCtx.Step(`^the user gets the article$`, a.theUserGetsTheArticle)
	scCtx.Step(`^the user search for the article$`, a.theUserSearchForTheArticle)
	scCtx.Step(`^there is article with id "([^"]*)"$`, a.thereIsArticleWithId)
}
