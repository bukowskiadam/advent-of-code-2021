const fs = require("fs");

const content = fs.readFileSync("input.txt");

const lines = content
  .toString()
  .split("\n")
  .map((x) => parseInt(x, 10));

const sum = (idx) => lines[idx] + lines[idx - 1] + lines[idx - 2];
let last = sum(2);
let increased = 0;

for (let i = 3; i < lines.length; i += 1) {
  const current = sum(i);
  if (current > last) {
    increased += 1;
  }
  last = current;
}
console.log(increased);

// answers: 1301
// answers: 1346
