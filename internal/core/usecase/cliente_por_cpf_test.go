package usecase

import (
	"context"
	"database/sql"
	"fiap-tech-challenge-api/internal/core/domain"
	mock_repo "fiap-tech-challenge-api/test/mock/repository"

	. "github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
	"go.uber.org/mock/gomock"
)

var _ = Describe("pesquisa cliente use case testes", func() {
	var (
		ctx            = context.Background()
		repo           *mock_repo.MockClienteRepo
		pesquisaPorCpf PesquisarCliente
	)

	BeforeEach(func() {
		mockCtrl := gomock.NewController(GinkgoT())
		repo = mock_repo.NewMockClienteRepo(mockCtrl)
		pesquisaPorCpf = NewPesquisarCliente(repo)
	})

	Context("pesquisa cliente", func() {
		clienteDTO := &domain.ClienteRequest{
			Nome:  "Mock",
			Cpf:   "20815919018",
			Email: "mock@gmail.com",
		}
		dto := &domain.Cliente{
			Nome:  sql.NullString{String: "Mock"},
			Cpf:   sql.NullString{String: "20815919018"},
			Email: sql.NullString{String: "mock@gmail.com"},
		}
		It("pesquisa por cpf com sucesso", func() {
			repo.EXPECT().PesquisaPorCPF(gomock.Any(), gomock.Any()).Return(dto, nil)
			cli, err := pesquisaPorCpf.PesquisaPorCPF(ctx, clienteDTO)

			gomega.Expect(err).To(gomega.BeNil())
			gomega.Expect(cli).ToNot(gomega.BeNil())
		})
		It("pesquisa por id com sucesso", func() {
			repo.EXPECT().PesquisaPorId(ctx, gomock.Any()).Return(dto, nil)
			cli, err := pesquisaPorCpf.PesquisaPorID(ctx, 1)

			gomega.Expect(err).To(gomega.BeNil())
			gomega.Expect(cli).ToNot(gomega.BeNil())
		})
	})
})
