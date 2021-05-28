/**
 * @param {number[]} digits
 * @return {number[]}
 */
var plusOne = function (digits) {
  // 这里不能用Number, 有数字上限, 必须使用BigInt
  var num = BigInt(digits.join(''));
  // console.log(num);

  var newNum = num + BigInt(1);
  // console.log(newNum);

  var newArr = String(newNum)
    .split('')
    .map(function (n) {
      return Number(n);
    });

  // console.log(newArr);

  return newArr;
};

plusOne([6, 1, 4, 5, 3, 9, 0, 1, 9, 5, 1, 8, 6, 7, 0, 5, 5, 4, 3]);
