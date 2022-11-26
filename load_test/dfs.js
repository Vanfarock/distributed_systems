import http from 'k6/http';

export default function () {
  const parameters = {
    cities: [
      {
        id: 1,
        neighbors: [
          {
            id: 2
          },
          {
            id: 3
          }
        ]
      },
      {
        id: 2,
        neighbors: [
          {
            id: 1
          },
          {
            id: 3
          },
          {
            id: 4
          }
        ]
      },
      {
        id: 3,
        neighbors: [
          {
            id: 1
          },
          {
            id: 2
          },
          {
            id: 4
          },
          {
            id: 5
          }
        ]
      },
      {
        id: 4,
        neighbors: [
          {
            id: 2
          },
          {
            id: 3
          },
          {
            id: 5
          }
        ]
      },
      {
        id: 5,
        neighbors: [
          {
            id: 3
          },
          {
            id: 4
          }
        ]
      }
    ],
    startIndex: 0,
    endIndex: 4
  }

  const json = JSON.stringify(parameters)

  http.post('http://localhost/dfs/dfs', json);
}