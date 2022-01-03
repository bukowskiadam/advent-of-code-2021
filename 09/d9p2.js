const fs = require("fs");

const lines = fs.readFileSync("input.txt").toString().split("\n");

const visited = lines.map((line) => new Array(line.length).fill(false));

const v = (i, j) => {
  if (i < 0 || j < 0 || i >= lines.length || j >= lines[0].length) {
    return 10;
  }
  return lines[i][j];
};

const basins = [];

const wasVisited = (i, j) => {
  if (i < 0 || j < 0 || i >= lines.length || j >= lines[0].length) {
    return true;
  }
  return visited[i][j];
};

const markVisited = (i, j) => {
  visited[i][j] = true;
};

const calculateBasinSize = (i, j) => {
  markVisited(i, j);
  let size = 1;
  const step = (x, y) => {
    if (!wasVisited(x, y) && v(x, y) < 9) {
      size += calculateBasinSize(x, y);
    }
  };

  step(i - 1, j);
  step(i + 1, j);
  step(i, j - 1);
  step(i, j + 1);
  return size;
};

for (let i = 0; i < lines.length; i++) {
  for (let j = 0; j < lines[i].length; j++) {
    if (wasVisited(i, j) || lines[i][j] === "9") {
      continue;
    }
    const basinSize = calculateBasinSize(i, j);
    basins.push(basinSize);
  }
}

const sorted = basins.sort((a, b) => b - a);

const ret = sorted.slice(0, 3).reduce((prev, now) => prev * now, 1);

console.log(ret);
