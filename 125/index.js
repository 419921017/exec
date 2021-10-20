/**
 * @param {string} s
 * @return {boolean}
 */
var isPalindrome = function (s) {
  s = s.toLocaleLowerCase();
  let left = 0;
  let right = s.length - 1;
  while (left < right) {
    while (left < right && !isalnum(s[left])) {
      left++;
    }
    while (left < right && !isalnum(s[right])) {
      right--;
    }
    if (left < right) {
      if (s[left] != s[right]) {
        return false;
      }
      left++;
      right--;
    }
  }
  return true;
};

function isalnum(s) {
  return /[a-zA-Z0-9]/.test(s);
}
