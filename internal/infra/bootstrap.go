package infra

import (
	"github.com/axel-andrade/opina-ai-api/internal/adapters/primary/handlers"
	"github.com/axel-andrade/opina-ai-api/internal/adapters/primary/http/controllers"
	"github.com/axel-andrade/opina-ai-api/internal/adapters/primary/http/presenters"
	common_ptr "github.com/axel-andrade/opina-ai-api/internal/adapters/primary/http/presenters/common"
	cockroach_repositories "github.com/axel-andrade/opina-ai-api/internal/adapters/secondary/database/cockroach/repositories"
	"github.com/axel-andrade/opina-ai-api/internal/core/usecases/voter/create_voter"
)

type Dependencies struct {
	BaseCockroachRepository  *cockroach_repositories.BaseCockroachRepository
	VoterCockroachRepositoty *cockroach_repositories.VoterCockroachRepository

	EncrypterHandler    *handlers.EncrypterHandler
	TokenManagerHandler *handlers.TokenManagerHandler

	CreateVoterController *controllers.CreateVoterController

	CreateVoterUC *create_voter.CreateVoterUC

	PaginationPresenter  *common_ptr.PaginationPresenter
	CreateVoterPresenter *presenters.CreateVoterPresenter
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
	d.VoterCockroachRepositoty = cockroach_repositories.BuildVoterRepository()
}

func loadHandlers(d *Dependencies) {
	d.EncrypterHandler = handlers.BuildEncrypterHandler()
	d.TokenManagerHandler = handlers.BuildTokenManagerHandler()
}

func loadPresenters(d *Dependencies) {
	d.PaginationPresenter = common_ptr.BuildPaginationPresenter()
	d.CreateVoterPresenter = presenters.BuildCreateVoterPresenter()
}

func loadUseCases(d *Dependencies) {
	d.CreateVoterUC = create_voter.BuildCreateVoterUC(struct {
		*cockroach_repositories.VoterCockroachRepository
	}{d.VoterCockroachRepositoty})
}

func loadControllers(d *Dependencies) {
	d.CreateVoterController = controllers.BuildCreateVoterController(d.CreateVoterUC, d.CreateVoterPresenter)
}
