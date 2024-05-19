package usecase

import (
	"context"
	"fiap-tech-challenge-api/internal/core/domain"
	mock_repo "fiap-tech-challenge-api/test/mock/repository"

	. "github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
	"go.uber.org/mock/gomock"
)

var _ = Describe("atualiza produto use case testes", func() {
	var (
		ctx               = context.Background()
		repo              *mock_repo.MockProdutoRepo
		atualizaProdutoUC AtualizarProduto
	)

	BeforeEach(func() {
		mockCtrl := gomock.NewController(GinkgoT())
		repo = mock_repo.NewMockProdutoRepo(mockCtrl)
		atualizaProdutoUC = NewAtualizaProduto(repo)
	})

	Context("atualiza produto", func() {
		produtoDTO := &domain.Produto{
			Id:        1,
			Categoria: "1",
			Descricao: "preparando",
		}
		It("atualiza com sucesso", func() {
			repo.EXPECT().Atualiza(ctx, produtoDTO, int64(1)).Return(nil)
			err := atualizaProdutoUC.Atualiza(ctx, produtoDTO, 1)

			gomega.Expect(err).To(gomega.BeNil())
		})
	})
})
