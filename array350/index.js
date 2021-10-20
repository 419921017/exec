// 第350题：两个数组的交集
// 给定两个数组，编写一个函数来计算它们的交集。

function intersect1(num1, num2) {
  let map = {};
  num1.forEach((item) => {
    if (map[item]) {
      map[item]++;
    } else {
      map[item] = 1;
    }
  });

  num2.forEach((item) => {
    if (map[item]) {
      map[item]++;
    } else {
      map[item] = 1;
    }
  });

  return Object.entries(map)
    .filter(([key, value]) => value > 1)
    .map((item) => item[0]);
}

// console.log('123', intersect1([1, 2, 2, 3], [1, 2, 3]));

function intersect2(num1, num2) {
  let x = 0,
    y = 0,
    z = 0;

  let res = [];
  num1 = num1.sort((a, b) => a - b);
  num2 = num2.sort((a, b) => a - b);

  while (x < num1.length && y < num2.length) {
    if (num1[x] > num2[y]) {
      y++;
    } else if (num1[x] < num2[y]) {
      x++;
    } else {
      res.push(num1[x]);
      x++;
      y++;
    }
  }
  return res;
}
console.log('123', intersect2([1, 2, 2, 3], [1, 2, 3]));
