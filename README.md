# Google Authenticator, Cron을 활용한 Scrapping 자동화 서버

새로운 QR인증 관리에 대해서 궁금한 부분이 있어서 작업이 진행이 되었다.

기본적으로 MySQL을 통해 3티어 아키텍처를 구성하고, Cron을 활용하여 주기적으로 MySQL에 있는 데이터를 긁어와 Scrapping을 사용하는 API서버.

# 기능 정리

다음과 같은 기능으로 정리가 된다.

1. Google Authenticator를 활용한 QR코드 인증 방식 도입
2. Cron을 활용하여, 백그라운드 서비스로직 구현
3. MySQL을 활용하여, API 데이터 관리

# MySQL Table

MySQL의 테이블은 다음과 같이 간단하게 구성이 되어 있다.

- Scrapping을 위한 데이터만을 저장하기 위한 테이블
