// @ts-check
const fs = require("fs");
const path = require("path");

/** @type {() => Record<string, string>} */
function createGrid() {
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

  //   for (let line of lines) {
  //     const match = line.matchAll(
  //       /Sensor at x=(-?\d+), y=(-?\d+): closest beacon is at x=(-?\d+), y=(-?\d+)/g
  //     );
  //     for (let m of match) {
  //       const sx = parseInt(m[1]);
  //       const sy = parseInt(m[2]);
  //       const bx = parseInt(m[3]);
  //       const by = parseInt(m[4]);

  //       grid[`${sx},${sy}`] = "S";
  //       grid[`${bx},${by}`] = "B";

  //       // calculate manhatten length
  //       const beaconDistance = Math.abs(sx - bx) + Math.abs(sy - by);

  //       //   for (let x = sx + beaconDistance; x >= sx - beaconDistance; x--) {
  //       //     for (let y = sy + beaconDistance; y >= sy - beaconDistance; y--) {
  //       //       const distance = Math.abs(sx - x) + Math.abs(sy - y);
  //       //       if (distance <= beaconDistance && !grid[`${x},${y}`]) {
  //       //         grid[`${x},${y}`] = "#";
  //       //       }
  //       //     }
  //       //   }
  //       const y = 2000000;
  //       for (let x = sx + beaconDistance; x >= sx - beaconDistance; x--) {
  //         const distance = Math.abs(sx - x) + Math.abs(sy - y);
  //         if (distance <= beaconDistance) {
  //           grid[`${x},${y}`] = "#";
  //         }
  //       }
  //     }
  //   }

  return grid;
}

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

// starOne();

function starTwo() {
  /** @type {Record<string, string>} */
  const grid = {};

  const maxPos = 4000000;
  //   const maxPos = 20;
  const lines = fs
    .readFileSync(path.join(__dirname, "input.txt"), { encoding: "utf-8" })
    .split("\n");

  let top = Number.MIN_VALUE,
    bot = Number.MAX_VALUE,
    left = Number.MAX_VALUE,
    right = Number.MIN_VALUE;

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

      if (sx + range > right) {
        right = sx + range;
      }

      if (sx - range < left) {
        left = sx - range;
      }

      if (sy + range > top) {
        top = sy + range;
      }

      if (sy - range < bot) {
        bot = sy - range;
      }

      data.push({ sensor: [sx, sy], beacon: [bx, by], range });
    }
  }

  // shrink search area
  if (top > maxPos) {
    top = maxPos;
  }
  if (bot < 0) {
    bot = 0;
  }
  if (left < 0) {
    left = 0;
  }

  if (right > maxPos) {
    right = maxPos;
  }

  console.log({ top, bot, left, right });

  /** @type {number[][]} */
  const canidate = [];
  for (let y = bot; y <= top; y++) {
    for (let x = left; x <= right; x++) {
      let maybe = true;
      for (let d of data) {
        const range = Math.abs(d.sensor[0] - x) + Math.abs(d.sensor[1] - y);
        if (range <= d.range) {
          maybe = false;
          break;
        }
      }

      if (maybe) {
        canidate.push([x, y]);
      }
    }
  }

  console.log(canidate);

  // setup sensors and beacons
  //   for (let coord of coords) {
  //     let {sensor, beacon, range} = coord;
  //     grid[`${sx},${sy}`] = "S";
  //     grid[`${bx},${by}`] = "B";
  //   }

  //   let total = 0;
  //   // send out beams
  //   for (let coord of coords) {
  //     let [sx, sy, bx, by] = coord;
  //     // calculate manhatten length
  //     const beaconDistance = Math.abs(sx - bx) + Math.abs(sy - by);

  //     for (let x = sx + beaconDistance; x >= sx - beaconDistance; x--) {
  //       for (let y = sy + beaconDistance; y >= sy - beaconDistance; y--) {
  //         const distance = Math.abs(sx - x) + Math.abs(sy - y);
  //         if (distance <= beaconDistance && !grid[`${x},${y}`]) {
  //           grid[`${x},${y}`] = "#";
  //         }
  //       }
  //     }
  //   }

  //   printGrid(grid);

  //   console.log(total);
}

starTwo();
