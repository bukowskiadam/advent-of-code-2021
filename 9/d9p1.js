const fs = require("fs");

const lines = fs.readFileSync("input.txt").toString().split("\n");

let riskLevel = 0;

const v = (i, j) => {
  if (i < 0 || j < 0 || i >= lines.length || j >= lines[0].length) {
    return 10;
  }
  return lines[i][j];
};

for (let i = 0; i < lines.length; i++) {
  for (let j = 0; j < lines[i].length; j++) {
    if (
      v(i - 1, j) > lines[i][j] &&
      v(i, j - 1) > lines[i][j] &&
      v(i + 1, j) > lines[i][j] &&
      v(i, j + 1) > lines[i][j]
    ) {
      console.log(`x=${j} y=${i} val=${parseInt(lines[i][j], 10)}`);
      riskLevel += parseInt(lines[i][j], 10) + 1;
    }
  }
}

console.log(riskLevel);
