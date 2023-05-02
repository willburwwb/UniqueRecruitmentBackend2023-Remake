FROM redis:latest
RUN --mount=type=bind,source="../../../hrdata/redis",target="/data" \ 
  && --network=uniquerecruitmentbackend2023-remake_database
EXPOSE 10001
