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

  for (let line of lines) {
    const match = line.matchAll(/(\d+),(\d+)/g);
    let previousCoords = null;
    for (let m of match) {
      const x = parseInt(m[1]);
      const y = parseInt(m[2]);

      grid[`${x},${y}`] = "#";

      if (previousCoords) {
        let [px, py] = previousCoords;

        const dx = Math.sign(x - px);
        const dy = Math.sign(y - py);

        while (px !== x || py !== y) {
          grid[`${px},${py}`] = "#";

          px += dx;
          py += dy;
        }
      }
      previousCoords = [x, y];
    }
  }

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
    for (let x = left; x <= right; x++) {
      process.stdout.write(grid[`${x},${y}`] || ".");
    }
    process.stdout.write("\n");
  }
}

function starOne() {
  const grid = createGrid();
  printGrid(grid);

  let overflow = false;

  let grains = 0;

  const { up } = getBounds(grid);

  while (!overflow) {
    let canMove = true;
    let sx = 500,
      sy = 0;
    while (canMove) {
      if (sy > up) {
        overflow = true;
        canMove = false;
      }

      //drop sand
      if (!grid[`${sx},${sy + 1}`]) {
        sy++;
        continue;
      }

      // down to the left
      if (!grid[`${sx - 1},${sy + 1}`]) {
        sx--;
        sy++;
        continue;
      }

      // down right
      if (!grid[`${sx + 1},${sy + 1}`]) {
        sx++;
        sy++;
        continue;
      }

      canMove = false;
      grid[`${sx},${sy}`] = "O";
      grains++;
    }
  }

  printGrid(grid);
  console.log(grains);
}

starOne();

function starTwo() {
  const grid = createGrid();
  printGrid(grid);

  let full = false;

  let grains = 0;

  const { up } = getBounds(grid);

  while (!full) {
    let canMove = true;
    let sx = 500,
      sy = 0;
    while (canMove) {
      // all full
      if (grid[`${sx},${sy}`]) {
        full = true;
        canMove = false;
        break;
      }

      // hit the floor
      if (sy + 1 >= up + 2) {
        canMove = false;
        grid[`${sx},${sy}`] = "O";
        grains++;
        continue;
      }

      //drop sand
      if (!grid[`${sx},${sy + 1}`]) {
        sy++;
        continue;
      }

      // down to the left
      if (!grid[`${sx - 1},${sy + 1}`]) {
        sx--;
        sy++;
        continue;
      }

      // down right
      if (!grid[`${sx + 1},${sy + 1}`]) {
        sx++;
        sy++;
        continue;
      }

      canMove = false;
      grid[`${sx},${sy}`] = "O";
      grains++;
    }
  }

  printGrid(grid);
  console.log(grains);
}
starTwo();
