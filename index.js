const input_a = document.querySelector('#seq_A');
const input_b = document.querySelector('#seq_B');
const input_head_prob = document.querySelector('#head_prob');
const input_rep_times = document.querySelector('#rep_times');
const start_btn = document.querySelector('.start_btn');
const result_wrap = document.querySelector('.result_wrap');
const result_display = document.querySelector('.result_display');

start_btn.addEventListener('click', countProb);
const randomToss = (headProb) => (Math.random() < headProb ? 'H' : 'T');

function countProb() {
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
  console.log(seqA, seqB, headProb, repTimes);
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
    result_display.innerHTML = `${((aWiningCnt / repTimes) * 100).toFixed(1)}%`;
  }, 1);
}
