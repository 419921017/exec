//
let graph = {};

graph['start'] = {};
graph['start']['a'] = 6;
graph['start']['b'] = 2;

graph['a'] = {};
graph['a']['fin'] = 1;

graph['b'] = {};
graph['b']['a'] = 3;
graph['b']['fin'] = 5;

graph['fin'] = {};

//
let costs = {};
costs['a'] = 6;
costs['b'] = 2;
costs['fin'] = Infinity;

let parents = {};

parents['a'] = 'start';
parents['b'] = 'start';
parents['fin'] = null;

let processed = [];

let node;

node = findLowestCostNode(costs);

while (node) {
  let cost = costs[node];
  let neighbors = graph[node];
  Object.keys(neighbors).forEach((n) => {
    let newCost = cost + neighbors[n];
    if (costs[n] > newCost) {
      costs[n] = newCost;
      parents[n] = node;
    }
  });
  processed.push(node);
  node = findLowestCostNode(costs);
}

console.log('costs', costs);
console.log('parents', parents);

function findLowestCostNode(costs) {
  let lowestCost = Infinity;
  let lowestCostNode = null;
  for (const node in costs) {
    let cost = costs[node];
    if (cost < lowestCost && !processed.includes(node)) {
      lowestCost = cost;
      lowestCostNode = node;
    }
  }
  return lowestCostNode;
}
