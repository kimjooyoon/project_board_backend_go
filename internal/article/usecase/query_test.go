package usecase

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"project_board_backend/internal/app/article"
	"project_board_backend/internal/article/domain"
	"project_board_backend/internal/mocks/infrastructure/out/maria/article/query"
)

var _ = Describe("articleQueryUseCase", func() {
	var (
		usecase article.QueryService
		query   *mocks_query.MockArticleQuery
	)

	BeforeEach(func() {
		query = mocks_query.NewMockArticleQuery(GinkgoT())
		usecase = NewArticleQueryUseCase(query)
	})
	Describe("SearchUser 메소드는", func() {
		var (
			res article.SearchUserResponse
			err error
		)
		When("고객의 이름이 주어지면", func() {
			BeforeEach(func() {
				query.EXPECT().FindByName("홍길동").
					Return([]domain.Article{
						domain.NewArticleBuilder().
							ID(101).
							Title("홍길동전 1화").
							UserId(2000).
							Content("아버지를 아버지라 부르지 못하고\n형을 형이라 부르지 못했다 ...").
							Build(),
					}).Once()
				res, err = usecase.SearchUser("홍길동")
			})

			It("고객의 게시판을 반환한다", func() {
				Expect(len(res.List)).To(Equal(1), "게시글은 1개다")
				Expect(res.List[0].ID).To(Equal("101"), "조회된 첫번째 게시글은 식별자를 가진다")
				Expect(res.List[0].UserId).To(Equal("2000"), "조회된 첫번째 게시글은 작성자 식별자를 가진다")
				Expect(res.List[0].Title).To(Equal("홍길동전 1화"), "조회된 첫번째 게시글은 제목을 가진다")
				Expect(res.List[0].Content).To(Equal("아버지를 아버지라 부르지 못하고\n형을 형이라 부르지 못했다 ..."), "조회된 첫번째 게시글은 내용을 가진다")
				Expect(err).To(BeNil(), "에러는 없다")
			})
		})
	})
})
