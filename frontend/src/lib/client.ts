import { GraphQLClient } from 'graphql-request'

const ENDPOINT =
  import.meta.env.VITE_GRAPHQL_ENDPOINT ?? 'http://localhost:4350/graphql'

export const gqlClient = new GraphQLClient(ENDPOINT)
export const GRAPHQL_ENDPOINT = ENDPOINT
