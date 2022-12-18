// @ts-check
const fs = require("fs");
const path = require("path");

const planes = {};

const offset = 0.5;

function addPlane(x, y, z) {
  if (planes[`${x},${y},${z}`]) {
    planes[`${x},${y},${z}`]++;
  } else {
    planes[`${x},${y},${z}`] = 1;
  }
}

fs.readFileSync(path.join(__dirname, "input.txt"), { encoding: "utf-8" })
  .split("\n")
  .forEach((line) => {
    const [x, y, z] = line.split(",").map((n) => parseInt(n));

    //x
    addPlane(x + offset, y, z);
    addPlane(x - offset, y, z);

    // y
    addPlane(x, y + offset, z);
    addPlane(x, y - offset, z);

    // z
    addPlane(x, y, z + offset);
    addPlane(x, y, z - offset);
  });

let surfaceArea = 0;

for (let num of Object.values(planes)) {
  if (num === 1) {
    surfaceArea++;
  }
}

console.log(surfaceArea);
