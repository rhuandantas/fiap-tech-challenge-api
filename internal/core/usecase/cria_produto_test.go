package usecase

import (
	"context"
	"fiap-tech-challenge-api/internal/core/domain"
	mock_repo "fiap-tech-challenge-api/test/mock/repository"

	. "github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
	"go.uber.org/mock/gomock"
)

var _ = Describe("cadastra produto use case testes", func() {
	var (
		ctx               = context.Background()
		repo              *mock_repo.MockProdutoRepo
		cadastraProdutoUC CadastrarProduto
	)

	BeforeEach(func() {
		mockCtrl := gomock.NewController(GinkgoT())
		repo = mock_repo.NewMockProdutoRepo(mockCtrl)
		cadastraProdutoUC = NewCadastraProduto(repo)
	})

	Context("cadastra produto", func() {
		produtoDTO := &domain.Produto{
			Categoria: "bebida",
			Descricao: "coca",
		}
		It("cadastra com sucesso", func() {
			repo.EXPECT().Insere(ctx, produtoDTO).Return(produtoDTO, nil)
			prod, err := cadastraProdutoUC.Cadastra(ctx, produtoDTO)

			gomega.Expect(err).To(gomega.BeNil())
			gomega.Expect(prod).ToNot(gomega.BeNil())
		})
	})
})
