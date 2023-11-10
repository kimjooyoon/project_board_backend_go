package usecase

import (
	"github.com/kimjooyoon/project_board_backend_go/internal/app/article"
	"github.com/kimjooyoon/project_board_backend_go/internal/article/domain"
	"github.com/kimjooyoon/project_board_backend_go/internal/mocks/app/article/repo"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/mock"
)

// BDD Use Case Adapter
var _ = Describe("articleQueryUseCase", func() {
	// UseCase Adapter 를 구성하는 환경 변수를 작성합니다
	var (
		usecase article.QueryService
		query   *mocksrepo.MockArticleQuery
	)

	// mockery(make mock 명령어)를 통해 작성된 mock 파일을 제공합니다
	// usecase 를 초기화 합니다
	BeforeEach(func() {
		query = mocksrepo.NewMockArticleQuery(GinkgoT())
		usecase = NewArticleQueryUseCase(query)
	})
	// 정의할 행위를 결정합니다
	Describe("SearchUser 메소드는", func() {
		// 해당 행위를 통해 반환될 값을 정의합니다
		var (
			res article.SearchUserResponse
			err error
		)
		// 어떤 상황이 주어졌는지 결정합니다
		When("고객의 이름이 주어지면", func() {
			// 결정된 행위에서 실행할 기능을 작성합니다
			BeforeEach(func() {
				// query 환경 변수를 수정(mock)하여 원하는 범위는 [단위]에서 분리합니다
				query.EXPECT().FindByName(mock.Anything).
					Return([]domain.Article{
						domain.NewArticleBuilder().
							ID(101).
							Title("홍길동전 1화").
							UserId(2000).
							Content("아버지를 아버지라 부르지 못하고\n형을 형이라 부르지 못했다 ...").
							Build(),
					}, 10,
					).Maybe()
				// 실행을 통해 반환 값을 재할당합니다
				res, err = usecase.SearchUser("홍길동")
			})

			// 행위를 검증합니다
			It("고객의 게시판을 반환한다", func() {
				// Expect: 검증 대상 / .To: 예상 값, + 문자열 (상세 설명)
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
