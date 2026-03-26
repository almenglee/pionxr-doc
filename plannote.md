
1. 문서 사양 => DRAFT, OUTPUT
2. 생성기 에이전트 => 단계별 인스트럭션
3. 출력기
p1. ui/ux
p2. 캐싱
p3. 그래픽 모듈 => asm, os의 경우 가상 아키텍쳐, init status, steps 나열

* 코드 예제 및 실행 기능

태그, 블록
@, ~@, 명시적 출력 비출력 설정
블록 id, keywords

draft -> 전처리 -> [draft blocks] -> ai -> doc
-> rendered output

draft 전처리기:
@ 위치 수집, 위치정보를 기반으로 내용 분량 추정(줄)
블록 id,class, hints 파싱


1. 교육 전반 (가장 큰 시장)

핵심 기능
	•	choice, short_answer (이미 있음) *
	•	fill_blank
	•	matching (연결 문제) 
	•	flashcard *
	•	quiz_group

예시 (영어 학습)


@[flashcard]
+++
front = "abandon"
back = "버리다"
+++

왜 중요한가
	•	CS 말고도 영어, 역사, 과학 전부 적용 가능
	•	“문서 → 학습 콘텐츠”로 확장

👉 이건 개발 외 영역에서도 그대로 통합니다

⸻

2. 지식 정리 / 개인 노트 (Obsidian 영역)

핵심 기능
	•	collapsible *
	•	note, highlight *
	•	link_preview
	•	backlink 
	•	tag

예시

@[note]
+++
type = "idea"
+++
이건 중요한 개념이다.

왜 중요한가
	•	Obsidian, Notion 사용자층 흡수 가능
	•	“정리” + “구조화” 기능

⸻

3. 스토리 / 콘텐츠 (미디어, 블로그)

핵심 기능
	•	spoiler
	•	timeline
	•	character_card
	•	scene

예시

@[spoiler]
이 장면에서 범인은 사실 주인공이다.

왜 중요한가
	•	콘텐츠 제작자용
	•	웹소설, 리뷰, 분석글 등

⸻

4. 의사결정 / 설문 / 선택

핵심 기능
	•	poll
	•	vote
	•	decision_tree
	•	survey

예시

@[poll]
+++
question = "Which option?"
options = ["A", "B", "C"]
+++

왜 중요한가
	•	커뮤니티, 블로그, 교육에서 활용 가능
	•	단순 읽기 → 참여

⸻

5. 생산성 / 체크 / 계획

핵심 기능
	•	todo
	•	checklist
	•	progress
	•	habit_tracker

예시

@[checklist]
+++
items = ["운동", "공부", "독서"]
+++

왜 중요한가
	•	Notion, Todo 앱 영역
	•	개인 생산성 시장

⸻

6. 데이터 설명 / 비즈니스 문서

핵심 기능
	•	chart
	•	metric
	•	kpi
	•	report_section

예시

@[metric]
+++
label = "Revenue"
value = "$120k"
trend = "up"
+++

왜 중요한가
	•	보고서, 프레젠테이션 대체 가능
	•	문서 + 데이터 결합

⸻

7. 인터랙티브 설명 (모든 분야 공통 핵심)

핵심 기능
	•	tabs
	•	toggle
	•	stepper
	•	compare

예시

@[tabs]
+++
labels = ["개념", "예시"]
+++

왜 중요한가
	•	UX 개선
	•	복잡한 설명을 분할

⸻

8. 시각화 (비개발에서도 중요)

핵심 기능
	•	diagram
	•	flow
	•	mindmap

예시
	•	역사 흐름
	•	경제 구조
	•	조직 구조

⸻

핵심 공통 패턴

모든 영역에서 반복되는 구조는 동일합니다.

1. 질문 / 응답
	•	choice
	•	poll
	•	quiz

2. 구조화
	•	collapsible
	•	tabs
	•	steps

3. 시각화
	•	diagram
	•	chart

4. 상태 변화
	•	progress
	•	timeline

5. 참여
	•	vote
	•	input

⸻

중요한 인사이트

이 시스템은 사실 “Markdown 확장”이 아니라

**“문서 기반 UI 정의 언어”**입니다

그래서 적용 분야는:
	•	개발 → 교육 → 콘텐츠 → 생산성 → 비즈니스

까지 자연스럽게 확장됩니다.

⸻

그런데 현실적으로 중요한 것

모든 분야를 동시에 잡으면 망합니다.

전략적으로는 이렇게 가야 합니다

1단계 (초기)
	•	개발 + CS + 학습
→ runnable, choice, collapsible

2단계
	•	일반 교육
→ flashcard, quiz, steps

3단계
	•	생산성 / 콘텐츠
→ note, poll, checklist

⸻

한 줄 결론

개발 영역이 아니어도 핵심은 동일합니다.
사용자가 원하는 건 “읽는 문서”가 아니라 “상호작용하는 문서”입니다.

⸻

원하면 다음으로:
	•	비개발 포함 MVP 블록 15개 세트
	•	실제 사용자 시나리오 기반 문서 예시

까지 구체적으로 설계해드리겠습니다.핵심부터 말하겠습니다.

사용자는 “문법”을 원하는 게 아니라
**“문서 안에서 바로 동작하는 기능(UI + 인터랙션)”**을 원합니다.

그래서 기능 설계는 “키워드” 기준이 아니라
사용자 시나리오 기준으로 나눠야 합니다.

⸻

가장 수요가 확실한 기능 카테고리

1. 학습 / 평가 (가장 중요)

이건 거의 필수입니다.

핵심 블록
	•	choice (객관식)
	•	short_answer
	•	fill_blank
	•	ordering (순서 맞추기)

왜 중요한가
	•	기존 Markdown은 학습 콘텐츠를 못 만듦
	•	이게 있으면 “문서 → 코스/강의”로 확장됨

👉 이건 pionxr의 킬러 기능 후보입니다

⸻

2. 코드 실행 / 실습 (두 번째 핵심)

이미 당신이 runnable을 생각한 이유입니다.

핵심 블록
	•	runnable
	•	terminal
	•	diff_view
	•	step_exec

기능 예
	•	코드 실행
	•	입력/출력 확인
	•	단계별 실행

왜 중요한가
	•	CS/OS/프로그래밍 콘텐츠에서 필수
	•	그냥 설명 → 실습으로 전환

⸻

3. 구조화된 설명 UI

이건 docs 시스템에서 이미 검증된 영역입니다.

핵심 블록
	•	collapsible
	•	note
	•	warning
	•	tip
	•	important

역할
	•	강조
	•	숨김/펼침
	•	시각적 구조화

👉 이건 기본 UX 레이어입니다 (없으면 문서가 밋밋함)

⸻

4. 단계형 설명 / 흐름 제어

이건 의외로 강력합니다.

핵심 블록
	•	steps
	•	step
	•	timeline
	•	progress

예
	•	알고리즘 단계 설명
	•	OS boot 과정
	•	실행 흐름

👉 당신 OS/low-level 콘텐츠랑 특히 잘 맞습니다

⸻

5. 시각화 / 다이어그램

텍스트만으로는 한계가 있습니다.

핵심 블록
	•	diagram
	•	memory_view
	•	graph
	•	state_machine

예
	•	메모리 레이아웃
	•	CPU 상태
	•	데이터 흐름

👉 이건 pionxr 차별화 포인트가 될 수 있습니다

⸻

6. 비교 / diff / before-after

이건 개발자 문서에서 매우 자주 필요합니다.

핵심 블록
	•	diff
	•	compare
	•	before_after

예
	•	코드 변경
	•	버전 비교
	•	최적화 전/후

⸻

7. 데이터 / 표 / 실험 결과

단순 테이블보다 더 강력한 것

핵심 블록
	•	table (강화된)
	•	chart
	•	dataset

예
	•	성능 비교
	•	실험 결과 시각화

⸻

8. 콘텐츠 조합 / 재사용 (중요)

이건 나중에 매우 중요해집니다.

핵심 블록
	•	include
	•	reference
	•	snippet
	•	block reuse

👉 당신이 생각한 block system과 매우 잘 맞음

⸻

9. 인터랙티브 UI 요소

작지만 체감 큼

핵심 블록
	•	tabs
	•	accordion
	•	toggle
	•	switch

예
	•	언어별 코드
	•	OS별 설명

⸻

10. 메타 / 시스템 기능

사용자는 직접 인식하지 않지만 중요

핵심 블록
	•	toc
	•	anchor
	•	index
	•	search_hint

⸻

우선순위 (현실적인 추천)

모든 걸 한 번에 만들면 망합니다.
우선순위는 이렇게 가는 게 맞습니다.

Tier 1 (반드시)
	•	text (기본)
	•	collapsible
	•	choice
	•	short_answer
	•	runnable

👉 이걸로 이미 “기존 Markdown보다 확실히 나은 것”이 됩니다

⸻

Tier 2 (빠르게 추가)
	•	note / warning / tip
	•	tabs
	•	steps

👉 UX가 확 올라갑니다

⸻

Tier 3 (차별화)
	•	diagram
	•	memory_view
	•	diff
	•	terminal

👉 여기서 pionxr만의 색이 생깁니다

⸻

Tier 4 (플랫폼화)
	•	include
	•	dataset
	•	chart
	•	reference

👉 여기서부터는 문서 시스템 → 플랫폼으로 넘어갑니다

⸻

중요한 인사이트

사용자들이 진짜 원하는 것

문법이 아니라 이겁니다:
	1.	설명 + 인터랙션을 한 문서에 넣기
	2.	코드를 “읽는 것”이 아니라 “실행”하기
	3.	학습 콘텐츠를 문서 형태로 만들기
	4.	복잡한 개념을 시각적으로 표현하기

⸻

당신 설계가 좋은 이유

지금 구조는 이걸 자연스럽게 가능하게 합니다:

@[choice]
+++
...
+++

@[runnable]
+++
...
+++

@[collapsible]
...

즉:

문서 = UI 블록 트리

이게 핵심입니다.

⸻

한 줄 결론

사용자가 원하는 건 “마크다운 확장”이 아니라
“문서 안에서 바로 동작하는 인터랙티브 블록들”입니다.

⸻

원하면 다음 단계로:
	•	실제 pionxr 기본 블록 10개 스펙 정의
	•	각 블록의 metadata schema + UX 동작
	•	MVP 기능 세트

까지 바로 구체화해드리겠습니다.

collapsible
note
choice
short_answer

+++ collapsible
#!body

+++

@[keyword]
+++ toml:data

meta = "metadata"

+++ markdown:body



@@@keyword
@@@
#!meta
