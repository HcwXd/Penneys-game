const input_a = document.querySelector('#seq_A');
const input_b = document.querySelector('#seq_B');
const input_head_prob = document.querySelector('#head_prob');
const input_rep_times = document.querySelector('#rep_times');
const start_btn = document.querySelector('.start_btn');
const result_wrap = document.querySelector('.result_wrap');
const result_display = document.querySelector('.result_display');

const randomToss = (headProb) => (Math.random() < headProb ? 'H' : 'T');

const opB = (flipFirst = true) => {
  const seqA = input_a.value.toUpperCase();
  let seqB = flipFirst ? (seqA[seqA.length - 1] === 'H' ? 'T' : 'H') : seqA[seqA.length - 1];
  seqB += seqA.slice(0, seqA.length - 1);
  input_b.value = seqB;
};
let bestB = [];

const gen = (len = 5) => {
  let ans = [];
  per('');
  function per(par) {
    if (par.length === len) {
      ans.push(par);
      return;
    } else {
      per(par + 'T');
      per(par + 'H');
    }
  }
  for (let a of ans) {
    let data = [];
    for (let b of ans) {
      data.push(playSingleMatch(a, b));
    }
    data.sort((a, b) => a[0] - b[0]);
    bestB.push(data[0]);
  }
  bestB.sort((a, b) => a[0] - b[0]);
};

function playSingleMatch(seqA, seqB, repTimes = 1000, headProb = 0.5) {
  let aWiningCnt = 0;
  for (let idx = 0; idx < repTimes; idx++) {
    let curSeq = [];
    while (true) {
      curSeq.push(randomToss(headProb));
      if (curSeq.length > seqA.length) curSeq.shift();
      if (curSeq.join('') === seqA) {
        aWiningCnt++;
        break;
      }
      if (curSeq.join('') === seqB) break;
    }
  }
  return [aWiningCnt / repTimes, `${seqA} vs ${seqB}`];
}

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

function Con(a, b) {
  let i = L(b, b) - L(b, a);
  let j = L(a, a) - L(a, b);
  return i / (i + j);
}

start_btn.addEventListener('click', () => {
  const seqA = input_a.value.toUpperCase();
  const seqB = input_b.value.toUpperCase();
  const headProb = +input_head_prob.value;
  const repTimes = +input_rep_times.value;
  if (seqA.length !== seqB.length) {
    alert('Sequence length must be the same');
    return;
  }
  if (headProb >= 1 || headProb <= 0) {
    alert('Head prob should be smaller than 1, bigger than 0');
    return;
  }
  result_display.innerHTML = 'Running';
  let aWiningCnt = 0;
  window.setTimeout(() => {
    for (let idx = 0; idx < repTimes; idx++) {
      let curSeq = [];
      while (true) {
        curSeq.push(randomToss(headProb));
        if (curSeq.length > seqA.length) curSeq.shift();
        if (curSeq.join('') === seqA) {
          aWiningCnt++;
          break;
        }
        if (curSeq.join('') === seqB) break;
      }
    }
    console.log(Con(seqA, seqB));
    result_display.innerHTML = `${((aWiningCnt / repTimes) * 100).toFixed(1)}%`;
  }, 1);
});
