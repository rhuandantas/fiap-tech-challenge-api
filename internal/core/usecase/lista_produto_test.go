package usecase

import (
	"context"
	"fiap-tech-challenge-api/internal/core/domain"
	mock_repo "fiap-tech-challenge-api/test/mock/repository"

	. "github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
	"go.uber.org/mock/gomock"
)

var _ = Describe("lista produto use case testes", func() {
	var (
		ctx          = context.Background()
		repo         *mock_repo.MockProdutoRepo
		pegarProduto PegaPorIDS
	)

	BeforeEach(func() {
		mockCtrl := gomock.NewController(GinkgoT())
		repo = mock_repo.NewMockProdutoRepo(mockCtrl)
		pegarProduto = NewPegaPorIDS(repo)
	})

	Context("lista produto", func() {
		produtoDTO := &domain.Produto{
			Categoria: "bebida",
			Descricao: "coca",
		}
		It("lista com sucesso", func() {
			repo.EXPECT().PesquisaPorIDS(ctx, gomock.Any()).Return([]*domain.Produto{produtoDTO}, nil)
			prod, err := pegarProduto.PegaPorIDS(ctx, []string{"1"})

			gomega.Expect(err).To(gomega.BeNil())
			gomega.Expect(prod).ToNot(gomega.BeNil())
		})
	})
})
