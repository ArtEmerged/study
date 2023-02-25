curl -s https://01.alem.school/api/graphql-engine/v1/graphql --data '{"query":"{user(where:{login:{_eq:\"aecheist\"}}){id}}"}' | tr -d -c 0-9


