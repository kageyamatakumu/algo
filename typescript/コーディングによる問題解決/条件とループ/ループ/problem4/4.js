// 問題文
// カメのアルルは卓球を 5 試合行いました。
// 試合結果は o と x のみからなる長さ 5 の文字列 S で表され、S の各文字は各試合の勝敗を表します。
// たとえば S が oxxox の場合、アルルの試合結果が順に 勝ち・負け・負け・勝ち・負け であったことを意味します。
// アルルの勝利数を求めてください。
var victoryCount = function (result) {
    var count = 0;
    for (var _i = 0, result_1 = result; _i < result_1.length; _i++) {
        var char = result_1[_i];
        if (char === 'o') {
            count++;
        }
    }
    return count;
    return count;
};
console.log(victoryCount("oxxox"));
console.log(victoryCount("ooooo"));
