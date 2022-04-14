# test

FROM gradle:7.3.2-jdk17-alpine@sha256:8f7ef83b44bd3874d4ec93fefdc3854f144796b9791a0bbe8d643f4387828c43 as build
COPY --chown=gradle:gradle . /home/gradle/src
WORKDIR /home/gradle/src
RUN gradle bootJar --no-daemon

FROM azul/zulu-openjdk-alpine:17.0.1-17.30.15-jre-headless
RUN mkdir /app
RUN addgroup --system javauser && adduser -S -s /bin/false -G javauser javauser
COPY --from=build /home/gradle/src/build/libs/profiles-demo-0.0.1-SNAPSHOT.jar /app/profiles-demo-application.jar
WORKDIR /app
RUN chown -R javauser:javauser /app
USER javauser
CMD [ "/usr/bin/java", "-jar", "profiles-demo-application.jar" ]
