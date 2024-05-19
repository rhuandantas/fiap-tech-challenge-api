package usecase

import (
	"context"
	"fiap-tech-challenge-api/internal/core/domain"
	mock_repo "fiap-tech-challenge-api/test/mock/repository"

	. "github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
	"go.uber.org/mock/gomock"
)

var _ = Describe("lista produto por categoria use case testes", func() {
	var (
		ctx                      = context.Background()
		repo                     *mock_repo.MockProdutoRepo
		pegarProdutoPorCategoria PegarProdutoPorCategoria
	)

	BeforeEach(func() {
		mockCtrl := gomock.NewController(GinkgoT())
		repo = mock_repo.NewMockProdutoRepo(mockCtrl)
		pegarProdutoPorCategoria = NewPegaProdutoPorCategoria(repo)
	})

	Context("lista produto", func() {
		produtoDTO := &domain.Produto{
			Categoria: "bebida",
			Descricao: "coca",
		}
		It("lista com sucesso", func() {
			repo.EXPECT().PesquisaPorCategoria(ctx, produtoDTO).Return([]*domain.Produto{produtoDTO}, nil)
			prod, err := pegarProdutoPorCategoria.Pega(ctx, produtoDTO)

			gomega.Expect(err).To(gomega.BeNil())
			gomega.Expect(prod).ToNot(gomega.BeNil())
		})
	})
})
