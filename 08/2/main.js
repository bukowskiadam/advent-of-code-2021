const fs = require("fs");

const lines = fs.readFileSync("8/input.txt").toString().split("\n");

const sort = (s) => s.split("").sort().join("");
const includes = (s1, s2) =>
  s2.split("").every((letter) => s1.includes(letter));

let sum = 0;
for (let line of lines) {
  if (!line) continue;
  const [left, right] = line.split(" | ").map((x) => x.split(" "));
  const dict = {};
  // built dict of know digits
  left.forEach((digit) => {
    if (digit.length === 2) {
      dict[1] = sort(digit);
    } else if (digit.length === 3) {
      dict[7] = sort(digit);
    } else if (digit.length === 4) {
      dict[4] = sort(digit);
    } else if (digit.length === 7) {
      dict[8] = sort(digit);
    }
  });
  left.forEach((digit) => {
    const sorted = sort(digit);
    if (digit.length === 6) {
      // case for 6, 9, 0
      if (includes(sorted, dict[4])) {
        dict[9] = sorted;
      }
    }
  });
  left.forEach((digit) => {
    const sorted = sort(digit);
    if (digit.length == 5) {
      // case for 2, 5, 3
      if (includes(sorted, dict[1])) {
        dict[3] = sorted;
      } else if (includes(dict[9], sorted)) {
        dict[5] = sorted;
      } else {
        dict[2] = sorted;
      }
    }
  });
  left.forEach((digit) => {
    const sorted = sort(digit);
    if (digit.length == 6) {
      // case for 6, 9, 0
      if (!includes(sorted, dict[4])) {
        if (includes(sorted, dict[5])) {
          dict[6] = sorted;
        } else {
          dict[0] = sorted;
        }
      }
    }
  });

  console.log(
    dict,
    right.map((x) => sort(x))
  );

  invertedDict = Object.fromEntries(
    Object.entries(dict).map(([k, v]) => [v, k])
  );

  const str = right.map((letters) => invertedDict[sort(letters)]).join("");
  const val = parseInt(str, 10);
  console.log(str, val);
  sum += val;
}

console.log(sum);
