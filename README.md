# 게시판 서비스 백엔드 프로젝트 
#### 카카오 로그인, 게시판 검색, 페이징 기능, 대댓글 기능 등등이 있습니다.

## 레이아웃 구조

<img src="https://docs.google.com/drawings/d/e/2PACX-1vQ5ps72uaZcEJzwnJbPhzUfEeBbN6CJ04j7hl2i3K2HHatNcsoyG2tgX2vnrN5xxDKLp5Jm5bzzmZdv/pub?w=960&h=657" alt="Hexagonal architecture layout image">

- 참조: [프로젝트 레이아웃](https://github.com/golang-standards/project-layout/blob/master/README_ko.md)
- 참조: [golang ddd 디자인](https://programmingpercy.tech/blog/how-to-domain-driven-design-ddd-golang/)

### /build
빌드를 위한 파일을 작성합니다

### /cmd
실행을 위한 파일을 작성합니다

### /deploy
배포를 위한 파일을 작성합니다

### /docs
관련 문서를 작성합니다

### /internal

#### [app](internal%2Fapp)
어플리케이션 인터페이스를 정의합니다
clean architecture 의 use-case, out port 인터페이스를 제공합니다

#### [article](internal%2Farticle)
게시판 도메인을 정의합니다

#### [comment](internal%2Fcomment)
댓글 도메인을 정의합니다

#### [hashtag](internal%2Fhashtag)
해시태그 도메인을 정의합니다

#### [infrastructure](internal%2Finfrastructure)
system data layer 를 정의합니다

#### [service](internal%2Fservice)
어플리케이션 기능을 정의합니다

#### [user](internal%2Fuser)
사용자 도메인을 정의합니다

## 구조

* Use-case 다이어그램
<img width="768" alt="스크린샷 2023-06-24 오전 12 40 32" src="https://github.com/myoungsuk/Emsys-hompage/assets/81986479/18f28c98-a8fa-4594-aa66-e65702db5e03">



* ER 다이어그램
<img width="959" alt="스크린샷 2023-06-24 오전 12 40 11" src="https://github.com/myoungsuk/Emsys-hompage/assets/81986479/400cb3fe-114d-46e8-b46e-3781809bca05">


## 개발 환경

* Goland
* golang
* echo

## 기술 세부 스택

echo

* gorm
