# GOLANG MICROSERVICE E COMMERCE PROJECT

## APP PORTS
- 3001 userservice
- 3002 mailerservice
- 3003 loggerservice

## OTHER PORTS
- 5432 postgres
- 5050 pgadmin
- 8080 jenkins
- 27017 mongo
- 8081 mongo-express
- 9200 elasticsearch
- 5601 kibana

# Topic: E-Commerce Project
- User can verify their account through the email received upon registration.
- User can reset their password via email if they forget it.
- Logging is performed when a user logs in.
- Logging is performed when a user makes an incorrect login attempt.
- If a user makes more than 3 incorrect login attempts, their account will be locked. It can be unlocked by an admin.
- User passwords must comply with specified rules.
- Verification codes are deleted by a cron job.
- Incorrect login attempts are reset by a cron job.


- TODO
- [] grpc
- [] rabbitmq
- [] session for user
- [] fiber cookie for user

