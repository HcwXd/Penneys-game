let a = 'HHHHTH';
let b = 'THHHHH';

function L(seqA, seqB) {
  let ans = 0;
  let len = seqA.length;
  for (let idx = 0; idx < len; idx++) {
    if (seqA.slice(idx, len) === seqB.slice(0, len - idx)) {
      ans += 2 ** (len - idx - 1);
    }
  }
  return ans;
}

function Conway(a, b) {
  let i = L(b, b) - L(b, a);
  let j = L(a, a) - L(a, b);
  return i / (i + j);
}

console.log(Conway(a, b));
