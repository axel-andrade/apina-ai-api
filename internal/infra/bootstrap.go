package infra

import (
	"github.com/axel-andrade/opina-ai-api/internal/adapters/primary/handlers"
	"github.com/axel-andrade/opina-ai-api/internal/adapters/primary/http/controllers"
	"github.com/axel-andrade/opina-ai-api/internal/adapters/primary/http/presenters"
	common_ptr "github.com/axel-andrade/opina-ai-api/internal/adapters/primary/http/presenters/common"
	mongo_repositories "github.com/axel-andrade/opina-ai-api/internal/adapters/secondary/database/mongo/repositories"
	create_content "github.com/axel-andrade/opina-ai-api/internal/core/usecases/content/create"
	"github.com/axel-andrade/opina-ai-api/internal/core/usecases/content/process_audio"
)

type Dependencies struct {
	BaseMongoRepository    *mongo_repositories.BaseMongoRepository
	ContentMongoRepository *mongo_repositories.ContentMongoRepository

	EncrypterHandler    *handlers.EncrypterHandler
	TokenManagerHandler *handlers.TokenManagerHandler
	FileHandler         *handlers.FileHandler
	DeepSpeechHandler   *handlers.DeepSpeechHandler
	LanguageToolHandler *handlers.LanguageToolHandler

	CreateContentController *controllers.CreateContentController
	ProcessAudioController  *controllers.ProcessAudioController

	CreateContentUC *create_content.CreateContentUC
	ProcessAudioUC  *process_audio.ProcessAudioUC

	PaginationPresenter    *common_ptr.PaginationPresenter
	CreateContentPresenter *presenters.CreateContentPresenter
	ProcessAudioPresenter  *presenters.ProcessAudioPresenter
}

func LoadDependencies() *Dependencies {
	d := &Dependencies{}

	loadRepositories(d)
	loadHandlers(d)
	loadPresenters(d)
	loadUseCases(d)
	loadControllers(d)

	return d
}

func loadRepositories(d *Dependencies) {
	d.ContentMongoRepository = mongo_repositories.BuildContentMongoRepository()
}

func loadHandlers(d *Dependencies) {
	d.EncrypterHandler = handlers.BuildEncrypterHandler()
	d.TokenManagerHandler = handlers.BuildTokenManagerHandler()
	d.FileHandler = handlers.BuildFileHandler()
	d.DeepSpeechHandler = handlers.BuildDeepSpeechHandler()
	d.LanguageToolHandler = handlers.BuildLanguageToolHandler()
}

func loadPresenters(d *Dependencies) {
	d.PaginationPresenter = common_ptr.BuildPaginationPresenter()
	d.CreateContentPresenter = presenters.BuildCreateContentPresenter()
	d.ProcessAudioPresenter = presenters.BuildProcessAudioPresenter()
}

func loadUseCases(d *Dependencies) {
	d.CreateContentUC = create_content.BuildCreateContentUC(struct {
		*mongo_repositories.ContentMongoRepository
		*handlers.FileHandler
	}{d.ContentMongoRepository, d.FileHandler})

	d.ProcessAudioUC = process_audio.BuildProcessAudioUC(struct {
		*mongo_repositories.ContentMongoRepository
		*handlers.DeepSpeechHandler
		*handlers.LanguageToolHandler
	}{d.ContentMongoRepository, d.DeepSpeechHandler, d.LanguageToolHandler})
}

func loadControllers(d *Dependencies) {
	d.CreateContentController = controllers.BuildCreateContentController(d.CreateContentUC, d.CreateContentPresenter)
	d.ProcessAudioController = controllers.BuildProcessAudioController(d.ProcessAudioUC, d.ProcessAudioPresenter)
}
