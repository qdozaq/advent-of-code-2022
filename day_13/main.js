// @ts-check
const fs = require("fs");
const path = require("path");

/**
 * @typedef {Array<number|Packet>} Packet
 * @typedef {{left:Packet, right:Packet}[]} PacketPairs
 */

/** @type {(left: Packet, right: Packet) => number} */
function inOrder(left, right) {
  for (let i = 0; i < left.length; i++) {
    const l = left[i];
    const r = right[i];

    // ran out of items
    if (r === undefined) {
      return -1;
    }

    if (typeof l === "number" && typeof r === "number") {
      if (l === r) {
        continue;
      } else {
        return r - l;
      }
    }

    if (Array.isArray(l) && Array.isArray(r)) {
      const result = inOrder(l, r);
      if (result === 0) continue;
      return result;
    }

    // mixed types
    if (typeof l === "number" && Array.isArray(r)) {
      const result = inOrder([l], r);
      if (result === 0) continue;
      return result;
    }

    if (Array.isArray(l) && typeof r === "number") {
      const result = inOrder(l, [r]);
      if (result === 0) continue;
      return result;
    }
  }

  if (right.length > left.length) {
    return 1;
  }

  return 0;
}

function starOne() {
  /** @type {PacketPairs} */
  const pairs = fs
    .readFileSync(path.join(__dirname, "./input.txt"), { encoding: "utf-8" })
    .split("\n\n")
    .map((pair) => {
      const [left, right] = pair.split("\n").map((arr) => {
        return JSON.parse(`{"data": ${arr}}`).data;
      });
      return { left, right };
    });

  let i = 1;
  let total = 0;
  for (let { left, right } of pairs) {
    const result = inOrder(left, right);
    if (result > 0) {
      //   console.log(i);
      total += i;
    }

    i++;
  }

  console.log(total);
}

starOne();

function starTwo() {
  /** @type {Packet[]} */
  const packets = fs
    .readFileSync(path.join(__dirname, "./input.txt"), { encoding: "utf-8" })
    .split("\n")
    .filter((line) => line.length > 0)
    .map((line) => {
      return JSON.parse(`{"data": ${line}}`).data;
    });

  // add dividers
  packets.push([[2]], [[6]]);

  packets.sort((left, right) => inOrder(right, left));

  let total = 1;
  let i = 1;
  for (let packet of packets) {
    const s = JSON.stringify(packet);
    if (s === "[[2]]" || s === "[[6]]") {
      total *= i;
    }
    i++;
  }

  console.log(total);
}

starTwo();
