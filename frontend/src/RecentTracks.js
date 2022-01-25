import { useEffect, useState } from 'react';
import { useQuery, gql } from '@apollo/client';
import Metadata from './Metadata.js';

export default function RecentTracks({ stationId }) {
  const { loading, error, data } = useQuery(gql`
    query RecentlyPlayed($stationId: Int!) {
      allTrackEvents(condition: { stationId: $stationId, action: "played"}, orderBy: CREATED_AT_DESC, first: 5) {
        edges {
          node {
            createdAt
            stationId
            id
            action
            trackByTrackId {
              id
              filename
              metadata
            }
          }
        }
      }
    }
  `, {
    pollInterval: 10000,
    variables: { stationId }
  });

  if (loading) {
    return 'spinner'
  }

  const { edges } = data.allTrackEvents

  if (edges.length) {
    return (
      <div>
        <b>
          <Metadata metadata={edges[0].node.trackByTrackId?.metadata} />
        </b>
        {edges.slice(1).map(({ node }) => (
          <div key={node.id}>
            <Metadata metadata={node.trackByTrackId?.metadata} />
          </div>
        ))}
      </div>
    )
  } else {
    return 'no edges'
  }
}

