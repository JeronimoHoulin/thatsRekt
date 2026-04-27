import { gqlClient } from './client'

// ---- shared types (mirror schema.graphql) ----

export type VoteDirection = 'None' | 'Upvote' | 'Downvote'
export type EditKind = 'AmendNote' | 'AddAttackers' | 'AddVictims'

export interface AddressEntity {
  id: string
  attackerScore: string
  attackerAppearances?: number
  isVictim?: boolean
}

export interface PostAttackerLink {
  address: AddressEntity
}

export interface PostVictimLink {
  address: AddressEntity
}

export interface FeedPost {
  id: string
  poster: { id: string }
  attackedAt: string
  note: string
  upvotes: number
  downvotes: number
  netScore: number
  createdAtTimestamp: string
  attackerLinks: PostAttackerLink[]
  victimLinks: PostVictimLink[]
}

export interface VoteEntity {
  id: string
  voter: { id: string }
  oldDirection: VoteDirection
  newDirection: VoteDirection
  blockNumber: number
  timestamp: string
}

export interface EditEntity {
  id: string
  kind: EditKind
  newNote: string | null
  addedAttackers: string[] | null
  addedVictims: string[] | null
  blockNumber: number
  timestamp: string
}

export interface PostDetail {
  id: string
  poster: { id: string }
  attackedAt: string
  lastUpdatedAt: string
  note: string
  upvotes: number
  downvotes: number
  netScore: number
  removed: boolean
  createdAtTimestamp: string
  removedAtTimestamp: string | null
  attackerLinks: PostAttackerLink[]
  victimLinks: PostVictimLink[]
  votes: VoteEntity[]
  edits: EditEntity[]
}

// ---- queries ----

const FEED_QUERY = /* GraphQL */ `
  query Feed($limit: Int!) {
    posts(orderBy: createdAtBlock_DESC, limit: $limit, where: { removed_eq: false }) {
      id
      poster {
        id
      }
      attackedAt
      note
      upvotes
      downvotes
      netScore
      createdAtTimestamp
      attackerLinks {
        address {
          id
          attackerScore
        }
      }
      victimLinks {
        address {
          id
        }
      }
    }
  }
`

const POST_DETAIL_QUERY = /* GraphQL */ `
  query PostDetail($id: String!) {
    postById(id: $id) {
      id
      poster {
        id
      }
      attackedAt
      lastUpdatedAt
      note
      upvotes
      downvotes
      netScore
      netScore
      removed
      createdAtTimestamp
      removedAtTimestamp
      attackerLinks {
        address {
          id
          attackerScore
          attackerAppearances
        }
      }
      victimLinks {
        address {
          id
          isVictim
        }
      }
      votes(orderBy: blockNumber_ASC) {
        id
        voter {
          id
        }
        oldDirection
        newDirection
        blockNumber
        timestamp
      }
      edits(orderBy: blockNumber_ASC) {
        id
        kind
        newNote
        addedAttackers
        addedVictims
        blockNumber
        timestamp
      }
    }
  }
`

export async function fetchFeed(limit = 50): Promise<FeedPost[]> {
  const data = await gqlClient.request<{ posts: FeedPost[] }>(FEED_QUERY, { limit })
  return data.posts
}

export async function fetchPostDetail(id: string): Promise<PostDetail | null> {
  const data = await gqlClient.request<{ postById: PostDetail | null }>(POST_DETAIL_QUERY, { id })
  return data.postById
}
