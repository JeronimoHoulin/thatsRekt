import { gqlClient } from './client'
import { mockFetchFeed, mockFetchPostDetail } from './mock'

const USE_MOCK = import.meta.env.VITE_USE_MOCK_DATA === 'true'
export const IS_MOCK_MODE = USE_MOCK

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

// ---- sort options exposed in the UI ----

export type SortOption = 'newest' | 'oldest'

const SORT_TO_ORDER_BY: Record<SortOption, string> = {
  newest: 'createdAtBlock_DESC',
  oldest: 'createdAtBlock_ASC',
}

const FEED_QUERY = /* GraphQL */ `
  query Feed($limit: Int!, $orderBy: [PostOrderByInput!]!) {
    posts(orderBy: $orderBy, limit: $limit, where: { removed_eq: false }) {
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

export async function fetchFeed(
  limit = 50,
  sort: SortOption = 'newest',
): Promise<FeedPost[]> {
  if (USE_MOCK) return mockFetchFeed(limit, sort)
  const data = await gqlClient.request<{ posts: FeedPost[] }>(FEED_QUERY, {
    limit,
    orderBy: [SORT_TO_ORDER_BY[sort]],
  })
  return data.posts
}

export async function fetchPostDetail(id: string): Promise<PostDetail | null> {
  if (USE_MOCK) return mockFetchPostDetail(id)
  const data = await gqlClient.request<{ postById: PostDetail | null }>(POST_DETAIL_QUERY, { id })
  return data.postById
}
