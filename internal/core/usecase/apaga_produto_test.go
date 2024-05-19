package usecase

import (
	"context"
	"fiap-tech-challenge-api/internal/core/domain"
	mock_repo "fiap-tech-challenge-api/test/mock/repository"

	. "github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
	"go.uber.org/mock/gomock"
)

var _ = Describe("apaga produto use case testes", func() {
	var (
		ctx            = context.Background()
		repo           *mock_repo.MockProdutoRepo
		apagaProdutoUC ApagarProduto
	)

	BeforeEach(func() {
		mockCtrl := gomock.NewController(GinkgoT())
		repo = mock_repo.NewMockProdutoRepo(mockCtrl)
		apagaProdutoUC = NewApagaProduto(repo)
	})

	Context("apaga produto", func() {
		produtoDTO := &domain.Produto{
			Id:        1,
			Categoria: "1",
			Descricao: "preparando",
		}
		It("apaga com sucesso", func() {
			repo.EXPECT().Apaga(ctx, produtoDTO).Return(nil)
			err := apagaProdutoUC.Apaga(ctx, produtoDTO)

			gomega.Expect(err).To(gomega.BeNil())
		})
	})
})
