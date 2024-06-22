package usecase

import (
	"context"
	"database/sql"
	"errors"
	"fiap-tech-challenge-api/internal/core/domain"
	mock_repo "fiap-tech-challenge-api/test/mock/repository"

	. "github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
	"go.uber.org/mock/gomock"
)

var _ = Describe("cadastra cliente use case testes", func() {
	var (
		ctx             = context.Background()
		repo            *mock_repo.MockClienteRepo
		cadastraCliente CadastrarClienteUseCase
	)

	BeforeEach(func() {
		mockCtrl := gomock.NewController(GinkgoT())
		repo = mock_repo.NewMockClienteRepo(mockCtrl)
		cadastraCliente = NewCadastraCliente(repo)
	})

	Context("cadastra cliente", func() {
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
		It("cadastra cliente com sucesso", func() {
			repo.EXPECT().PesquisaPorCPF(gomock.Any(), gomock.Any()).Return(nil, nil)
			repo.EXPECT().Insere(ctx, gomock.Any()).Return(dto, nil)
			cli, err := cadastraCliente.Cadastra(ctx, clienteDTO)

			gomega.Expect(err).To(gomega.BeNil())
			gomega.Expect(cli).ToNot(gomega.BeNil())
		})
		It("falha ao cadastrar cliente", func() {
			repo.EXPECT().PesquisaPorCPF(gomock.Any(), gomock.Any()).Return(nil, nil)
			repo.EXPECT().Insere(ctx, gomock.Any()).Return(nil, errors.New("mock error"))
			cli, err := cadastraCliente.Cadastra(ctx, clienteDTO)

			gomega.Expect(err.Error()).To(gomega.Equal("mock error"))
			gomega.Expect(cli).To(gomega.BeNil())
		})
	})
})
