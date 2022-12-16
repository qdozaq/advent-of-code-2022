// @ts-check
const fs = require("fs");
const path = require("path");

/** @type {(grid: Record<string, string>) => {left: number, right: number, up:number, down:number}} */
function getBounds(grid) {
  return Object.keys(grid).reduce(
    (bounds, coord) => {
      const [x, y] = coord.split(",").map((n) => parseInt(n));

      let { left, right, up, down } = bounds;

      if (x < left) {
        left = x;
      }
      if (x > right) {
        right = x;
      }

      if (y < down) {
        down = y;
      }

      if (y > up) {
        up = y;
      }
      return { up, down, left, right };
    },
    {
      up: Number.MIN_VALUE,
      down: Number.MAX_VALUE,
      left: Number.MAX_VALUE,
      right: Number.MIN_VALUE,
    }
  );
}

/** @type {(grid: Record<string, string>) => void} */
function printGrid(grid) {
  const { up, down, left, right } = getBounds(grid);

  for (let y = down; y <= up; y++) {
    process.stdout.write(y.toString());
    for (let x = left; x <= right; x++) {
      process.stdout.write(grid[`${x},${y}`] || ".");
    }
    process.stdout.write("\n");
  }
}

function starOne() {
  /** @type {Record<string, string>} */
  const grid = {};

  const lines = fs
    .readFileSync(path.join(__dirname, "input.txt"), { encoding: "utf-8" })
    .split("\n");

  /** @type {number[][]} */
  const coords = [];
  for (let line of lines) {
    const match = line.matchAll(
      /Sensor at x=(-?\d+), y=(-?\d+): closest beacon is at x=(-?\d+), y=(-?\d+)/g
    );
    for (let m of match) {
      const [_, ...coord] = m;
      coords.push(coord.map((n) => parseInt(n)));
    }
  }

  // setup sensors and beacons
  for (let coord of coords) {
    let [sx, sy, bx, by] = coord;
    grid[`${sx},${sy}`] = "S";
    grid[`${bx},${by}`] = "B";
  }

  let total = 0;
  // send out beams
  for (let coord of coords) {
    let [sx, sy, bx, by] = coord;
    // calculate manhatten length
    const beaconDistance = Math.abs(sx - bx) + Math.abs(sy - by);
    const y = 2000000;
    for (let x = sx + beaconDistance; x >= sx - beaconDistance; x--) {
      const distance = Math.abs(sx - x) + Math.abs(sy - y);
      if (distance <= beaconDistance) {
        const pos = grid[`${x},${y}`];
        if (!pos || pos === "S") {
          total++;
          grid[`${x},${y}`] = "#";
        }
      }
    }
  }

  console.log(total);
}

starOne();

function starTwo() {
  /** @type {Record<string, string>} */
  const grid = {};

  const maxPos = 4000000;
  // const maxPos = 20;
  const lines = fs
    .readFileSync(path.join(__dirname, "input.txt"), { encoding: "utf-8" })
    .split("\n");

  /** @type {Array<{sensor:number[], beacon: number[], range: number}>} */
  const data = [];
  for (let line of lines) {
    const match = line.matchAll(
      /Sensor at x=(-?\d+), y=(-?\d+): closest beacon is at x=(-?\d+), y=(-?\d+)/g
    );
    for (let m of match) {
      const [_, ...coord] = m;
      const [sx, sy, bx, by] = coord.map((n) => parseInt(n));
      const range = Math.abs(sx - bx) + Math.abs(sy - by);

      data.push({ sensor: [sx, sy], beacon: [bx, by], range });
    }
  }

  outer: for (let d of data) {
    // go around perimeter
    let top = Math.min(d.sensor[1] + d.range, maxPos);
    let bot = Math.max(d.sensor[1] - d.range, 0);

    // check perimeter
    for (let y = bot; y <= top; y++) {
      const distance = Math.abs(d.sensor[1] - y);
      const distFromCenter = d.range - distance;
      const x1 = d.sensor[0] - distFromCenter - 1;
      const x2 = d.sensor[0] + distFromCenter + 1;

      for (let x of [x1, x2]) {
        let maybe = true;
        for (let d of data) {
          let range = Math.abs(d.sensor[0] - x) + Math.abs(d.sensor[1] - y);
          if (range <= d.range || x < 0 || x > maxPos) {
            maybe = false;
            break;
          }
        }

        if (maybe) {
          const freq = x * maxPos + y;
          console.log(x, y, freq);
          break outer;
        }
      }
    }
  }
}

starTwo();
