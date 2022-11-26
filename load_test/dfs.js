import http from 'k6/http';

export default function () {
  const MAX_CITIES = 200
  const MIN_CITIES = 80
  const MIN_NEIGHBORS = 1

  const cities = []
  const maxCities = Math.min(Math.floor(Math.random() * MAX_CITIES), MIN_CITIES);
  for (let i = 1; i <= maxCities; i++) {
    const neighbors = []
    const maxNeighbors = Math.min(Math.floor(Math.random() * maxCities), MIN_NEIGHBORS)

    for (let j = 1; j <= maxNeighbors; j++) {
      if (i == j) {
        continue
      }

      neighbors.push({ id: j })
    }

    cities.push({
      id: i,
      neighbors: neighbors
    })
  }

  const parameters = {
    cities: cities,
    startIndex: Math.floor(Math.random() * maxCities),
    endIndex: Math.floor(Math.random() * maxCities)
  }

  // const parameters = {
  //   cities: [
  //     {
  //       id: 1,
  //       neighbors: [
  //         {
  //           id: 2
  //         },
  //         {
  //           id: 3
  //         }
  //       ]
  //     },
  //     {
  //       id: 2,
  //       neighbors: [
  //         {
  //           id: 1
  //         },
  //         {
  //           id: 3
  //         },
  //         {
  //           id: 4
  //         }
  //       ]
  //     },
  //     {
  //       id: 3,
  //       neighbors: [
  //         {
  //           id: 1
  //         },
  //         {
  //           id: 2
  //         },
  //         {
  //           id: 4
  //         },
  //         {
  //           id: 5
  //         }
  //       ]
  //     },
  //     {
  //       id: 4,
  //       neighbors: [
  //         {
  //           id: 2
  //         },
  //         {
  //           id: 3
  //         },
  //         {
  //           id: 5
  //         }
  //       ]
  //     },
  //     {
  //       id: 5,
  //       neighbors: [
  //         {
  //           id: 3
  //         },
  //         {
  //           id: 4
  //         }
  //       ]
  //     }
  //   ],
  //   startIndex: 0,
  //   endIndex: 4
  // }

  const json = JSON.stringify(parameters)

  http.post('http://localhost/dfs/dfs', json);
}