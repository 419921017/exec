let stations = {};

const statesNeeded = new Set([1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13]);

stations['kone'] = new Set([1, 2, 3]);
stations['ktwo'] = new Set([1, 2, 3]);
stations['kthree'] = new Set([1, 2, 3]);
stations['kfour'] = new Set([1, 2, 3]);
stations['kive'] = new Set([1, 2, 3]);

const finalStations = new Set();

const indStations = new Set();

while (statesNeeded.size > 0) {
  const bestStations = new Set();
  const statesCovered = new Set();

  for (const station in stations) {
    const stateForStation = stations[station];
    // console.log('station', station);
    // console.log('stateForStation', stateForStation);
    const covered = statesNeed & stateForStation; // 交集
    if (covered.size > statesCovered.size) {
      bestStations = station;
      statesCovered = covered;
    }
  }
  // 删除已有的
  for (const station in statesCovered) {
    statesNeeded.delete(station);
  }
  // statesNeeded -= statesCovered;
  finalStations.add(bestStations);
}
