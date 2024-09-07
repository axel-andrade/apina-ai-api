package infra

import (
	"github.com/axel-andrade/opina-ai-api/internal/adapters/primary/handlers"
	"github.com/axel-andrade/opina-ai-api/internal/adapters/primary/http/controllers"
	"github.com/axel-andrade/opina-ai-api/internal/adapters/primary/http/presenters"
	common_ptr "github.com/axel-andrade/opina-ai-api/internal/adapters/primary/http/presenters/common"
	cockroach_repositories "github.com/axel-andrade/opina-ai-api/internal/adapters/secondary/database/cockroach/repositories"
	"github.com/axel-andrade/opina-ai-api/internal/core/usecases/import_voters"
	"github.com/axel-andrade/opina-ai-api/internal/core/usecases/voter/create_voter"
)

type Dependencies struct {
	BaseCockroachRepository   *cockroach_repositories.BaseCockroachRepository
	VoterCockroachRepositoty  *cockroach_repositories.VoterCockroachRepository
	ImportCockroachRepository *cockroach_repositories.ImportCockroachRepository

	EncrypterHandler *handlers.EncrypterHandler

	CreateVoterController *controllers.CreateVoterController

	CreateVoterUC  *create_voter.CreateVoterUC
	ImportVotersUC *import_voters.ImportVotersUC

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
	d.VoterCockroachRepositoty = cockroach_repositories.BuildCockroachVoterRepository()
	d.ImportCockroachRepository = cockroach_repositories.BuildCockroachImportRepository()
}

func loadHandlers(d *Dependencies) {
	d.EncrypterHandler = handlers.BuildEncrypterHandler()
}

func loadPresenters(d *Dependencies) {
	d.PaginationPresenter = common_ptr.BuildPaginationPresenter()
	d.CreateVoterPresenter = presenters.BuildCreateVoterPresenter()
}

func loadUseCases(d *Dependencies) {
	d.CreateVoterUC = create_voter.BuildCreateVoterUC(struct {
		*cockroach_repositories.VoterCockroachRepository
	}{d.VoterCockroachRepositoty})

	//	d.ImportVotersUC = import_voters.BuildImportVotersUC(struct {
	//		*cockroach_repositories.ImportCockroachRepository
	//	}{d.ImportCockroachRepository})
}

func loadControllers(d *Dependencies) {
	d.CreateVoterController = controllers.BuildCreateVoterController(d.CreateVoterUC, d.CreateVoterPresenter)
}
